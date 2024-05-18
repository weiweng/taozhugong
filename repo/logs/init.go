package logs

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/weiweng/taozhugong/conf"
)

var output *rotatelogs.RotateLogs

func Init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	var err error
	// Read more at: https://github.com/lestrrat-go/file-rotatelogs#synopsis
	pathToNormalLog := conf.ConfigHandler.Path + "/log.%Y%m%d%H"
	output, err = rotatelogs.New(
		pathToNormalLog,
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour))
	if err != nil {
		log.Fatalf("log init failed: %v", err)
	}
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	// log.SetOutput(os.Stdout)
	logrus.SetOutput(output)

	ll := logrus.DebugLevel
	switch conf.ConfigHandler.Log.Level {
	case "debug":
		ll = logrus.DebugLevel
	case "info":
		ll = logrus.InfoLevel
	case "warn":
		ll = logrus.WarnLevel
	case "error":
		ll = logrus.ErrorLevel
	case "fatal":
		ll = logrus.FatalLevel
	case "panic":
		ll = logrus.PanicLevel
	}
	// Only log the warning severity or above.
	logrus.SetLevel(ll)
}
