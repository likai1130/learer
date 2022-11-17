package main

import "time"

// Session 用户会话
type Session struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	//用户ID
	UserID int64 `json:"user_id"`
	//用户Token
	Token string `json:"token" gorm:"type:varchar(32)"`
	// 终端代理(useragent)
	UserAgent string `json:"user_agent" gorm:"type:text"`
	// ip
	IP string `json:"ip" gorm:"type:varchar(15)"`
	//国家
	AgentCountry int `json:"agent_country" gorm:"type:int(4)"`
	//状态(可强制下线)
	AgentStatus int `json:"agent_status" gorm:"type:int(2)"`
	//创建时间
	CreateTime time.Time `json:"create_time"`
}

// User
type User struct {
	//用户ID
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	//邮箱
	Email string `json:"email,omitempty" gorm:"type:varchar(48)"`
	//用户名
	UserName string `json:"user_name,omitempty" gorm:"type:varchar(20)"`
	//电话
	PhoneNumber string `json:"phone_number,omitempty" gorm:"type:varchar(14)"`
	//认证KEY
	AuthKey string `json:"auth_key,omitempty" gorm:"type:varchar(64)"`
	//用户主KEY
	MasterKey string `json:"master_key,omitempty" gorm:"type:varchar(64)"`
	//RSA 公钥
	RsaPubKey string `json:"rsa_pub_key,omitempty" gorm:"type:text"`
	//RSA私钥
	RsaPriKey string `json:"rsa_pri_key,omitempty" gorm:"type:text"`
	//用户状态[0-正常，1-禁用	]
	Status int `json:"status,omitempty" gorm:"type:int(2)"`
	//salt随机数
	SaltRandom string `json:"salt_random,omitempty" gorm:"type:varchar(128)"`
	//注册时间
	RegisterTime time.Time `json:"register_time,omitempty" gorm:"type:datetime"`
}

type LoginReq struct {
	//手机号
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	//认证key
	AuthKey string `json:"authKey" binding:"required"`
}
