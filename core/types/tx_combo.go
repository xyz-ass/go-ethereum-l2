// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// ComboTx is the transaction data of the original Ethereum transactions.
type ComboTx struct {
	BlockHash        *common.Hash
	BlockNumber      uint64
	From             *common.Address
	Nonce            uint64   // nonce of sender account
	GasPrice         *big.Int // wei per gas
	Gas              uint64   // gas limit
	Hash             *common.Hash
	Input            []byte
	TransactionIndex uint64
	To               *common.Address `rlp:"nil"` // nil means contract creation
	Value            *big.Int        // wei amount
	Data             []byte          // contract invocation input data
	V, R, S          *big.Int        // signature values
	SourceHash       *common.Hash
	Mint             *big.Int
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *ComboTx) copy() TxData {
	cpy := &ComboTx{
		Nonce: tx.Nonce,
		To:    copyAddressPtr(tx.To),
		Data:  common.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are initialized below.
		Value:    new(big.Int),
		GasPrice: new(big.Int),
		V:        new(big.Int),
		R:        new(big.Int),
		S:        new(big.Int),
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

// accessors for innerTx.
func (tx *ComboTx) txType() byte           { return ComboTxType }
func (tx *ComboTx) chainID() *big.Int      { return deriveChainId(tx.V) }
func (tx *ComboTx) accessList() AccessList { return nil }
func (tx *ComboTx) data() []byte           { return tx.Data }
func (tx *ComboTx) gas() uint64            { return tx.Gas }
func (tx *ComboTx) gasPrice() *big.Int     { return tx.GasPrice }
func (tx *ComboTx) gasTipCap() *big.Int    { return tx.GasPrice }
func (tx *ComboTx) gasFeeCap() *big.Int    { return tx.GasPrice }
func (tx *ComboTx) value() *big.Int        { return tx.Value }
func (tx *ComboTx) nonce() uint64          { return tx.Nonce }
func (tx *ComboTx) to() *common.Address    { return tx.To }

func (tx *ComboTx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	return dst.Set(tx.GasPrice)
}

func (tx *ComboTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *ComboTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.V, tx.R, tx.S = v, r, s
}

func (tx *ComboTx) encode(*bytes.Buffer) error {
	panic("encode called on ComboTx")
}

func (tx *ComboTx) decode([]byte) error {
	panic("decode called on ComboTx)")
}
