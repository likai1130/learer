package typed

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"learner/example/mogodemo/basic"
	"learner/example/mogodemo/typed/model"
)

type NFTsGetter interface {
	NFTs(chain uint) NFtsInterface
}

type NFtsInterface interface {
	List(ctx context.Context, filter Filter) (*model.EthNftMetaList, error)
	Get(ctx context.Context, filter Filter) (*model.EthNftMeta, error)
	Count(ctx context.Context, filter Filter) (int64, error)
}

type Filter = bson.D

type nfts struct {
	client  basic.DataSource
	chainId uint
	collectName string
}

func (n *nfts) List(ctx context.Context, filter Filter) (*model.EthNftMetaList, error) {
	metaList := &model.EthNftMetaList{}
	err := n.client.List(ctx, "", filter).Into(metaList.Items)
	metaList.Filter = filter
	return metaList, err

}

func (n *nfts) Get(ctx context.Context, filter Filter) (*model.EthNftMeta, error) {
	return nil,nil
}

func (n *nfts) Count(ctx context.Context, filter Filter) (int64, error) {
	return 0, nil
}

func newNFTs(d *DataSourceClient, chainId uint) *nfts {
	return &nfts{
		client:  d.DataClient(),
		chainId: chainId,
		collectName: "dasda",
	}
}
