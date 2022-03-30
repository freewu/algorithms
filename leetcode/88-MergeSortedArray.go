package main

import "fmt"

/**
88. Merge Sorted Array
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

Constraints:

	nums1.length == m + n
	nums2.length == n
	0 <= m, n <= 200
	1 <= m + n <= 200
	-10^9 <= nums1[i], nums2[j] <= 10^9

Example 1:

	Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	Output: [1,2,2,3,5,6]
	Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
	The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

Example 2:

	Input: nums1 = [1], m = 1, nums2 = [], n = 0
	Output: [1]
	Explanation: The arrays we are merging are [1] and [].
	The result of the merge is [1].

Example 3:

	Input: nums1 = [0], m = 0, nums2 = [1], n = 1
	Output: [1]
	Explanation: The arrays we are merging are [] and [1].
	The result of the merge is [1].
	Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

 */

func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] <= nums2[n-1] {
			nums1[p-1] = nums2[n-1]
			n--
		} else {
			nums1[p-1] = nums1[m-1]
			m--
		}
	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}

// best solution
func mergeBest(nums1 []int, m int, nums2 []int, n int) {
	mergeIdx := len(nums1)- 1
	num1Idx := m - 1
	num2Idx := n - 1
	for mergeIdx >= 0 && num2Idx >= 0 {
		if (num1Idx < 0 ) ||  (nums1[num1Idx] < nums2[num2Idx]) {
			nums1[mergeIdx] = nums2[num2Idx]
			num2Idx--
			mergeIdx--
		} else {
			nums1[mergeIdx] = nums1[num1Idx]
			num1Idx--
			mergeIdx--
		}
	}
}

func main() {
	num1 := []int{1,2,3,0,0,0}
	num2 := []int{2,5,6}
	fmt.Printf("before num1 = %v\n",num1) // [1 2 3 0 0 0]
	fmt.Printf("before num2 = %v\n",num2)
	//merge(num1,3,num2,3)
	mergeBest(num1,3,num2,3)
	fmt.Printf("after num1 = %v\n",num1) // [1 2 2 3 5 6]
	fmt.Printf("after num2 = %v\n",num2)
}
