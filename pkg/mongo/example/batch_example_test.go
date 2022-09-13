package example

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"learner/pkg/mongo/constants"
	"learner/pkg/mongo/moclient"
	"learner/pkg/mongo/moconfig"
	"learner/pkg/mongo/model"
	"log"
	"sync"
	"testing"
)

func TestBatchInsert(t *testing.T) {
	moconfig.InitConfig("/Users/likai/hisun/gospace/src/learner/pkg/mongo/application.yaml")
	wg := sync.WaitGroup{}
	for i:=0;i<1000;i++ {
		wg.Add(1)
		go insert(&wg)
	}
	wg.Wait()
	log.Printf("insert ent. count is %d",1000)
}

func TestGet(t *testing.T)  {
	for i:=0;i<1000;i++ {
		get("kli")
	}
	log.Printf("insert ent. count is %d",1000)
}
func insert(wg *sync.WaitGroup) {
	defer wg.Done()
	cli := moclient.NewMongoCliInstance().MongoCli
	coll := cli.Database(constants.MONGODB_DATABASE).Collection(constants.MONGODB_DATABASE_COLLECT)
	ctx, cancel := context.WithTimeout(context.Background(), constants.MONGODB_CONNECT_TIMEOUT)
	defer cancel()

	userModel := model.UserData{
		Name:       "kli",
		Number:     152664,
		Age:        18,
		BirthMonth: 1852,
	}
	insertOneResult, err := coll.InsertOne(ctx, userModel)
	if err != nil {
		log.Fatalf("Insert user data is error : %s",err.Error())
		return
	}
//	log.Printf("Insert user data is success! id = %v\n",insertOneResult.InsertedID)
	docId := insertOneResult.InsertedID.(primitive.ObjectID)
	log.Printf("insert one ID str is :%s", docId.String())
}

func get(name string) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.MONGODB_CONNECT_TIMEOUT)
	defer cancel()

	cli := moclient.NewMongoCliInstance().MongoCli
	coll := cli.Database(constants.MONGODB_DATABASE).Collection(constants.MONGODB_DATABASE_COLLECT)
	//过滤条件
	filter := bson.M{"name": name}
	singleResult := coll.FindOne(ctx, filter)
	if singleResult == nil || singleResult.Err() != nil {
		log.Fatalf("Find one user is error")
		if singleResult.Err() != nil {
			log.Fatalf(singleResult.Err().Error())
		}
		return
	}
	userData := &model.UserData{}
	err := singleResult.Decode(userData)
	if err != nil {
		log.Fatalf("Find one user is error:%s\n",err.Error())
		return
	}
	marshal, _ := json.Marshal(userData)
	log.Printf("Find one user is success!\n %s",string(marshal))
}

func TestSub(t *testing.T)  {
	moconfig.InitConfig("/Users/likai/hisun/gospace/src/learner/pkg/mongo/application.yaml")
	cli := moclient.NewMongoCliInstance().MongoCli
	collection := cli.Database(constants.MONGODB_DATABASE).Collection(constants.MONGODB_DATABASE_COLLECT)
	changeStream, err := collection.Watch(context.Background(), mongo.Pipeline{},options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		panic(err)
	}

	for changeStream.Next(context.TODO()) {
		fmt.Println(changeStream.Current)
	}
}