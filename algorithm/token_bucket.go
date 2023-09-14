package main

import (
	"log"
	"time"
)

type TokenBucket struct {
	capacity      int           //令牌桶的容量
	tokens        chan struct{} // 令牌通道
	tokenRate     time.Duration //令牌生成速率
	lastTokenTime time.Time     //上次生成令牌的时间
}

// 实例化令牌桶
func NewTokenBucket(capacity int, rate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:      capacity,
		tokens:        make(chan struct{}, capacity),
		tokenRate:     rate,
		lastTokenTime: time.Now(),
	}
}

// 获取一个令牌
func (tb *TokenBucket) getToken() bool {
	select {
	case <-tb.tokens:
		return true
	default:
		return false
	}
}

// 放回一个令牌
func (tb *TokenBucket) putToken() {
	tb.tokens <- struct{}{}
}

// 等待获取一个令牌
func (tb *TokenBucket) Wait() {
	for {
		if tb.getToken() {
			return
		}
		time.Sleep(tb.tokenRate)
	}
}

// 周期性向令牌桶中添加令牌
func (tb *TokenBucket) Refill() {
	for {
		elapsed := time.Since(tb.lastTokenTime)
		if elapsed >= tb.tokenRate {
			numTokens := int(elapsed / tb.tokenRate)
			for i := 0; i < numTokens; i++ {
				if len(tb.tokens) < tb.capacity {
					tb.putToken()
				}
			}
			tb.lastTokenTime = tb.lastTokenTime.Add(time.Duration(numTokens) * tb.tokenRate)
			time.Sleep(tb.tokenRate)

		}
	}
}

func main() {
	// 创建一个容量为 1，速率为 1 秒/个的令牌桶
	tb := NewTokenBucket(1, time.Second)

	// 启动周期性添加令牌的协程
	go tb.Refill()

	// 模拟请求
	for i := 0; i < 10; i++ {
		log.Printf("Request %d: ", i)
		if tb.getToken() {
			log.Println("OK")
		} else {
			log.Println("Failed")
		}
		time.Sleep(2000 * time.Millisecond)
	}
}
