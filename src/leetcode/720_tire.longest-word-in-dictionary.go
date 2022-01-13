package main

import (
	"fmt"
)

type Tire struct {
	end      bool
	word     string
	children [26]*Tire
}

func (t *Tire) Insert(word string) {
	for _, c := range word {
		if t.children[c-'a'] == nil {
			t.children[c-'a'] = &Tire{}
		}
		fmt.Printf("word %v, ch %v, node %v\n", word, string(c), t)
		t = t.children[c-'a']
	}
	t.word = word
	t.end = true
}

// O(n)
func longestWord(words []string) string {
	maxDepth := 0
	ans := ""
	var dfs func(root *Tire, depth int)
	dfs = func(root *Tire, depth int) {
		if depth > 0 && !root.end {
			return
		}
		// 保留第一个最大长度的
		if depth > maxDepth {
			ans = root.word
			maxDepth = depth
		}
		for i := 0; i < len(root.children); i++ {
			if root.children[i] != nil {
				fmt.Println(i)
				dfs(root.children[i], depth+1)
			}
		}
	}
	root := &Tire{}
	for _, w := range words {
		root.Insert(w)
	}
	dfs(root, 0)
	return ans
}

func main() {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))                    // world
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"})) // apple
}
