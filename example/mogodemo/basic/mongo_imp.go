package basic

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoImp struct {
	dbName      string
	mongoClient *mongo.Client
}

func (m *mongoImp) Update(ctx context.Context, collectName string, filter bson.D, entityType interface{}) (int64, error) {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	result, err := collection.UpdateOne(ctx, filter, entityType)
	if err != nil {
		return 0, errors.Wrap(err, "update db data is error")
	}
	if result.ModifiedCount == 0 {
		return 0, errors.Wrap(errors.New("result.ModifiedCount is 0"), "update db data is error")
	}
	return result.ModifiedCount, nil
}

func (m *mongoImp) UpdateBatch(ctx context.Context, collectName string, filter bson.D, typeSlices interface{}) (int64, error) {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	results, err := collection.UpdateMany(ctx, filter, typeSlices)
	if err != nil {
		return 0, errors.Wrap(err, "update db many data is error ")
	}
	return results.ModifiedCount, nil
}

func (m *mongoImp) DeleteBatch(ctx context.Context, collectName string, filter bson.D) (int64, error) {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return 0, errors.Wrap(err, "delete db data is error ")
	}
	return result.DeletedCount, nil
}

func (m *mongoImp) Delete(ctx context.Context, collectName string, filter bson.D) (int64, error) {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, errors.Wrap(err, "delete db data is error ")
	}
	return result.DeletedCount, nil
}

func (m *mongoImp) Insert(ctx context.Context, collectName string, entityType interface{}) (string, error) {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	insertOneResult, err := collection.InsertOne(ctx, entityType)
	if err != nil {
		return "", errors.Wrap(err, "insert db  data is error")
	}
	insertId := insertOneResult.InsertedID.(primitive.ObjectID).Hex()
	return insertId, nil
}

func (m *mongoImp) InsertBatch(ctx context.Context, collectName string, typeSlices []interface{}) (int, error) {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	insertManyResult, err := collection.InsertMany(ctx, typeSlices)
	if err != nil {
		return 0, errors.Wrap(err, "insert db many data is error ")
	}
	return len(insertManyResult.InsertedIDs), nil
}

func (m *mongoImp) List(ctx context.Context, collectName string, filter bson.D, options ...*options.FindOptions) ([]map[string]interface{}, error) {
	if filter == nil {
		filter = bson.D{}
	}
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	cursors, err := collection.Find(ctx, filter, options...)
	if err != nil || cursors.Err() != nil {
		return nil, errors.Wrap(err, "find db  data is error")
	}
	defer func(cursors *mongo.Cursor, ctx context.Context) {
		err := cursors.Close(ctx)
		if err != nil {
			logger.GetLogger().Errorf("List context error: %+v", errors.Wrap(err, "mongoImp list context close error"))
		}
	}(cursors, ctx)
	typesSlice := make([]map[string]interface{}, 0)
	err = cursors.All(ctx, &typesSlice)
	return typesSlice, err
}

func (m *mongoImp) Get(ctx context.Context, collectName string, filter bson.D) (*mongo.SingleResult, error) {
	if filter == nil {
		filter = bson.D{}
	}
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return nil, errors.Wrap(result.Err(), "find db data is error")
	}
	return result, nil
}

func (m *mongoImp) FindOneAndUpdate(ctx context.Context, collectName string, filter, update interface{}) (*mongo.SingleResult, error) {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	result := collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return nil, errors.Wrap(result.Err(), "insertInc db data is error")
	}
	return result, nil
}

func (m *mongoImp) Count(ctx context.Context, collectName string, filter interface{}) (int64, error) {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	result, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, errors.Wrap(err, "count is error")
	}
	return result, nil
}

// Aggregate 聚合查询/**
func (m *mongoImp) Aggregate(ctx context.Context, collectName string, filter interface{}, opts ...*options.AggregateOptions) ([]map[string]interface{}, error) {
	collection := m.mongoClient.Database(m.dbName).Collection(collectName)
	cursors, err := collection.Aggregate(ctx, filter, opts...)
	if err != nil || cursors.Err() != nil {
		return nil, errors.Wrap(err, "aggregate db is error")
	}
	defer func(cursors *mongo.Cursor, ctx context.Context) {
		err := cursors.Close(ctx)
		if err != nil {
			logger.GetLogger().Errorf("%+v", errors.Wrap(err, "complexImp list context close error"))
		}
	}(cursors, ctx)
	typesSlice := make([]map[string]interface{}, 0)
	err = cursors.All(ctx, &typesSlice)
	return typesSlice, err
}

func newDefaultMongoInstance(chainName string) *mongoImp {
	instance, _ := mongodb.GetMongoFactory().GetMongoInstance(chainName)
	return &mongoImp{
		dbName:      instance.DbName,
		mongoClient: instance.MongoCli,
	}
}
