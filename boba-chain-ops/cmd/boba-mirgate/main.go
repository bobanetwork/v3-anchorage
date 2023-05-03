package main

import (
	"fmt"
	"os"

	"github.com/ethereum-optimism/optimism/boba-chain-ops/genesis"
	"github.com/ethereum-optimism/optimism/op-chain-ops/db"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mattn/go-isatty"
	"github.com/urfave/cli"
)

func main() {
	log.Root().SetHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(isatty.IsTerminal(os.Stderr.Fd()))))

	app := &cli.App{
		Name:  "boba-migrate",
		Usage: "Write allocation data from the legacy data to a json file for erigon",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "db-path",
				Usage:    "Path to database",
				Required: true,
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
			&cli.StringFlag{
				Name:     "output-path",
				Usage:    "Path to output file",
				Required: true,
			},
			&cli.Int64Flag{
				Name:  "hardfork-block",
				Usage: "Block number of the hardfork",
				Value: 0,
			},
		},
		Action: func(ctx *cli.Context) error {
			dbPath := ctx.String("db-path")
			dbCache := ctx.Int("db-cache")
			dbHandles := ctx.Int("db-handles")
			outputPath := ctx.String("output-path")
			hardforkBlock := ctx.Int64("hardfork-block")

			db, err := db.Open(dbPath, dbCache, dbHandles)
			if err != nil {
				return fmt.Errorf("error opening db: %s", err)
			}
			defer db.Close()
			genesisCfg := genesis.NewGenesisConfig(db, outputPath, hardforkBlock)
			if err := genesisCfg.Dump(); err != nil {
				return fmt.Errorf("error dumping genesis config: %s", err)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Crit("critical error exits", "err", err)
	}
}
