package main

import "fmt"

// O(m+n)
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	a, b := 0, 0
	m, n := len(firstList), len(secondList)
	ans := make([][]int, 0)
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for a < m && b < n {
		// fmt.Printf("a %v, b %v\n", firstList[a], secondList[b])
		if firstList[a][1] < secondList[b][0] {
			// no overlap, b after a, move a
			a++
		} else if firstList[a][0] > secondList[b][1] {
			// no overlap, a after b, move b
			b++
		} else {
			// overlapping
			ans = append(ans, []int{max(firstList[a][0], secondList[b][0]), min(firstList[a][1], secondList[b][1])})
			if firstList[a][1] > secondList[b][1] {
				// a is longer, move b
				b++
			} else if firstList[a][1] < secondList[b][1] {
				// b is longer, move a
				a++
			} else {
				// same ending, both move
				a++
				b++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}})) // [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
	fmt.Println(intervalIntersection([][]int{{1, 3}, {5, 9}}, [][]int{}))                                                         // []
	fmt.Println(intervalIntersection([][]int{}, [][]int{{4, 8}, {10, 12}}))                                                       // []
}
