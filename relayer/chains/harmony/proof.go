package harmony

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	eth "github.com/harmony-one/go-sdk/pkg/rpc/eth"
)

const AccountStorageRootIndex = 2

type ETHProof struct {
	AccountProofRLP []byte
	StorageProofRLP [][]byte
}

func (cl Client) GetETHProof(address common.Address, storageKeys [][]byte, blockNumber *big.Int) (*ETHProof, error) {
	bz, err := cl.getProof(address, storageKeys, "0x"+blockNumber.Text(16))
	if err != nil {
		return nil, err
	}
	var proof struct {
		AccountProof []string `json:"accountProof"`
		StorageProof []struct {
			Proof []string `json:"proof"`
		} `json:"storageProof"`
	}
	if err := json.Unmarshal(bz, &proof); err != nil {
		return nil, err
	}

	var encodedProof ETHProof
	encodedProof.AccountProofRLP, err = encodeRLP(proof.AccountProof)
	if err != nil {
		return nil, err
	}
	for _, p := range proof.StorageProof {
		bz, err := encodeRLP(p.Proof)
		if err != nil {
			panic(err)
		}
		encodedProof.StorageProofRLP = append(encodedProof.StorageProofRLP, bz)
	}

	return &encodedProof, nil
}

func (cl Client) getProof(address common.Address, storageKeys [][]byte, blockNumber string) ([]byte, error) {
	hashes := []common.Hash{}
	for _, k := range storageKeys {
		var h common.Hash
		if err := h.UnmarshalText(k); err != nil {
			return nil, err
		}
		hashes = append(hashes, h)
	}
	rep, err := cl.messenger.SendRPC(eth.Method.GetProof, []interface{}{
		address, hashes, blockNumber,
	})
	if err != nil {
		return nil, err
	}
	val, ok := rep["result"]
	if !ok {
		return nil, errors.New("invalid response")
	}
	return json.Marshal(val)
}

// decodeRLP decodes the proof according to the IBFT2.0 client proof format implemented by yui-ibc-solidity
// and formats it for Ethereum's Account/Storage Proof.
func decodeRLP(proof []byte) ([][]byte, error) {
	var val [][][]byte
	if err := rlp.DecodeBytes(proof, &val); err != nil {
		return nil, err
	}

	var res [][]byte
	for _, v := range val {
		bz, err := rlp.EncodeToBytes(v)
		if err != nil {
			return nil, err
		}
		res = append(res, bz)
	}
	return res, nil
}

func encodeRLP(proof []string) ([]byte, error) {
	var target [][][]byte
	for _, p := range proof {
		bz, err := hex.DecodeString(p[2:])
		if err != nil {
			panic(err)
		}
		var val [][]byte
		if err := rlp.DecodeBytes(bz, &val); err != nil {
			panic(err)
		}
		target = append(target, val)
	}
	bz, err := rlp.EncodeToBytes(target)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

func decodeStorageHash(accountRLP []byte) ([]byte, error) {
	var account [][]byte
	if err := rlp.DecodeBytes(accountRLP, &account); err != nil {
		return nil, err
	}
	if len(account) <= AccountStorageRootIndex {
		return nil, errors.New("invalid decode account")
	}
	return account[AccountStorageRootIndex], nil
}
