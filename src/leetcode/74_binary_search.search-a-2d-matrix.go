package main

import "fmt"

// O(logn)
func searchMatrix(matrix [][]int, target int) bool {
	rlo, rhi, rmid := 0, len(matrix)-1, 0
	for rlo < rhi {
		// 向上取整
		rmid = rlo + (rhi-rlo)/2 + 1
		if matrix[rmid][0] <= target {
			rlo = rmid
		} else {
			rhi = rmid - 1
		}
	}
	clo, chi, cmid := 0, len(matrix[0])-1, 0
	for clo <= chi {
		cmid = clo + (chi-clo)/2
		if matrix[rlo][cmid] == target {
			return true
		}
		if matrix[rlo][cmid] < target {
			clo = cmid + 1
		} else {
			chi = cmid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	fmt.Println(searchMatrix([][]int{{1}}, 1))
	fmt.Println(searchMatrix([][]int{{1}, {3}}, 3))
}
