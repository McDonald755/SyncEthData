package utils

import (
	"SyncEthData/db"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
	"time"
)

func TransformData(block *types.Block) {
	if block == nil {
		return
	}
	b := transferBlock(block)
	header := transferHeader(block.Header())

	transactions := []db.TRANSACTION{}
	if len(block.Transactions()) > 0 {
		for _, transaction := range block.Transactions() {
			trx := transferTrx(transaction, block.Header().Number)
			if trx != nil {
				transactions = append(transactions, *trx)
			}
		}
	}

	db.SaveData(b, header, &transactions)
}

func transferHeader(header *types.Header) *db.HEADER {
	var baseFee = big.NewInt(0)
	if header.BaseFee != nil {
		baseFee = header.BaseFee
	}
	result := db.HEADER{
		ParentHash:  header.ParentHash.Hex(),
		UncleHash:   header.UncleHash.Hex(),
		CoinHase:    header.Coinbase.Hex(),
		Root:        header.Root.Hex(),
		TxHash:      header.TxHash.Hex(),
		ReceiptHash: header.ReceiptHash.Hex(),
		//BLOOM:       header.Bloom.Big().String(),
		Difficulty:  header.Difficulty.Int64(),
		BlockNumber: header.Number.Int64(),
		GasLimit:    header.GasLimit,
		GasUsed:     header.GasUsed,
		Time:        header.Time,
		Extra:       hex.Dump(header.Extra),
		Nonce:       strconv.Itoa(int(header.Nonce.Uint64())),
		Basefee:     baseFee.Int64(),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}

	return &result
}

func transferTrx(trx *types.Transaction, num *big.Int) *db.TRANSACTION {
	var (
		toAccount, value string
	)
	if trx.To() != nil {
		toAccount = trx.To().Hex()
	}

	if trx.Value() != nil {
		value = trx.Value().String()
	}
	result := db.TRANSACTION{
		TxData: "0x" + hex.EncodeToString(trx.Data()),
		Hash:   trx.Hash().Hex(),
		Size:   trx.Size().String(),
		//FROMACCOUNT: msg.From().Hex(),
		ToAccount:   toAccount,
		Value:       value,
		TxnType:     int64(trx.Type()),
		BlockNumber: num.Int64(),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	return &result
}

func transferBlock(block *types.Block) *db.BLOCK {
	result := db.BLOCK{
		BlockNum:   block.Number().Int64(),
		BlockHash:  block.Hash().Hex(),
		BlockSize:  block.Size().String(),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	return &result
}
