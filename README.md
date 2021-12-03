# SyncData

##数据库脚本
use eth_data;
DROP TABLE IF EXISTS BLOCK;
CREATE TABLE `BLOCK` (
`ID` int(11) NOT NULL AUTO_INCREMENT,
`create_time` datetime DEFAULT NULL,
`update_time` datetime DEFAULT NULL,
`block_num` bigint(20) DEFAULT NULL,
`block_hash` varchar(255) DEFAULT NULL,
`BLOCK_SIZE` varchar(255) DEFAULT NULL,
PRIMARY KEY (`ID`),
KEY `block_num` (`block_num`)
) ENGINE=InnoDB AUTO_INCREMENT=7401 DEFAULT CHARSET=utf8mb4 COMMENT='BLOCK'
#
DROP TABLE IF EXISTS HEADER;
CREATE TABLE `HEADER` (
`ID` int(11) NOT NULL AUTO_INCREMENT,
`create_time` datetime DEFAULT NULL,
`update_time` datetime DEFAULT NULL,
`parent_hash` varchar(255) DEFAULT NULL,
`uncle_hash` varchar(255) DEFAULT NULL,
`coin_base` varchar(255) DEFAULT NULL,
`root` varchar(255) DEFAULT NULL,
`tx_hash` varchar(255) DEFAULT NULL,
`receipt_hash` varchar(255) DEFAULT NULL,
`difficulty` bigint(20) DEFAULT NULL,
`block_number` bigint(20) DEFAULT NULL,
`gas_limit` bigint(20) DEFAULT NULL,
`gas_used` bigint(20) DEFAULT NULL,
`time` int(11) DEFAULT NULL,
`extra` longtext,
`nonce` varchar(255) DEFAULT NULL,
`basefee` bigint(20) DEFAULT NULL,
PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=7401 DEFAULT CHARSET=utf8mb4 COMMENT='HEADER'
#
DROP TABLE IF EXISTS TRANSACTION;
CREATE TABLE `TRANSACTION` (
`ID` int(11) NOT NULL AUTO_INCREMENT,
`create_time` datetime DEFAULT NULL,
`update_time` datetime DEFAULT NULL,
`tx_data` longtext,
`hash` varchar(255) DEFAULT NULL,
`size` varchar(255) DEFAULT NULL,
`to_account` varchar(255) DEFAULT NULL,
`txn_type` int(11) DEFAULT NULL,
`value` varchar(255) DEFAULT NULL,
`block_number` bigint(20) DEFAULT NULL,
PRIMARY KEY (`ID`),
UNIQUE KEY `tx_hash` (`hash`),
KEY `block_num` (`block_number`)
) ENGINE=InnoDB AUTO_INCREMENT=554071 DEFAULT CHARSET=utf8mb4 COMMENT='TRANSACTION'