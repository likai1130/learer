package basic

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoImp struct {
	dbName      string
	mongoClient *mongo.Client
}

func (m *MongoImp) Update(ctx context.Context, collectName string, filter bson.D, entityType interface{}) Result {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	updateResult, err := collection.UpdateOne(ctx, filter, entityType)
	if err != nil {
		return Result{err: errors.Wrap(err, "update db data is error")}
	}
	if updateResult.ModifiedCount == 0 {
		return Result{err: errors.Wrap(errors.New("result.ModifiedCount is 0"), "update db data is error")}
	}
	return Result{}
}

func (m *MongoImp) UpdateBatch(ctx context.Context, collectName string, filter bson.D, typeSlices interface{}) Result {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	results, err := collection.UpdateMany(ctx, filter, typeSlices)
	if err != nil {
		return Result{err: errors.Wrap(err, "update db many data is error ")}
	}
	return Result{count: results.ModifiedCount}
}

func (m *MongoImp) DeleteBatch(ctx context.Context, collectName string, filter bson.D) Result {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return Result{err: errors.Wrap(err, "delete db data is error ")}
	}
	return Result{count: result.DeletedCount}
}

func (m *MongoImp) Delete(ctx context.Context, collectName string, filter bson.D) Result {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return Result{err: errors.Wrap(err, "delete db data is error ")}
	}
	return Result{count: result.DeletedCount}
}

func (m *MongoImp) Insert(ctx context.Context, collectName string, entityType interface{}) Result {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	insertOneResult, err := collection.InsertOne(ctx, entityType)
	if err != nil {
		return Result{err: errors.Wrap(err, "insert db data is error ")}
	}
	insertId := insertOneResult.InsertedID.(primitive.ObjectID).Hex()
	return Result{id: insertId}
}

func (m *MongoImp) InsertBatch(ctx context.Context, collectName string, typeSlices []interface{}) Result {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	insertManyResult, err := collection.InsertMany(ctx, typeSlices)
	if err != nil {
		return Result{err: errors.Wrap(err, "insert db many data is error ")}
	}
	return Result{count: int64(len(insertManyResult.InsertedIDs))}
}

func (m *MongoImp) List(ctx context.Context, collectName string, filter bson.D, options ...*options.FindOptions) Result {
	if filter == nil {
		filter = bson.D{}
	}
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	cursors, err := collection.Find(ctx, filter, options...)
	if err != nil || cursors.Err() != nil {
		return Result{err: errors.Wrap(err, "find db  data is error")}
	}
	defer func(cursors *mongo.Cursor, ctx context.Context) {
		err := cursors.Close(ctx)
		if err != nil {
			log.Printf("List context error: %+v", errors.Wrap(err, "mongoImp list context close error"))
		}
	}(cursors, ctx)
	var typesSlice []bson.M
	err = cursors.All(ctx, &typesSlice)
	if err != nil {
		return Result{err: err}
	}
	marshal, err := bson.Marshal(typesSlice)
	if err != nil {
		return Result{err: err}
	}
	return Result{data: marshal}
}

func (m *MongoImp) Get(ctx context.Context, collectName string, filter bson.D) Result {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return Result{err: errors.Wrap(result.Err(), "find db data is error")}
	}
	bytes, err := result.DecodeBytes()
	if err != nil {
		return Result{err: err}
	}
	return Result{data: bytes}
}

func (m *MongoImp) FindOneAndUpdate(ctx context.Context, collectName string, filter, update interface{}) Result {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	result := collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return Result{err: errors.Wrap(result.Err(), "insertInc db data is error")}
	}
	bytes, err := result.DecodeBytes()
	if err != nil {
		return Result{err: err}
	}
	return Result{data: bytes}
}

func (m *MongoImp) Count(ctx context.Context, collectName string, filter interface{}) Result {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	result, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return Result{err: errors.Wrap(err, "count is error")}
	}
	return Result{count: result}
}

// Aggregate 聚合查询/**
func (m *MongoImp) Aggregate(ctx context.Context, collectName string, filter interface{}, opts ...*options.AggregateOptions) Result {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	cursors, err := collection.Aggregate(ctx, filter, opts...)
	if err != nil || cursors.Err() != nil {
		return Result{err: errors.Wrap(err, "aggregate db is error")}
	}
	defer func(cursors *mongo.Cursor, ctx context.Context) {
		err := cursors.Close(ctx)
		if err != nil {
			log.Printf("%+v", errors.Wrap(err, "complexImp list context close error"))
		}
	}(cursors, ctx)
	var typesSlice = []bson.M{}
	err = cursors.All(ctx, &typesSlice)
	if err != nil {
		return Result{err: err}
	}
	marshal, err := bson.Marshal(typesSlice)
	if err != nil {
		return Result{err: err}
	}
	return Result{data: marshal}
}

func New(conf *MongoConf) (*MongoImp, error) {
	instance, err := conf.NewMongoCliInstance()
	if err != nil {
		return nil, err
	}

	return &MongoImp{
		dbName:      conf.DbName,
		mongoClient: instance,
	}, err
}
