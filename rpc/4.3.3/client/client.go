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
	"net"
	"net/rpc"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("can't listen tcp, error:", err)
	}

	clientChan := make(chan *rpc.Client)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("can't accept connection, err:", err)
			}
			clientChan <- rpc.NewClient(conn)
		}
	}()

	client := <- clientChan
	var reply string
	if err := client.Call("HelloService.Hello", "World!", &reply); err != nil {
		log.Fatal("call a remote function error:", err)
	}
	fmt.Println(reply)
}
