package main

/*
1. Two Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

# 解题思路
最优的做法时间复杂度是 O(n)。
1 声明一个 map[值] = 下标
2 顺序扫描数组,
3 计算出 目标值 - 当前值 的 差值
4 查找 差值 是否能在 map 里找到
5 能找到 返回 [ 当前下标 , map[差值] ]
6 没找到  map[当前值] = 当前下票
7 回到第3步
*/

import "fmt"

// O(n^2) 自己的解法
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
func twoSumBest(nums []int, target int) []int {
	m := make(map[int]int, len(nums)) // 优先设定好map长度,避免扩容产生性能波动
	for i, num := range nums {
		if idx, ok := m[target - num]; ok { // 少个中间变量
			return []int{idx, i}
		}
		m[num] = i
	}
	return []int{}
}

func twoSum1(nums []int, target int) []int {
	res := make(map[int]int) // 使用一个map 来存 map[值]= 位置
	for i := 0; i < len(nums); i++ {
		another := target - nums[i] // 得到 当前数的差值
		if _, ok := res[another]; ok { // 差值存在
			return []int{res[another], i}
		}
		res[nums[i]] = i //
	}
	return nil
}

func main() {
	fmt.Printf("twoSum([]int{ 2, 7, 11, 15 }, 9) = %v\n",twoSum([]int{ 2, 7, 11, 15 }, 9))
	fmt.Printf("twoSumBest([]int{ 2, 7, 11, 15 }, 9) = %v\n",twoSumBest([]int{ 2, 7, 11, 15 }, 9))
	fmt.Printf("twoSum1([]int{ 2, 7, 11, 15 }, 9) = %v\n",twoSum1([]int{ 2, 7, 11, 15 }, 9))
	fmt.Printf("twoSum1([]int{ 2, 4, 4, 11, 15 }, 8) = %v\n",twoSum1([]int{ 2, 4, 4, 11, 15 }, 8))
}
