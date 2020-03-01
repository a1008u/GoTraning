package FizzBuzz_test

import (
	"github.com/a1008u/GoTraning/Algorithm/FizzBuzz"
	"testing"
)

func Test_FizzBuzz(t *testing.T){
	nums := []int{1,2,3,4,5,15}
	expectNumbs := []string{"1","2","Fizz","4","Buzz","FizzBuzz"}
	resultNums := FizzBuzz.FizzBuzz(nums)

	for i := 0; i < len(expectNumbs); i++ {
		println(expectNumbs[i], resultNums[i])
		if expectNumbs[i] != resultNums[i] {
			t.Errorf("Error expect is %v, result is %v", expectNumbs[i], resultNums[i])
		}
	}
}
