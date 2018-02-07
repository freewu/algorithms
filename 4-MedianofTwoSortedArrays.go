package main

import (
	"fmt"
)

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l1 = len(nums1)
	var l2 = len(nums2)
	var l3 = l1 + l2
	var t = make([]int, l3)
	var step = 0
	var s1 = 0
	var s2 = 0

	// 合并成一个数组
	for {
		if s1 == l1 || s2 == l2 {
			break
		}
		if nums1[s1] < nums2[s2] {
			t[step] = nums1[s1]
			s1++
		} else {
			t[step] = nums2[s2]
			s2++
		}
		step++
	}
	for {
		if s1 == l1 {
			break
		}
		t[step] = nums1[s1]
		step++
		s1++
	}
	for {
		if s2 == l2 {
			break
		}
		t[step] = nums2[s2]
		step++
		s2++
	}
	// 计算中位数
	if 0 == l3%2 {
		return (float64(t[l3/2]) + float64(t[l3/2-1])) / 2.0
	} else {
		return float64(t[l3/2])
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))    // 2.0
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // 2.0
}
