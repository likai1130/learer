package model

import "time"

/**
 * 本 Log_Event 主要集中以太坊的各种事件：铸币NFT(ERC721)，铸币ERC20,转移ERC20(直接转移，授权转移)，
 */

type BaseEthLog struct {
	Id          string `bson:"_id,omitempty" json:"id"`          //mongoDB 持久化主键
	Address     string `bson:"address" json:"address"`           //合约地址
	BlockNumber uint64 `bson:"block_number" json:"block_number"` //区块号
	TxHash      string `bson:"tx_hash" json:"tx_hash"`           // 交易hash
	TxIndex     uint `bson:"tx_index" json:"tx_index"`         //交易内部序号
	BlockHash   string `bson:"block_hash" json:"block_hash"`     //区块hash
	Index       uint `bson:"index" json:"index"`               //区块内部序号
	Removed     bool   `bson:"removed" json:"removed"`           //状态 //如果本次交易不成功，有日志，需要重新提交在下一个块执行merge，产生新的日志
	CreatedAt time.Time `bson:"created_at" json:"created_at"` // 创建时间 -------------------需要重新定义
}

// EthLogCoinSolar /** 铸币，直接转移，授权转移
type EthLogCoinSolar struct {
	//这里主要解析topic，
	BaseEthLog BaseEthLog `bson:"baseethlog" json:"baseethlog"`
	EventName  string     `bson:"event_name" json:"event_name"` //事件名称
	EventHash  string     `bson:"event_hash" json:"event_hash"` //事件hash
	From       string     `bson:"from" json:"from"`             // 从owner 转出。铸币ERC20，默认0地址
	To         string     `bson:"to" json:"to"`                 // 铸币的接收方
	Amount     uint64     `bson:"amount" json:"amount"`         //数量
}

// EthLogPurchasePromise /**
type EthLogPurchasePromise struct {
	//解析topic、data 数据
	BaseEthLog   BaseEthLog `bson:"baseethlog" json:"baseethlog"`
	EventName    string     `bson:"event_name" json:"event_name"`       //事件名称
	EventHash    string     `bson:"event_hash" json:"event_hash"`       //事件hash
	Buyer        string     `bson:"buyer" json:"buyer"`                 //买方
	SpenderProxy string     `bson:"spender_proxy" json:"spender_proxy"` //授权代理方，这里一般是，交易市场
	TokenId      string     `bson:"token_id" json:"token_id"`           //待售物 nft 的 Token ID
	Price        uint64     `bson:"price" json:"price"`                 // 待售物的 售价
	AddedValue   uint64     `bson:"added_value" json:"added_value"`     //授权可操作的代币数量。一般情况下，要比price 大，不然后期转移数据将会受限
}

// EthLogMintNFT /**
type EthLogMintNFT struct {
	//这里主要解析topic，Data 数据信息
	BaseEthLog BaseEthLog `bson:"baseethlog" json:"baseethlog"`
	EventName  string     `bson:"event_name" json:"event_name"` //事件名称
	EventHash  string     `bson:"event_hash" json:"event_hash"` //事件hash
	To         string     `bson:"to" json:"to"`                 //topic
	TokenId    string     `bson:"token_id" json:"token_id"`     //topic
	TokenURI   string     `bson:"token_uri" json:"token_uri"`   // data
}

// EthLogSaleOrder /挂单，改单，撤单
type EthLogSaleOrder struct {
	BaseEthLog   BaseEthLog `bson:"baseethlog" json:"baseethlog"`
	EventName    string     `bson:"event_name" json:"event_name"`       //事件名称
	EventHash    string     `bson:"event_hash" json:"event_hash"`       //事件hash topic[0]
	Operator     string     `bson:"operator" json:"operator"`           //申请挂单的owner topic[1]
	TokenId      string     `bson:"token_id" json:"token_id"`           // 待售 NFT topic[2]
	OrderId      string     `bson:"order_id" json:"order_id"`           //订单ID topic[3]
	FeePercent   uint64     `bson:"fee_percent" json:"fee_percent"`     //抽成 data
	SellingPrice uint64     `bson:"selling_price" json:"selling_price"` //售价 data
}

// EthLogTransferNFT /**
type EthLogTransferNFT struct {
	BaseEthLog BaseEthLog `bson:"baseethlog" json:"baseethlog"`
	EventName  string     `bson:"event_name" json:"event_name"` //事件名称
	EventHash  string     `bson:"event_hash" json:"event_hash"` //事件hash topic[0]
	Buyer      string     `bson:"buyer" json:"buyer"`           // 买家 topic[1]
	Seller     string     `bson:"seller" json:"seller"`         //卖家 topic[2]
	TokenId    string     `bson:"token_id" json:"token_id"`     // 交易的NFT Token  topic[3]
	TradeId    string     `bson:"trade_id" json:"trade_id"`     // 交易ID data
	OrderId    string     `bson:"order_id" json:"order_id"`     // 订单ID data
}

// EthLogTradeTickets /**
type EthLogTradeTickets struct {
	BaseEthLog   BaseEthLog `bson:"baseethlog" json:"baseethlog"`
	EventName    string     `bson:"event_name" json:"event_name"`             //事件名称
	EventHash    string     `bson:"event_hash" json:"event_hash"`             //事件hash topic[0]
	Buyer        string     `bson:"buyer" json:"buyer"`                       //买家 topic[1]
	Seller       string     `bson:"seller" json:"seller"`                     //卖家 topic[2]
	TokenId      string     `bson:"token_id" json:"token_id"`                 //待售 NFT topic[3]
	TradeId      string     `bson:"trade_id" json:"trade_id"`                 // 交易ID
	OrderId      string     `bson:"order_id" json:"order_id"`                 //订单ID
	Price        uint64     `bson:"price" json:"price"`                       // 售价
	FeePercent   uint64     `bson:"fee_percent" json:"fee_percent"`           // 抽成
	SellerIncome uint64     `bson:"seller_income" json:"seller_income"`       //卖家收益
	OrganiserIncome uint64 `bson:"organiser_income" json:"organiser_income"` //平台收益
}
