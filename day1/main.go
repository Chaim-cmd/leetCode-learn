package main

//两数之和
// 暴力写法
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func NewTwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, v := range nums {

		if j, ok := seen[target-v]; ok {
			return []int{j, i}
		}
		seen[v] = i
	}
	return nil
}
