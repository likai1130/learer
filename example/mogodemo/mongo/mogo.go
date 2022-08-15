package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataSource interface {
	Insert(ctx context.Context, collectName string, obj interface{}) (string, error)
	InsertBatch(ctx context.Context, collectName string, objSlices []interface{}) (int, error)

	Get(ctx context.Context, collectName string, filter bson.D, options ...*options.FindOneOptions) (*mongo.SingleResult, error)
	List(ctx context.Context, collectName string, filter bson.D, options ...*options.FindOptions) ([]bson.M, error)

	Update(ctx context.Context, collectName string, filter bson.D, obj interface{}) (int64, error)
	UpdateBatch(ctx context.Context, collectName string, filter bson.D, objSlices interface{}) (int64, error)
	FindOneAndUpdate(ctx context.Context, collectName string, filter, update interface{}) (*mongo.SingleResult, error)

	Delete(ctx context.Context, collectName string, filter bson.D) (int64, error)
	DeleteBatch(ctx context.Context, collectName string, filter bson.D) (int64, error)

	Aggregate(ctx context.Context, collectName string, pipeline interface{}) ([]bson.M, error)
	Count(ctx context.Context, collectName string, filter interface{}) (int64, error)
}
