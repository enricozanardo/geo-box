package main

import (
	"github.com/goinggo/tracelog"
	"net"
	"os"
	"google.golang.org/grpc"
	pb_geo "github.com/onezerobinary/geo-box/proto"
	"github.com/onezerobinary/geo-box/geo"
)

const (
	GRPC_PORT = ":1973"
)

func main() {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	// Start the Push Service
	listen, err := net.Listen("tcp", GRPC_PORT)
	if err != nil {
		tracelog.Errorf(err, "app", "main", "Failed to start the GEO service")
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	// Add to the grpcServer the Service
	pb_geo.RegisterGeoServiceServer(grpcServer, &geo.GeoServiceServer{})

	tracelog.Trace("main", "main", "Grpc Server Listening on port " + GRPC_PORT)

	grpcServer.Serve(listen)
}
