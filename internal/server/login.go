package server

import (
	"context"
	"errors"

	userpkg "github.com/0loff/gophkeeper_server/internal/user"
	pb "github.com/0loff/gophkeeper_server/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UserLogin(ctx context.Context, in *pb.UserLoginRequest) (*empty.Empty, error) {
	token, err := s.UP.Login(ctx, in.Email, in.Password)
	if err != nil {
		if errors.Is(err, userpkg.ErrWrongCreds) {
			return &emptypb.Empty{}, status.Errorf(codes.NotFound, "Wrong login or password")
		}
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "Internal server error")
	}

	authHeader := metadata.New(map[string]string{"token": token})
	if err := grpc.SetHeader(ctx, authHeader); err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot send token as auth header")
	}
	return &emptypb.Empty{}, nil
}
