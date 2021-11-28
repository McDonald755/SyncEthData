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
		PARENTHASH:  header.ParentHash.Hex(),
		UNCLEHASH:   header.UncleHash.Hex(),
		COINBASE:    header.Coinbase.Hex(),
		ROOT:        header.Root.Hex(),
		TXHASH:      header.TxHash.Hex(),
		RECEIPTHASH: header.ReceiptHash.Hex(),
		//BLOOM:       header.Bloom.Big().String(),
		DIFFICULTY:  header.Difficulty.Int64(),
		BLOCKNUMBER: header.Number.Int64(),
		GASLIMIT:    header.GasLimit,
		GASUSED:     header.GasUsed,
		TIME:        header.Time,
		EXTRA:       hex.Dump(header.Extra),
		NONCE:       strconv.Itoa(int(header.Nonce.Uint64())),
		BASEFEE:     baseFee.Int64(),
		CREATETIME:  time.Now(),
		UPDATETIME:  time.Now(),
	}

	return &result
}

func transferTrx(trx *types.Transaction, num *big.Int) *db.TRANSACTION {
	//msg, err := trx.AsMessage(types.NewLondonSigner(trx.ChainId()), nil)
	//if err != nil {
	//	log.Error("Get trx message Error TRX hash is:", trx.Hash().Hex(), err)
	//	return nil
	//}
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
		TXDATA: "0x" + hex.EncodeToString(trx.Data()),
		HASH:   trx.Hash().Hex(),
		SIZE:   trx.Size().String(),
		//FROMACCOUNT: msg.From().Hex(),
		TOACCOUNT:   toAccount,
		VALUE:       value,
		TXNTYPE:     int64(trx.Type()),
		BLOCKNUMBER: num.Int64(),
		CREATETIME:  time.Now(),
		UPDATETIME:  time.Now(),
	}
	return &result
}

func transferBlock(block *types.Block) *db.BLOCK {
	result := db.BLOCK{
		BLOCKNUM:   block.Number().Int64(),
		BLOCKHASH:  block.Hash().Hex(),
		BLOCKSIZE:  block.Size().String(),
		CREATETIME: time.Now(),
		UPDATETIME: time.Now(),
	}
	return &result
}
