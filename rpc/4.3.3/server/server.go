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
	"net"
	"net/rpc"
	"time"
)

type HelloService struct {}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello, " + request
	return nil
}

func main() {
	_ = rpc.RegisterName("HelloService", new(HelloService))

	for {
		conn, _ := net.Dial("tcp", ":1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		rpc.ServeConn(conn)
		conn.Close()
	}
}
