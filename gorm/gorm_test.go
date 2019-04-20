package gorm

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var cfgSuccess = `
    [callerGorm.default]
		debug = true
        level = "panic"
        network = "tcp"
        dialect = "mysql"
        addr = "127.0.0.1:3306"
        username = "root"
        password = ""
        name = "oauth"
        charset = "utf8"
        parseTime = "True"
        loc = "Local"
        timeout = "1s"
        readTimeout = "1s"
        writeTimeout = "1s"
        maxOpenConns = 30
        maxIdleConns = 10
        connMaxLifetime = "300s"
`

func TestParseConfig(t *testing.T) {
	Convey("test parse config", t, func() {
		obj := Cfg{}
		err := parseConfig([]byte(cfgSuccess), &obj)
		So(err, ShouldBeNil)
		Convey("解析出 gorm defalut 不能为空", func() {
			_, ok := obj.CallerGorm["default"]
			So(ok, ShouldBeTrue)
		})
	})
}
