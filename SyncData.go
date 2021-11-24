package main

import (
	"SyncEthData/config"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/big"
)

func getBlockInfo(num *big.Int) {
	block, err := config.CLIENT.BlockByNumber(context.Background(), num)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(block)

	TransformData(block)
}
