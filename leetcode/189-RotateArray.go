package main

import "fmt"

/**
189. Rotate Array
Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:

	Input: nums = [1,2,3,4,5,6,7], k = 3
	Output: [5,6,7,1,2,3,4]
	Explanation:
	rotate 1 steps to the right: [7,1,2,3,4,5,6]
	rotate 2 steps to the right: [6,7,1,2,3,4,5]
	rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:

	Input: nums = [-1,-100,3,99], k = 2
	Output: [3,99,-1,-100]
	Explanation:
	rotate 1 steps to the right: [99,-1,-100,3]
	rotate 2 steps to the right: [3,99,-1,-100]


Constraints:

	1 <= nums.length <= 10^5
	-2^31 <= nums[i] <= 2^31 - 1
	0 <= k <= 10^5


Follow up:

	Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
	Could you do it in-place with O(1) extra space?
 */

//  时间复杂度 O(n)，空间复杂度 O(1)
// 末尾 k mod n 个元素移动至了数组头部，剩下的元素右移 k mod n 个位置至最尾部。确定了最终态以后再变换就很容易。
// 先将数组中所有元素从头到尾翻转一次，尾部的所有元素都到了头部，然后再将 [0,(k mod n) − 1] 区间内的元素翻转一次，
// 最后再将 [k mod n, n − 1] 区间内的元素翻转一次，
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums) // 先将数组中所有元素从头到尾翻转一次，尾部的所有元素都到了头部，
	fmt.Printf("reverse(nums) = %v\n",nums)
	reverse(nums[:k]) // 再将 [0,(k mod n) − 1] 区间内的元素翻转一次
	fmt.Printf("reverse(nums[:k]) = %v\n",nums)
	reverse(nums[k:]) // 再将 [k mod n, n − 1] 区间内的元素翻转一次
	fmt.Printf("reverse(nums[k:]) = %v\n",nums)
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[ n-1-i ] = a[ n-1-i ], a[i]
	}
}

//  时间复杂度 O(n)，空间复杂度 O(n)
//  使用一个额外的数组，先将原数组下标为 i 的元素移动到 (i+k) mod n 的位置，再将剩下的元素拷贝回来即可。
func rotate1(nums []int, k int) {
	newNums := make([]int, len(nums))
	// 先将原数组下标为 i 的元素移动到 (i+k) mod n 的位置
	for i, v := range nums {
		fmt.Printf("(i+k) %% len(nums) = %v\n", (i+k) % len(nums))
		newNums[(i+k) % len(nums)] = v
	}
	copy(nums, newNums) // 再将剩下的元素拷贝回来
}

func main() {
	nums := []int{ 1,2,3,4,5,6,7 }
	fmt.Printf("before: nums = %v\n",nums) // nums = [1,2,3,4,5,6,7], k = 3
	rotate(nums,3)
	//rotate1(nums,3)
	fmt.Printf("after: nums = %v\n",nums) // [5,6,7,1,2,3,4]

	nums1 := []int{ -1,-100,3,99 }
	fmt.Printf("before: nums = %v\n",nums1) // nums = [-1,-100,3,99], k = 2
	rotate(nums1,3)
	//rotate1(nums1,3)
	fmt.Printf("after: nums = %v\n",nums1) // [3,99,-1,-100]
}