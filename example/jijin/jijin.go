package jijin

import (
	"fmt"
	"strconv"
)

// 基金
type Fund struct {
	Name               string `json:"name"`                //名称
	MaximumDrawdown    int    `json:"maximum_drawdown"`    //最大回撤
	ReplenishmentTimes int    `json:"replenishment_times"` //补仓次数
	Cover              *Cover
}

func NewFund(name string, maximumDrawdown int) *Fund {
	f := &Fund{
		Name:               name,
		MaximumDrawdown:    maximumDrawdown,
		ReplenishmentTimes: maximumDrawdown / 4,
	}
	fmt.Printf("基金名称：%s----最大回撤百分之%d-----可补仓%d次\n", f.Name, f.MaximumDrawdown, f.ReplenishmentTimes)
	return f
}

// 补仓
type Cover struct {
	costPrice               float64 `json:"cost_price"`                //成本单价
	initialPurchaseQuantity float64 `json:"initial_purchase_quantity"` // 初买入数量
	netReplenishmentValue   float64 `json:"net_replenishment_value"`   //补仓净值
	replenishmentQuantity   float64 `json:"replenishment_quantity"`    // 补仓数量
	replenishmentMoney      float64 `json:"replenishment_money"`       // 补仓金额
	finalCost               float64 `json:"final_cost"`                // 最终成本
	finalQuantity           float64 `json:"final_quantity"`            //最终数量
	netWorth4               float64 `json:"net_worth_4"`               // %4净值即下次补仓净值
}

// 补仓后的成本价=（首次买入的总成本+补仓的总成本）÷持有的股票总量。
func newCover(costPrice, initialPurchaseQuantity, netReplenishmentValue, replenishmentMoney float64) *Cover {
	replenishmentQuantity := Decimal(float64(replenishmentMoney / netReplenishmentValue))

	finalQuantity := Decimal(initialPurchaseQuantity + replenishmentQuantity) // 总数量
	costMoney := costPrice * initialPurchaseQuantity                          // 成本金额

	finalCost := (costMoney + replenishmentMoney) / finalQuantity

	netWorth4 := netReplenishmentValue * 0.96

	return &Cover{
		costPrice:               costPrice,
		initialPurchaseQuantity: initialPurchaseQuantity,
		netReplenishmentValue:   Decimal(netReplenishmentValue),
		replenishmentQuantity:   replenishmentQuantity,
		replenishmentMoney:      replenishmentMoney,
		finalCost:               Decimal(finalCost),
		finalQuantity:           finalQuantity,
		netWorth4:               Decimal(netWorth4),
	}
}

func (c *Cover) ToString() {
	out := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v",
		c.costPrice, c.initialPurchaseQuantity, c.netReplenishmentValue, c.replenishmentMoney,
		c.replenishmentQuantity, c.finalCost, c.finalQuantity, c.netWorth4)
	fmt.Println(out)
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return value
}
