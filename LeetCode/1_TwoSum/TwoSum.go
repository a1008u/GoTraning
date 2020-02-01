package TwoSum

// main method
//func main() {
//	twoSumResult := TwoSum();
//	log.Printf("result is %v", twoSumResult)
//}

//THE PURPOSE
//　Approach 1: Brute Force
//INTENTION
//
//arg:
//
//return:
//
func TwoSum(nums []int, target int) []int{
	var (
		resultIndex1 int
		resultIndex2 int
	)
	for index := range nums {
		for i := 1 + index; i < len(nums); i++ {
			if nums[index] + nums[i] == target {
				resultIndex1 = index
				resultIndex2 = i
				break
			}
		}
	}
	return []int{resultIndex1, resultIndex2}
}

//THE PURPOSE
//  Approach 2: Two-pass Hash Table
//INTENTION
//　Mapを利用して、取得したいindexを素早く取得する
//arg:
//
//return:
//
func TwoSumMap(nums []int, target int) []int{
	var (
		resultIndex1 int
		resultIndex2 int
	)
	m  := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i:= 0; i < len(m); i++ {
		targetNum := target - nums[i]
		if value, ok := m[targetNum]; ok {
			if value != i {
				resultIndex1 = i
				resultIndex2 = m[targetNum]
				break
			}
		}
	}
	return []int{resultIndex1, resultIndex2}
}

//THE PURPOSE
//	Approach 3: One-pass Hash Table
//INTENTION
//
//arg:
//
//return:
//
func TwoSumOnePass(nums []int, target int) []int{
	var (
		resultIndex1 int
		resultIndex2 int
	)
	m  := map[int]int{}
	for i := 0; i < len(nums); i++ {
		targetNum := target - nums[i]
		if valueKey, ok := m[targetNum]; ok {
			if valueKey != i {
				resultIndex1 = valueKey
				resultIndex2 = i
				break
			}
		}
		m[nums[i]] = i
	}
	return []int{resultIndex1, resultIndex2}
}
