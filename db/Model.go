package db

import (
	"math/big"
	"time"
)

type TRANSACTION struct {
	ID          int64     `orm:"ID" json:"ID"`
	CREATETIME  time.Time `orm:"CREATE_TIME" json:"CREATE_TIME"`
	UPDATETIME  time.Time `orm:"UPDATE_TIME" json:"UPDATE_TIME"`
	TXDATA      string    `orm:"TX_DATA" json:"TX_DATA"`
	HASH        string    `orm:"HASH" json:"HASH"`
	SIZE        string    `orm:"SIZE" json:"SIZE"`
	FROMACCOUNT string    `orm:"FROM_ACCOUNT" json:"FROM_ACCOUNT"`
	BLOCKNUMBER *big.Int  `orm:"BLOCK_NUMBER" json:"BLOCK_NUMBER"`
}

func (TRANSACTION) TableName() string {
	return "TRANSACTION"
}

type BLOCK struct {
	ID         int64     `orm:"ID" json:"ID"`
	CREATETIME time.Time `orm:"CREATE_TIME" json:"CREATE_TIME"`
	UPDATETIME time.Time `orm:"UPDATE_TIME" json:"UPDATE_TIME"`
	BLOCKNUM   *big.Int  `orm:"BLOCK_NUM" json:"BLOCK_NUM"`
	BLOCKHASH  string    `orm:"BLOCK_HASH" json:"BLOCK_HASH"`
	BLOCKSIZE  string    `orm:"BLOCK_SIZE" json:"BLOCK_SIZE"`
}

func (BLOCK) TableName() string {
	return "BLOCK"
}

type HEADER struct {
	ID          int64     `orm:"ID" json:"ID"`
	CREATETIME  time.Time `orm:"CREATE_TIME" json:"CREATE_TIME"`
	UPDATETIME  time.Time `orm:"UPDATE_TIME" json:"UPDATE_TIME"`
	PARENTHASH  string    `orm:"PARENT_HASH" json:"PARENT_HASH"`
	UNCLEHASH   string    `orm:"UNCLE_HASH" json:"UNCLE_HASH"`
	COINBASE    string    `orm:"COIN_BASE" json:"COIN_BASE"`
	ROOT        string    `orm:"ROOT" json:"ROOT"`
	TXHASH      string    `orm:"TXHASH" json:"TXHASH"`
	RECEIPTHASH string    `orm:"RECEIPTHASH" json:"RECEIPTHASH"`
	BLOOM       *big.Int  `orm:"BLOOM" json:"BLOOM"`
	DIFFICULTY  *big.Int  `orm:"DIFFICULTY" json:"DIFFICULTY"`
	BLOCKNUMBER *big.Int  `orm:"BLOCK_NUMBER" json:"BLOCK_NUMBER"`
	GASLIMIT    uint64    `orm:"GAS_LIMIT" json:"GAS_LIMIT"`
	GASUSED     uint64    `orm:"GAS_USED" json:"GAS_USED"`
	TIME        uint64    `orm:"TIME" json:"TIME"`
	EXTRA       string    `orm:"EXTRA" json:"EXTRA"`
	NONCE       uint64    `orm:"NONCE" json:"NONCE"`
	BASEFEE     *big.Int  `orm:"BASEFEE" json:"BASEFEE"`
}

func (HEADER) TableName() string {
	return "HEADER"
}
