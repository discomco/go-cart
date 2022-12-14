package mongo_db

import "github.com/discomco/go-cart/sdk/schema"

type doc[T schema.ISchema] struct {
	_id  string `bson:"_id"`
	data T      `bson:"data"`
}

func newDoc[T schema.ISchema](id string, data T) *doc[T] {
	return &doc[T]{
		_id:  id,
		data: data,
	}
}
