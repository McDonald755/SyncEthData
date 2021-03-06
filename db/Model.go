package db

import (
	"time"
)

type TRANSACTION struct {
	ID         int64     `gorm:"primaryKey;column:ID" json:"ID"`
	CreateTime time.Time `gorm:"column:create_time;default:null" json:"CREATE_TIME"`
	UpdateTime time.Time `gorm:"column:update_time;default:null" json:"UPDATE_TIME"`
	TxData     string    `gorm:"column:tx_data" json:"TX_DATA"`
	Hash       string    `gorm:"column:hash" json:"HASH"`
	Size       string    `gorm:"column:size" json:"SIZE"`
	//FROMACCOUNT string    `gorm:"column:FROM_ACCOUNT" json:"FROM_ACCOUNT"`
	ToAccount   string `gorm:"column:to_account" json:"TO_ACCOUNT"`
	TxnType     int64  `gorm:"column:txn_type" json:"TXNTYPE"`
	Value       string `gorm:"column:value" json:"VALUE"`
	BlockNumber int64  `gorm:"column:block_number" json:"BLOCK_NUMBER"`
}

func (TRANSACTION) TableName() string {
	return "TRANSACTION"
}

type BLOCK struct {
	ID         int64     `gorm:"primaryKey;column:ID" json:"ID"`
	CreateTime time.Time `gorm:"column:create_time;default:null" json:"CREATE_TIME"`
	UpdateTime time.Time `gorm:"column:update_time;default:null" json:"UPDATE_TIME"`
	BlockNum   int64     `gorm:"column:block_num" json:"BLOCK_NUM"`
	BlockHash  string    `gorm:"column:block_hash" json:"BLOCK_HASH"`
	BlockSize  string    `gorm:"column:block_size" json:"BLOCK_SIZE"`
}

func (BLOCK) TableName() string {
	return "BLOCK"
}

type HEADER struct {
	ID          int64     `gorm:"primaryKey;column:ID" json:"ID"`
	CreateTime  time.Time `gorm:"column:create_time;default:null" json:"CREATE_TIME"`
	UpdateTime  time.Time `gorm:"column:update_time;default:null" json:"UPDATE_TIME"`
	ParentHash  string    `gorm:"column:parent_hash" json:"PARENT_HASH"`
	UncleHash   string    `gorm:"column:uncle_hash" json:"UNCLE_HASH"`
	CoinHase    string    `gorm:"column:coin_base" json:"COIN_BASE"`
	Root        string    `gorm:"column:root" json:"ROOT"`
	TxHash      string    `gorm:"column:tx_hash" json:"TXHASH"`
	ReceiptHash string    `gorm:"column:receipt_hash" json:"RECEIPTHASH"`
	//BLOOM       string     `gorm:"column:BLOOM" json:"BLOOM"`
	Difficulty  int64  `gorm:"column:difficulty" json:"DIFFICULTY"`
	BlockNumber int64  `gorm:"column:block_number" json:"BLOCK_NUMBER"`
	GasLimit    uint64 `gorm:"column:gas_limit" json:"GAS_LIMIT"`
	GasUsed     uint64 `gorm:"column:gas_used" json:"GAS_USED"`
	Time        uint64 `gorm:"column:time" json:"TIME"`
	Extra       string `gorm:"column:extra" json:"EXTRA"`
	Nonce       string `gorm:"column:nonce" json:"NONCE"`
	Basefee     int64  `gorm:"column:basefee" json:"BASEFEE"`
}

func (HEADER) TableName() string {
	return "HEADER"
}
