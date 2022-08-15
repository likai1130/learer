package typed

import (
	"context"
	"learner/example/func/obj/mogodemo/typed/model"
)

type NFTsGetter interface {
	NFTs(chain uint) NFtsInterface
}

type NFtsInterface interface {
	List(ctx context.Context,filter map[string]interface{}) ([]model.EthNftMeta, error)
	Get(ctx context.Context,filter map[string]interface{}) (*model.EthNftMeta, error)
	Count(ctx context.Context,filter map[string]interface{}) (int64, error)
}