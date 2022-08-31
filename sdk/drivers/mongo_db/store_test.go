package mongo_db

import (
	"context"
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/test/bogus"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

func TestThatWeCanCreateAReadStore(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var clt *mongo.Client
	testEnv.Invoke(func(newMongo MongoDbClientFtor) {
		clt = newMongo()
	})
	assert.NotNil(t, clt)
	store := newStore[bogus.Root](clt, testDbName, testCollectionName)
	// THEN
	assert.NotNil(t, store)
}

func TestThatWeCanInjectASpecializedMongoReadStore(t *testing.T) {
	// GIVEN
	localEnv := builder.InjectCoLoMed(ConfigPath)
	assert.NotNil(t, localEnv)
	// AND
	localEnv.Inject(localEnv,
		SingletonMongoClient,
	).Inject(testEnv,
		BogusRootMongoStore,
	)
	// WHEN
	var store IStore[bogus.Root]
	localEnv.Invoke(func(newStore MongoDbStoreFtor[bogus.Root]) {
		store = newStore()
	})
	// THEN
	assert.NotNil(t, store)
}

func TestThatWeCanCheckIfAnItemExistsInTheSpecializedMongoReadStore(t *testing.T) {
	//GIVEN
	assert.NotNil(t, testEnv)
	// AND
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	assert.NotNil(t, ctx)
	// AND
	ID, err := bogus.NewRootIdentity()
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	// AND
	root := bogus.NewRoot(ID)
	assert.NotNil(t, root)
	// AND
	var store IStore[bogus.Root]
	testEnv.Invoke(func(newStore MongoDbStoreFtor[bogus.Root]) {
		store = newStore()
	})
	assert.NotNil(t, store)
	// WHEN
	exists, err := store.Exists(ctx, root.ID.Id())
	// THEN
	assert.Error(t, err)
	assert.False(t, exists)
}

func TestThatWeCanInsertADocumentIntoTheSpecializedMongoDbStore(t *testing.T) {
	// AND
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	assert.NotNil(t, ctx)
	// AND
	ID, err := bogus.NewRootIdentity()
	assert.NotNil(t, ID)
	// AND
	root := bogus.NewRoot(ID)
	assert.NoError(t, err)
	root.Car = bogus.NewCar("Ford", "Desire", 25, 48.5)
	assert.NotNil(t, root)
	// AND
	var store IStore[bogus.Root]
	testEnv.Invoke(func(newStore MongoDbStoreFtor[bogus.Root]) {
		store = newStore()
	})
	assert.NotNil(t, store)
	// WHEN
	id := root.ID.Id()
	res, err := store.Set(ctx, id, *root)
	// THEN
	assert.NoError(t, err)
	assert.Equal(t, id, res)
}

//func TestThatWeCanGetARootFromTheSpecializedMongoDbStore(t *testing.T) {
//	// AND
//	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
//	defer cancel()
//	assert.NotNil(t, ctx)
//	// AND
//	Id := bogus.NewRootIdentity()
//	assert.NotNil(t, Id)
//	// AND
//	root := bogus.NewRoot(Id)
//	root.Car = bogus.NewCar("Toyota", "Zulma", 72, 1500.20)
//	assert.NotNil(t, root)
//	// AND
//	var store IStore[bogus.Root]
//	testEnv.Invoke(func(newStore MongoDbStoreFtor[bogus.Root]) {
//		store = newStore()
//	})
//	assert.NotNil(t, store)
//	// AND
//	res, err := store.Set(ctx, root.Id.Id(), *root)
//	// AND
//	assert.NoError(t, err)
//	assert.Equal(t, root.Id.Id(), res)
//
//	id := root.Id.Id()
//	// WHEN
//	res1, err := store.Get(ctx, id)
//	assert.NoError(t, err)
//	assert.NotNil(t, res1)
//
//	assert.Equal(t, root, res1)
//
//}
