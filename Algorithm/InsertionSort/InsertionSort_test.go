package InsertionSort_test

import (
	"github.com/a1008u/GoTraning/Algorithm/InsertionSort"
	"reflect"
	"sort"
	"testing"
)

func Test_InsertionSort(t *testing.T){
	nums := []int{6, 15, 4, 2, 8, 5, 11, 9, 7, 13}
	//expectNumbs := []int{6, 15, 4, 2, 8, 5, 11, 9, 7, 13}
	resultNums := InsertionSort.InsertionSort(nums)

	sort.Ints(nums)
	if !reflect.DeepEqual(nums, resultNums) {
		t.Errorf("Error expect is %v, result is %v", nums, resultNums)
	}
}

