package server

// import (
// 	"context"
// 	"fmt"

// 	ch "github.com/0loff/gophkeeper_server/internal/context_helpers"
// 	"github.com/0loff/gophkeeper_server/internal/logger"
// 	pb "github.com/0loff/gophkeeper_server/proto"
// 	"github.com/golang/protobuf/ptypes/empty"
// 	"go.uber.org/zap"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// func (s *Server) UserdataGet(ctx context.Context, in *empty.Empty) (*pb.UserdataGetResponse, error) {
// 	var TextdataEntries []pb.TextdataEntry

// 	uuid, ok := ch.GetUserIDFromContext(ctx)
// 	if !ok {
// 		logger.Log.Error("Cannot get UserID from context")
// 	}

// 	user_id, err := s.UP.GetUserID(ctx, uuid)
// 	if err != nil {
// 		logger.Log.Error("Cannot recognize user", zap.Error(err))
// 		return nil, status.Errorf(codes.Internal, "Internal server error")
// 	}

// 	textdataEntries, err := s.DP.GetTextdata(ctx, user_id)
// 	if err != nil {
// 		logger.Log.Error("Error during get text data by user", zap.Error(err))
// 		return nil, status.Errorf(codes.Internal, "Internal server error")
// 	}

// 	for _, entry := range textdataEntries {
// 		TextdataEntries = append(TextdataEntries, pb.TextdataEntry{
// 			ID:      int64(entry.ID),
// 			Title:   entry.Title,
// 			Article: entry.Article,
// 		})
// 	}

// 	fmt.Print(pb.UserdataGetResponse{Data: TextdataEntries})

// 	return &pb.UserdataGetResponse{Data: TextdataEntries}, nil
// }
