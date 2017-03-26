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
	Convey("Type of slice and first address", t, func() {
		target := "./iplist.txt"
		hosts, _ := ReadIPs(target)
		Convey("So iplist.txt should be moved to slice of strings", func() {
			So(hosts[0], ShouldHaveSameTypeAs,"8.8.8.8")
		})
		Convey("and first address in iplist.txt should be 8.8.8.8", func() {
			So(hosts[0], ShouldEqual,"8.8.8.8")
		})
	})
	Convey("Type of map and first address pair", t, func() {
		target := "./iplist2.txt"
		hostsmap, _ := ReadIPMaps(target)
		Convey("So iplist2.txt should be moved to map of strings", func() {
			So(hostsmap["10.199.207.1"], ShouldHaveSameTypeAs,"98.175.15.13")
		})
		Convey("and first entry in iplist2.txt should be (10.199.207.1, 98.175.15.13)", func() {
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

func TestRemoveDuplicates(t *testing.T) {
	input := `1.1.1.1 2.2.2.2	22222	App test	file	low	test	This is test only
	         1.1.1.1 2.2.2.2	22222	App test	file	low	test	This is test only
	         1.1.1.1 2.2.2.2	22222	App test	file	low	test	This is test only`
	removed := `1.1.1.1 2.2.2.2	22222	App test	file	low	test	This is test only`

	output, _ := RemoveDuplicates(input)
	Convey("Remove Duplicates from input", t, func() {
		So(removed, ShouldEqual, output)
	})

}
