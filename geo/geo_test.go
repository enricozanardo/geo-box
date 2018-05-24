package geo

import (
	"testing"
	pb_geo "github.com/onezerobinary/geo-box/proto"
	"github.com/goinggo/tracelog"
	"strconv"
	"fmt"
)

func TestCalculatePoint(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	fakeAddress := pb_geo.Address{}
	fakeAddress.Address = "Via Brennero"
	fakeAddress.AddressNumber = "16"
	fakeAddress.PostalCode = "39100"
	fakeAddress.Place = "Bolzano"
	fakeAddress.Country = "IT"

	point, err := CalculatePoint(fakeAddress)

	if err != nil {
		t.Error("It was not possible to calculate the Point")
	}

	la := strconv.FormatFloat(float64(point.Latitude), 'G', -1, 64)
	lo := strconv.FormatFloat(float64(point.Longitude), 'G', -1, 64)

	coordinates := "Lat: " + la  + " Long: " + lo + " Geohash: " + point.GeoHash

	tracelog.Trace("geo_test", "TestCalculatePoint", coordinates)
}


func TestGetDevices(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	fakeAddress := pb_geo.Address{}
	fakeAddress.Address = "Triq Il San Pawl"
	fakeAddress.AddressNumber = "493"
	fakeAddress.PostalCode = "SPB3416"
	fakeAddress.Place = "San Pawl Il-Bahar"
	fakeAddress.Country = "MT"

	fakePoint, err := CalculatePoint(fakeAddress)

	if err != nil {
		t.Error("It was not possible to calculate the Point")
	}

	fakeResearchArea := pb_geo.ResearchArea{}
	fakeResearchArea.Precision = 5
	fakeResearchArea.Point = fakePoint

	devices, err := GetDevices(fakeResearchArea)

	if err != nil {
		t.Error("It was not possible to get the devices")
	}

	for _, token := range devices.Expopushtoken {
		fmt.Println("Device: " + token)
	}
}