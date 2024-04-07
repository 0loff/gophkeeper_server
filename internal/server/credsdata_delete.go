package server

import (
	"context"

	pb "github.com/0loff/gophkeeper_server/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CredsdataDelete(ctx context.Context, in *pb.CredsdataDeleteRequest) (*pb.CallbackStatusResponse, error) {
	if err := s.DP.DelCredsdata(ctx, int(in.ID)); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	return &pb.CallbackStatusResponse{Status: statusSuccess}, nil
}
