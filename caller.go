package caller

import (
	"fmt"
	"io/ioutil"

	"github.com/yitume/caller/pkg/common"
)

func Init(cfg interface{}, callers ...common.CallerFunc) (err error) {
	var cfgByte []byte
	switch cfg.(type) {
	case string:
		cfgByte, err = parseFile(cfg.(string))
		if err != nil {
			return
		}
	case []byte:
		cfgByte = cfg.([]byte)
	default:
		return fmt.Errorf("type is error %s", cfg)
	}

	for _, caller := range callers {
		obj := caller()
		if err = obj.InitCfg(cfgByte); err != nil {
			return
		}
	}
	return nil
}

// Init from file.
func parseFile(path string) ([]byte, error) {
	// read file to []byte
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return b, err
	}
	return b, nil
}
