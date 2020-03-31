package MergeSort_test

import (
	"github.com/a1008u/GoTraning/Algorithm/MergeSort"
	"reflect"
	"sort"
	"testing"
)

func Test_MergeSort(t *testing.T){
	nums := []int{6, 15, 4, 2, 8, 5, 11, 9, 7, 13}
	//expectNumbs := []int{6, 15, 4, 2, 8, 5, 11, 9, 7, 13}
	resultNums := MergeSort.MergeSort(nums)

	sort.Ints(nums)
	if !reflect.DeepEqual(nums, resultNums) {
		t.Errorf("Error expect is %v, result is %v", nums, resultNums)
	}
}
