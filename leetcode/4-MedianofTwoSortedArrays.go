package main

// 4. Median of Two Sorted Arrays
// Given two sorted arrays nums1 and nums2 of size m and n respectively, 
// return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).

// Example 1:
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.

// Example 2:
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

// Constraints:
//     nums1.length == m
//     nums2.length == n
//     0 <= m <= 1000
//     0 <= n <= 1000
//     1 <= m + n <= 2000
//     -10^6 <= nums1[i], nums2[i] <= 10^6

import "fmt"
import "math"

// by self
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    l1, l2 := len(nums1),len(nums2)
    l3 := l1 + l2
    t := make([]int, l3)
    step, s1, s2 := 0, 0, 0
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
    if 0 == l3 % 2 {
        return (float64(t[l3 / 2]) + float64(t[l3 / 2 - 1])) / 2.0
    } else {
        return float64(t[ l3 / 2])
    }
}

// best solution
func findMedianSortedArraysBest(a []int, b []int) float64 {
    m,n := len(a), len(b)
    if m > n {
        a, b, m,n = b,a,n,m
    }
    mid := (m + n - 1) / 2 // 两个数组的中位位置
    l,r := 0,m-1
    for l <= r {
        midA := l + (r-l)/2
        midB := mid - midA
        if a[midA] < b[midB] {
            l = midA + 1
        } else {
            r = midA - 1
        }
    }

    left := 0
    if l > 0 {
        left = a[l-1]
    } else {
        left = math.MinInt32
    }

    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    if mid-l >= 0 {
        left = max(left, b[mid-l])
    }

    if (m+n)%2 ==1 {
        return float64(left)
    }

    right := 0  //right median, max of a[l] and b[mid-l+1]
    if l < m {
        right= a[l]
    } else {
        right= math.MaxInt32
    }
    if mid - l + 1 < n {
        right = min(right,b[mid-l+1])
    }
    return float64(left+right) / 2.0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))    // 2.0
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // 2.5

	fmt.Println(findMedianSortedArraysBest([]int{1, 3}, []int{2}))    // 2.0
	fmt.Println(findMedianSortedArraysBest([]int{1, 2}, []int{3, 4})) // 2.5
}
