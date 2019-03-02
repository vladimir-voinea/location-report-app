package main

import (
	"context"
	"log"
	"net"

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
	log.Printf("Received from: %s", in.GetLocations()[0].Latitude, in.GetLocations()[0].Longitude, in.GetLocations()[0].Bearing)
	response := &pb.PushLocationsResponse{}
	response.Success = true
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	certFile := "server1.pem"
	keyFile := "server1.key"

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
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
