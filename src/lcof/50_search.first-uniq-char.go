package main

import (
	"fmt"
)

func firstUniqChar(s string) byte {
	letter := [26]int{}
	var ans byte = ' '
	for i := 0; i < len(s); i++ {
		letter[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if letter[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(firstUniqChar("abaccdeff")) // 'b'
	fmt.Println(firstUniqChar(""))          // ' '
}
