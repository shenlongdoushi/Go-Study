package main

import "fmt"

// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

// Solution1: Loop the array element to find the sum = 9

func main() {
	// fmt.Print(twoSum_loop([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum_map([]int{2, 11, 11, 7}, 9))
}

func twoSum_loop(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[i]+nums[j] == target {
				// fmt.Println(i,j)
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum_map(nums []int, target int) []int {
	m := make(map[int]int)
	for k,v := range nums {
		if idx, remain := m[target-v]; remain {
			fmt.Println(remain)
			return []int{idx,k}
		}
		m[v]=k
	}
	return nil
}
