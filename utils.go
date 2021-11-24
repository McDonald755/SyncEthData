package main

import (
	"SyncEthData/db"
	"github.com/ethereum/go-ethereum/core/types"
)

func TransformData(block *types.Block) {

	if block != nil {
		b := transferBlock(block)
		db.SaveBlock(b)
	}

	if block.Header() != nil {
		header := transferHeader(block.Header())
		db.SaveHeader(header)
	}

	if len(block.Transactions()) > 0 {
		for _, transaction := range block.Transactions() {
			trx := transferTrx(transaction)
			//这里可能可以优化程批量保存
			db.SaveTrx(trx)
		}

	}

}

func transferHeader(header *types.Header) db.HEADER {

	//TODO 解析数据撞到db对象里面
}

func transferTrx(header *types.Transaction) db.TRANSACTION {
	//TODO 解析数据撞到db对象里面
}

func transferBlock(block *types.Block) db.BLOCK {
	//TODO 解析数据撞到db对象里面
}
