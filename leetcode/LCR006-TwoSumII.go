package main

// LCR 006. 两数之和 II - 输入有序数组
// 给定一个已按照 升序排列  的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。
// 函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 0 开始计数 ，所以答案数组应当满足 0 <= answer[0] < answer[1] < numbers.length 。
// 假设数组中存在且只存在一对符合条件的数字，同时一个数字不能使用两次。

// 示例 1：
// 输入：numbers = [1,2,4,6,10], target = 8
// 输出：[1,3]
// 解释：2 与 6 之和等于目标数 8 。因此 index1 = 1, index2 = 3 。

// 示例 2：
// 输入：numbers = [2,3,4], target = 6
// 输出：[0,2]

// 示例 3：
// 输入：numbers = [-1,0], target = -1
// 输出：[0,1]
 
// 提示：
//         2 <= numbers.length <= 3 * 10^4
//         -1000 <= numbers[i] <= 1000
//         numbers 按 非递减顺序 排列
//         -1000 <= target <= 1000
//         仅存在一个有效答案

import "fmt"

// O(N^2) 暴力解法
func twoSum(numbers []int, target int) []int {
    for i := 0; i < len(numbers); i++ {
        for j := i + 1; j < len(numbers); j++ {
            if numbers[i] + numbers[j] == target {
                return []int{i,j}
            }
        }
    }
    return nil
}

// 使用 map 缓存之的计算结果
func twoSum1(numbers []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		another := target - numbers[i]
		if index, ok := m[another]; ok {
			return []int{ index , i }
		}
		m[numbers[i]] = i
	}
	return nil
}

// 利用数组有序的特性的解法
func twoSum2(numbers []int, target int) []int {
	i, j := 0, len(numbers) - 1
	for i < j { // 从外向内缩进
		val := numbers[i] + numbers[j]
		if val == target { // 如果刚好匹配则返回
			return []int{i, j}
		}
		if val < target { // 如果过小，说明开头需要向里走 ->
			i++
		} else { // 过大, 从尾部向头收缩 <-
			j--
		}
	}
	return nil
}

func main() {
    fmt.Println(twoSum([]int{1,2,4,6,10}, 8)) // [1,3]
    fmt.Println(twoSum([]int{2,3,4}, 6)) // [0,2]
    fmt.Println(twoSum([]int{-1,0}, -1)) // [0,1]

    fmt.Println(twoSum1([]int{1,2,4,6,10}, 8)) // [1,3]
    fmt.Println(twoSum1([]int{2,3,4}, 6)) // [0,2]
    fmt.Println(twoSum1([]int{-1,0}, -1)) // [0,1]

    fmt.Println(twoSum2([]int{1,2,4,6,10}, 8)) // [1,3]
    fmt.Println(twoSum2([]int{2,3,4}, 6)) // [0,2]
    fmt.Println(twoSum2([]int{-1,0}, -1)) // [0,1]
}