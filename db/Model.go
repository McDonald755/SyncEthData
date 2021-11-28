package db

import (
	"time"
)

type TRANSACTION struct {
	ID         int64     `gorm:"primaryKey;column:ID" json:"ID"`
	CREATETIME time.Time `gorm:"column:CREATE_TIME;default:null" json:"CREATE_TIME"`
	UPDATETIME time.Time `gorm:"column:UPDATE_TIME;default:null" json:"UPDATE_TIME"`
	TXDATA     string    `gorm:"column:TX_DATA" json:"TX_DATA"`
	HASH       string    `gorm:"column:HASH" json:"HASH"`
	SIZE       string    `gorm:"column:SIZE" json:"SIZE"`
	//FROMACCOUNT string    `gorm:"column:FROM_ACCOUNT" json:"FROM_ACCOUNT"`
	TOACCOUNT   string `gorm:"column:TO_ACCOUNT" json:"TO_ACCOUNT"`
	TXNTYPE     int64  `gorm:"column:TXNTYPE" json:"TXNTYPE"`
	VALUE       string `gorm:"column:VALUE" json:"VALUE"`
	BLOCKNUMBER int64  `gorm:"column:BLOCK_NUMBER" json:"BLOCK_NUMBER"`
}

func (TRANSACTION) TableName() string {
	return "TRANSACTION"
}

type BLOCK struct {
	ID         int64     `gorm:"primaryKey;column:ID" json:"ID"`
	CREATETIME time.Time `gorm:"column:CREATE_TIME;default:null" json:"CREATE_TIME"`
	UPDATETIME time.Time `gorm:"column:UPDATE_TIME;default:null" json:"UPDATE_TIME"`
	BLOCKNUM   int64     `gorm:"column:BLOCK_NUM" json:"BLOCK_NUM"`
	BLOCKHASH  string    `gorm:"column:BLOCK_HASH" json:"BLOCK_HASH"`
	BLOCKSIZE  string    `gorm:"column:BLOCK_SIZE" json:"BLOCK_SIZE"`
}

func (BLOCK) TableName() string {
	return "BLOCK"
}

type HEADER struct {
	ID          int64     `gorm:"primaryKey;column:ID" json:"ID"`
	CREATETIME  time.Time `gorm:"column:CREATE_TIME;default:null" json:"CREATE_TIME"`
	UPDATETIME  time.Time `gorm:"column:UPDATE_TIME;default:null" json:"UPDATE_TIME"`
	PARENTHASH  string    `gorm:"column:PARENT_HASH" json:"PARENT_HASH"`
	UNCLEHASH   string    `gorm:"column:UNCLE_HASH" json:"UNCLE_HASH"`
	COINBASE    string    `gorm:"column:COIN_BASE" json:"COIN_BASE"`
	ROOT        string    `gorm:"column:ROOT" json:"ROOT"`
	TXHASH      string    `gorm:"column:TX_HASH" json:"TXHASH"`
	RECEIPTHASH string    `gorm:"column:RECEIPT_HASH" json:"RECEIPTHASH"`
	//BLOOM       string     `gorm:"column:BLOOM" json:"BLOOM"`
	DIFFICULTY  int64  `gorm:"column:DIFFICULTY" json:"DIFFICULTY"`
	BLOCKNUMBER int64  `gorm:"column:BLOCK_NUMBER" json:"BLOCK_NUMBER"`
	GASLIMIT    uint64 `gorm:"column:GAS_LIMIT" json:"GAS_LIMIT"`
	GASUSED     uint64 `gorm:"column:GAS_USED" json:"GAS_USED"`
	TIME        uint64 `gorm:"column:TIME" json:"TIME"`
	EXTRA       string `gorm:"column:EXTRA" json:"EXTRA"`
	NONCE       string `gorm:"column:NONCE" json:"NONCE"`
	BASEFEE     int64  `gorm:"column:BASEFEE" json:"BASEFEE"`
}

func (HEADER) TableName() string {
	return "HEADER"
}
