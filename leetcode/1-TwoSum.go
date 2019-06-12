package main

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

import "fmt"

// O(n^2)
func twoSum(nums []int, target int) []int {
	var l = len(nums)

	if l < 2 {
		return nil
	}
	// 如果数据只有两个值的情况
	if 2 == l && ((nums[0] + nums[1]) != target) {
		return nil
	}

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			// 冒泡的方式遍历
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// best speed solution O(n)
func twoSum1(nums []int, target int) []int {
	m := map[int]int{}
	r := []int{}

	for i, num := range nums {
		// 计算出差值
		compliment := target - num

		// 如果存在 差值的返回 该值所在list的位置(第一次肯定的是失败的)
		if _, ok := m[compliment]; ok {
			r = []int{m[compliment], i}
			break
		} else {
			// 写入map key为值 value为list所在的位置
			m[num] = i
		}
	}

	return r
}

func main() {
	var list = []int{2, 7, 11, 15}
	fmt.Println(twoSum(list, 9))
	fmt.Println(twoSum1(list, 9))
}
