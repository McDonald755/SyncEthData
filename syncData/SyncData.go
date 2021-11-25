package syncData

import (
	"SyncEthData/config"
	"SyncEthData/utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
)

func GetBlockInfo() {
	//TODO 改成批量获取,看看一个连接多少次请求合适
	for _, client := range config.CLIENT {
		block, err := client.BlockByNumber(context.Background(), num)
		if err != nil {
			log.Error(err)
		}
		fmt.Println(block)

		utils.TransformData(block)
	}

}

func GetHeadByNum(client *ethclient.Client, num *big.Int) *types.Header {
	header, err := client.HeaderByNumber(context.Background(), num)
	if err != nil {
		fmt.Println(err)
	}
	return header
}

func GetHeadByHash(client *ethclient.Client, h string) {
	hash, err := client.HeaderByHash(context.Background(), common.HexToHash(h))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v \n", hash)
}

func GetBlockByNum(client *ethclient.Client, num *big.Int) *types.Block {
	block, err := client.BlockByNumber(context.Background(), num)
	if err != nil {
		fmt.Println(err)
	}
	return block
}

func GetBlockHeight(client *ethclient.Client) int {
	return 0
}
