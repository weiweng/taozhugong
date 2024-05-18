package conf

import (
	"log"
	"time"

	"github.com/spf13/viper"
	"github.com/weiweng/taozhugong/helper/env"
	"github.com/weiweng/taozhugong/helper/path"
)

var ConfigHandler Config

type Config struct {
	v *viper.Viper
	ServiceWeb
	Log
	Mysql
	Redis
}

type ServiceWeb struct {
	Level string
	Addr  string
}

type Log struct {
	Level string
	Path  string
}

type Mysql struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	Port            string
	User            string
	Host            string
	Password        string
	DbName          string
}

type Redis struct {
	MaxIdle        int
	MaxActive      int
	Addr           string
	IdleTimeout    time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	ConnectTimeout time.Duration
}

func (c *Config) ReadConfig() {
	c.ServiceWeb.Addr = c.v.GetString("service.addr")
	c.ServiceWeb.Level = c.v.GetString("service.level")

	c.Log.Path = c.v.GetString("log.path")
	c.Log.Level = c.v.GetString("log.level")

	c.Mysql.Port = c.v.GetString("mysql.port")
	c.Mysql.Host = c.v.GetString("mysql.host")
	c.Mysql.User = c.v.GetString("mysql.user")
	c.Mysql.Password = c.v.GetString("mysql.password")
	c.Mysql.DbName = c.v.GetString("mysql.db")
	c.Mysql.MaxIdleConns = c.v.GetInt("mysql.maxIdleConns")
	c.Mysql.MaxOpenConns = c.v.GetInt("mysql.maxOpenConns")
	c.Mysql.ConnMaxLifetime = c.v.GetDuration("mysql.connMaxLifetime")

	c.Redis.Addr = c.v.GetString("redis.addr")
	c.Redis.MaxActive = c.v.GetInt("redis.maxActive")
	c.Redis.MaxIdle = c.v.GetInt("redis.maxIdle")
	c.Redis.IdleTimeout = c.v.GetDuration("redis.idleTimeout")
	c.Redis.ReadTimeout = c.v.GetDuration("redis.readTimeout")
	c.Redis.WriteTimeout = c.v.GetDuration("redis.writeTimeout")
	c.Redis.ConnectTimeout = c.v.GetDuration("redis.connectTimeout")
}

func init() {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("toml")
	v.AddConfigPath(path.Root + "/conf/" + env.Env)
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	ConfigHandler = Config{v: v}
	ConfigHandler.ReadConfig()
}
