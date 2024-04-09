package server

import (
	"context"

	pb "github.com/0loff/gophkeeper_server/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) TextdataDelete(ctx context.Context, in *pb.TextDataDeleteRequest) (*pb.CallbackStatusResponse, error) {

	if err := s.DP.DelTextdata(ctx, int(in.ID)); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	return &pb.CallbackStatusResponse{Status: statusSuccess}, nil
}
