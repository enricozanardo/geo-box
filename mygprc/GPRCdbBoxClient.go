package mygprc


import (
	pb_device "github.com/onezerobinary/db-box/proto/device"
	"github.com/goinggo/tracelog"
	"os"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)


const (
	ADDRESS = "localhost:1982"    // Development address of db-box
	//ADDRESS = "172.104.230.81:1982" // Staging environment of db-box
)

func StartGRPCConnection() (connection *grpc.ClientConn){
	// set up connection to the gRPC server
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "StartGRPCConnection", "Did not open the connection")
		os.Exit(1)
	}

	return conn
}

func StopGRPCConnection(connection *grpc.ClientConn){
	// set up connection to the gRPC server
	err := connection.Close()

	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "StopGRPCConnection", "Did not close the connection")
		os.Exit(1)
	}
}


func GetExpoPushTokensByGeoHash(geoHash *pb_device.GeoHash) (*pb_device.ExpoPushTokens) {
	tracelog.Trace("GPRCdbBoxClient","GetExpoPushTokensByGeoHash","Search devices in a nearest area")

	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)

	client := pb_device.NewDeviceServiceClient(conn)

	expoPushTokens, _ := client.GetExpoPushTokensByGeoHash(context.Background(), geoHash)

	if len(expoPushTokens.Token) == 0 {
		expoPushTokens.Token = []string{}
	}

	return expoPushTokens
}