package main

import "fmt"

/**
229. Majority Element II

Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Example 1:

	Input: nums = [3,2,3]
	Output: [3]

Example 2:

	Input: nums = [1]
	Output: [1]

Example 3:

	Input: nums = [1,2]
	Output: [1,2]

Constraints:

	1 <= nums.length <= 5 * 10^4
	-10^9 <= nums[i] <= 10^9


Follow up: Could you solve the problem in linear time and in O(1) space?

给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)
 */

// 时间复杂度 O(n) 空间复杂度 O(n)
func majorityElement(nums []int) []int {
	result, m := make([]int, 0), make(map[int]int)
	for _, val := range nums {
		if v, ok := m[val]; ok {
			m[val] = v + 1
		} else {
			m[val] = 1
		}
	}
	for k, v := range m { // 循环 map 把数据值 大于 数组长度 1 / 3 的 添加到返回集中
		if v > len(nums) / 3 {
			result = append(result, k)
		}
	}
	return result
}

// 时间复杂度 O(n) 空间复杂度 O(1)
func majorityElement1(nums []int) []int {
	// 超过 1 / 3 最多存在两个相关的数据
	// since we are checking if a num appears more than 1/3 of the time
	// it is only possible to have at most 2 nums (>1/3 + >1/3 = >2/3)
	count1, count2, candidate1, candidate2 := 0, 0, 0, 1
	// Select Candidates
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 <= 0 {
			// We have a bad first candidate, replace!
			candidate1, count1 = num, 1
		} else if count2 <= 0 {
			// We have a bad second candidate, replace!
			candidate2, count2 = num, 1
		} else {
			// Both candidates suck, boo!
			count1--
			count2--
		}
	}
	// Recount! // 得到 这两个值后 重新统计两个值的出现次数
	count1, count2 = 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}
	length := len(nums)
	if count1 > length/3 && count2 > length/3 {
		return []int{candidate1, candidate2}
	}
	if count1 > length/3 {
		return []int{candidate1}
	}
	if count2 > length/3 {
		return []int{candidate2}
	}
	return []int{}
}

// best sulotion
// 就是使用 map的解法 把需要计算的都算好 不要在循环里计算
func majorityElementBest(nums []int) []int {
	minFreq := len(nums) / 3
	numsByFreq := make(map[int]int)
	for _, num := range nums {
		if _, numExists := numsByFreq[num]; !numExists {
			numsByFreq[num] = 1
		} else {
			numsByFreq[num]++
		}
	}
	var res []int
	for num, freq := range numsByFreq {
		if freq > minFreq {
			res = append(res, num)
		}
	}
	return res
}

func main() {
	fmt.Printf("majorityElement([]int{ 3,2,3 } = %v\n",majorityElement([]int{ 3,2,3 })) // [3]
	fmt.Printf("majorityElement([]int{ 1 } = %v\n",majorityElement([]int{ 1 })) // [1]
	fmt.Printf("majorityElement([]int{ 1,2 } = %v\n",majorityElement([]int{ 1,2 })) // [1,2]

	fmt.Printf("majorityElement1([]int{ 3,2,3 } = %v\n",majorityElement1([]int{ 3,2,3 })) // [3]
	fmt.Printf("majorityElement1([]int{ 1 } = %v\n",majorityElement1([]int{ 1 })) // [1]
	fmt.Printf("majorityElement1([]int{ 1,2 } = %v\n",majorityElement1([]int{ 1,2 })) // [1,2]

	fmt.Printf("majorityElementBest([]int{ 3,2,3 } = %v\n",majorityElementBest([]int{ 3,2,3 })) // [3]
	fmt.Printf("majorityElementBest([]int{ 1 } = %v\n",majorityElementBest([]int{ 1 })) // [1]
	fmt.Printf("majorityElementBest([]int{ 1,2 } = %v\n",majorityElementBest([]int{ 1,2 })) // [1,2]
}
