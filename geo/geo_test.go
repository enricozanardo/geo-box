package geo

import (
	"testing"
	pb_geo "github.com/onezerobinary/geo-box/proto"
	"github.com/goinggo/tracelog"
	"strconv"
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

	la := strconv.FormatFloat(float64(point.Latitude), 'E', -1, 64)
	lo := strconv.FormatFloat(float64(point.Longitude), 'E', -1, 64)

	coordinates := "Lat: " + la  + " Long: " + lo

	tracelog.Trace("", "", coordinates)
}