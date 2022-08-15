package mongodb

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const ConnectTimeout = 5 * time.Second //连接超时时间

type MongoConf struct {
	Hosts          []string `json:"hosts"`
	UserName       string   `json:"user_name"`
	Password       string   `json:"password"`
	MaxPoolSize    uint64   `json:"max_pool_size"`
	DbName         string   `json:"db_name"`
	ReplicaSet     string   `json:"replica_set"`
	ReadPreference string   `json:"read_preference"`
}

func (m *MongoConf) NewMongoCliInstance() (mongoInstance *mongo.Client, err error) {
	return m.setUp()
}

func (m *MongoConf) setUp() (*mongo.Client, error) {
	ctx := context.TODO()
	clientOptions := options.Client().SetHosts(m.Hosts).
		SetMaxPoolSize(m.MaxPoolSize).            //最大连接数量
		SetServerSelectionTimeout(ConnectTimeout) //连接超时10s

	if m.UserName != "" && m.Password != "" {
		clientOptions.SetAuth(options.Credential{Username: m.UserName, Password: m.Password})
	}

	if m.DbName != "" {
		clientOptions = clientOptions.SetAppName(m.DbName)
	}

	if m.ReplicaSet != "" {
		clientOptions = clientOptions.SetReplicaSet(m.ReplicaSet)
	}

	if m.ReadPreference != "" {
		model, err := readpref.ModeFromString(m.ReadPreference)
		if err != nil {
			return nil, errors.WithMessage(err, fmt.Sprintf("MongoDB set [%s] is error", m.ReadPreference))
		}
		pref, err := readpref.New(model)
		if err != nil {
			return nil, errors.WithMessage(err, "MongoDB set readPreference is error")
		}
		clientOptions = clientOptions.SetReadPreference(pref)
	}

	mongoInstance, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.WithMessage(err, "MongoDB connect fail")
	}

	err = mongoInstance.Ping(ctx, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "MongoDB ping fail")
	}
	return mongoInstance, nil
}
