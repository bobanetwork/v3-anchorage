package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mattn/go-isatty"
	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum-optimism/optimism/boba-chain-ops/ether"
	"github.com/ethereum-optimism/optimism/boba-chain-ops/genesis"
	"github.com/ethereum-optimism/optimism/op-chain-ops/db"
)

func main() {
	log.Root().SetHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(isatty.IsTerminal(os.Stderr.Fd()))))

	app := &cli.App{
		Name:  "eth-crawler",
		Usage: "Make RPC calls to an Ethereum node to get all addresses that have transferred from or to ETH",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "rpc-url",
				Usage:    "RPC URL for an Ethereum Node",
				Required: true,
				EnvVars:  []string{"RPC_URL"},
			},
			&cli.Int64Flag{
				Name:    "end-block",
				Value:   0,
				Usage:   "Block to end at",
				EnvVars: []string{"END_BLOCK"},
			},
			&cli.StringFlag{
				Name:  "output-path",
				Value: "eth-account.json",
				Usage: "Path to the output file",
			},
			&cli.DurationFlag{
				Name:  "rpc-timeout",
				Value: 5 * time.Second,
				Usage: "Timeout for RPC calls",
			},
			&cli.DurationFlag{
				Name:  "rpc-poll-interval",
				Value: 1 * time.Second,
				Usage: "Interval between RPC calls",
			},
			&cli.BoolFlag{
				Name:     "post-check-only",
				Usage:    "Only perform sanity checks",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "db-path",
				Usage:    "Path to database",
				Required: false,
			},
			&cli.IntFlag{
				Name:  "db-cache",
				Usage: "LevelDB cache size in mb",
				Value: 1024,
			},
			&cli.IntFlag{
				Name:  "db-handles",
				Usage: "LevelDB number of handles",
				Value: 60,
			},
		},
		Action: func(ctx *cli.Context) error {
			rpcURL := ctx.String("rpc-url")
			endBlock := ctx.Int64("end-block")
			rpcTimeout := ctx.Duration("rpc-timeout")
			rpcPollingInterval := ctx.Duration("rpc-poll-interval")
			outputPath := ctx.String("output-path")

			// post check only
			postCheckOnly := ctx.Bool("post-check-only")
			dbPath := ctx.String("db-path")
			if postCheckOnly {
				if len(dbPath) == 0 {
					return fmt.Errorf("Must specify a db-path if post-check-only is true")
				}
				if endBlock == 0 {
					return fmt.Errorf("Must specify an end-block if post-check-only is true")
				}

				dbCache := ctx.Int("db-cache")
				dbHandles := ctx.Int("db-handles")
				db, err := db.Open(dbPath, dbCache, dbHandles)
				if err != nil {
					return err
				}
				defer db.Close()
				log.Info("starting post check", "dbPath", dbPath, "eth addresses file", outputPath)
				postCheckcfg := genesis.NewEthAddressesCfg(outputPath, db, endBlock)
				if err := postCheckcfg.Check(); err != nil {
					return err
				}
				log.Info("post check complete! all good!")
				return nil
			}

			rpcClient, err := rpc.Dial(rpcURL)
			if err != nil {
				return err
			}
			defer rpcClient.Close()

			ethClient, err := ethclient.Dial(rpcURL)
			if err != nil {
				return err
			}
			defer ethClient.Close()

			log.Info("starting crawler", "rpcURL", rpcURL, "endBlock", fmt.Sprintf("%d", endBlock), "rpcTimeout", rpcTimeout, "rpcPollingInterval", rpcPollingInterval, "out", outputPath)
			c := ether.NewEthCrawler(rpcClient, ethClient, endBlock, rpcTimeout, rpcPollingInterval, outputPath)
			if err := c.Start(); err != nil {
				log.Error("error starting crawler", "err", err)
				return err
			}
			c.Wait()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Crit("critical error exits", "err", err)
	}
}
