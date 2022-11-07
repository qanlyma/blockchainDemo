//block.go

package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"qanly_chain/merkletree"
	"qanly_chain/transaction"
	"qanly_chain/utils"
	"time"
)

type Block struct { //区块结构体
	Timestamp    int64
	Hash         []byte
	PrevHash     []byte
	Target       []byte
	Nonce        int64
	Transactions []*transaction.Transaction
	MTree        *merkletree.MerkleTree
}

func (b *Block) BackTrasactionSummary() []byte {
	txIDs := make([][]byte, 0)
	for _, tx := range b.Transactions {
		txIDs = append(txIDs, tx.ID)
	}
	summary := bytes.Join(txIDs, []byte{})
	return summary
}

func (b *Block) SetHash() {
	information := bytes.Join([][]byte{utils.ToHexInt(b.Timestamp), b.PrevHash, b.Target, utils.ToHexInt(b.Nonce), b.BackTrasactionSummary(), b.MTree.RootNode.Data}, []byte{})
	hash := sha256.Sum256(information)
	b.Hash = hash[:]
}

func CreateBlock(prevhash []byte, txs []*transaction.Transaction) *Block {
	block := Block{time.Now().Unix(), []byte{}, prevhash, []byte{}, 0, txs, merkletree.CrateMerkleTree(txs)}
	block.Target = block.GetTarget()
	block.Nonce = block.FindNonce()
	block.SetHash()
	return &block
}

func GenesisBlock(address []byte) *Block {
	tx := transaction.BaseTx(address)
	genesis := CreateBlock([]byte("MQL is awesome!"), []*transaction.Transaction{tx})
	genesis.SetHash()
	return genesis
}

func (b *Block) Serialize() []byte { //序列化区块生成字节串,Badger的键值对只支持字节串存储形式
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
	utils.Handle(err)
	return res.Bytes()
}

func DeSerializeBlock(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	utils.Handle(err)
	return &block
}
