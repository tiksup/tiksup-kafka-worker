/*
* This file initializes a new gRPC client and a connection closing function.
* Copyright (C) 2024-2025 jsusmachaca
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program. If not, see <https://www.gnu.org/licenses/>.
 */

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
