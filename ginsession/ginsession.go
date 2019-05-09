package ginsession

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/yitume/caller/common"
)

var defaultCaller *callerStore

type callerStore struct {
	caller gin.HandlerFunc
	cfg    Cfg
}

func New() common.Caller {
	defaultCaller = &callerStore{}
	return defaultCaller
}

func Caller() gin.HandlerFunc {
	return defaultCaller.caller
}

func (c *callerStore) InitCfg(cfg []byte) error {
	if err := parseConfig(cfg, &c.cfg); err != nil {
		return err
	}
	c.initCaller()
	return nil
}

func (c *callerStore) Get(key string) interface{} {
	return c.caller
}

func (c *callerStore) Set(key string, val interface{}) {
	c.caller = val.(gin.HandlerFunc)
}

func (c *callerStore) initCaller() {
	caller, err := provider(c.cfg.CallerGinSession)
	if err != nil {
		panic(err.Error())
	}
	c.Set("", caller)
}

func parseConfig(cfg []byte, value interface{}) error {
	var err error
	if err = toml.Unmarshal(cfg, value); err != nil {
		return err
	}
	return nil
}

func provider(cfg CallerCfg) (session gin.HandlerFunc, err error) {
	var store redis.Store
	store, err = redis.NewStore(cfg.Size, cfg.Network, cfg.Addr, cfg.Pwd, []byte(cfg.Keypairs))
	session = sessions.Sessions(cfg.Name, store)
	return
}
