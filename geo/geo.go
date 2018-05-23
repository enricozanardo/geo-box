package geo



import (
	pb_geo "github.com/onezerobinary/geo-box/proto"
	"strings"
	"github.com/goinggo/tracelog"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/onezerobinary/geo-box/model"
	"github.com/mmcloughlin/geohash"
)

const (
	URL = "https://maps.googleapis.com/maps/api/geocode/json?address="
	GMAP_KEY = "AIzaSyAsNbeCd9XNLrAFx3_ErXg4J6jzVfv_dgo"
)


func CalculatePoint(address pb_geo.Address)(point *pb_geo.Point, error error){

	address.Address = strings.Replace(address.Address, " ", "+", -1)
	address.AddressNumber = strings.Replace(address.AddressNumber, " ", "", -1)
	address.PostalCode = strings.Replace(address.PostalCode, " ", "", -1)
	address.Place = strings.Replace(address.Place, " ", "+", -1)
	address.Country = strings.Replace(address.Country, " ", "+", -1)

	if address.Country == "" {
		address.Country = "IT"
	}

	addressToSearch := address.Address + "+" + address.AddressNumber + "+" + address.PostalCode + "+" + address.Place + "+" + address.Country

	searchURL := URL + addressToSearch + "&key=" + GMAP_KEY

	tracelog.Trace("route", "CalculatePoint", "Search for: " + searchURL)

	res, err := http.Get(searchURL)
	defer res.Body.Close()

	if err != nil {
		tracelog.Error(err, "route", "It was not possible to get the point calling the API")
		return &pb_geo.Point{}, err
	}

	body, err := ioutil.ReadAll(res.Body)

	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(body, &objmap)

	if err != nil {
		tracelog.Error(err, "route", "It was not possible to read the information of the address")
		return &pb_geo.Point{}, err
	}

	var s []model.GeoInfo

	err = json.Unmarshal(*objmap["results"], &s)

	if err != nil {
		tracelog.Error(err, "route", "No results found")
		return &pb_geo.Point{}, err
	}

	if len(s) > 0 {
		point.Latitude = float32(s[0].Geometry.Location.Lat)
		point.Longitude = float32(s[0].Geometry.Location.Lng)
		point.GeoHash = geohash.Encode(s[0].Geometry.Location.Lat, s[0].Geometry.Location.Lng)
	}

	return point, nil
}


//rpc GetPoint (Address) returns (Point) {}


//rpc RecordRoute(stream Point) returns (RouteSummary) {}
//
//rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}