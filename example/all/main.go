package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yitume/caller"
	"github.com/yitume/caller/ginsession"
	"github.com/yitume/caller/gorm"
	"github.com/yitume/caller/zap"
)

var cfg = `
[callerGinSession]
    name = "mysession"
    size = 10
    network = "tcp"
    addr = "127.0.0.1:6379"
    pwd = ""
    keypairs = "secret"
[callerGorm.default]
    debug = true
    level = "panic"
    network = "tcp"
    dialect = "mysql"
    addr = "127.0.0.1:3306"
    username = "root"
    password = ""
    db = "default"
    charset = "utf8"
    parseTime = "True"
    loc = "Local"
    timeout = "1s"
    readTimeout = "1s"
    writeTimeout = "1s"
    maxOpenConns = 30
    maxIdleConns = 10
    connMaxLifetime = "300s"
[callerZap.default]
    debug = true
    level = "debug"
    path = "./system.log"
`
var (
	Db      *gorm.Client
	Logger  *zap.Client
	Session gin.HandlerFunc
)

func main() {
	if err := caller.Init(
		cfg,
		zap.New,
		gorm.New,
		ginsession.New,
	); err != nil {
		panic(err)
	}

	initModel()
	type User struct {
		Uid  int
		Name string
	}
	u := User{}
	Db.Table("user").Where("uid=?", 1).Find(&u)
	Logger.Info("hello world")
	r := gin.New()
	r.Use(ginsession.Caller())
}

func initModel() {
	Db = gorm.Caller("default")
	Logger = zap.Caller("default")
	Session = ginsession.Caller()
}
