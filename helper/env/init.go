package env

import (
	"flag"
	"os"

	"github.com/weiweng/taozhugong/helper/constant"
)

// Env defines the environment string: development/testing/production/preview
var Env string

func init() {
	var defaultEnv = constant.EnvDevelopment
	if len(os.Getenv("ENV")) > 0 {
		defaultEnv = os.Getenv("ENV")
	}
	flag.StringVar(&Env, "env", defaultEnv, "set the environment")
}
