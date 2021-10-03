package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question1 struct {
	para1
	ans1
}

type para1 struct {
	nums   []int
	target int
}

type ans1 struct {
	one []int
}

func TestTwoSum(t *testing.T) {
	qs := []question1{
		{
			para1{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans1{[]int{1, 3}},
		},
		{
			para1{[]int{2, 7, 11, 15}, 9},
			ans1{[]int{0, 2}},
		},
	}

	fmt.Printf("------------------------TwoSum Testing------------------------\n")

	for _, testCase := range qs {
		testInp, testRes := testCase.para1, testCase.ans1
		actualRes := twoSum(testInp.nums, testInp.target)
		fmt.Printf("Actual: [Input]: %v   [Output]: %v\n", testInp, actualRes)
		fmt.Printf("TestFile: [Input]: %v   [Output]: %v\n", testInp, testRes.one)
		fmt.Println(reflect.DeepEqual(actualRes, testRes.one))
		if !reflect.DeepEqual(actualRes, testRes.one) {
			t.Errorf("[Input]: %v should be [Output]: %v",testInp, testRes.one)
		}

	}
	fmt.Printf("------------------------Test Finished------------------------\n")
}
