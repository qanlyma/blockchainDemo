//transactionpool.go

package blockchain

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
	"qanly_chain/constcoe"
	"qanly_chain/transaction"
	"qanly_chain/utils"
)

type TransactionPool struct {
	PubTx []*transaction.Transaction
}

func (tp *TransactionPool) AddTransaction(tx *transaction.Transaction) {
	tp.PubTx = append(tp.PubTx, tx)
}

func (tp *TransactionPool) SaveFile() { //将交易信息池保存到constcoe.TransactionPoolFile这个地址中
	var content bytes.Buffer
	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(tp)
	utils.Handle(err)
	err = ioutil.WriteFile(constcoe.TransactionPoolFile, content.Bytes(), 0644)
	utils.Handle(err)
}

func (tp *TransactionPool) LoadFile() error {
	if !utils.FileExists(constcoe.TransactionPoolFile) {
		return nil
	}

	var transactionPool TransactionPool

	fileContent, err := ioutil.ReadFile(constcoe.TransactionPoolFile)
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(bytes.NewBuffer(fileContent))
	err = decoder.Decode(&transactionPool)

	if err != nil {
		return err
	}

	tp.PubTx = transactionPool.PubTx
	return nil
}

func CreateTransactionPool() *TransactionPool {
	transactionPool := TransactionPool{}
	err := transactionPool.LoadFile()
	utils.Handle(err)
	return &transactionPool
}

func RemoveTransactionPoolFile() error {
	err := os.Remove(constcoe.TransactionPoolFile)
	return err
}
