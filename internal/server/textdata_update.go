package server

import (
	"context"

	pb "github.com/0loff/gophkeeper_server/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) TextdataUpdate(ctx context.Context, in *pb.TextDataUpdateRequest) (*pb.CallbackStatusResponse, error) {
	if err := s.DP.UpdTextdata(ctx, int(in.ID), in.Text, in.Metainfo); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	return &pb.CallbackStatusResponse{Status: statusSuccess}, nil
}
