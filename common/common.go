package common

import "time"

type Caller interface {
	// Init cfg returns parse cfg error.
	InitCfg(cfg []byte) error
	// Get returns the invoker value associated to the given key.
	Get(key string) interface{}
	// Set sets the invoker value associated to the given key.
	Set(key string, val interface{})
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

type CallerFunc func() Caller
