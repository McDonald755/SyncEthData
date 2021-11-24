package db

type TRANSACTION struct {
	ID          int64  `orm:"ID" json:"ID"`
	CREATETIME  string `orm:"CREATE_TIME" json:"CREATE_TIME"`
	UPDATETIME  string `orm:"UPDATE_TIME" json:"UPDATE_TIME"`
	TXDATA      string `orm:"TX_DATA" json:"TX_DATA"`
	HASH        string `orm:"HASH" json:"HASH"`
	SIZE        string `orm:"SIZE" json:"SIZE"`
	FROMACCOUNT string `orm:"FROM_ACCOUNT" json:"FROM_ACCOUNT"`
	BLOCKNUMBER string `orm:"BLOCK_NUMBER" json:"BLOCK_NUMBER"`
}

func (TRANSACTION) TableName() string {
	return "TRANSACTION"
}

type BLOCK struct {
	ID         int64  `orm:"ID" json:"ID"`
	CREATETIME string `orm:"CREATE_TIME" json:"CREATE_TIME"`
	UPDATETIME string `orm:"UPDATE_TIME" json:"UPDATE_TIME"`
	BLOCKNUM   int64  `orm:"BLOCK_NUM" json:"BLOCK_NUM"`
	BLOCKHASH  string `orm:"BLOCK_HASH" json:"BLOCK_HASH"`
	BLOCKSIZE  string `orm:"BLOCK_SIZE" json:"BLOCK_SIZE"`
}

func (BLOCK) TableName() string {
	return "BLOCK"
}

type HEADER struct {
	ID          int64  `orm:"ID" json:"ID"`
	CREATETIME  string `orm:"CREATE_TIME" json:"CREATE_TIME"`
	UPDATETIME  string `orm:"UPDATE_TIME" json:"UPDATE_TIME"`
	PARENTHASH  string `orm:"PARENT_HASH" json:"PARENT_HASH"`
	UNCLEHASH   string `orm:"UNCLE_HASH" json:"UNCLE_HASH"`
	COINBASE    string `orm:"COIN_BASE" json:"COIN_BASE"`
	ROOT        string `orm:"ROOT" json:"ROOT"`
	TXHASH      string `orm:"TXHASH" json:"TXHASH"`
	RECEIPTHASH string `orm:"RECEIPTHASH" json:"RECEIPTHASH"`
	BLOOM       string `orm:"BLOOM" json:"BLOOM"`
	DIFFICULTY  int64  `orm:"DIFFICULTY" json:"DIFFICULTY"`
	BLOCKNUMBER int64  `orm:"BLOCK_NUMBER" json:"BLOCK_NUMBER"`
	GASLIMIT    string `orm:"GAS_LIMIT" json:"GAS_LIMIT"`
	GASUSED     string `orm:"GAS_USED" json:"GAS_USED"`
	TIME        string `orm:"TIME" json:"TIME"`
	EXTRA       string `orm:"EXTRA" json:"EXTRA"`
	NONCE       string `orm:"NONCE" json:"NONCE"`
	BASEFEE     string `orm:"BASEFEE" json:"BASEFEE"`
}

func (HEADER) TableName() string {
	return "HEADER"
}
