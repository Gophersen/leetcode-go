package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mapNums := make(map[int]int)
	for i, j := range nums {
		if num, ok := mapNums[target-nums[i]]; ok {
			return []int{num, i}
		}
		mapNums[j] = i
	}
	return []int{}
}

func main() {
	result := twoSum([]int{3, 2, 3}, 6)
	fmt.Println(result)
}
