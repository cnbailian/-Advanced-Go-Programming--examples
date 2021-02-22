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
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
	conn net.Conn
	isLogin bool
}

func (s *HelloService) Hello(request string, reply *string) error {
	if !s.isLogin {
		return fmt.Errorf("unauthorized")
	}
	*reply = "Hello, " + request + ", from " + s.conn.RemoteAddr().String()
	return nil
}

func (s *HelloService) Login(auth string, reply *string) error {
	if auth != "user:password" {
		return errors.New("the user or password is invalid")
	}
	s.isLogin = true
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("can't listen tcp, error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("can't accept connection, err:", err)
		}

		go func() {
			defer conn.Close()

			s := rpc.NewServer()
			_ = s.Register(&HelloService{conn: conn})
			s.ServeConn(conn)
		}()
	}
}
