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
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dial connects error:", err)
	}

	done := make(chan string)
	go func() {
		var keyChanged string
		if err := client.Call("KVStoreService.Watch", 30, &keyChanged); err != nil {
			log.Fatal(err)
		}
		done <- keyChanged
	}()

	// 原文示例中没有在此处 sleep，会导致 watch 和 set 的执行顺序不固定，从而引起无法 watch 到 abc 改变的情况。
	time.Sleep(1 * time.Second)

	if err := client.Call("KVStoreService.Set", [2]string{"abc", "abc-value"}, new(string)); err != nil {
		log.Fatal(err)
	}

	res := <- done
	fmt.Println("watch:", res)
}
