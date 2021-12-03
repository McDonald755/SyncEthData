package db

import (
	"SyncEthData/config"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

func SaveData(block *BLOCK, header *HEADER, trx *[]TRANSACTION) {
	dbErr := gorm.DB{}
	tx := config.DB.Begin()
	dbErr = *tx.Create(block)
	dbErr = *tx.Create(header)
	if len(*trx) > 0 {
		dbErr = *tx.Create(trx)
	}

	if dbErr.Error != nil {
		log.Error(dbErr.Error)
		log.Error("----------------------Error Num is:", block.BlockNum)
		tx.Rollback()
	} else {
		//fmt.Println("save:", block.BLOCKNUM)
		tx.Commit()
	}
}

func FixToAccount(from int, sw *sync.WaitGroup) {

	type T struct {
		ID   int64
		Hash string
	}
	var f []T
	transaction := TRANSACTION{}
	config.DB.Table("TRANSACTION").Select("ID,hash").Where("to_account = ''").Limit(100).Offset(from).Find(&f)
	for _, v := range f {
		receipt, err := config.CLIENT[0].TransactionReceipt(context.Background(), common.HexToHash(v.Hash))
		if err != nil {
			fmt.Print(err)
		}
		config.DB.Model(&transaction).Where("ID = ?", v.ID).Update("to_account", receipt.ContractAddress.String())
	}
	sw.Done()
}

/**

 */
func GetOracleAddrAll() map[string]byte {
	var (
		addres []string
		result map[string]byte
	)
	config.DB.Table("ORACLE_DATA").Select("address").Find(&addres)
	for _, addre := range addres {
		result[addre] = byte(0)
	}
	return result
}

func SaveOracle(oracle *ORACLE_DATA) {
	config.DB.Create(oracle)
}

func SaveOrUpdateNftData(nft *NFT_DATA) {
	var id int64
	result := config.DB.Table("NFT_DATA").Select("ID").Where("oracle_add,token_id", nft.OracleAdd, nft.TokenId).Find(&id)
	if result.RowsAffected == 0 {
		config.DB.Create(nft)
	} else {
		nft.ID = id
		config.DB.Save(nft)
	}
}
