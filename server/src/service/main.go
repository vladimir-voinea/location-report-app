package main

import (
	"context"
	pb "location_reporting_service"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) PushLocations(ctx context.Context, in *pb.PushLocationsRequest) (*pb.PushLocationsResponse, error) {
	log.Printf("Received from: %s", in.GetUser().Name)
	response := &pb.PushLocationsResponse{}
	response.Success = true
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLocationReportingServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
