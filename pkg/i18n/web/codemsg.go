package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"learner/pkg/i18n/ressponse"
)

func main() {

	c := &gin.Context{}
	/*c.Request.Header.Add("lang","zh-cn")
	c.Request.Header.Add("Accept-Language","zh-CN,zh;q=0.9,en;q=0.8")*/

	responseParams := response.ResponseI18nMsgParams{
		c,
		response.InvalidParams,
		nil,
		errors.New("测试使用"),
		nil,
		nil,
	}

	response.NewResponse(responseParams)

}
