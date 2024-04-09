package server

import (
	"context"

	pb "github.com/0loff/gophkeeper_server/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CardsdataUpdate(ctx context.Context, in *pb.CardsdataUpdateRequest) (*pb.CallbackStatusResponse, error) {
	if err := s.DP.UpdCardsdata(ctx, int(in.ID), in.Pan, in.Expiry, in.Holder, in.Metainfo); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	return &pb.CallbackStatusResponse{Status: statusSuccess}, nil
}
