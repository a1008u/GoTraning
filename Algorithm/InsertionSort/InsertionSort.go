package InsertionSort

import "log"

func InsertionSort(nums []int)  []int{
	for index := 1; index < len(nums); index++ {
		temp := nums[index]
		targetIndex := index - 1
		for{
			if!((targetIndex >= 0) && (nums[targetIndex] > temp)){
				break
			}
			nums[targetIndex + 1] = nums[targetIndex]
			targetIndex -= 1
		}
		log.Println("=======================")
		log.Printf("before nums is %v", nums)
		log.Printf("index is %v", index)
		log.Printf("targetIndex + 1 is %v", targetIndex + 1)
		log.Printf("nums[index] is %v", nums[index])
		log.Printf("nums[targetIndex + 1] is %v", nums[targetIndex + 1])
		nums[targetIndex + 1] = temp
		log.Printf("after nums is %v", nums)
		log.Println("=======================")
	}
	return nums
}
