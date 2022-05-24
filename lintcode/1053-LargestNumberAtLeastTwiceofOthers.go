package main

/**
1053 · Largest Number At Least Twice of Others

# Description
In a given integer array nums, there is always exactly one largest element.
Find whether the largest element in the array is at least twice as much as every other number in the array.
If it is, return the index of the largest element, otherwise return -1.

nums will have a length in the range [1, 50].
Every nums[i] will be an integer in the range [0, 99].

Example 1:

	Input: nums = [3, 6, 1, 0]
	Output: 1
	Explanation: 6 is the largest integer, and for every other number in the array x,
	6 is more than twice as big as x.  The index of value 6 is 1, so we return 1.

Example 2:

	Input: nums = [1, 2, 3, 4]
	Output: -1
	Explanation: 4 isn't at least as big as twice the value of 3, so we return -1.

*/

import "fmt"

/**
 * @param nums: a integer array
 * @return: the index of the largest element
 */
 func DominantIndex(nums []int) int {
    // Write your code here
	max := 0 // 保存最大值的下标
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[max] {
			// 判断最大元素是否至少是数组中每个其他数字的两倍
			if nums[i] < 2 * nums[max] {
				return -1
			}
			max = i
		} else {
			// 判断最大元素是否至少是数组中每个其他数字的两倍
			if nums[max] < 2 * nums[i] {
				return -1
			}
		}	
	}
	return max
}

func main() {
	fmt.Printf("DominantIndex([]int{ 3, 6, 1, 0 }) = %v\n",DominantIndex([]int{ 3, 6, 1, 0 })) // 1
	fmt.Printf("DominantIndex([]int{ 1, 2, 3, 4 }) = %v\n",DominantIndex([]int{ 1, 2, 3, 4 })) // -1
}