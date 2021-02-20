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
	"fmt"
	"log"
	"net/rpc"

	"github.com/cnbailian/Advanced-Go-Programming-examples/rpc/4.1.2/types"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ types.HelloServiceInterface = (*HelloServiceClient)(nil)

func (c *HelloServiceClient) Hello(request string, reply *string) error {
	return c.Call(types.HelloServiceName + ".Hello", request, reply)
}

func DialHelloService(network string, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, fmt.Errorf("dial connects error: %w", err)
	}
	return &HelloServiceClient{c}, nil
}

func main() {
	client, err := DialHelloService("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	var reply string
	if err := client.Hello("World!", &reply); err != nil {
		log.Fatal("call a remote function error:", err)
	}

	fmt.Println(reply)
}
