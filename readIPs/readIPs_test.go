package readIPs

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReadTargets(t *testing.T) {
	Convey("File iplist.txt exists in the same directory ", t, func() {
		target := "./iplist.txt"
		hosts, _ := ReadIPs(target)
		Convey("So iplist.txt should contain 4 valid IPs and more than one empty line", func() {
			So(len(hosts), ShouldEqual,4)
		})
	})
	Convey("File iplist.txt does not exists", t, func() {
		target := "./tmp/xxx/iplist.txt"
		_, err := ReadIPs(target)
		Convey("So file iplist.txt does not exists", func() {
			So(err.Error(), ShouldContainSubstring,"does not exist")
		})
	})
}
