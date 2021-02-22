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
	"log"
	"net"

	"github.com/cnbailian/Advanced-Go-Programming-examples/rpc/4.5.3/types"
	"google.golang.org/grpc"
)

type HelloService struct{}

func (s *HelloService) Hello(ctx context.Context, message *types.HelloMessage) (*types.HelloMessage, error) {
	return &types.HelloMessage{Value: "Hello " + message.GetValue()}, nil
}

func filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("method:", info.FullMethod)
	return handler(ctx, req)
}

func main() {
	server := grpc.NewServer(grpc.UnaryInterceptor(filter))
	types.RegisterHelloServiceServer(server, new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("can't listen tcp, error:", err)
	}
	if err := server.Serve(listener); err != nil {
		log.Fatal("can't listen grpc server:", err)
	}
}
