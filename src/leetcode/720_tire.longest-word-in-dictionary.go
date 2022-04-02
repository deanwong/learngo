package main

import (
	"fmt"
)

type TireNode struct {
	end      bool
	children [26]*TireNode
}

func (this *TireNode) Insert(word string) {
	node := this
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &TireNode{}
		}
		node = node.children[c-'a']
	}
	node.end = true
}

func (this *TireNode) Search(word string) bool {
	node := this
	for _, c := range word {
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
		if !node.end {
			return false
		}
	}
	return true
}

func longestWord(words []string) string {
	root := &TireNode{}
	for _, word := range words {
		root.Insert(word)
	}
	ans := ""
	for _, word := range words {
		n := len(word)
		max := len(ans)
		if n < max {
			continue
		}
		if n == max && word > ans {
			continue
		}
		if root.Search(word) {
			ans = word
		}
	}
	return ans
}

func main() {
	// fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))                    // world
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"})) // apple
}
