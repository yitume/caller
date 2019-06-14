package zap

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/yitume/caller/pkg/common"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultCaller *callerStore

type callerStore struct {
	caller map[string]*Client
	cfg    Cfg
}

type Client struct {
	*zap.Logger
}

func New() common.Caller {
	defaultCaller = &callerStore{
		caller: make(map[string]*Client, 0),
	}
	return defaultCaller
}

func Caller(name string) *Client {
	return defaultCaller.caller[name]
}

func (c *callerStore) InitCfg(cfg []byte) error {
	if err := parseConfig(cfg, &c.cfg); err != nil {
		return err
	}
	c.initCaller()
	return nil
}

func (c *callerStore) Get(key string) interface{} {
	return c.caller[key]
}

func (c *callerStore) Set(key string, val interface{}) {
	c.caller[key] = val.(*Client)
}

func (c *callerStore) initCaller() {
	for name, cfg := range c.cfg.CallerZap {
		db := provider(cfg)
		c.Set(name, db)
	}
}

func parseConfig(cfg []byte, value interface{}) error {
	var err error
	if err = toml.Unmarshal(cfg, value); err != nil {
		return err
	}
	return nil
}

func provider(cfg CallerCfg) (db *Client) {
	var js string
	if cfg.Debug {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stdout"]
      }`, cfg.Level)
	} else {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["%s"],
      "errorOutputPaths": ["%s"]
      }`, cfg.Level, cfg.Path, cfg.Path)
	}

	var zcfg zap.Config
	if err := json.Unmarshal([]byte(js), &zcfg); err != nil {
		panic(err)
	}
	zcfg.EncoderConfig = zap.NewProductionEncoderConfig()
	zcfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	l, err := zcfg.Build()
	if err != nil {
		log.Fatal("init logger error: ", err)
	}
	return &Client{l}
}
