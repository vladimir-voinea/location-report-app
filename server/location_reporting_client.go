package main

/*p
import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	pb "github.com/vladimir-voinea/location-report-app/server/location_reporting_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("Getting credentials")
	tls := &tls.Config{InsecureSkipVerify: true}
	// creds, err := credentials.NewClientTLSFromFile(cert)
	// if err != nil {
	// 	fmt.Errorf("could not load tls cert: %s", err)
	// 	os.Exit(1)
	// }

	fmt.Println("Dialing")
	conn, err := grpc.Dial("128.199.48.78:50051", grpc.WithTransportCredentials(credentials.NewTLS(tls)))
	if err != nil {
		fmt.Errorf("could not dial. %s", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Creating client")
	client := pb.NewLocationReportingServiceClient(conn)

	fmt.Println("Requesting")
	fmt.Println(conn.GetState())

	response, err := client.PushLocations(context.Background(), &pb.PushLocationsRequest{Locations: []*pb.Location{{Latitude: 23.34,
		Longitude: 34.53,
		Bearing:   342.1,
		Timestamp: 22222}}})
	if err != nil {
		fmt.Printf("Could not call. %s", err)
	}

	if response != nil {
		if response.Success {
			fmt.Println("Succes")
		} else {
			fmt.Println("Fail")
		}
	}
}
*/
