package main

import "fmt"

// O(logn)
func firstBadVersion(n int) int {
	lo, hi, mid := 1, n, 1
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if isBadVersion(mid) {
			if !isBadVersion(mid - 1) {
				return mid
			}
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}

func isBadVersion(version int) bool {
	if version == 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(firstBadVersion(1))
}
