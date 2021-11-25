package cmd

import (
	"SyncEthData/config"
	"SyncEthData/syncData"
	"SyncEthData/utils"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"math/big"
	"time"
)

func ScanCmd() *cobra.Command {
	var blockNum int
	scanCmd := &cobra.Command{
		Use:   "scan",
		Short: "s",
		Long:  "It will sync the latest block ",
		RunE: func(cmd *cobra.Command, args []string) error {
			scanBlock(blockNum)
			return nil
		},
	}
	scanCmd.Flags().IntVarP(&blockNum, "blockNum", "b", 0, "input blockNum")
	return scanCmd
}

func scanBlock(blockNum int) {
	height := syncData.GetBlockHeight(config.CLIENT[0])
	distance := (height - blockNum) / len(config.CLIENT)

	for i := 0; i < len(config.CLIENT)-1; i++ {
		go getBlock(config.CLIENT[i], i, distance, blockNum)
	}

	go scanNewBlock(config.CLIENT[len(config.CLIENT)-1], (len(config.CLIENT)-1)*distance+blockNum)
}

func getBlock(client *ethclient.Client, i int, distance int, blockNum int) {
	from := distance*i + blockNum
	end := from + distance
	for from < end {
		block, err := syncData.GetBlockByNum(client, big.NewInt(int64(from)))
		if err != nil {
			time.Sleep(time.Hour)
		} else {
			utils.TransformData(block)
			from += 1
		}

	}
}

//同步最新区块
func scanNewBlock(client *ethclient.Client, from int) {
	for true {
		block, err := syncData.GetBlockByNum(client, big.NewInt(int64(from)))
		if err != nil {
			time.Sleep(time.Hour)
		} else {
			utils.TransformData(block)
			from += 1
		}

	}
}
