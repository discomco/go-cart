package builder

import "github.com/discomco/go-cart/core/ioc"

const (
	CFG_PATH = "../../config/config.yaml"
)

var (
	testEnv ioc.IDig
)

func buildTestEnv() ioc.IDig {
	return InjectCoLoMed(CFG_PATH)
}

func init() {
	testEnv = buildTestEnv()
}
