package redigo

import (
	"github.com/yitume/caller/common"
)

type Cfg struct {
	CallerRedigo map[string]CallerCfg
}

type CallerCfg struct {
	Debug bool

	Network        string //tcp
	Addr           string // 127.0.0.1:6379
	DB             int
	Password       string
	ConnectTimeout common.Duration
	ReadTimeout    common.Duration
	WriteTimeout   common.Duration

	MaxIdle     int
	MaxActive   int
	IdleTimeout common.Duration
	Wait        bool
}
