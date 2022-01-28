package main

import (
	"fmt"
)

// O(n+m+∣Σ∣) -> O(n+m+26)
func checkInclusion(s1 string, s2 string) bool {
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := 0; i < len(s1); i++ {
		cnt1[s1[i]-'a']++
	}
	n := len(s2)
	// left pointer
	i := 0
	// right pointer
	for j := 0; j < n; j++ {
		// move right pointer and add
		cnt2[s2[j]-'a']++
		for i <= j && cnt2[s2[i]-'a'] > cnt1[s2[i]-'a'] {
			// move left pointer and subtract
			cnt2[s2[i]-'a']--
			i++
		}
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))        // true
	fmt.Println(checkInclusion("ab", "eidboaoo"))        // false
	fmt.Println(checkInclusion("hello", "ooolleoooleh")) // false
	fmt.Println(checkInclusion("adc", "dcda"))           // true
}
