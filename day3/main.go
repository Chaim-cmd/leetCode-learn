package main

import "sort"

//多数元素
//方法一：排序
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

//方法二：摩尔
func MajorityElement(nums []int) int {
	count, num := 0, 0
	for _, v := range nums {
		if count == 0 {
			num = v
		}
		if v == num {
			count++
		} else {
			count--
		}
	}
	return num
}
