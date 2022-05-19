package main

import "fmt"

/**
220. Contains Duplicate III
Given an integer array nums and two integers k and t,
return true if there are two distinct indices i and j in the array such that abs(nums[i] - nums[j]) <= t and abs(i - j) <= k.

Example 1:

	Input: nums = [1,2,3,1], k = 3, t = 0
	Output: true

Example 2:

	Input: nums = [1,0,1,1], k = 1, t = 2
	Output: true

Example 3:

	Input: nums = [1,5,9,1,5,9], k = 2, t = 3
	Output: false

Constraints:

	1 <= nums.length <= 2 * 10^4
	-2^31 <= nums[i] <= 2^31 - 1
	0 <= k <= 10^4
	0 <= t <= 2^31 - 1

在 num 中能否找到一组 i 和 j，使得 num[i] 和 num[j] 的绝对差值最大为 t，并且 i 和 j 之前的绝对差值最大为 k。
 */

// 解法一 桶排序
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k <= 0 || t < 0 || len(nums) < 2 {
		return false
	}
	buckets := map[int]int{}
	for i := 0; i < len(nums); i++ {
		// Get the ID of the bucket from element value nums[i] and bucket width t + 1
		key := nums[i] / (t + 1)
		// -7/9 = 0, but need -7/9 = -1
		if nums[i] < 0 {
			key--
		}
		if _, ok := buckets[key]; ok {
			return true
		}
		// check the lower bucket, and have to check the value too
		if v, ok := buckets[key-1]; ok && nums[i]-v <= t {
			return true
		}
		// check the upper bucket, and have to check the value too
		if v, ok := buckets[key+1]; ok && v-nums[i] <= t {
			return true
		}
		// maintain k size of window
		if len(buckets) >= k {
			delete(buckets, nums[i-k]/(t+1))
		}
		buckets[key] = nums[i]
	}
	return false
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// 解法二 滑动窗口 + 剪枝
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	if k <= 0 {
		return false
	}
	for i := 0; i < n; i++ {
		count := 0
		for j := i + 1; j < n && count < k; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
			count++
		}
	}
	return false
}

func main() {
	fmt.Printf("containsNearbyAlmostDuplicate([]int{1,2,3,1},3,0) = %v\n",containsNearbyAlmostDuplicate([]int{1,2,3,1},3,0)) // true
	fmt.Printf("containsNearbyAlmostDuplicate([]int{1,0,1,1},1,2) = %v\n",containsNearbyAlmostDuplicate([]int{1,0,1,1},1,2)) // true
	fmt.Printf("containsNearbyAlmostDuplicate([]int{1,5,9,1,5,9},2,3) = %v\n",containsNearbyAlmostDuplicate([]int{1,5,9,1,5,9},2,3)) // false

	fmt.Printf("containsNearbyAlmostDuplicate1([]int{1,2,3,1},3,0) = %v\n",containsNearbyAlmostDuplicate1([]int{1,2,3,1},3,0)) // true
	fmt.Printf("containsNearbyAlmostDuplicate1([]int{1,0,1,1},1,2) = %v\n",containsNearbyAlmostDuplicate1([]int{1,0,1,1},1,2)) // true
	fmt.Printf("containsNearbyAlmostDuplicate1([]int{1,5,9,1,5,9},2,3) = %v\n",containsNearbyAlmostDuplicate1([]int{1,5,9,1,5,9},2,3)) // false
}