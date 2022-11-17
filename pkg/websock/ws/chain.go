package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Resp struct {
	Code    int
	Message string
	Data    interface{}
}

var database map[string]string

func init() {
	database = make(map[string]string)
}

func MintNft() Resp {
	fmt.Println("正在铸币")
	tokenId := "31e4338a201049599ae7634d95a476f7"
	go exec(tokenId)

	return Resp{
		Code:    0,
		Message: "success",
		Data:    tokenId,
	}
}

func exec(tokenId string) {
	time.Sleep(2 * time.Second)
	database[tokenId] = "777777777777777777777777777777"
	fmt.Println("铸币成功")

	//todo 写到socck
	WebsocketManager.SendAll([]byte("777777777777777777777777777777"))
}

func Mint(ctx *gin.Context) {
	nft := MintNft()
	ctx.JSON(200, nft)
}
