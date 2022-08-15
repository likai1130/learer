package basic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataSource interface {
	Insert(ctx context.Context, collectName string, entityType interface{}) (string, error)
	InsertBatch(ctx context.Context, collectName string, typeSlices []interface{}) (int, error)
	Get(ctx context.Context, collectName string, filter bson.D) (*mongo.SingleResult, error)
	List(ctx context.Context, collectName string, filter bson.D, options ...*options.FindOptions) ([]map[string]interface{}, error)
	Update(ctx context.Context, collectName string, filter bson.D, entityType interface{}) (int64, error)
	UpdateBatch(ctx context.Context, collectName string, filter bson.D, typeSlices interface{}) (int64, error)
	Delete(ctx context.Context, collectName string, filter bson.D) (int64, error)
	DeleteBatch(ctx context.Context, collectName string, filter bson.D) (int64, error)
	Aggregate(ctx context.Context, collectName string, pipeline interface{},options ...*options.AggregateOptions) ([]map[string]interface{}, error)
	Count(ctx context.Context, collectName string, filter interface{}) (int64, error)
	FindOneAndUpdate(ctx context.Context, collectName string, filter, update interface{}) (*mongo.SingleResult, error)
}
