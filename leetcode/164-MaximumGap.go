package main

import (
	"fmt"
)

/**
164. Maximum Gap
Given an integer array nums, return the maximum difference between two successive elements in its sorted form.
If the array contains less than two elements, return 0.
You must write an algorithm that runs in linear time and uses linear extra space.

Example 1:

	Input: nums = [3,6,9,1]
	Output: 3
	Explanation: The sorted form of the array is [1,3,6,9], either (3,6) or (6,9) has the maximum difference 3.

Example 2:

	Input: nums = [10]
	Output: 0
	Explanation: The array contains less than 2 elements, therefore return 0.

Constraints:

	1 <= nums.length <= 10^5
	0 <= nums[i] <= 10^9
 */

//
//func maximumGap2(nums []int) int {
//	if len(nums) < 2 {
//		return 0
//	}
//	res := 0
//	for i := 0; i < len(nums) - 1; i++ {
//		fmt.Printf("i + 1 = %v,i = %v \n",i+1, i)
//		fmt.Printf("nums[i+1]= %v,nums[i] = %v \n",nums[i+1], nums[i])
//		val := abs(nums[i+1] - nums[i])
//		if val > res { // 循环计算出差值
//			res = val
//		}
//	}
//	return res
//}
//
//func abs(a int) int {
//	if a > 0 {
//		return a
//	}
//	return -a
//}


// 解法1 快排
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	// 先排序
	quickSort(nums, 0, len(nums)-1) // successive elements 这也是为啥要先排序的原因
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		if (nums[i+1] - nums[i]) > res { // 循环计算出差值
			res = nums[i+1] - nums[i]
		}
	}
	return res
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

// 解法2 基数排序
func maximumGap1(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	// m is the maximal number in nums
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		m = max(m, nums[i])
	}
	exp := 1 // 1, 10, 100, 1000 ...
	R := 10  // 10 digits
	aux := make([]int, len(nums))
	for (m / exp) > 0 { // Go through all digits from LSB to MSB
		count := make([]int, R)
		for i := 0; i < len(nums); i++ {
			count[(nums[i]/exp)%10]++
		}
		for i := 1; i < len(count); i++ {
			count[i] += count[i-1]
		}
		for i := len(nums) - 1; i >= 0; i-- {
			tmp := count[(nums[i]/exp)%10]
			tmp--
			aux[tmp] = nums[i]
			count[(nums[i]/exp)%10] = tmp
		}
		for i := 0; i < len(nums); i++ {
			nums[i] = aux[i]
		}
		exp *= 10
	}
	maxValue := 0
	for i := 1; i < len(aux); i++ {
		maxValue = max(maxValue, aux[i]-aux[i-1])
	}
	return maxValue
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("maximumGap([]int{ 3,6,9,1 }) = %v\n",maximumGap([]int{ 3,6,9,1 })) // 3
	fmt.Printf("maximumGap([]int{ 10 }) = %v\n",maximumGap([]int{ 10 })) // 0

	fmt.Printf("maximumGap1([]int{ 3,6,9,1 }) = %v\n",maximumGap1([]int{ 3,6,9,1 })) // 3
	fmt.Printf("maximumGap1([]int{ 10 }) = %v\n",maximumGap1([]int{ 10 })) // 0

	fmt.Printf("maximumGap2([]int{ 3,6,9,1 }) = %v\n",maximumGap2([]int{ 3,6,9,1 })) // 3
	fmt.Printf("maximumGap2([]int{ 10 }) = %v\n",maximumGap2([]int{ 10 })) // 0
}