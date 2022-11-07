//inoutput.go

package transaction

import (
	"bytes"
	"qanly_chain/utils"
)

type TxOutput struct {
	Value      int
	HashPubKey []byte
}

type TxInput struct {
	TxID   []byte //支持本次交易的前置交易信息
	OutIdx int    //指明是前置交易信息中的第几个Output
	PubKey []byte
	Sig    []byte
}

func (in *TxInput) FromAddressRight(address []byte) bool {
	return bytes.Equal(in.PubKey, address)
}

func (out *TxOutput) ToAddressRight(address []byte) bool {
	return bytes.Equal(out.HashPubKey, utils.PublicKeyHash(address))
}
