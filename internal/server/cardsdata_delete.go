package server

import (
	"context"

	pb "github.com/0loff/gophkeeper_server/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CardsdataDelete(ctx context.Context, in *pb.CardsdataDeleteRequest) (*pb.CallbackStatusResponse, error) {
	if err := s.DP.DelCardsdata(ctx, int(in.ID)); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	return &pb.CallbackStatusResponse{Status: statusSuccess}, nil
}
