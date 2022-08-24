package mongo_db

import "github.com/discomco/go-cart/model"

type doc[T model.IReadModel] struct {
	_id  string `bson:"_id"`
	data T      `bson:"data"`
}

func newDoc[T model.IReadModel](id string, data T) *doc[T] {
	return &doc[T]{
		_id:  id,
		data: data,
	}
}
