package server

import (
	"context"

	ch "github.com/0loff/gophkeeper_server/internal/context_helpers"
	"github.com/0loff/gophkeeper_server/internal/logger"
	pb "github.com/0loff/gophkeeper_server/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) BindataCreate(ctx context.Context, in *pb.BindataStoreRequest) (*pb.CallbackStatusResponse, error) {
	uuid, ok := ch.GetUserIDFromContext(ctx)
	if !ok {
		logger.Log.Error("Cannot get UserID from context")
	}

	user_id, err := s.UP.GetUserID(ctx, uuid)
	if err != nil {
		logger.Log.Error("Cannot recognize user", zap.Error(err))
		return &pb.CallbackStatusResponse{Status: statusFail}, status.Errorf(codes.Internal, "Internal server error")
	}

	if err = s.DP.StoreBindata(ctx, user_id, in.Binary, in.Metainfo); err != nil {
		return &pb.CallbackStatusResponse{Status: statusFail}, status.Errorf(codes.Internal, "Internal server error")
	}

	return &pb.CallbackStatusResponse{Status: statusSuccess}, nil
}
