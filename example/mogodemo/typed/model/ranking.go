package model

import "time"

/**
 * 交易涨跌幅：
 *    按照 nft 的 tokenId ，记录累计成交量，以及与上一次的成交的涨跌幅，以及未来累计涨跌幅
 */
type NFTTradePriceLimit struct {
	Id                        string    `bson:"_id,omitempty" json:"id"`
	ContractAddress           string    `bson:"contract_address" json:"contract_address"` // 应用铸币 合约地址
	TokenId                   string    `bson:"token_id" json:"token_id"`                 //nft 的token
	LastPrice                 uint64    `bson:"last_price" json:"last_price"`             //上次售价（标价）
	LastSaleTime              time.Time `bson:"last_sale_time" json:"last_sale_time"`
	LastTradeId               string    `bson:"last_trade_id" json:"last_trade_id"`
	CurrentPrice              uint64    `bson:"current_price" json:"current_price"` //本地售价
	CurrentSaleTime           time.Time `bson:"current_sale_time" json:"current_sale_time"`
	CurrentTradeId            string    `bson:"current_trade_id" json:"current_trade_id"`
	PriceGap                  int64     `bson:"price_gap" json:"price_gap"`                       //差价；当前价格-上次价格，有可能是负数
	GapRangePercent           float64   `bson:"gap_range_percent" json:"gap_range_percent"`       //涨跌幅：百分比；PriceGap/lastPrice*100%
	TradeTotalAmount          uint64    `bson:"trade_total_amount" json:"trade_total_amount"`     //交易总量
	GenesisPrice              uint64    `bson:"genesis_price" json:"genesis_price"`               //创世交易价格，即第一次价格
	GenesisTradeId            string    `bson:"genesis_trade_id" json:"genesis_trade_id"`         //创世交易Id，即第一次交易Id
	GenesisSaleTime           time.Time `bson:"genesis_sale_time" json:"genesis_sale_time"`       //创世交易时间，即第一次交易时间
	AccumulatePriceGap        int64     `bson:"accumulate_price_gap" json:"accumulate_price_gap"` //累积差价：当前价格-创世价格
	AccumulateGapRangePercent float64   `bson:"accumulate_gap_range_percent" json:"accumulate_gap_range_percent"`
}
