package geoip
//http://www.devdungeon.com/content/ip-geolocation-go
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GeoIP struct {
	// The right side is the name of the JSON variable
	Ip          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name""`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	Zipcode     string  `json:"zip_code"`
	Lat         float32 `json:"latitude"`
	Lon         float32 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
	AreaCode    int     `json:"area_code"`
}

var (
	address  string
	err      error
	geo      GeoIP
	response *http.Response
	body     []byte
)

func Geo(address string, url string) (geo GeoIP, err error) {

	// Use freegeoip.net to get a JSON response
	// There is also /xml/ and /csv/ formats available
	response, err = http.Get(url + address)  //"https://freegeoip.net/json/"
	if err != nil {
		return GeoIP{}, err
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return GeoIP{}, err
	}

	// Unmarshal the JSON byte slice to a GeoIP struct
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return GeoIP{}, err
	}

	return geo, nil

}



