package SelectionSort

func Selectionsort(nums []int)  []int{
	for index, _ := range nums{
		minIdx := index
		for  cki := index+1; cki < len(nums); cki++ {
			if nums[minIdx] > nums[cki]{
				minIdx = cki
			}
		}
		nums[minIdx], nums[index] = nums[index], nums[minIdx]
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
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums
}
