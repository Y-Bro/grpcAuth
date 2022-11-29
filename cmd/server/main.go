package main

import (
	"log"
	"net"

	pb "github.com/Y-bro/go-grpcAuth/cmd/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
}

func main() {

	net, err := net.Listen("tcp", "0.0.0.0:80")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on port 80")

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &Server{})

	if err := s.Serve(net); err != nil {
		log.Fatal(err)
	}
}
