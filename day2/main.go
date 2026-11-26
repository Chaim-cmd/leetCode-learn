package day2

import "sort"

//方法一：排序
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func ContainsDuplicate(nums []int) bool {
	//提问：为什么使用struct{} ：因为既然value 不承载信息，就该选一个“最不像信息”的类型，而且他还不占任何字节
	set := map[int]struct{}{}

	for _, v := range nums {
		if _, has := set[v]; has {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
