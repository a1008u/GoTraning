package PalindromeNumber

func IsPalindrome(x int) bool {
	// ReverseIntegerする必要ないのはfalse
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}

	// ReverseInteger
	revertedNumber := 0
	for {
		if x <= revertedNumber {
			break
		}
		revertedNumber = revertedNumber * 10 + x % 10
		x /= 10
	}
	return  x == revertedNumber || x == revertedNumber/10
}
