package main

import (
	"fmt"
	"math"
)

func findRestaurant(list1 []string, list2 []string) []string {
	lookup := make(map[string]int)
	for i := 0; i < len(list1); i++ {
		lookup[list1[i]] = i
	}
	ans := make([]string, 0)
	min := math.MaxInt32
	for i := 0; i < len(list2); i++ {
		if idx, ok := lookup[list2[i]]; ok {
			sum := idx + i
			if sum < min {
				min = sum
				ans = nil
				ans = append(ans, list2[i])
			} else if sum == min {
				ans = append(ans, list2[i])
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}))
}
