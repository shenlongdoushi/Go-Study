package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type twoSumStruct struct {
	twoSumInp
	twoSumAns
}

type twoSumInp struct {
	nums []int
	target int
}

type twoSumAns struct {
	ans1 []int
}

// Test func has to start with capital case to use.
func TestTwoSum(t *testing.T)  {
	twoSumTestCases := []twoSumStruct{
		{
			twoSumInp{[]int{2,7,11,15},9},
			twoSumAns{[]int{0,1}},
		},
		{
			twoSumInp{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			twoSumAns{[]int{1, 3}},
		},
	}

	fmt.Printf("------------------------TwoSum Testing------------------------\n")

	for idx,testCase := range twoSumTestCases {
		testInp, testRes := testCase.twoSumInp, testCase.twoSumAns
		actualRes := twoSum(testInp.nums,testInp.target)
		fmt.Printf("TestCase%d [Input]: %v - [Output]: %v\n",idx,testInp, testRes.ans1)
		fmt.Printf("ActualExec%d [Input]: %v - [Output]: %v\n",idx,testInp, actualRes)

		if !reflect.DeepEqual(testRes.ans1,actualRes) {
			t.Errorf("TestCase%d [Input]: %v should be [Output]: %v\n",idx,testInp, testRes.ans1)
		}
	}

	fmt.Printf("------------------------Test Finished------------------------\n")
}