package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Login 登录
// @Tags 用户管理
// @Summary 用户登录
// @Description 使用手机号和AuthKey登录
// @Accept  json
// @Produce  json
// @Param loginVO body LoginReq true "手机号和认证key"
// @Success 200 {object} core.Response{} "{"code": 200, "message":"Ok"}"
// @Failure 400 {object} core.Response
// @Failure 401 {object} core.Response
// @Failure 403 {object} core.Response
// @Failure 500 {object} core.Response
// @Router /v1/login [post]
func Login(c *gin.Context) (interface{}, error) {
	var loginVO LoginReq
	if err := c.ShouldBindJSON(&loginVO); err != nil {
		return nil, err
	}
	//todo 读库，校验 取user

	return User{
		Id:           1,
		Email:        "150",
		UserName:     "",
		PhoneNumber:  "",
		AuthKey:      "",
		MasterKey:    "",
		RsaPubKey:    "",
		RsaPriKey:    "",
		Status:       0,
		SaltRandom:   "",
		RegisterTime: time.Time{},
	}, nil
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}
