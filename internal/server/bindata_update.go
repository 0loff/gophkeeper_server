package server

import (
	"context"

	pb "github.com/0loff/gophkeeper_server/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) BindataUpdate(ctx context.Context, in *pb.BindataUpdateRequest) (*pb.CallbackStatusResponse, error) {
	if err := s.DP.UpdBindata(ctx, int(in.ID), in.Binary, in.Metainfo); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	return &pb.CallbackStatusResponse{Status: statusSuccess}, nil
}
