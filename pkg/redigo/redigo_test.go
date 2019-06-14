package redigo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var cfgSuccess = `
    [callerRedigo.default]
        debug = true
        addr = "127.0.0.1:6379"
        network = "tcp"
        db = 0
        password = ""
        connectTimeout = "1s"
        readTimeout = "1s"
        writeTimeout = "1s"

        maxIdle = 5
        maxActive = 20
        idleTimeout = "60s"
        wait = false
`

func TestParseConfig(t *testing.T) {
	Convey("test parse config", t, func() {
		obj := Cfg{}
		err := parseConfig([]byte(cfgSuccess), &obj)
		So(err, ShouldBeNil)
		Convey("解析出 redigo defalut 不能为空", func() {
			_, ok := obj.CallerRedigo["default"]
			So(ok, ShouldBeTrue)
		})
	})
}
