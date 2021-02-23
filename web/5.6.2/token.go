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

var tokenBucket = make(chan struct{}, 100)

func main() {
	// 写入令牌
	go func() {
		for {
			ticker := time.NewTicker(1 * time.Second)
			select {
			case <-ticker.C:
				tokenBucket <- struct{}{}
			}
		}
	}()

	// 读取令牌
	for {
		<-tokenBucket
		fmt.Println(time.Now().Second(), "running...")
	}

	// 现实世界中的令牌桶要复杂得多，比如对于 HTTP 请求限流的令牌桶来说，不同的请求根据调用时间和相应大小来计算要使用多少令牌。
}
