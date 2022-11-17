package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var ErrInvalidVerificationCode = errors.New("invalid verification code")

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(User)
		//r, _ := v["role"].(SysRole)
		identity := map[string]interface{}{
			"UserId":      u.Id,
			"Username":    u.UserName,
			"PhoneNumber": u.PhoneNumber,
			"AuthKey":     u.AuthKey,
			"MasterKey":   u.MasterKey,
			"Status":      u.Status,
			"RsaPubKey":   u.RsaPubKey,
		}
		return jwt.MapClaims{
			jwt.IdentityKey: identity,
		}
	}
	return jwt.MapClaims{}
}

func Authorizator(data interface{}, c *gin.Context) bool {
	//identity := data.(map[string]interface{})
	//_, ok := identity["UserId"]
	return true
}
