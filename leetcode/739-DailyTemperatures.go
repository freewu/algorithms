package main

import "fmt"

// 739. Daily Temperatures
// Given an array of integers temperatures represents the daily temperatures, 
// return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
// If there is no future day for which this is possible, keep answer[i] == 0 instead.

// Example 1:
// Input: temperatures = [73,74,75,71,69,72,76,73]
// Output: [1,1,4,2,1,1,0,0]

// Example 2:
// Input: temperatures = [30,40,50,60]
// Output: [1,1,1,0]

// Example 3:
// Input: temperatures = [30,60,90]
// Output: [1,1,0]
 

// Constraints:

// 		1 <= temperatures.length <= 10^5
// 		30 <= temperatures[i] <= 100

// 暴力解法
func dailyTemperatures(T []int) []int {
	res, j := make([]int, len(T)), 0
	for i := 0; i < len(T); i++ {
		// 从 i + 1 开始找
		for j = i + 1; j < len(T); j++ {
			// 如果发现后面温度大于该天，计算之间相差天数
			if T[j] > T[i] {
				res[i] = j - i
				break
			}
		}
	}
	return res
}

// 使用 stack 
func dailyTemperatures1(T []int) []int {
	res := make([]int, len(T))
	var toCheck []int
	for i, t := range T {
		for len(toCheck) > 0 && T[toCheck[len(toCheck)-1]] < t {
			// 维护一个单调递减的单调栈
			idx := toCheck[len(toCheck)-1]
			res[idx] = i - idx
			toCheck = toCheck[:len(toCheck)-1]
		}
		toCheck = append(toCheck, i)
	}
	return res
}

func main() {
	fmt.Println(dailyTemperatures([]int{73,74,75,71,69,72,76,73})) // [1,1,4,2,1,1,0,0]
	fmt.Println(dailyTemperatures([]int{30,40,50,60})) // [1,1,1,0]
	fmt.Println(dailyTemperatures([]int{30,60,90})) // [1,1,0]

	fmt.Println(dailyTemperatures1([]int{73,74,75,71,69,72,76,73})) // [1,1,4,2,1,1,0,0]
	fmt.Println(dailyTemperatures1([]int{30,40,50,60})) // [1,1,1,0]
	fmt.Println(dailyTemperatures1([]int{30,60,90})) // [1,1,0]
}