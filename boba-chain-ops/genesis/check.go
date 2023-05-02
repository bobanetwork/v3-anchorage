package genesis

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum-optimism/optimism/boba-chain-ops/ether"
	"github.com/ethereum-optimism/optimism/boba-chain-ops/utils"
	opether "github.com/ethereum-optimism/optimism/op-chain-ops/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/trie"
)

type EthAddressesCfg struct {
	EthAddressesDir string
	db              ethdb.Database
	hardforkBlock   int64
}

func NewEthAddressesCfg(ethAddressesDir string, db ethdb.Database, hardforkBlock int64) *EthAddressesCfg {
	return &EthAddressesCfg{
		EthAddressesDir: ethAddressesDir,
		db:              db,
		hardforkBlock:   hardforkBlock,
	}
}

func (cfg *EthAddressesCfg) Check() error {
	file, err := os.Open(cfg.EthAddressesDir)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	if len(bytes) == 0 {
		return errors.New("Invalid eth addresses directory. The directory is empty.")
	}
	var addresses ether.EthAddresses
	if err := json.Unmarshal(bytes, &addresses); err != nil {
		return err
	}

	var header *types.Header
	if hash := rawdb.ReadCanonicalHash(cfg.db, uint64(cfg.hardforkBlock)); hash != (common.Hash{}) {
		header = rawdb.ReadHeader(cfg.db, hash, uint64(cfg.hardforkBlock))
	} else {
		return errors.New("hard fork block not found")
	}

	conf := &state.DumpConfig{
		SkipCode:          false,
		SkipStorage:       false,
		OnlyWithAddresses: false,
		Start:             common.Hash{}.Bytes(),
		Max:               uint64(0),
	}
	statedb, err := state.New(header.Root, state.NewDatabaseWithConfig(cfg.db, &trie.Config{Preimages: true}), nil)
	state := statedb.RawDump(conf)
	ovmEthDump := state.Accounts[utils.OVM_ETH_ADDRESS]
	if ovmEthDump.Code == nil {
		return errors.New("OVM_ETH dump not found")
	}
	storage := ovmEthDump.Storage
	return ValidAddressInStorage(storage, addresses.Addresses)
}

func ValidAddressInStorage(storage map[common.Hash]string, addresses []*common.Address) error {
	var (
		err              error
		validAddrCount   = 0
		validAddrCountCh = make(chan int)
	)

	for _, address := range addresses {
		go func(addr common.Address) {
			storageKey := opether.CalcOVMETHStorageKey(addr)
			if storage[storageKey] != "" {
				validAddrCountCh <- 1
			} else {
				validAddrCountCh <- 0
			}
		}(*address)
	}

	for range addresses {
		validAddrCount += <-validAddrCountCh
	}

	dataStorageKey := []common.Hash{
		common.BytesToHash([]byte{2}),
		common.BytesToHash([]byte{3}),
		common.BytesToHash([]byte{4}),
		common.BytesToHash([]byte{6}),
	}
	for _, key := range dataStorageKey {
		if storage[key] != "" {
			validAddrCount++
		}
	}

	if len(storage) != validAddrCount {
		log.Warn("Some addresses in eth addresses file are not valid", "valid", validAddrCount, "total", len(storage))
		err = fmt.Errorf("Some addresses in eth addresses file are not valid, valid: %d, total: %d", validAddrCount, len(storage))
	} else {
		log.Info("All addresses in eth addresses file are valid", "valid", validAddrCount, "total", len(storage))
	}

	return err
}
