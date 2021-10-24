package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"log"


)

/**
	NFT信息
 */
type NFT struct {
	Name string `json:"name,omitempty"` 							//NFT名字
	ImageUrl string `json:"image_url,omitempty"` 					//缩图地址
	ImgePreviewUrl string `json:"image_preview_url,omitempty"`		//缩图地址，预览图（s250）
	ImageThumbnailUrl string `json:"image_thumbnail_url,omitempty"`	//缩图地址（s128）
	ImageOriginalUrl string `json:"image_original_url,omitempty"` 	//原图地址
	NumSales int `json:"num_sales,omitempty"`						//销售数量
	Description string `json:"description,omitempty"` 				//描述
	ExternalLink string `json:"eexternal_link,omitempty"`			//NFT相关外部链接
	TokenId string `json:"token_id,omitempty"`						//NFT tokenID
	ChainId string `json:"chain_id,omitempty"`						//上链地址
	Status int `json:"status,omitempty"`							//上链状态 0为未上链，1已上链
	Price float64 `json:"price,omitempty"` 							//NFT价格
	AssetContract AssetContract `json:"asset_contract,omitempty"`	//合约信息
	DataInfo DataDescipt `json:"data_info,omitempty"`				//数据信息
	CreateTime string `json:"create_time,omitempty"`				//创建时间

}

/**
	合约信息
 */
type AssetContract struct {
	Address string `json:"address,omitempty"`						//钱包地址
	AssetContractType string `json:"asset_contract_type,omitempty"` //合约类型
	CreatedDate string `json:"created_date,omitempty"`				//创建时间
	Name string `json:"name,omitempty"`								//合约名称
	NftVersion string `json:"nft_version,omitempty"`				//NFT版本
	Owner int `json:"owner,omitempty"`								//拥有者ID
	SchemaName string `json:"schema_name,omitempty"`				//模式
	Symbol string `json:"symbol,omitempty"`
	Description string `json:"description,omitempty"`				//合约描述
	ExternalLink string `json:"external_link,omitempty"`			//合约相关链接
	ImageURL string `json:"image_url,omitempty"`					//token地址
}

/**
	数据描述信息
 */
type DataDescipt struct {
	HashAndSuffixUrl string `json:"hash_and_suffix_url,omitempty"`	//hash.后缀
	DataFingerprint string `json:"data_fingerprint,omitempty"`		//数据指纹
	DataPreview string `json:"data_preview,omitempty"`				//数据预览
	Size int64 `json:"size,omitempty"`								//数据大小
	CreateTime string `json:"create_time,omitempty"`				//创建时间
	ExpirationTime string `json:"expiration_time,omitempty"`		//过期时间
}


type Student struct {
	Name string
	Age int
	Sex int
}

type RedisConfig struct {
	Addrs []string `json:"addrs",env:"REDIS_ADDRS"`
	Pwd string `json:"pwd"`
	PoolSize int `json:"pool_size"`
	DB int `json:"db"` //单机下选择DB，集群模式无效
}


var myNft = `{
    "name": "MyNFT",
    "description": "demo",
    "token_id": "240440361554912285924556170495673876140",
    "chain_id": "123",
    "price": 1,
    "asset_contract": {
      "address": "0xaDA389f8b4bE8D9890dA6E4443079dC04C16f0ce",
      "nft_version": "v0.1"
    },
    "data_info": {
      "hash_and_suffix_url": "https://pnode.solarfs.io/dn/short/b4e321144937ca55e85e79b76241beac-程序员5.jpg",
      "data_fingerprint": "b4e321144937ca55e85e79b76241beac",
      "size": 68073,
      "create_time": "20210621141912",
      "expiration_time": "20220621235959"
    },
    "create_time": "2021-06-22 19:23:47"
  }`

var RedisClient redis.Cmdable

func (c RedisConfig) Connect()(redis.Cmdable, error)  {
	addrNum := len(c.Addrs)
	if addrNum == 0{
		return nil, errors.New("redis addr is absent")
	}

	if addrNum > 1{

		RedisClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    c.Addrs,
			Password: c.Pwd,
			PoolSize: c.PoolSize,
		})
		_, err := RedisClient.Ping().Result()
		if err != nil {
			return nil,errors.New(fmt.Sprintf("redis connection failed: ", err.Error()))
		}
		return RedisClient,nil

	}

	RedisClient =  redis.NewClient(&redis.Options{
		Addr: c.Addrs[0],
		Password: c.Pwd,
		PoolSize: c.PoolSize,
		DB: c.DB,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		return nil,errors.New(fmt.Sprintf("redis connection failed: ", err.Error()))
	}
	return RedisClient,nil


}

func main() {
	config := RedisConfig{

		[]string{"localhost:6379"},
		"",
		1000,
		0,
	}
	client, e := config.Connect()
	if e != nil {
		panic(e)
	}
	RedisClient = client
	//ExampleClient_List()
	//RedisClient.Del("hash_test")
	//ExampleClient_Hash()
	//ExampleClient_Set()
	//RedisClient.Set("key","aa",10000000)

	/*s := RedisClient.Exists("key").Val()
	fmt.Println(s)*/


	//count, _ := RedisClient.SCard("set_test").Result()

	retAll, err := RedisClient.HGetAll("/nft/0x8fD86498ff1B750454f5187cc0B2bC6b51d0E07B/4").Result()
	log.Println("retAll",retAll,err)
}

func ExampleClient_List() {
	log.Println("ExampleClient_List")
	defer log.Println("ExampleClient_List")

	kli := Student{
		Name: "kli",
		Age:  18,
		Sex:  0,
	}
	zhangsan := Student{
		Name: "zhangsan",
		Age:  20,
		Sex:  0,
	}
	lisi := Student{
		Name: "lisi",
		Age: 22,
		Sex:  0,
	}

	wangwu := Student{
		Name: "wangwu",
		Age: 23,
		Sex:  0,
	}

	bytesKli, _ := json.Marshal(kli)
	bytesZhangsan,_ := json.Marshal(zhangsan)
	bytesLisi, _ := json.Marshal(lisi)
	bytesWangwu, _ := json.Marshal(wangwu)


	//添加
	RedisClient.RPush("list_test", bytesKli,bytesZhangsan,bytesLisi,bytesWangwu)

	//设置
	//log.Println(RedisClient.LSet("list_test", 2, "kli").Err())

	//remove
	/*ret, err := RedisClient.LRem("list_test", 3, "message2").Result()
	log.Println(ret, err)*/

	//遍历
	rLen, err := RedisClient.LLen("list_test").Result()
	log.Println(rLen, err)
	lists, err := RedisClient.LRange("list_test", 0, rLen-1).Result()
	log.Println("LRange", lists, err)



	//pop没有时阻塞
	/*result, err := RedisClient.BLPop(1*time.Second, "list_test").Result()
	log.Println("result:", result, err, len(result))*/
}

func ExampleClient_Hash() {
	log.Println("ExampleClient_Hash")
	defer log.Println("ExampleClient_Hash")
	tokenId := "0x8fD86498ff1B750454f5187cc0B2bC6b51d0E07B"
	var datas  = map[string]interface{}{}
	datas["0x8fD86498ff1B750454f5187cc0B2bC6b51d0E07B"] = myNft

	var datas2  = map[string]interface{}{}
	datas2["0x8fD86498ff1B750454f5187cc0B2bC6b51d0E07C"] = myNft

	//添加
	if err := RedisClient.HMSet("hash_test",datas).Err(); err != nil{
		log.Fatal(err)
	}
	if err := RedisClient.HMSet("hash_test",datas2).Err(); err != nil{
		log.Fatal(err)
	}

	//获取
	rets, err := RedisClient.HMGet("hash_test", tokenId).Result()
	log.Println("rets:", rets, err)

	//成员
	retAll, err := RedisClient.HGetAll("hash_test").Result()
	log.Println("retAll",retAll,err)

	/*//存在
	bExist, err := RedisClient.HExists("hash_test", "tel").Result()
	log.Println(bExist, err)

	bRet, err := RedisClient.HSetNX("hash_test", "id", 100).Result()
	log.Println(bRet, err)

	//删除
	log.Println(RedisClient.HDel("hash_test", "age").Result())*/
}

func ExampleClient_Set() {
	log.Println("ExampleClient_Set")
	defer log.Println("ExampleClient_Set")

	kli := Student{
		Name: "kli",
		Age:  18,
		Sex:  0,
	}
	zhangsan := Student{
		Name: "zhangsan",
		Age:  20,
		Sex:  0,
	}
	lisi := Student{
		Name: "lisi",
		Age: 23,
		Sex:  0,
	}

	wangwu := Student{
		Name: "wangwu",
		Age: 23,
		Sex:  0,
	}
	wangwu2 := Student{
		Name: "wangwu",
		Age: 24,
		Sex:  1,
	}

	bytesKli, _ := json.Marshal(kli)
	bytesZhangsan,_ := json.Marshal(zhangsan)
	bytesLisi, _ := json.Marshal(lisi)
	bytesWangwu, _ := json.Marshal(wangwu)
	bytesWangwu2, _ := json.Marshal(wangwu2)

	//添加
	ret, err := RedisClient.SAdd("set_test", bytesKli, bytesZhangsan, bytesLisi,bytesWangwu2,bytesWangwu).Result()
	log.Println(ret, err)

	//数量
	count, err := RedisClient.SCard("set_test").Result()
	log.Println(count, err)


	//成员
	members, err := RedisClient.SMembers("set_test").Result()
	log.Println(members, err)

	//删除
	/*ret, err = RedisClient.SRem("set_test",bytesKli, bytesZhangsan, bytesLisi,bytesWangwu2,bytesWangwu).Result()
	log.Println(ret, err)*/


	bret, err := RedisClient.SIsMember("set_test", bytesWangwu).Result()
	log.Println(bret, err)



}