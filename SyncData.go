package main

import (
	"SyncEthData/config"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/big"
)

func getBlockInfo(num *big.Int) {
	//TODO 改成批量获取,看看一个连接多少次请求合适
	for i, client := range config.CLIENT {
		block, err := client.BlockByNumber(context.Background(), num)
		if err != nil {
			log.Error(err)
		}
		fmt.Println(block)

		TransformData(block)
	}

}
