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
				Name:  "out",
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
		},
		Action: func(ctx *cli.Context) error {
			rpcURL := ctx.String("rpc-url")
			endBlock := ctx.Int64("end-block")
			rpcTimeout := ctx.Duration("rpc-timeout")
			rpcPollingInterval := ctx.Duration("rpc-poll-interval")
			out := ctx.String("out")

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

			log.Info("starting crawler", "rpcURL", rpcURL, "endBlock", fmt.Sprintf("%d", endBlock), "rpcTimeout", rpcTimeout, "rpcPollingInterval", rpcPollingInterval, "out", out)
			c := ether.NewEthCrawler(rpcClient, ethClient, endBlock, rpcTimeout, rpcPollingInterval, out)
			if err := c.Start(); err != nil {
				log.Error("error starting crawler", "err", err)
				return err
			}
			c.Wait()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Crit("error fetching addresses", "err", err)
	}
}
