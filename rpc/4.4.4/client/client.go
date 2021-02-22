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
	"io"
	"log"
	"time"

	"github.com/cnbailian/Advanced-Go-Programming-examples/rpc/4.4.4/types"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal("dial connects error:", err)
	}
	defer conn.Close()

	client := types.NewPubSubServiceClient(conn)
	var topic = "topic"
	go func() {
		for {
			if _, err := client.Publish(context.Background(), &types.PublishMessage{Topic: topic, Value: "World!"}); err != nil {
				log.Fatal("publish error:", err)
			}
			time.Sleep(3 * time.Second)
		}
	}()

	go func() {
		stream, err := client.Subscribe(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		if err := stream.Send(&types.Subscriber{Topic: topic}); err != nil {
			log.Fatal(err)
		}
		for {
			reply, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			fmt.Println(fmt.Sprintf("subscriber 0: %s", reply.GetValue()))
		}
	}()

	go func() {
		stream, err := client.Subscribe(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		if err := stream.Send(&types.Subscriber{Topic: topic}); err != nil {
			log.Fatal(err)
		}
		for {
			reply, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			fmt.Println(fmt.Sprintf("subscriber 1: %s", reply.GetValue()))
		}
	}()

	select {}
}
