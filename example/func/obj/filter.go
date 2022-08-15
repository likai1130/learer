package main

import (
	"context"
	"fmt"
	"io"
)

/**
模拟构造数据库查询过滤条件
*/

type DbFilter struct {
	ctx    context.Context        // 上下文
	dbName string                 //数据库名称
	db     interface{}            //数据库实例
	filter map[string]interface{} //过滤条件
}

// FooOption 代表可选参数
type Option func(filter *DbFilter)

func WithTokenId(tokenId string) Option {
	return func(filter *DbFilter) {
		filter.filter["token_id"] = tokenId
	}
}

func WithBuyer(buyer string) Option {
	return func(filter *DbFilter) {
		filter.filter["buyer"] = buyer
	}
}

func WithOwner(owner string) Option {
	return func(filter *DbFilter) {
		filter.filter["owner"] = owner
	}
}

func WithPrice(lowPrice, highPrice float64) Option {
	m := make(map[string]float64)
	m["$gte"] = lowPrice
	m["$lte"] = highPrice
	return func(filter *DbFilter) {
		filter.filter["price"] = m
	}
}

func NewFilter(ctx context.Context, dbName string, filterOpts ...Option) *DbFilter {
	m := make(map[string]interface{})
	filter := &DbFilter{
		ctx:    ctx,
		dbName: dbName,
		db:     nil,
		filter: m,
	}
	for _, option := range filterOpts {
		option(filter)
	}
	return filter
}

func main() {
	filter := NewFilter(context.Background(), "kli",
		WithOwner("123"),
		WithPrice(1, 3),
		WithTokenId("123432"))
	fmt.Println(filter.filter)
}

type ReadWriter interface {

	io.Reader
	io.Writer
}

