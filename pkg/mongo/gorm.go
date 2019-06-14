package mongo

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/globalsign/mgo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yitume/caller/pkg/common"
)

var defaultCaller *callerStore

type callerStore struct {
	caller map[string]*Client
	cfg    Cfg
}

type Client struct {
	*mgo.Database
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
	for name, cfg := range c.cfg.CallerMongo {
		db, err := provider(cfg)
		if err != nil {
			continue
		}
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

func provider(cfg CallerCfg) (resp *Client, err error) {
	session, err := mgo.Dial(cfg.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	fmt.Println("cfg.debug", cfg.Debug)
	mgo.SetDebug(cfg.Debug)
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return &Client{session.DB(cfg.Database)}, err
}
