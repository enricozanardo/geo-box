package geo

import (
	pb_geo "github.com/onezerobinary/geo-box/proto"
	"github.com/goinggo/tracelog"
)



type GeoServiceServer struct {

}


func (s *GeoServiceServer) GetPoint(address *pb_geo.Address) (point *pb_geo.Point, err error) {

		point, err = CalculatePoint(*address)

		if err != nil {
			tracelog.Error(err, "geoServiceServer", "GetPoint")
			return &pb_geo.Point{}, err
		}

		return point, nil
}

func (s *GeoServiceServer) GetDeviceList(researchArea *pb_geo.ResearchArea) (devices *pb_geo.Devices, err error) {

	//TODO: get the devices
	fakeDevices := []pb_geo.Device{}
	fakeDevices = append(fakeDevices, pb_geo.Device{"ABC-123"})
	fakeDevices = append(fakeDevices, pb_geo.Device{"FGH-456"})

	for _, device := range fakeDevices {
		if err := stream.Send(&device); err != nil {
			return err
		}
	}

	return nil
}



//
//rpc RecordRoute(stream Point) returns (RouteSummary) {}
//
//// A Bidirectional streaming RPC.
////
//// Accepts a stream of RouteNotes sent while a route is being traversed,
//// while receiving other RouteNotes (e.g. from other users).
//rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}