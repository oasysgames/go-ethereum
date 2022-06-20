package oasys

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

const (
	// Address of initial wallet in genesis.
	mainnetGenesisWallet = "0xfc302c2252a69003b3f03750564393924d2c96ae"
	testnetGenesisWallet = "0x2929efbad86a7989879fc7650d1e30c02a488660"

	hexPrefix         = "0x"
	storageSlotLength = 64
)

var (
	GenesisHash common.Hash

	builtinContracts = map[uint64][]deployable{
		1: {
			tokenContractSet,
			rollupContractSet,
			tokenFactoryContractSet,
			verseBuilderContractSet,
			nftBridgeContractSet,
		},
	}
)

// StateDB is an interface of state.StateDB.
type StateDB interface {
	GetCode(addr common.Address) []byte
	SetCode(addr common.Address, code []byte)
	SetState(addr common.Address, key common.Hash, value common.Hash)
}

// Deploy oasys built-in contracts.
func Deploy(config *params.ChainConfig, state StateDB, block uint64) {
	if config == nil || config.Oasys == nil || state == nil {
		return
	}
	if contracts, ok := builtinContracts[block]; ok {
		for _, c := range contracts {
			c.deploy(state)
		}
	}
}

// deployable
type deployable interface {
	deploy(state StateDB)
}

// contractSet
type contractSet []*contract

func (p contractSet) deploy(state StateDB) {
	for _, c := range p {
		c.deploy(state)
	}
}

// contract
type contract struct {
	name           string
	address        string
	code           string
	fixedStorage   map[string]interface{}
	dynamicStorage map[string]string
}

func (c *contract) deploy(state StateDB) {
	address := common.HexToAddress(c.address)
	if len(state.GetCode(address)) != 0 {
		panic(fmt.Errorf("%s contract already exists", c.name))
	}

	bytecode, err := c.hexcode()
	if err != nil {
		panic(fmt.Errorf("failed to decode %s contract code: %s", c.name, err.Error()))
	}
	storage, err := c.storage()
	if err != nil {
		panic(fmt.Errorf("failed to create %s contract storage map: %s", c.name, err.Error()))
	}

	state.SetCode(address, bytecode)
	for key, val := range storage {
		state.SetState(address, key, val)
	}
	log.Info("Deploy contract", "name", c.name, "address", c.address)
}

// ByteCodes returns the contract byte codes.
func (c *contract) hexcode() ([]byte, error) {
	bytecode, err := hex.DecodeString(strings.TrimPrefix(c.code, hexPrefix))
	if err != nil {
		return nil, err
	}
	return bytecode, nil
}

// Storage returns the contract storage slot map.
func (c *contract) storage() (map[common.Hash]common.Hash, error) {
	storage := make(map[common.Hash]common.Hash)

	if c.fixedStorage != nil {
		for key, val := range c.fixedStorage {
			slot := common.HexToHash(key)
			switch t := val.(type) {
			case common.Hash:
				storage[slot] = t
			case common.Address:
				storage[slot] = t.Hash()
			case *big.Int:
				storage[slot] = common.BigToHash(t)
			case string:
				if !strings.HasPrefix(t, hexPrefix) {
					if len(t) > 31 {
						return nil, fmt.Errorf("strings longer than 32 bytes must be set to DynamicStorages len: %d, key: %s", len(t), key)
					}
					t = toHex(t)
				}
				storage[slot] = common.HexToHash(t)
			default:
				return nil, fmt.Errorf("unsupported type: %s, key: %s", t, key)
			}
		}
	}

	if c.dynamicStorage != nil {
		for key, val := range c.dynamicStorage {
			val = strings.TrimPrefix(val, hexPrefix)

			rootSlot := common.HexToHash(key)
			storage[rootSlot] = common.BigToHash(big.NewInt(int64(len(val) + 1)))

			chunkStartPos := crypto.Keccak256Hash(rootSlot.Bytes()).Big()
			for i, chunk := range toChunks(val, storageSlotLength) {
				chunkSlot := common.BigToHash(new(big.Int).Add(chunkStartPos, big.NewInt(int64(i))))
				storage[chunkSlot] = common.HexToHash(chunk)
			}
		}
	}

	return storage, nil
}

// Returns the copied instance.
func (c *contract) copy() *contract {
	cpy := &contract{
		name:           c.name,
		address:        c.address,
		code:           c.code,
		fixedStorage:   make(map[string]interface{}),
		dynamicStorage: make(map[string]string),
	}
	for key, value := range c.fixedStorage {
		cpy.fixedStorage[key] = value
	}
	for key, value := range c.dynamicStorage {
		cpy.dynamicStorage[key] = value
	}
	return cpy
}

func toChunks(s string, l int) []string {
	slen := len(s)
	chunks := make([]string, 0)
	for i := 0; i < slen; i += l {
		end := i + l
		if end > slen {
			end = slen
		}
		slice := s[i:end]
		chunks = append(chunks, rightZeroPad(slice, l))
	}
	return chunks
}

func toHex(s string) string {
	hexs := hex.EncodeToString([]byte(s))
	hexlen := strconv.FormatInt(int64(len(s)*2), 16)
	return rightZeroPad(hexs, 62) + leftZeroPad(hexlen, 2)
}

func rightZeroPad(s string, l int) string {
	return s + strings.Repeat("0", l-len(s))
}

func leftZeroPad(s string, l int) string {
	return strings.Repeat("0", l-len(s)) + s
}
