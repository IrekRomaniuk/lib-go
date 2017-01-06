package slice2go

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func mockAction(ip string) (string, error) {
	return ip, nil //do Nothing
}

func TestReadTargets(t *testing.T) {
	p := Target {
		hosts: []string{"8.8.8.8", "1.1.1.1"},
		Action: mockAction,
	}
	Convey("mock Action ", t, func() {
		Convey("So Action is mocked", func() {
			result, _ := Slice2go(&p, 5)
			So(len(result), ShouldEqual,2)
		})
	})

}
