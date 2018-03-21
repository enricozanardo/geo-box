package main

import (
	"fmt"
	"github.com/goinggo/tracelog"
	"net"
	"os"
	"google.golang.org/grpc"
)

const (
	GRPC_PORT = ":1979"
)

func main() {
	// Start the Push Service
	listen, err := net.Listen("tcp", GRPC_PORT)
	if err != nil {
		tracelog.Errorf(err, "app", "main", "Failed to start the service")
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	// Add to the grpcServer the Service
	//pb_push.RegisterPushServiceServer(grpcServer, &PushServiceServer{})

	tracelog.Trace("main", "main", "Grpc Server Listening on port " + GRPC_PORT)

	grpcServer.Serve(listen)
}
