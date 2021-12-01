package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const uri = "mongodb://mongoadmin:secret@localhost"

func main() {
	//主上下文context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	//连接远程数据库
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	//创建集合(用于查询数据和插入文档)
	coll := client.Database("nftdisk").Collection("movies")
	res, err := coll.InsertOne(ctx, bson.M{"hello": "world"})
	if err != nil {
		panic(err)
	}
	id := res.InsertedID
	fmt.Printf("id:%s \n", id.(primitive.ObjectID))

	//查询集合
	find, err := coll.Find(ctx, bson.D{})
	if err != nil{
		panic(err)
	}
	var results []struct{
		Hello string
	}
	if err = find.All(ctx, &results); err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(results)
	fmt.Printf("%s \n",string(marshal))
	//查询单条(转换实体对象)
	var result struct{
		Hello string
	}
	filter := bson.D{{"hello", "world"}}
	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	marshal, _ = json.Marshal(result)
	fmt.Printf("%s \n",string(marshal))

}