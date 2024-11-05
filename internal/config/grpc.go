package config

import (
	pb "github.com/tiksup/tiksup-kafka-worker/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Conn   *grpc.ClientConn
	Client pb.EventTriggerServiceClient
}

func CreateEventClient(target string) (*Client, error) {
	conn, err := grpc.NewClient(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	client := pb.NewEventTriggerServiceClient(conn)
	return &Client{
		Conn:   conn,
		Client: client,
	}, nil
}

func (c *Client) Close() {
	c.Conn.Close()
}
