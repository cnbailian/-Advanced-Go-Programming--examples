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
	"time"
)

type lock struct {
	c chan struct{}
}

func NewLock() *lock {
	lock := &lock{
		c: make(chan struct{}, 1),
	}
	return lock
}

func (l *lock) Lock() bool {
	select {
	case l.c <- struct{}{}:
		return true
	default:
		return false
	}
}

func (l *lock) Unlock() {
	<-l.c
}

// 为什么叫尝试锁，不是自旋锁吗？
func main() {
	lock := NewLock()

	go func() {
		for {
			if lock.Lock() {
				fmt.Println("id0: locking")
				time.Sleep(1 * time.Second)
				lock.Unlock()
			}
		}
	}()

	for {
		if lock.Lock() {
			fmt.Println("id1: locking")
			time.Sleep(2 * time.Second)
			lock.Unlock()
		}
	}
}
