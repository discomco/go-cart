package mongo_db

import (
	"fmt"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type MongoDbStoreFtor[TModel model.IReadModel] func() IStore[TModel]

type IStore[T model.IReadModel] interface {
	domain.IReadModelStore[T]
}

type store[T model.IReadModel] struct {
	mongo   *mongo.Client
	dbName  string
	colName string
}

func (s *store[T]) Exists(ctx context.Context, key string) (bool, error) {
	col := s.mongo.Database(s.dbName).Collection(s.colName)
	res := col.FindOne(ctx, bson.M{"_id": key})
	err := res.Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *store[T]) Get(ctx context.Context, key string) (*T, error) {
	col := s.mongo.Database(s.dbName).Collection(s.colName)
	res := col.FindOne(ctx, bson.M{"_id": key})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var err error
	var doc2 doc[T]
	err = res.Decode(&doc2)

	var raw bson.Raw
	if raw, err = res.DecodeBytes(); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	var doc doc[T]
	err = bson.Unmarshal(raw, &doc)
	if err != nil {
		return nil, err
	}
	return &doc.data, nil
}

func (s *store[T]) Set(ctx context.Context, key string, data T) (string, error) {
	col := s.mongo.Database(s.dbName).Collection(s.colName)
	//	doc := newDoc[T](key, data)
	res, err := col.InsertOne(ctx, bson.D{{"_id", key}, {"data", data}})
	//		bson.M{"_id": doc._id, "data": doc.data})
	if err != nil {
		return "", err
	}
	ret := fmt.Sprintf("%v", res.InsertedID)
	return ret, nil
}

func (s *store[T]) Delete(ctx context.Context, key string) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func newStore[T model.IReadModel](mongo *mongo.Client, dbName string, colName string) *store[T] {
	return &store[T]{
		mongo:   mongo,
		dbName:  dbName,
		colName: colName,
	}
}

func NewMongoStore[TModel model.IReadModel](newMongoDb MongoDbClientFtor, dbName string, colName string) MongoDbStoreFtor[TModel] {
	return func() IStore[TModel] {
		clt := newMongoDb()
		return newStore[TModel](clt, dbName, colName)
	}
}
