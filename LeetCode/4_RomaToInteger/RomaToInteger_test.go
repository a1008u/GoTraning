package RomaToInteger_test

import (
	RomaToInteger "github.com/a1008u/GoTraning/LeetCode/4_RomaToInteger"
	"testing"
)

func Test_romanToInt_OK_1(t *testing.T) {

	targetNumber := "III"
	expect := 3
	actual := RomaToInteger.RomanToInt(targetNumber)
	if expect != actual {
		t.Errorf("Error %v", targetNumber)
	}
}

func Test_romanToInt_OK_2(t *testing.T) {

	targetNumber := "IV"
	expect := 4
	actual := RomaToInteger.RomanToInt(targetNumber)
	if expect != actual {
		t.Errorf("Error %v", targetNumber)
	}
}

func Test_romanToInt_OK_3(t *testing.T) {

	targetNumber := "IX"
	expect := 9
	actual := RomaToInteger.RomanToInt(targetNumber)
	if expect != actual {
		t.Errorf("Error %v", targetNumber)
	}
}

func Test_romanToInt_OK_4(t *testing.T) {

	targetNumber := "LVIII"
	expect := 58
	actual := RomaToInteger.RomanToInt(targetNumber)
	if expect != actual {
		t.Errorf("Error %v", targetNumber)
	}
}

func Test_romanToInt_OK_5(t *testing.T) {

	targetNumber := "MCMXCIV"
	expect := 1994
	actual := RomaToInteger.RomanToInt(targetNumber)
	if expect != actual {
		t.Errorf("Error %v", targetNumber)
	}
}
