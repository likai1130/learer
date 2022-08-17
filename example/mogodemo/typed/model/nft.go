package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type EthNftMeta struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"_id"` // 主键
	ChainId         uint               `bson:"chain_id" json:"chain_id"`
	ContractAddress string             `bson:"contract_address" json:"contract_address"` // 应用铸币 合约地址
	ContractSymbol  string             `bson:"contract_symbol" json:"contract_symbol"`   // symbol
	ContractName    string             `bson:"contract_name" json:"contract_name"`       // name
	ContractType    string             `bson:"contract_type" json:"contract_type"`       // 合约类型 （ERC721/1155)
	TokenId     string      `bson:"token_id" json:"token_id"`                            //	nft 的token
	Amount      uint64      `bson:"amount" json:"amount"`                                //  nft 数量
	TokenURI    string      `bson:"token_uri" json:"token_uri"`                          //  nft的 metadata的 uri
	Creator     string      `bson:"creator" json:"creator"`                              // 谁创建的
	ToSpender   string      `bson:"to_spender" json:"to_spender"`                        // 初次的拥有者
	BlockNumber uint64      `bson:"block_number" json:"block_number"`                    // 区块高度
	BlockHash   string      `bson:"block_hash" json:"block_hash"`                        // 块的Hash地址
	Index       uint        `bson:"index" json:"index"`                                  // 块内序号
	TxHash      string      `bson:"tx_hash" json:"tx_hash"`                              // 本次交易的hash 地址
	TxIndex     uint        `bson:"tx_index" json:"tx_index" `                           // 本次交易的序号
	MetaSpec    SpecDetails `bson:"meta_spec" json:"meta_spec"`                          // nft 元数据
	CreatedAt   time.Time   `bson:"created_at" json:"created_at"`                        // 创建时间 -------------------需要重新定义
}

type SpecDetails struct {
	Name         string            `bson:"name" json:"name"`                   //NFT名称
	Image        string            `bson:"image" json:"image"`                 //缩图地址（s128）
	Description  string            `bson:"description" json:"description"`     //描述
	Author       string            `bson:"author" json:"author"`               //作者
	Labels       map[string]string `bson:"labels" json:"labels"`               //标签
	Attributes   map[string]string `bson:"attributes" json:"attributes"`       //自定义元数据，某个子类型特有的属性。如：商品类.美术(画).尺寸，类型等
	ReleaseTotal uint64            `bson:"release_total" json:"release_total"` //发布数量
}

type EthNftMetaList struct {
	Filter bson.D `json:"filter"`
	Items []EthNftMeta `json:"items"`
}
