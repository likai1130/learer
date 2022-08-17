package basic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataSource interface {
	Insert(ctx context.Context, collectName string, entityType interface{}) Result
	InsertBatch(ctx context.Context, collectName string, typeSlices []interface{}) Result
	Get(ctx context.Context, collectName string, filter bson.D) Result
	List(ctx context.Context, collectName string, filter bson.D, options ...*options.FindOptions) Result
	Update(ctx context.Context, collectName string, filter bson.D, entityType interface{}) Result
	UpdateBatch(ctx context.Context, collectName string, filter bson.D, typeSlices interface{}) Result
	Delete(ctx context.Context, collectName string, filter bson.D) Result
	DeleteBatch(ctx context.Context, collectName string, filter bson.D) Result
	Aggregate(ctx context.Context, collectName string, pipeline interface{}, options ...*options.AggregateOptions) Result
	Count(ctx context.Context, collectName string, filter interface{}) Result
	FindOneAndUpdate(ctx context.Context, collectName string, filter, update interface{}) Result
}

