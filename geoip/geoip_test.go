package geoip

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRunIPs(t *testing.T) {
	Convey("geoip 68.14.2.46", t, func() {
		Convey("So geoip with http://10.29.21.200:8080/json/", func() {
			geo, _ := Geo("68.14.2.46","http://10.29.21.200:8080/json/")
			So(geo.Ip, ShouldEqual,"68.14.2.46")
			So(geo.CountryCode, ShouldEqual,"US")
			So(geo.City, ShouldEqual,"Woonsocket")
			So(geo.RegionName, ShouldEqual,"Rhode Island")
			So(geo.Zipcode, ShouldEqual,"02895")
		})
		Convey("So geoip with https://freegeoip.net/json/", func() {
			geo, _ := Geo("68.14.2.46","https://freegeoip.net/json/")
			So(geo.Ip, ShouldEqual,"68.14.2.46")
			So(geo.CountryCode, ShouldEqual,"US")
			So(geo.City, ShouldEqual,"Woonsocket")
			So(geo.RegionName, ShouldEqual,"Rhode Island")
			So(geo.Zipcode, ShouldEqual,"02895")
		})
	})

}

/*
{
  "ip": "68.14.2.46",
  "country_code": "US",
  "country_name": "United States",
  "region_code": "RI",
  "region_name": "Rhode Island",
  "city": "Woonsocket",
  "zip_code": "02895",
  "time_zone": "America/New_York",
  "latitude": 42.0029,
  "longitude": -71.5148,
  "metro_code": 521
}
 */
