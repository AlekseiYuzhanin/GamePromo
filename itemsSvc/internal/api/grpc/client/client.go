package grpcclient

import (
	"fmt"

	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/config"
	logger "github.com/AlekseiYuzhanin/GamePromoLogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
	log  logger.Logger
}

func New(cfg *config.Config, log logger.Logger) (*Client, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.GrpcClient.Ip, cfg.GrpcClient.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: conn,
		log:  log,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
