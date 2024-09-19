// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserLaunchTx = "user_launch_tx"

// UserLaunchTx mapped from table <user_launch_tx>
type UserLaunchTx struct {
	ID              int32     `gorm:"column:Id;primaryKey;autoIncrement:true" json:"Id"`
	UID             int64     `gorm:"column:uid;not null;comment:用户id" json:"uid"`                                        // 用户id
	ContractAddress string    `gorm:"column:contract_address;not null;comment:合约地址" json:"contract_address"`              // 合约地址
	ChainID         int32     `gorm:"column:chain_id;not null;comment:链Id" json:"chain_id"`                               // 链Id
	PairAddress     string    `gorm:"column:pair_address;not null;comment:对币地址" json:"pair_address"`                      // 对币地址
	Fee             string    `gorm:"column:fee;not null;comment:交易手续费" json:"fee"`                                       // 交易手续费
	FeeTokenSymbol  string    `gorm:"column:fee_token_symbol;not null;comment:交易手续费token的symbol" json:"fee_token_symbol"` // 交易手续费token的symbol
	Timestamp       time.Time `gorm:"column:timestamp;not null;comment:交易时间" json:"timestamp"`                            // 交易时间
	TypeName        string    `gorm:"column:type_name;not null;comment:类别:  luanch" json:"type_name"`                     // 类别:  luanch
	ChainName       string    `gorm:"column:chain_name" json:"chain_name"`
	Tx              string    `gorm:"column:tx;not null;comment:交易哈希值" json:"tx"`                                          // 交易哈希值
	Owner           string    `gorm:"column:owner;comment:token owner address, eth:42|solana:44|ton:33|btc:" json:"owner"` // token owner address, eth:42|solana:44|ton:33|btc:
	Level           string    `gorm:"column:level;comment:等级" json:"level"`                                                // 等级
	CreatedAt       time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	BlockNumber     int64     `gorm:"column:block_number;not null;comment:区块数" json:"block_number"` // 区块数
}

// TableName UserLaunchTx's table name
func (*UserLaunchTx) TableName() string {
	return TableNameUserLaunchTx
}
