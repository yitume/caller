package redigo

import (
	"github.com/BurntSushi/toml"
	"github.com/yitume/caller/common"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
	"time"
)

var defaultCaller *callerStore

type callerStore struct {
	caller map[string]*RedigoClient
	cfg    Cfg
}

type RedigoClient struct {
	pool *redis.Pool
}

func New() common.Caller {
	defaultCaller = &callerStore{
		caller: make(map[string]*RedigoClient, 0),
	}
	return defaultCaller
}

func Caller(name string) *RedigoClient {
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
	c.caller[key] = val.(*RedigoClient)
}

func (c *callerStore) initCaller() {
	for name, cfg := range c.cfg.CallerRedigo {
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

func provider(cfg CallerCfg) (resp *RedigoClient) {
	dialOptions := []redis.DialOption{
		redis.DialConnectTimeout(cfg.ConnectTimeout.Duration),
		redis.DialReadTimeout(cfg.ReadTimeout.Duration),
		redis.DialWriteTimeout(cfg.WriteTimeout.Duration),
		redis.DialDatabase(cfg.DB),
		redis.DialPassword(cfg.Password),
	}

	resp = &RedigoClient{
		&redis.Pool{
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", cfg.Addr, dialOptions...)
				if err != nil {
					return nil, err
				}

				if cfg.Debug {
					return redis.NewLoggingConn(c, log.New(os.Stderr, "", log.LstdFlags), "redis"), nil
				}
				return c, nil
			},
			// Use the TestOnBorrow function to check the health of an idle connection
			// before the connection is returned to the application. This example PINGs
			// connections that have been idle more than a minute:
			//
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
			MaxIdle:     cfg.MaxActive,
			MaxActive:   cfg.MaxActive,
			IdleTimeout: cfg.IdleTimeout.Duration,
			Wait:        cfg.Wait, // wait until getting connection from pool
		},
	}
	return
}
