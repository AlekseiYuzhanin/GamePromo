package grpcserver

import (
	"fmt"
	"net"

	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/config"
	logger "github.com/AlekseiYuzhanin/GamePromoLogger"
	"google.golang.org/grpc"
)

type Server struct {
	conn *grpc.Server
	ln   net.Listener
	log  logger.Logger
}

func New(cfg *config.Config, log logger.Logger) (*Server, error) {
	srv := grpc.NewServer()
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcServer.Port))
	if err != nil {
		return nil, err
	}
	return &Server{
		conn: srv,
		ln:   ln,
		log:  log,
	}, nil
}

func (s *Server) Start() error {
	return s.conn.Serve(s.ln)
}

func (s *Server) Close() {
	s.conn.GracefulStop()
	if err := s.ln.Close(); err != nil {
		s.log.Error("Error close connection", logger.Field{Key: "Err", Value: err})
	}
	s.log.Info("GRPC server closed...")
}
