package QuickSort

import "log"

func QuickSort(nums []int) []int{

	if len(nums) <= 1{
		log.Printf("nums is %v", nums)
		return nums
	}

	pivot := nums[0]
	var left []int
	var right []int
	same := 0
	for _, num := range nums {
		if num < pivot{
			left = append(left, num)
		} else if num > pivot{
			right = append(right, num)
		} else {
			same += 1
		}
	}

	log.Println("---------------------")
	log.Printf("left is %v", left)
	log.Printf("pivot==nums[0] is %v", pivot)
	log.Printf("same is %v", same)
	log.Printf("right is %v", right)
	log.Println("---------------------")

	left = QuickSort(left)
	right = QuickSort(right)

	var result []int
	result = append(result, left...)
	result = append(result, pivot * same )
	result = append(result, right...)

	log.Println("=======================")
	log.Printf("left is %v", left)
	log.Printf("pivot * same is %v", pivot * same)
	log.Printf("right is %v", right)
	log.Printf("result is %v", result)
	log.Println("=======================")
	return result
}


func QuickSort2(nums []int) []int{
	if len(nums) <= 1{
		log.Printf("nums is %v", nums)
		return nums
	}
	pivot := nums[0]

	var left []int
	var right []int
	for _, num := range nums[1:]{
		if num <= pivot{
			left = append(left, num)
		}
		if num > pivot{
			right = append(right, num)
		}
	}

	log.Println("---------------------")
	log.Printf("left is %v", left)
	log.Printf("pivot==nums[0] is %v", pivot)
	log.Printf("right is %v", right)
	log.Println("---------------------")

	left = QuickSort2(left)
	right = QuickSort2(right)


	var result []int
	result = append(result, left...)
	result = append(result, pivot)
	result = append(result, right...)

	log.Println("=======================")
	log.Printf("left is %v", left)
	log.Printf("pivot is %v", pivot)
	log.Printf("right is %v", right)
	log.Printf("result is %v", result)
	log.Println("=======================")
	return result
}


