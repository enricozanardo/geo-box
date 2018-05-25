package geo

import (
	pb_geo "github.com/onezerobinary/geo-box/proto"
	"github.com/goinggo/tracelog"
	"context"
)


type GeoServiceServer struct {

}


func (s *GeoServiceServer) GetPoint(ctx context.Context, address *pb_geo.Address) (point *pb_geo.Point, err error) {

		point, err = CalculatePoint(*address)

		if err != nil {
			tracelog.Error(err, "geoServiceServer", "GetPoint")
			return &pb_geo.Point{}, err
		}

		return point, nil
}

func (s *GeoServiceServer) GetDeviceList(ctx context.Context, researchArea *pb_geo.ResearchArea) (devices *pb_geo.Devices, err error) {

	devices, err = GetDevices(*researchArea)

	if err != nil {
		tracelog.Error(err, "geoServiceServer", "GetDeviceList")
		return &pb_geo.Devices{}, err
	}

	return devices, nil
}