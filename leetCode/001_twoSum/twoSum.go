package leetcode

import "fmt"

// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

// Solution1: Loop the array element to find the sum = 9

func main() {
	fmt.Print(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
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
