package main

import (
	"context"
	"errors"
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

	if cert_file == "" {
		return "", "", errors.New("Certificate file path is empty")
	}

	if key_file == "" {
		return "", "", errors.New("Key file path is empty")
	}

	return cert_file, key_file, nil
}

func main() {
	cert_file, key_file, err := GetCredentials()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}
	defer lis.Close()
	fmt.Println("Almost ready")

	var opts []grpc.ServerOption
	creds, err := credentials.NewServerTLSFromFile(cert_file, key_file)
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
		os.Exit(1)
	}

	fmt.Println("Got credentials: " + creds.Info().SecurityProtocol)

	opts = []grpc.ServerOption{grpc.Creds(creds)}
	s := grpc.NewServer(opts...)
	pb.RegisterLocationReportingServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		os.Exit(1)
	}

	fmt.Println("Listening!")
}
