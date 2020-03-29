package InsertionSort

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
		nums[targetIndex + 1] = temp
	}
	return nums
}
