package geoip

import (
	"testing"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRunIPs(t *testing.T) {
	Convey("geoip ", t, func() {
		Convey("So geoip", func() {
			result := Geo("www.devdungeon.com")
			fmt.Println(result)
			//So(len(result), ShouldEqual,2)
		})
	})

}
