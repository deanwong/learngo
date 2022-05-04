package main

import "fmt"

func replaceSpace(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ans += "%20"
		} else {
			ans += string(s[i])
		}
	}
	return ans
}

func main() {
	fmt.Println(replaceSpace("We are happy."))
}
