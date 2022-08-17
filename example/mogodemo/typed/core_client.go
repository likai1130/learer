package typed

import (
	"learner/example/mogodemo/basic"
)

type DataSourceClient struct {
	client basic.DataSource
}

func (d *DataSourceClient) NFTs(chainId uint) NFtsInterface {
	return newNFTs(d,chainId)
}

func (c *DataSourceClient) DataClient() basic.DataSource {
	if c == nil {
		return nil
	}
	return c.client
}

func NewForConfigAndClient(conf *basic.MongoConf) (*DataSourceClient, error) {
	client, err := basic.New(conf)
	if err != nil {
		return nil, err
	}
	return &DataSourceClient{client}, nil
}