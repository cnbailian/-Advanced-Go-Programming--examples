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
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type KVStoreService struct {
	m map[string]string
	filter map[string]func(key string)
	mu sync.Mutex
}

func (s *KVStoreService) Get(key string, value *string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if v, ok := s.m[key]; ok {
		*value = v
		return nil
	}
	return fmt.Errorf("not found")
}

func (s *KVStoreService) Set(kv [2]string, reply *string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	key, value := kv[0], kv[1]
	if old := s.m[key]; old != value {
		for _, fn := range s.filter {
			fn(key)
		}
	}

	s.m[key] = value
	return nil
}

func (s *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10)

	s.mu.Lock()
	s.filter[id] = func(key string) { ch <- key }
	s.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
}

func main() {
	service := &KVStoreService{
		m:      map[string]string{},
		filter: map[string]func(key string){},
		mu:     sync.Mutex{},
	}
	_ = rpc.RegisterName("KVStoreService", service)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("can't listen tcp, error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("can't accept connection, err:", err)
		}
		rpc.ServeConn(conn)
	}
}
