package main

import (
	"fmt"
	"strconv"
)

/**
	多态

	1. 不同商品有相同的行为
	2. 商品的不同类别有相同行为，产生不同结果
 */

type Good interface {
	settleAccount() int
	orderInfo() string
}

type Phone struct {
	name string
	quantity int
	price int
}

/**
	手机结算账户
 */
func (phone Phone) settleAccount() int {
	return phone.quantity * phone.price
}

/**
	手机订单信息
 */
func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity)+ "个" +
		phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}


type FreeGift struct {
	name string
	quantity int
	price int
}

/**
	赠品结算
 */
func (FreeGift) settleAccount() int {
	return 0
}

/**
	赠品定单
 */
func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity)+ "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

var phoneInstance *Phone
var freeGiftInstance *FreeGift

func NewPhone() Good {
	if phoneInstance ==nil {
		phoneInstance = &Phone{}
	}
	return phoneInstance
}


func NewFreeGift() Good {
	if freeGiftInstance ==nil {
		freeGiftInstance = &FreeGift{}
	}
	return freeGiftInstance
}



/**
	结算
 */

func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}

func main() {
	iPhone := Phone{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}

	earphones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}

	goods := []Good{iPhone, earphones}
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单总共需要支付 %d 元",allPrice)
}