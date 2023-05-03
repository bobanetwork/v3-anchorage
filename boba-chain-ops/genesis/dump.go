package genesis

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/trie"
)

type GenesisConfig struct {
	db            ethdb.Database
	outputPath    string
	hardforkBlock int64
}

func NewGenesisConfig(db ethdb.Database, outputPath string, hardforkBlock int64) *GenesisConfig {
	return &GenesisConfig{
		db:            db,
		outputPath:    outputPath,
		hardforkBlock: hardforkBlock,
	}
}

func (cfg *GenesisConfig) Dump() error {
	hash := rawdb.ReadHeadHeaderHash(cfg.db)
	if cfg.hardforkBlock == 0 {
		hash = rawdb.ReadCanonicalHash(cfg.db, uint64(cfg.hardforkBlock))
	}

	num := rawdb.ReadHeaderNumber(cfg.db, hash)
	log.Info("Dumping genesis state", "hash", hash, "number", num)

	header := rawdb.ReadHeader(cfg.db, hash, *num)
	conf := &state.DumpConfig{
		SkipCode:          false,
		SkipStorage:       false,
		OnlyWithAddresses: false,
		Start:             common.Hash{}.Bytes(),
		Max:               uint64(0),
	}
	statedb, err := state.New(header.Root, state.NewDatabaseWithConfig(cfg.db, &trie.Config{Preimages: true}), nil)
	if err != nil {
		return err
	}
	state := statedb.RawDump(conf)
	alloc, err := json.Marshal(state.Accounts)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(cfg.outputPath, alloc, 0644); err != nil {
		log.Warn("Failed to write genesis alloc", "err", err)
		return err
	}
	return nil
}
