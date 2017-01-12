package readIPs

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	_ "fmt"
)

func TestReadTargets(t *testing.T) {
	Convey("File iplist.txt exists in the same directory ", t, func() {
		target := "./iplist.txt"
		hosts, _ := ReadIPs(target)
		Convey("So iplist.txt should contain 4 valid IPs and more than one empty line", func() {
			So(len(hosts), ShouldEqual,4)
		})
	})
	Convey("File iplist2.txt exists in the same directory ", t, func() {
		target := "./iplist2.txt"
		hostsmap, _ := ReadIPMaps(target)
		Convey("So iplist2.txt should contain 5 valid IP pairs", func() {
			So(len(hostsmap), ShouldEqual,5)
		})
	})
	Convey("Type of output slice and first entry value to be 8.8.8.8 ", t, func() {
		target := "./iplist.txt"
		hosts, _ := ReadIPs(target)
		Convey("So iplist.txt should be moved to slice of integers", func() {
			So(hosts[0], ShouldHaveSameTypeAs,"8.8.8.8")
		})
		Convey("anfd first entry in iplist.txt should is 8.8.8.8", func() {
			So(hosts[0], ShouldEqual,"8.8.8.8")
		})
	})
	Convey("Type of output and first pair ", t, func() {
		target := "./iplist2.txt"
		hostsmap, _ := ReadIPMaps(target)
		Convey("So iplist.txt should be moved to slice of integers", func() {
			So(hostsmap["10.199.207.1"], ShouldHaveSameTypeAs,"98.175.15.13")
		})
		Convey("anfd first entry in iplist.txt should is 8.8.8.8", func() {
			So(hostsmap["10.199.207.1"], ShouldEqual,"98.175.15.13")
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
