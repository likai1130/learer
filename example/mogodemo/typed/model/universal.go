package model

/**
 * 混合关注通用数据
 */
type NftUniversal struct {
	Id               string      `bson:"_id,omitempty" json:"id"`                    // 主键
	ContractAddress  string      `bson:"contract_address" json:"contract_address"`   // 应用铸币 合约地址
	TokenId          string      `bson:"token_id" json:"token_id"`                   //nft 的token
	OwnerOf          string      `bson:"owner_of" json:"owner_of"`                   //用户地址
	BornType         uint8       `bson:"born_type" json:"born_type"`                 //来源：铸币，交易
	SaleAddress      string      `bson:"sale_address" json:"sale_address"`           // 创建订单的合约地址
	OrderId          string      `bson:"order_id" json:"order_id"`                   //订单Id
	OrderOperator    string      `bson:"order_operator" json:"order_operator"`       // 订单拥有者地址
	FeePercent       uint64      `bson:"fee_percent" json:"fee_percent"`             //抽成比例
	SellingPrice     uint64      `bson:"selling_price" json:"selling_price"`         //售价
	Status           uint8       `bson:"status" json:"status"`                       //订单状态
	MetaWineSpec     SpecDetails `bson:"meta_wine_spec" json:"meta_wine_spec"`       //nft 元数据
	CollectionAmount uint64      `bson:"collection_amount" json:"collection_amount"` //收藏量
}

type CollectionQuantity struct {
	CollectionAmount uint64 `bson:"collection_amount" json:"collection_amount"`
}
