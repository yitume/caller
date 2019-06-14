package echosession

import (
	"github.com/BurntSushi/toml"
	"github.com/ipfans/echo-session"
	"github.com/labstack/echo"

	"github.com/yitume/caller/pkg/common"
)

var defaultCaller *callerStore

type callerStore struct {
	caller echo.HandlerFunc
	cfg    Cfg
}

func New() common.Caller {
	defaultCaller = &callerStore{}
	return defaultCaller
}

func Caller() echo.HandlerFunc {
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
	c.caller = val.(echo.HandlerFunc)
}

func (c *callerStore) initCaller() {
	caller, err := provider(c.cfg.CallerEchoSession)
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

func provider(cfg CallerCfg) (s echo.MiddlewareFunc, err error) {
	store, err := session.NewRedisStore(cfg.Size, cfg.Network, cfg.Addr, cfg.Pwd, []byte(cfg.Keypairs))
	s = session.Sessions(cfg.Name, store)
	return
}
