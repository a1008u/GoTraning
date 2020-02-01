package ReverseInteger_test

import (
	"github.com/a1008u/GoTraning/LeetCode/2_ReverseInteger"
	"testing"
)

func Test_Reverse_OK_1(t *testing.T) {
	targetInteger := 0
	reverseInteger := ReverseInteger.Reverse(targetInteger)

	expectInteger := 0
	if reverseInteger != expectInteger {
		t.Errorf("Error expect is %v, actual is %v", expectInteger, reverseInteger)
	}
}

func Test_Reverse_OK_2(t *testing.T) {
	targetInteger := 123
	reverseInteger := ReverseInteger.Reverse(targetInteger)

	expectInteger := 321
	if reverseInteger != expectInteger {
		t.Errorf("Error expect is %v, actual is %v", expectInteger, reverseInteger)
	}
}

func Test_Reverse_OK_3(t *testing.T) {
	targetInteger := -123
	reverseInteger := ReverseInteger.Reverse(targetInteger)

	expectInteger := -321
	if reverseInteger != expectInteger {
		t.Errorf("Error expect is %v, actual is %v", expectInteger, reverseInteger)
	}
}

func Test_Reverse_OK_4(t *testing.T) {
	targetInteger := 120
	reverseInteger := ReverseInteger.Reverse(targetInteger)

	expectInteger := 21
	if reverseInteger != expectInteger {
		t.Errorf("Error expect is %v, actual is %v", expectInteger, reverseInteger)
	}
}
