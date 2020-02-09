package RomaToInteger

import (
	"log"
)

var num = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}


func RomanToInt(s string) int {
	out := 0
	ln := len(s)
	for i := 0; i < ln; i++ {
		iNum := string(s[i])
		iResult := num[iNum]
		if i < ln -1 {
			iNextNum := string(s[i+1])
			iNextResult := num[iNextNum]

			if iResult < iNextResult {
				out += iNextResult - iResult
				i++
			} else {
				out += iResult
			}
		} else {
			out += iResult
		}
	}

	log.Println(s, out)
	return out
}

func RomantointV2(s string) int {
	var num int

	preNum := 5000
	for _, v := range []byte(s) {
		xx := tranferV2(v)
		if preNum < xx {
			num = num - preNum + xx - preNum
		} else {
			num += xx
			preNum = xx
		}
	}
	return num
}

func tranferV2(b byte) int {
	bricks := []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	bricksVal := []int{1, 5, 10, 50, 100, 500, 1000}
	for i, v := range bricks {
		if v == b {
			return bricksVal[i]
		}
	}
	return 0
}

func RomantointV3(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		sign := 1
		cur := num[string(s[i])]
		if i+1 < len(s) {
			next := num[string(s[i+1])]
			if next > cur {
				sign = -1
			}
		}
		result += cur * sign
	}
	return result
}

func RomantointV4(s string) int {
	var Dict = map[string]int{"I":1, "V":5, "X":10, "L":50, "C":100, "D":500, "M":1000}
	var prevNum, result int
	for i := len(s)-1; i >= 0; i-- {
		targetNum := Dict[string(s[i])]
		if targetNum < prevNum {
			result -= targetNum
		} else {
			result += targetNum
			prevNum = targetNum
		}
	}
	return result
}
