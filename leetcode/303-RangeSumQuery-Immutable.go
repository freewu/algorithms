package main

// 303. Range Sum Query - Immutable
// Given an integer array nums, handle multiple queries of the following type:
// 	Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.

// Implement the NumArray class:
// 	NumArray(int[] nums) Initializes the object with the integer array nums.
// 	int sumRange(int left, int right) Returns the sum of the elements of nums between indices left
// 	and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).

// Example 1:
// 	Input
// 		["NumArray", "sumRange", "sumRange", "sumRange"]
// 		[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
// 	Output
// 		[null, 1, -1, -3]
// 	Explanation
// 		NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
// 		numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
// 		numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
// 		numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3

// Constraints:
// 	1 <= nums.length <= 10^4
// 	-10^5 <= nums[i] <= 10^5
// 	0 <= left <= right < nums.length
// 	At most 104 calls will be made to sumRange.
	

// 给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

import "fmt"

// NumArray define
type NumArray struct {
	prefixSum []int
}

// Constructor define
func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1] // 累加好
	}
	return NumArray{prefixSum: nums}
}

// SumRange define
func (this *NumArray) SumRange(i int, j int) int {
	if i > 0 {
		return this.prefixSum[j] - this.prefixSum[i-1]
	}
	return this.prefixSum[j]
}

func main() {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1});
	fmt.Println(numArray.SumRange(0, 2)); // return (-2) + 0 + 3 = 1
	fmt.Println(numArray.SumRange(2, 5)); // return 3 + (-5) + 2 + (-1) = -1
	fmt.Println(numArray.SumRange(0, 5)); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3
}