package utils

import (
	"SyncEthData/config"
	"SyncEthData/db"
	"SyncEthData/oracle"
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func TransformData(block *types.Block) {
	if block == nil {
		return
	}
	b := transferBlock(block)
	header := transferHeader(block.Header())

	transactions := []db.TRANSACTION{}
	if len(block.Transactions()) > 0 {
		for _, transaction := range block.Transactions() {
			trx := transferTrx(transaction, block.Header().Number)
			if trx != nil {
				transactions = append(transactions, *trx)
			}
		}
	}
	db.SaveData(b, header, &transactions)
}

func transferHeader(header *types.Header) *db.HEADER {
	var baseFee = big.NewInt(0)
	if header.BaseFee != nil {
		baseFee = header.BaseFee
	}
	result := db.HEADER{
		ParentHash:  header.ParentHash.Hex(),
		UncleHash:   header.UncleHash.Hex(),
		CoinHase:    header.Coinbase.Hex(),
		Root:        header.Root.Hex(),
		TxHash:      header.TxHash.Hex(),
		ReceiptHash: header.ReceiptHash.Hex(),
		//BLOOM:       header.Bloom.Big().String(),
		Difficulty:  header.Difficulty.Int64(),
		BlockNumber: header.Number.Int64(),
		GasLimit:    header.GasLimit,
		GasUsed:     header.GasUsed,
		Time:        header.Time,
		Extra:       hex.Dump(header.Extra),
		Nonce:       strconv.Itoa(int(header.Nonce.Uint64())),
		Basefee:     baseFee.Int64(),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}

	return &result
}

func transferTrx(trx *types.Transaction, num *big.Int) *db.TRANSACTION {
	var (
		toAccount, value string
	)
	if trx.To() != nil && trx.To().Hex() != "" {
		toAccount = trx.To().Hex()
	} else {
		receipt, err := config.CLIENT[0].TransactionReceipt(context.Background(), trx.Hash())
		if err != nil {
			log.Error("Get contract addr error:", trx.Hash().String(), err)
		}
		toAccount = receipt.ContractAddress.String()
	}

	if trx.Value() != nil {
		value = trx.Value().String()
	}
	result := db.TRANSACTION{
		TxData: "0x" + hex.EncodeToString(trx.Data()),
		Hash:   trx.Hash().Hex(),
		Size:   trx.Size().String(),
		//FROMACCOUNT: msg.From().Hex(),
		ToAccount:   toAccount,
		Value:       value,
		TxnType:     int64(trx.Type()),
		BlockNumber: num.Int64(),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	return &result
}

func transferBlock(block *types.Block) *db.BLOCK {
	result := db.BLOCK{
		BlockNum:   block.Number().Int64(),
		BlockHash:  block.Hash().Hex(),
		BlockSize:  block.Size().String(),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	return &result
}

func FixToAccountData() {
	db.FixToAccount()
}

/**
////////////////////////////////////////////////////////////////nft function/////////////////////////////////
*/
func transferOracle(trx *types.Transaction) *db.ORACLE_DATA {
	symbol, name, _ := getTokenNameAndSymbol(trx.To().String(), nil)
	data := db.ORACLE_DATA{
		Address:     trx.To().String(),
		TokenSymbol: symbol,
		TokenName:   name,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	return &data
}
func getTokenNameAndSymbol(addr string, tokenId *big.Int) (string, string, string) {
	var s, n, i string
	address := common.HexToAddress(addr)
	newOracle, err := oracle.NewOracle(address, config.CLIENT[0])
	if err != nil {
		log.Error("Init Oracle Error:", err, "Oracle Addr Is :", addr)
	}
	symbol, err := newOracle.Symbol(nil)
	if err != nil {
		log.Error("Get Token Symbol Error:", err)
		s = "Undefined"
	} else {
		s = symbol
	}

	name, err := newOracle.Name(nil)
	if err != nil {
		log.Error("Get Token Name Error:", err)
		n = "Undefined"
	} else {
		n = name
	}

	if tokenId != nil {
		uri, err := newOracle.TokenURI(nil, tokenId)
		if err != nil {
			log.Error("Get Token Uri Error:", err)
			i = "Undefined"
		} else {
			i = uri
		}
	}
	return s, n, i
}

func CheckOracleType(trxs types.Transactions, oracles map[string]byte) map[string]byte {
	var (
		balanceOf         = "70a08231"
		ownerOf           = "6352211e"
		approve           = "095ea7b3"
		getApproved       = "081812fc"
		setApprovalForAll = "a22cb465"
		isApprovedForAll  = "e985e9c5"
		transferFrom      = "23b872dd"
		safeTransferFrom  = "42842e0e"
		safeTransferFrom2 = "b88d4fde"
	)
	for _, trx := range trxs {
		txData := hex.EncodeToString(trx.Data())
		b1 := strings.Contains(txData, balanceOf)
		b2 := strings.Contains(txData, ownerOf)
		b3 := strings.Contains(txData, approve)
		b4 := strings.Contains(txData, getApproved)
		b5 := strings.Contains(txData, setApprovalForAll)
		b6 := strings.Contains(txData, isApprovedForAll)
		b7 := strings.Contains(txData, transferFrom)
		b8 := strings.Contains(txData, safeTransferFrom)
		b9 := strings.Contains(txData, safeTransferFrom2)
		if b1 && b2 && b3 && b4 && b5 && b6 && b7 && b8 && b9 {
			if _, ok := oracles[trx.To().String()]; !ok {
				oracles[trx.To().String()] = byte(0)
				data := transferOracle(trx)
				db.SaveOracle(data)
			}
		}
	}
	return oracles
}
func TransferNftData(l types.Log) *db.NFT_DATA {
	parseInt, err := strconv.ParseInt(l.Topics[3].Hex(), 0, 16)
	if err != nil {
		log.Error(err)
	}
	symbol, name, uri := getTokenNameAndSymbol(l.Address.String(), big.NewInt(parseInt))
	data := db.NFT_DATA{
		TokenId:     parseInt,
		TokenSymbol: symbol,
		TokenName:   name,
		TokenUri:    uri,
		Owner:       common.HexToAddress(l.Topics[2].Hex()).String(),
		OracleAdd:   l.Address.String(),
	}
	return &data
}

func TransferAccounts(addrs *[]string) *[]common.Address {
	result := []common.Address{}
	for _, s := range *addrs {
		result = append(result, common.HexToAddress(s))
	}
	return &result
}
