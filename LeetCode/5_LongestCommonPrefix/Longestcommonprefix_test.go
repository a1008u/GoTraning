package LongestCommonPrefix_test

import (
	LongestCommonPrefix "github.com/a1008u/GoTraning/LeetCode/5_LongestCommonPrefix"
	"testing"
)

func Test_longestCommonPrefix_OK_1(t *testing.T) {

	targetNumber := []string{"flower","flow","flight"}
	expect := "fl"
	actual := LongestCommonPrefix.LongestCommonPrefix(targetNumber)
	if expect != actual {
		t.Errorf("Error1 %v", actual)
	}
	actual2 := LongestCommonPrefix.LongestCommonPrefixV2(targetNumber)
	if expect != actual2 {
		t.Errorf("Error2 %v", actual2)
	}
}

func Test_longestCommonPrefix_OK_2(t *testing.T) {

	targetNumber := []string{"dog","racecar","car"}
	expect := ""
	actual := LongestCommonPrefix.LongestCommonPrefix(targetNumber)
	if expect != actual {
		t.Errorf("Error1 %v", actual)
	}
	actual2 := LongestCommonPrefix.LongestCommonPrefixV2(targetNumber)
	if expect != actual2 {
		t.Errorf("Error2 %v", actual2)
	}
}
