package TwoSum_test

import (
	"github.com/a1008u/GoTraning/LeetCode/TwoSum"
	"testing"
)

//THE PURPOSE
//
//INTENTION
//
//arg:
//
//return:
//
func Test_TwoSum(t *testing.T) {

	sliceInt := []int{2, 3, 7, 111}
	target := 9
	twoSumResult := TwoSum.TwoSum(sliceInt, target)

	if twoSumResult[0] != 0 {
		t.Errorf("Error %v", twoSumResult)
	}
	if twoSumResult[1] != 2 {
		t.Errorf("Error %v", twoSumResult)
	}
}

//THE PURPOSE
//
//INTENTION
//
//arg:
//
//return:
//
func Test_TwoSumMap(t *testing.T) {
	sliceInt := []int{2, 3, 4}
	target := 6
	twoSumResult := TwoSum.TwoSumMap(sliceInt, target)

	if twoSumResult[0] != 0 {
		t.Errorf("Error %v", twoSumResult)
	}
	if twoSumResult[1] != 2 {
		t.Errorf("Error %v", twoSumResult)
	}
}

//THE PURPOSE
//
//INTENTION
//
//arg:
//
//return:
//
func TestTwoSumOnePass(t *testing.T) {
	sliceInt := []int{2, 3, 4}
	target := 6
	twoSumResult := TwoSum.TwoSumOnePass(sliceInt, target)

	if twoSumResult[0] != 0 {
		t.Errorf("Error %v", twoSumResult)
	}
	if twoSumResult[1] != 2 {
		t.Errorf("Error %v", twoSumResult)
	}
}
