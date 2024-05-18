package path

import (
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/weiweng/taozhugong/helper/constant"
	"github.com/weiweng/taozhugong/helper/env"
)

const appDir = "taozhugong"

var Root string

func init() {
	switch env.Env {
	case constant.EnvDevelopment:
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal("Initialize Root error")
		}
		// 读取
		index := strings.LastIndex(filename, appDir)
		if index >= 0 && index <= len(filename) {
			runes := []rune(filename)
			Root = string(runes[0:index]) + appDir
			log.Println("找到本地应用根目录 => ", Root)
		} else {
			log.Println("未找到根文件夹 => ", filename)
		}

	case constant.EnvProduction:
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Initialize Root error: %v", err)
		}
		Root = wd
	}
}
