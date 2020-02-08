package RomaToInteger

import (
	"log"
)

//var(
//	I int64 = 1
//	V int64 = 5
//	X int64 = 10
//	L int64 = 50
//	C int64 = 100
//	D int64 = 500
//	M int64 = 1000
//	IV int64 = 4
//	IX int64 = 9
//	XL int64 = 40
//	XC int64 = 90
//	CD int64 = 400
//	CM int64 = 900
//)

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
	// reserve := [{"I":1},{"V":5},{"X":10},{"L":50},{"C":100},{"D":500},{"M":1000}]
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
