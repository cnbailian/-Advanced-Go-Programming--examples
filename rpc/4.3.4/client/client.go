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
)

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dial connects error:", err)
	}

	if err := client.Call("HelloService.Login", "user:password", new(string)); err != nil {
		log.Fatal("login err:", err)
	}

	var reply string
	if err := client.Call("HelloService.Hello", "World!", &reply); err != nil {
		log.Fatal("call a remote function error:", err)
	}

	fmt.Println(reply)
}
