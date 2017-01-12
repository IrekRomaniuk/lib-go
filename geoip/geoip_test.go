package geoip

import (
	"testing"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRunIPs(t *testing.T) {

	}
	Convey("mock Action ", t, func() {
		Convey("So Action is mocked", func() {
			result, _ := RunIPs(&p, 5)
			fmt.Println(result)
			So(len(result), ShouldEqual,2)
		})
	})

}
