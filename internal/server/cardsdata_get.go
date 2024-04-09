package server

import (
	"context"

	ch "github.com/0loff/gophkeeper_server/internal/context_helpers"
	"github.com/0loff/gophkeeper_server/internal/logger"
	pb "github.com/0loff/gophkeeper_server/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) CardsdataGet(ctx context.Context, in *empty.Empty) (*pb.CardsdataEntriesResponse, error) {
	var response pb.CardsdataEntriesResponse
	uuid, ok := ch.GetUserIDFromContext(ctx)
	if !ok {
		logger.Log.Error("Cannot get UserID from context")
	}

	user_id, err := s.UP.GetUserID(ctx, uuid)
	if err != nil {
		logger.Log.Error("Cannot recognize user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	cardsdataEntries := s.DP.ReceiveCardsdata(ctx, user_id)

	for _, entry := range cardsdataEntries {
		response.Data = append(response.Data, &pb.CardsdataEntry{
			ID:       int64(entry.ID),
			Pan:      entry.Pan,
			Expiry:   entry.Expiry,
			Holder:   entry.Holder,
			Metainfo: entry.Metainfo,
		})
	}

	return &response, nil
}
