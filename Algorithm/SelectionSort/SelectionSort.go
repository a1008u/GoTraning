package SelectionSort

import "log"

func Selectionsort(nums []int)  []int{
	for index, _ := range nums{
		minIdx := index
		for  cki := index+1; cki < len(nums); cki++ {
			if nums[minIdx] > nums[cki]{
				minIdx = cki
			}
		}
		log.Println("=======================")
		log.Printf("before nums is %v", nums)
		log.Printf("index is %v", index)
		log.Printf("minIdx is %v", minIdx)
		log.Printf("nums[minIdx] is %v", nums[index])
		log.Printf("nums[index] is %v", nums[minIdx])
		nums[minIdx], nums[index] = nums[index], nums[minIdx]
		log.Printf("after nums is %v", nums)
		log.Println("=======================")
	}
	return nums
}


func Selectionsort2(nums []int)  []int{
	for i, _ := range nums{
		minIdx := i
		for cki := i; cki < len(nums); cki++ {
			if nums[minIdx] > nums[cki]{
				minIdx = cki
			}
		}
		log.Println("=======================")
		log.Printf("before nums is %v", nums)
		log.Printf("i is %v", i)
		log.Printf("minIdx is %v", minIdx)
		log.Printf("nums[minIdx] is %v", nums[i])
		log.Printf("nums[index] is %v", nums[minIdx])
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
		log.Printf("after nums is %v", nums)
		log.Println("=======================")
	}
	return nums
}
