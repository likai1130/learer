package jijin

import "testing"

/*
*
计算中欧时代先锋的基金定投多少次可以回本
*/
func TestZhongOuCount(t *testing.T) {
	var costPrice = 1.5234                          //成本单价
	var initialPurchaseQuantity = float64(34162.09) // 初买入数量
	var netReplenishmentValue = 1.2618              //补仓净值
	var replenishmentMoney = float64(1000)          // 补仓数量
	fund := NewFund("中欧时代先锋", 40)
	fund.ExecCover(costPrice, initialPurchaseQuantity, netReplenishmentValue, replenishmentMoney)
}
