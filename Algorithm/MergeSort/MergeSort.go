package MergeSort

import "log"

func MergeSort(nums []int) []int{

	if len(nums) <= 1{
		return nums
	}

	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])

	log.Println(left, right)

	return merge(left, right)
}

func merge(left []int, right []int) []int{
	var result []int
	i,j := 0, 0

	for{
		if(i < len(left)) && (j < len(right)) {
			if left[i] <= right[j] {
				result = append(result, left[i])
				i += 1
			} else {
				result = append(result, right[j])
				j += 1
			}
		} else {
			break
		}
	}

	if i < len(left) {
		log.Println("left ", left)
		result = append(result, left[i:]...)
	}

	if j < len(right) {
		log.Println("right ", right)
		result = append(result, right[j:]...)
	}

	return result
}
