package caller

import (
	"fmt"
	"github.com/yitume/caller/common"
	"io/ioutil"
)

func Init(cfg interface{}, callers ...common.CallerFunc) error {
	var cfgByte []byte
	var err error
	switch cfg.(type) {
	case string:
		cfgByte, err = parseFile(cfg.(string))
		if err != nil {
			return err
		}
	case []byte:
		cfgByte = cfg.([]byte)
	default:
		return fmt.Errorf("type is error %s", cfg)
	}

	for _, caller := range callers {
		obj := caller()
		obj.InitCfg(cfgByte)

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
