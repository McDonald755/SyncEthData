package db

import (
	"SyncEthData/config"
	log "github.com/sirupsen/logrus"
)

func SaveBlock(block BLOCK) error {
	res := config.DB.Save(block)
	if res.Error != nil {
		log.Error(res.Error)
		return res.Error
	}
	return nil
}

func SaveHeader(header HEADER) error {
	res := config.DB.Save(header)
	if res.Error != nil {
		log.Error(res.Error)
		return res.Error
	}
	return nil
}

func SaveTrx(trx TRANSACTION) error {
	res := config.DB.Save(trx)
	if res.Error != nil {
		log.Error(res.Error)
		return res.Error
	}
	return nil
}
