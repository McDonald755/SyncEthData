package syncData

import (
	"context"
	"fmt"
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
