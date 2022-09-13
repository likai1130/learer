package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

// 第二种
func main() {

	case2()
}

func case1() {
	ints := make([]int, 0)
	set3(&ints)
	fmt.Println(ints[0])
}
func set(num *[]int) {
	ints := make([]int, 0)
	ints = append(ints, 1)
	num = &ints
}


func set3(num *[]int) {
	*num = append(*num, 1)
}


//RangeRand 生成区间随机数
func rangeRand(min, max int64) string {
	if min > max {
		panic("the min is greater than max!")
	}
	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))
		r :=  result.Int64() - i64Min
		ds := fmt.Sprintf("%06d", r)
		return ds

	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		r :=  min + result.Int64()
		ds := fmt.Sprintf("%06d", r)
		return ds
	}
}

func case2()  {
	for i:=0;i<999999;i++ {
		s := rangeRand(1, 999999)
		if s[:2] == "00" {
			fmt.Println(s)
		}
	}
}