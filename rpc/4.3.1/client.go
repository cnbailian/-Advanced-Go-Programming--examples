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
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dial connects error:", err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	done := client.Go("HelloService.Hello", "World!", new(string), nil)
	// other work
	res := <- done.Done
	if res.Error != nil {
		log.Fatal("call a remote function error:", err)
	}
	fmt.Println(*res.Reply.(*string))
}

