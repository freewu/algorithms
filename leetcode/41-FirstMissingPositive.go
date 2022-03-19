package main

import "fmt"

/**
41. First Missing Positive
Given an unsorted integer array, find the smallest missing positive integer.

Note:
	Your algorithm should run in O(n) time and uses constant extra space.

Example 1:

	Input: [1,2,0]
	Output: 3

Example 2:

	Input: [3,4,-1,1]
	Output: 2

Example 3:

	Input: [7,8,9,11,12]
	Output: 1

解题思路:
	数组都装到 map 中
	然后 i 循环从 1 开始，依次比对 map 中是否存在 i，
	只要不存在 i 就立即返回结果 i
 */

func firstMissingPositive(nums []int) int {
	numMap := make(map[int]int, len(nums))
	for _, v := range nums { // 先把数组保存到 map 中
		numMap[v] = v
	}
	for index := 1; index < len(nums)+1; index++ {
		if _, ok := numMap[index]; !ok { // 如果数组中不存在 说明 index 是需要返回的值
			return index
		}
	}
	return len(nums) + 1
}

// best solution 思路一样就是做了 非合理值的判断
func firstMissingPositiveBest(nums []int) int {
	founds := make([]bool, 5*100000)
	for _, num := range nums {
		if num > 0 && num <= 5*100000{
			founds[num-1] = true
		}
	}
	for pos, found := range founds {
		if !found {
			return pos+1
		}
	}
	return 5*100000+1
}

func main() {
	fmt.Printf("firstMissingPositive([]int{1,2,0} = %v\n",firstMissingPositive([]int{1,2,0})) // 3
	fmt.Printf("firstMissingPositive([]int{3,4,-1,1} = %v\n",firstMissingPositive([]int{3,4,-1,1})) // 2
	fmt.Printf("firstMissingPositive([]int{7,8,9,11,12} = %v\n",firstMissingPositive([]int{7,8,9,11,12})) // 1
	fmt.Printf("firstMissingPositiveBest([]int{1,2,0} = %v\n",firstMissingPositiveBest([]int{1,2,0})) // 3
	fmt.Printf("firstMissingPositiveBest([]int{3,4,-1,1} = %v\n",firstMissingPositiveBest([]int{3,4,-1,1})) // 2
	fmt.Printf("firstMissingPositiveBest([]int{7,8,9,11,12} = %v\n",firstMissingPositiveBest([]int{7,8,9,11,12})) // 1
}