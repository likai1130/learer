package model

// MyCollection
//我的收藏中间关系/**
type MyCollection struct {
	Name              string `bson:"name" json:"name,required,omitempty"`                               //NFT名字
	ImageThumbnailUrl string `bson:"image_thumbnail_url" json:"image_thumbnail_url,required,omitempty"` //缩图地址（s128）
	WalletAddress     string `bson:"wallet_address" json:"wallet_address,required"`                     //钱包地址
	TokenId           string `bson:"token_id" json:"token_id,required"`
	ContractAddress   string `bson:"contract_address" json:"contract_address"`
	Rarity            string `bson:"rarity" json:"rarity"`
}
