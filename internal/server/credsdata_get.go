package server

import (
	"context"

	ch "github.com/0loff/gophkeeper_server/internal/context_helpers"
	"github.com/0loff/gophkeeper_server/internal/logger"
	pb "github.com/0loff/gophkeeper_server/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CredsdataGet(ctx context.Context, in *empty.Empty) (*pb.CredsdataEntriesResponse, error) {
	var response pb.CredsdataEntriesResponse
	uuid, ok := ch.GetUserIDFromContext(ctx)
	if !ok {
		logger.Log.Error("Cannot get UserID from context")
	}

	user_id, err := s.UP.GetUserID(ctx, uuid)
	if err != nil {
		logger.Log.Error("Cannot recognize user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	credsdataEntries := s.DP.ReceiveCredsdata(ctx, user_id)

	for _, entry := range credsdataEntries {
		response.Data = append(response.Data, &pb.CredsdataEntry{
			ID:       int64(entry.ID),
			Username: entry.Username,
			Password: entry.Password,
			Metainfo: entry.Metainfo,
		})
	}

	return &response, nil
}
