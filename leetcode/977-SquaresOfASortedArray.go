package main

// 977. Squares of a Sorted Array
// Given an integer array nums sorted in non-decreasing order, 
// return an array of the squares of each number sorted in non-decreasing order.

// Example 1:
// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Explanation: After squaring, the array becomes [16,1,0,9,100].
// After sorting, it becomes [0,1,9,16,100].

// Example 2:
// Input: nums = [-7,-3,2,3,11]
// Output: [4,9,9,49,121]
 
// Constraints:
//         1 <= nums.length <= 10^4
//         -10^4 <= nums[i] <= 10^4
//         nums is sorted in non-decreasing order.

// Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?

import "sort"
import "fmt"

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
    // 负数最大从 0 开始 i ,正数最大列表尾部开始 j
	for i, k, j := 0, len(nums)-1, len(ans)-1; i <= j; k-- {
        // 取平方最大的一个占坑
        s1, s2 := nums[i] * nums[i], nums[j] * nums[j]
		if s1 > s2 {
			ans[k] = s1
			i++
		} else {
			ans[k] = s2
			j--
		}
	}
	return ans
}

// 使用了 sort 达不到 O(n) 要求 
func sortedSquares1(nums []int) []int {
	for i, value := range nums {
		nums[i] = value * value
	}
	sort.Ints(nums)
	return nums
}

func main() {
    fmt.Println(sortedSquares([]int{ -4,-1,0,3,10 })) // [0,1,9,16,100]
    fmt.Println(sortedSquares([]int{ -7,-3,2,3,11 })) // [4,9,9,49,121]

    fmt.Println(sortedSquares1([]int{ -4,-1,0,3,10 })) // [0,1,9,16,100]
    fmt.Println(sortedSquares1([]int{ -7,-3,2,3,11 })) // [4,9,9,49,121]
}