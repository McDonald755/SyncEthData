package cmd

import (
	"SyncEthData/config"
	"SyncEthData/db"
	"SyncEthData/oracle"
	"SyncEthData/syncData"
	"SyncEthData/utils"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"math/big"
	"strings"
)

func SyncCmd() *cobra.Command {
	var blockNum int
	scanCmd := &cobra.Command{
		Use:   "sync",
		Short: "c",
		Long:  "It will sync the latest block ",
		RunE: func(cmd *cobra.Command, args []string) error {
			//sync data
			syncNewBlock(blockNum)
			return nil
		},
	}
	scanCmd.Flags().IntVarP(&blockNum, "blockNum", "b", 0, "input blockNum")
	return scanCmd
}

func syncNewBlock(from int) {
	i := 1
	//get all of nft contract addr data
	oracles := db.GetOracleAddrAll()

	//init standard ERC-721 contract data
	contractABI, err := abi.JSON(strings.NewReader(oracle.OracleABI))
	if err != nil {
		log.Error("Read Contract Error:", err)
	}

	//main loop
	for true {
		//init eth client

		client := config.CLIENT[i]
		number, err := client.BlockNumber(context.Background())
		if err != nil {
			log.Error("client is", i, err)
			//if error change other client
			i += 1
			if i > len(config.CLIENT) {
				i = 1
			}
		}

		for from < int(number) {
			block, err := syncData.GetBlockByNum(client, big.NewInt(int64(from)))
			if err != nil {
				log.Error(err)
				break
			}
			//Save trx ,header ,block data
			utils.TransformData(block)

			//deal and save nft contract data
			result := utils.CheckOracleType(block.Transactions(), oracles)
			//filter and save nft data
			syncData.ScanLog(client, contractABI, result)
			from += 1
		}
	}
}
