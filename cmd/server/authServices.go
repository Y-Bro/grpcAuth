package main

import (
	"context"

	pb "github.com/Y-bro/go-grpcAuth/cmd/proto"
	"github.com/Y-bro/go-grpcAuth/internal/pkg/jwt"
	"github.com/Y-bro/go-grpcAuth/internal/pkg/users"
)

func (s *Server) Signup(ctx context.Context, in *pb.Login) (*pb.Result, error) {
	user := &users.User{}

	user.Username = in.UserName
	user.Password = in.Password

	id, err := user.Create()

	if err != nil {
		return nil, err
	}

	return &pb.Result{
		Result: id,
	}, nil
}

func (s *Server) Login(ctx context.Context, in *pb.Login) (*pb.Result, error) {
	user := users.User{}

	user.Password = in.Password
	user.Username = in.UserName

	userCheck := user.ValidateUser()

	if !userCheck {
		return &pb.Result{
			Result: "Invalid creds",
		}, nil
	}

	token, err := jwt.GenerateToken(user.Username)

	if err != nil {
		return nil, err
	}

	return &pb.Result{
		Result: token,
	}, nil
}
