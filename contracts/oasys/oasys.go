package oasys

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

const (
	Genesis Namespace = iota
	WrappedOAS
	VerseBuilder
	Token
	Rollup

	hexPrefix         = "0x"
	storageSlotLength = 64
)

var (
	GenesisHash common.Hash

	// reserved address prefix list of built-in contracts.
	addressPrefixes = map[Namespace]string{
		Genesis:      "0x0000000000000000000000000000000000000000",
		WrappedOAS:   "0x0001000000000000000000000000000000000000",
		Rollup:       "0x0002000000000000000000000000000000000000",
		Token:        "0x0003000000000000000000000000000000000000",
		VerseBuilder: "0x0004000000000000000000000000000000000000",
	}

	builtinContracts = map[uint64][]deployable{
		1: {
			wrappedOAS,
			rollupContractSet,
			tokenContractSet,
			verseBuilderContractSet,
		},
	}
)

// Deploy oasys built-in contracts.
func Deploy(config *params.ChainConfig, state *state.StateDB, block uint64) {
	if config == nil || config.Oasys == nil || state == nil {
		return
	}
	if contracts, ok := builtinContracts[block]; ok {
		for _, c := range contracts {
			c.deploy(state)
		}
	}
}

// Namespace is a category of oasys built-in contracts.
type Namespace int

// GetAddress returns the address with add offset to the first address of namespace.
func (p Namespace) GetAddress(offset int64) common.Address {
	base := common.HexToAddress(addressPrefixes[p]).Hash().Big()
	return common.BigToAddress(new(big.Int).Add(base, big.NewInt(offset)))
}

// deployable
type deployable interface {
	deploy(state *state.StateDB)
}

// contractSet
type contractSet []*contract

func (p contractSet) deploy(state *state.StateDB) {
	for _, c := range p {
		c.deploy(state)
	}
}

// contract
type contract struct {
	name           string
	address        common.Address
	code           string
	fixedStorage   map[string]interface{}
	dynamicStorage map[string]string
}

func (c *contract) deploy(state *state.StateDB) {
	if len(state.GetCode(c.address)) != 0 {
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

	state.SetCode(c.address, bytecode)
	for key, val := range storage {
		state.SetState(c.address, key, val)
	}
	log.Info("Deploy contract", "name", c.name, "address", c.address.String())
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
