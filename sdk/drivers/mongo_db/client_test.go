package mongo_db

import (
	"context"
	"github.com/discomco/go-cart/config"
	"github.com/discomco/go-cart/core/builder"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

func TestThatWeCanRetrieveMongoDbConfig(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	//WHEN
	var mongoCfg config.IMongoDbConfig
	testEnv.Invoke(func(appConfig config.IAppConfig) {
		mongoCfg = appConfig.GetMongoDbConfig()
	})
	//THEN
	assert.NotNil(t, mongoCfg)
	assert.Equal(t, "mongodb://localhost:27017", mongoCfg.GetUri())
}

func TestThatWeCanCreateAMongoDbClient(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testMongoCfg)
	// AND
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	assert.NotNil(t, ctx)
	// WHEN
	client, err := newMongoDb(testMongoCfg)
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestThatWeCanInjectASingletonMongoDbFtor(t *testing.T) {
	// GIVEN
	localTestEnv := builder.InjectCoLoMed(ConfigPath)
	assert.NotNil(t, localTestEnv)
	// AND
	localTestEnv.Inject(localTestEnv,
		SingletonMongoClient,
	)
	var clt1 *mongo.Client
	var clt2 *mongo.Client
	// WHEN
	localTestEnv.Invoke(func(newMongo MongoDbClientFtor) {
		clt1 = newMongo()
		clt2 = newMongo()
	})
	// THEN
	assert.NotNil(t, clt1)
	assert.NotNil(t, clt2)
	// AND
	assert.Same(t, clt1, clt2)
}

//func TestThatWeCanInjectATransientMongoDbFtor(t *testing.T) {
//	// GIVEN
//	localTestEnv := builder.InjectCoLoMed(ConfigPath)
//	assert.NotNil(t, localTestEnv)
//	// AND
//	localTestEnv.Inject(localTestEnv,
//		TransientMongoClient,
//	)
//	var clt1 *mongo.Client
//	var clt2 *mongo.Client
//	// WHEN
//	localTestEnv.Invoke(func(newMongo MongoDbClientFtor) {
//		clt1 = newMongo()
//		clt2 = newMongo()
//	})
//	// THEN
//	assert.NotNil(t, clt1)
//	assert.NotNil(t, clt2)
//	// AND
//	assert.NotSame(t, clt1, clt2)
//}
