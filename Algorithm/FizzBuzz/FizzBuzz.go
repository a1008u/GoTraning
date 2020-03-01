package FizzBuzz

import "strconv"

func FizzBuzz(nums []int)[]string{
	result := []string{}
	for i := range nums{
		if nums[i]%3 == 0 && nums[i]%5 == 0{
			result = append(result, "FizzBuzz")
			continue
		}
		if nums[i]%3 == 0{
			println("fizz")
			result = append(result, "Fizz")
			continue
		}
		if nums[i]%5 == 0{
			result = append(result, "Buzz")
			continue
		}
		result = append(result, strconv.Itoa(nums[i]))
	}
	return result
}
