package moclient

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"learner/pkg/mongo/constants"
	"learner/pkg/mongo/moconfig"
	"log"
	"os"
	"sync"
	"time"
)

var MCli *MongoClient
var once sync.Once

type MongoClient struct {
	MongoCli *mongo.Client
}

func NewMongoCliInstance() *MongoClient {
	once.Do(func() {
		MCli = &MongoClient{
			MongoCli: setUp(),
		}
	})
	return MCli
}

func setUp() *mongo.Client{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mConf := moconfig.AppConfig.MongoConf
	clientOptions := options.Client().SetHosts(mConf.Hosts).
		SetMaxPoolSize(mConf.MaxPoolSize).                                 //最大连接数量
		SetConnectTimeout(constants.MONGODB_CONNECT_TIMEOUT) //连接超时20s
	if mConf.UserName != "" && mConf.Password != "" {
		clientOptions.SetAuth(options.Credential{Username: mConf.UserName, Password: mConf.Password})
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connect fail：%s",err.Error())
		os.Exit(1)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping fail：%s",err.Error())
		os.Exit(1)
	}
	log.Println("MongoDB connect success !")
	return client
}

func Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	MCli.MongoCli.Disconnect(ctx)
}