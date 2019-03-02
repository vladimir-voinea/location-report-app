package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/vladimir-voinea/location-report-app/server/location_reporting_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) PushLocations(ctx context.Context, in *pb.PushLocationsRequest) (*pb.PushLocationsResponse, error) {
	log.Printf("Received push request")
	for index, location := range in.GetLocations() {
		fmt.Printf("%d. Lat: %f Lon: %f Bearing: %f\n", index, location.Latitude, location.Longitude, location.Bearing)
	}
	response := &pb.PushLocationsResponse{}
	response.Success = true
	return response, nil
}

func GetCredentials() (cert string, key string, err error) {
	var cert_file string
	var key_file string

	flag.StringVar(&cert_file, "cert_file", "", "Certificate file in PEM format")
	flag.StringVar(&key_file, "key_file", "", "Key file")

	flag.Parse()

	fmt.Println("Cert file: " + cert_file)
	fmt.Println("Key file: " + key_file)

	return cert_file, key_file, nil
}

func main() {
	cert_file, key_file, err := GetCredentials()
	if err != nil {
		flag.PrintDefaults()
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	creds, err := credentials.NewServerTLSFromFile(cert_file, key_file)
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}

	s := grpc.NewServer(opts...)
	pb.RegisterLocationReportingServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
