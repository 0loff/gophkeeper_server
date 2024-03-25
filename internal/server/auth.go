package server

import (
	"context"

	"github.com/0loff/gophkeeper_server/internal/logger"
	pb "github.com/0loff/gophkeeper_server/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UserAuth(ctx context.Context, in *pb.UserAuthRequest) (*empty.Empty, error) {
	token, err := s.UP.Auth(ctx, in.Login, in.Password, in.Email)
	if err != nil {
		logger.Log.Error("User registration is failed", zap.Error(err))
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "Internal server error. Cannot create new user.")
	}

	authHeader := metadata.New(map[string]string{"token": token})
	if err := grpc.SetHeader(ctx, authHeader); err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot send token as auth header")
	}
	return &emptypb.Empty{}, nil
}
