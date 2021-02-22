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
	"io"
	"log"
	"net"

	"github.com/cnbailian/Advanced-Go-Programming-examples/rpc/4.4.4/types"
	"google.golang.org/grpc"
)

var topics = map[string][]types.PubSubService_SubscribeServer{}

type PubSubService struct {}

func (s *PubSubService) Publish(ctx context.Context, message *types.PublishMessage) (*types.PublishMessage, error) {
	if subscribers, ok := topics[message.GetTopic()]; ok {
		for _, subscriber := range subscribers {
			reply := &types.Subscriber{Value: "Hello " + message.GetValue()}
			_ = subscriber.Send(reply)
		}
	}
	return &types.PublishMessage{}, nil
}

func (s *PubSubService) Subscribe(stream types.PubSubService_SubscribeServer) error {
	message, err := stream.Recv()
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}
	topics[message.GetTopic()] = append(topics[message.GetTopic()], stream)
	select{}
}

func main() {
	server := grpc.NewServer()
	types.RegisterPubSubServiceServer(server, new(PubSubService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("can't listen tcp, error:", err)
	}
	if err := server.Serve(listener); err != nil {
		log.Fatal("can't listen grpc server:", err)
	}
}
