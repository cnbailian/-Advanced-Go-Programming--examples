/*
Copyright 2021 BaiLian.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cnbailian/Advanced-Go-Programming-examples/rpc/4.5.2/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type HelloService struct {
	auth *types.Authentication
}

func (s *HelloService) Hello(ctx context.Context, message *types.HelloMessage) (*types.HelloMessage, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing credentials")
	}

	var appid string
	var appkey string
	if val, ok := md["user"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}
	if appid != s.auth.User || appkey != s.auth.Password {
		return nil, fmt.Errorf("invalid token")
	}

	return &types.HelloMessage{Value: "Hello " + message.GetValue()}, nil
}

func main() {
	server := grpc.NewServer()
	service := &HelloService{
		auth: &types.Authentication{
			User:     "gopher",
			Password: "password",
		},
	}
	types.RegisterHelloServiceServer(server, service)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("can't listen tcp, error:", err)
	}
	if err := server.Serve(listener); err != nil {
		log.Fatal("can't listen grpc server:", err)
	}
}
