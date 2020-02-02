package PalindromeNumber_test

import (
	"github.com/a1008u/GoTraning/LeetCode/PalindromeNumber"
	"testing"
)

func Test_IsPalindrome_OK_1(t *testing.T) {

	targetNumber := 121
	actual := PalindromeNumber.IsPalindrome(targetNumber)
	if !actual {
		t.Errorf("Error %v", targetNumber)
	}
}

func Test_IsPalindrome_OK_2(t *testing.T) {

	targetNumber := 1221
	actual := PalindromeNumber.IsPalindrome(targetNumber)
	if !actual {
		t.Errorf("Error %v", targetNumber)
	}
}

func Test_IsPalindrome_OK_3(t *testing.T) {

	targetNumber := 12521
	actual := PalindromeNumber.IsPalindrome(targetNumber)
	if !actual {
		t.Errorf("Error %v", targetNumber)
	}
}

func Test_IsPalindrome_NG_1(t *testing.T) {

	targetNumber := -121
	actual := PalindromeNumber.IsPalindrome(targetNumber)
	if actual {
		t.Errorf("Error %v", targetNumber)
	}
}

func Test_IsPalindrome_NG_2(t *testing.T) {

	targetNumber := 10
	actual := PalindromeNumber.IsPalindrome(targetNumber)
	if actual {
		t.Errorf("Error %v", targetNumber)
	}
}

func Test_IsPalindrome_NG_3(t *testing.T) {

	targetNumber := 120
	actual := PalindromeNumber.IsPalindrome(targetNumber)
	if actual {
		t.Errorf("Error %v", targetNumber)
	}
}
