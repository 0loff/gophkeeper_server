package server

import (
	"github.com/0loff/gophkeeper_server/internal/data"
	"github.com/0loff/gophkeeper_server/internal/interceptors"
	"github.com/0loff/gophkeeper_server/internal/user"
	pb "github.com/0loff/gophkeeper_server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	statusSuccess string = "success"
	statusFail    string = "fail"
)

type Server struct {
	pb.UnimplementedGophkeeperServer

	Srv *grpc.Server
	UP  user.UserProcessor
	DP  data.DataProcessor
}

func NewServer(up user.UserProcessor, dp data.DataProcessor) *Server {
	return &Server{
		UP: up,
		DP: dp,
	}
}

func (s *Server) Init() {
	s.Srv = grpc.NewServer(grpc.UnaryInterceptor(interceptors.AuthInterceptor))

	reflection.Register(s.Srv)
	pb.RegisterGophkeeperServer(s.Srv, s)
}
