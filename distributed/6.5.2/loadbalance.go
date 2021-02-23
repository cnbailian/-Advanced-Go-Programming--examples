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
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func shuffle(indexes []int) {
	for i := len(indexes); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		indexes[lastIdx], indexes[idx] = indexes[idx], indexes[lastIdx]
	}
}

func shuffleUnsafe(indexes []int) {
	for i := 0; i < len(indexes); i++ {
		a := rand.Intn(len(indexes))
		b := rand.Intn(len(indexes))
		indexes[a], indexes[b] = indexes[b], indexes[a]
	}
}

func main() {
	var cnt = map[int]int{}
	for i := 0; i < 1000000; i++ {
		var s = []int{0, 1, 2, 3, 4, 5, 6}
		shuffle(s)
		cnt[s[0]]++
	}

	fmt.Println(cnt)
}
