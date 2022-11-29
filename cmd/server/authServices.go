package main

import (
	"context"

	pb "github.com/Y-bro/go-grpcAuth/cmd/proto"
)

func (s *Server) Signin(ctx context.Context, in *pb.Login) (*pb.Result, error) {
	// username := in.UserName
	// password := in.Password

	return nil, nil

}
