package test

import (
	"testing"

	"github.com/Wishrem/wuso/config"
	"github.com/Wishrem/wuso/pkg/utils"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetMysqlDSN(t *testing.T) {
	Convey("Test get MySQL's DSN", t, func() {
		config.Init(path)
		So(utils.GetMysqlDSN(), ShouldEqual, "wuso:wuso@tcp(127.0.0.1:3306)/wuso?charset=utf8mb4&parseTime=true")
	})
}
