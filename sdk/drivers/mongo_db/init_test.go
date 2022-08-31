package mongo_db

import (
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/test/bogus"
)

const (
	ConfigPath = "../../config/config.yaml"
)

var (
	testEnv            ioc.IDig
	testMongoCfg       config.IMongoDbConfig
	testDbName         string
	testCollectionName string
)

func init() {
	testEnv = buildTestEnv()
}

func buildTestEnv() ioc.IDig {
	testDbName = "test-root"
	testCollectionName = "test-cars"
	testEnv = builder.InjectCoLoMed(ConfigPath)
	testEnv.Invoke(func(appConfig config.IAppConfig) {
		testMongoCfg = appConfig.GetMongoDbConfig()
	})
	testEnv.Inject(testEnv,
		SingletonMongoClient,
	).Inject(testEnv,
		BogusRootMongoStore,
	)
	return testEnv
}

func BogusRootMongoStore(newClient MongoDbClientFtor) MongoDbStoreFtor[bogus.Root] {
	return NewMongoStore[bogus.Root](newClient, testDbName, testCollectionName)
}
