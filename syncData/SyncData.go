package syncData

import (
	"SyncEthData/db"
	"SyncEthData/utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
)

//func GetBlockInfo() {
//	//TODO 改成批量获取,看看一个连接多少次请求合适
//	for _, client := range config.CLIENT {
//		block, err := client.BlockByNumber(context.Background(), num)
//		if err != nil {
//			log.Error(err)
//		}
//		utils.TransformData(block)
//	}
//
//}

func GetHeadByNum(client *ethclient.Client, num *big.Int) *types.Header {
	header, err := client.HeaderByNumber(context.Background(), num)
	if err != nil {
		log.Error(err)
	}
	return header
}

func GetHeadByHash(client *ethclient.Client, h string) {
	hash, err := client.HeaderByHash(context.Background(), common.HexToHash(h))
	if err != nil {
		log.Error(err)
	}
	fmt.Printf("%#v \n", hash)
}

func GetBlockByNum(client *ethclient.Client, num *big.Int) (*types.Block, error) {
	block, err := client.BlockByNumber(context.Background(), num)
	if err != nil {
		log.Error("Get Block Num Error:", err)
	}
	return block, err
}

func GetBlockHeight(client *ethclient.Client) int {
	number, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Error("Get Block Height Error:", err)
	}
	return int(number)
}

func ScanLog(client *ethclient.Client, contractABI abi.ABI, addr map[string]byte) {
	var (
		from  *big.Int
		addrs []string
	)

	for k, _ := range addr {
		addrs = append(addrs, k)
	}
	accounts := utils.TransferAccounts(&addrs)
	query := ethereum.FilterQuery{
		Topics: [][]common.Hash{
			{contractABI.Events["Transfer"].ID},
		},
		FromBlock: from,
		ToBlock:   from,
		Addresses: *accounts,
	}

	filterLogs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Error("Get log error:", err)
	}

	for _, l := range filterLogs {
		data := utils.TransferNftData(l)
		db.SaveOrUpdateNftData(data)
	}
}
