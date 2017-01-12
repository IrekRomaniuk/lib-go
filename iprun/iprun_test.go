package iprun

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)
//i.e. GoSNMP(ip string, community string, oid string) (result map[string]string, err error)
func mockAction(ip ...string) (string, error) {
	//fmt.Printf("%s %T\n",string(ip[0]),string(ip[0]))
	return string(ip[0]), nil //do Nothing
}

func TestRunIPs(t *testing.T) {
	p := Targets {
		Hosts: []string{"8.8.8.8", "1.1.1.1"},
		Action: mockAction,
	}
	Convey("mock Action ", t, func() {
		Convey("So Action is mocked", func() {
			result, _ := RunIPs(&p, 5)
			fmt.Println(result)
			So(len(result), ShouldEqual,2)
		})
	})

}
