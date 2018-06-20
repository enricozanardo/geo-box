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
	"fmt"
	"github.com/onezerobinary/geo-box/mygprc"
	pb_device "github.com/onezerobinary/db-box/proto/device"
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
		address.Country = "MT"
		//address.Country = "IT"
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


	newPoint := pb_geo.Point{}

	if len(s) > 0 {
		newPoint.Latitude = float32(s[0].Geometry.Location.Lat)
		newPoint.Longitude = float32(s[0].Geometry.Location.Lng)
		newPoint.GeoHash = geohash.Encode(s[0].Geometry.Location.Lat, s[0].Geometry.Location.Lng)
	}

	return &newPoint, nil
}

func GetDevices(researchArea pb_geo.ResearchArea) (devices *pb_geo.Devices, err error) {

	// Trim the geoHash to increase the search area (Precision 5)
	runes := []rune(researchArea.Point.GeoHash)
	researchGeoHash := string(runes[0:researchArea.Precision])

	neighbours := geohash.Neighbors(researchGeoHash)

	// Add to the neighbours the trimmed geoHash
	neighbours = append(neighbours, researchGeoHash)

	fmt.Println(neighbours)

	//find all the companies that are in each geoHash that is coming from neighbours
	for _, geohash := range neighbours {
		// Get companies
		g := pb_device.GeoHash{}
		g.Geohash = geohash
		expoPushTokens := mygprc.GetExpoPushTokensByGeoHash(&g)

		if len(expoPushTokens.Token) == 0 {
			expoPushTokens.Token = []string{}
		} else {

			devices := pb_geo.Devices{}

			for _, device := range expoPushTokens.Token {
				devices.Expopushtoken  = append(devices.Expopushtoken, device)
			}

			return &devices, nil
		}
	}

	d := pb_geo.Devices{}

	return &d, nil
}