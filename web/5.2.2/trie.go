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

import "fmt"

type Trie struct {
	nodes []*TrieNode
}

type TrieNode struct {
	isEnd bool
	nodes []*TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		nodes: make([]*TrieNode, 26),
	}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	nodes := t.nodes
	for i, c := range word {
		cur := nodes[c-97]
		if cur == nil {
			nodes[c-97] = &TrieNode{
				nodes: make([]*TrieNode, 26),
			}
		}
		if i+1 == len(word) {
			nodes[c-97].isEnd = true
		}
		nodes = nodes[c-97].nodes
	}
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	nodes := t.nodes
	for i, c := range word {
		if nodes[c-97] == nil {
			return false
		}
		if i+1 == len(word) && !nodes[c-97].isEnd {
			return false
		}
		nodes = nodes[c-97].nodes
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	nodes := t.nodes
	for _, c := range prefix {
		if nodes[c-97] == nil {
			return false
		}
		nodes = nodes[c-97].nodes
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()
	trie.Insert("abcdefg")
	fmt.Println(trie.Search("abcd"))
	fmt.Println(trie.StartsWith("abcd"))
}
