package SelectionSort_test

import (
	"github.com/a1008u/GoTraning/Algorithm/SelectionSort"
	"reflect"
	"sort"
	"testing"
)

func Test_SelectionSort(t *testing.T){
	nums := []int{6, 15, 4, 2, 8, 5, 11, 9, 7, 13}
	//expectNumbs := []int{6, 15, 4, 2, 8, 5, 11, 9, 7, 13}
	resultNums := SelectionSort.Selectionsort(nums)

	sort.Ints(nums)
	if !reflect.DeepEqual(nums, resultNums) {
		t.Errorf("Error expect is %v, result is %v", nums, resultNums)
	}
}

func Test_SelectionSort2(t *testing.T){
	nums := []int{6, 15, 4, 2, 8, 5, 11, 9, 7, 13}
	//expectNumbs := []int{6, 15, 4, 2, 8, 5, 11, 9, 7, 13}
	resultNums := SelectionSort.Selectionsort2(nums)

	sort.Ints(nums)
	if !reflect.DeepEqual(nums, resultNums) {
		t.Errorf("Error expect is %v, result is %v", nums, resultNums)
	}
}
