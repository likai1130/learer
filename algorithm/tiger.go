package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/**
	白老虎id生成 tige_99_xxxxxx
 */
func main() {
	idM := make(map[string]bool,100000)
	for i:=0;i<100000;i++ {
		captcha := CreateCaptcha()
		idM[captcha] = true
	}
	fmt.Printf("总数量：%d\n", len(idM))
}

func CreateCaptcha() string{
	source := rand.NewSource(time.Now().UnixNano()).Int63()
	formatInt := strconv.FormatInt(source, 10)
	return formatInt[len(formatInt)-6:]
}