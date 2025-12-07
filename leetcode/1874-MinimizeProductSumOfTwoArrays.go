package main

// 1874. Minimize Product Sum of Two Arrays
// The product sum of two equal-length arrays a and b is equal to the sum of a[i] * b[i] for all 0 <= i < a.length (0-indexed).
//     For example, if a = [1,2,3,4] and b = [5,2,3,1], the product sum would be 1*5 + 2*2 + 3*3 + 4*1 = 22.

// Given two arrays nums1 and nums2 of length n, 
// return the minimum product sum if you are allowed to rearrange the order of the elements in nums1. 

// Example 1:
// Input: nums1 = [5,3,4,2], nums2 = [4,2,2,5]
// Output: 40
// Explanation: We can rearrange nums1 to become [3,5,4,2]. The product sum of [3,5,4,2] and [4,2,2,5] is 3*4 + 5*2 + 4*2 + 2*5 = 40.

// Example 2:
// Input: nums1 = [2,1,4,5,7], nums2 = [3,2,4,8,6]
// Output: 65
// Explanation: We can rearrange nums1 to become [5,7,4,1,2]. The product sum of [5,7,4,1,2] and [3,2,4,8,6] is 5*3 + 7*2 + 4*4 + 1*8 + 2*6 = 65.

// Constraints:
//     n == nums1.length == nums2.length
//     1 <= n <= 10^5
//     1 <= nums1[i], nums2[i] <= 100

import "fmt"
import "sort"

func minProductSum(nums1 []int, nums2 []int) int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    res, n := 0, len(nums1)
    for i, j := 0, n - 1; i < n; i++ {
        res += nums1[i] * nums2[j]
        j--
    }
    return res
}

func minProductSum1(nums1 []int, nums2 []int) int {
    countingSortHelper := func(nums []int, limit int) { // 计数排序
        counts := make([]int, 1 + limit) 
        for _, v := range nums {
            counts[v] += 1 
        }
        index := 0 
        for i := 1; i <= limit; i += 1 {
            for counts[i] > 0 {
                nums[index] = i
                index += 1 
                counts[i]--
            }
        }
    }
    res, n := 0, len(nums1) 
    countingSortHelper(nums1, 100)
    countingSortHelper(nums2, 100)
    for i, v := range nums1 {
        res += v * nums2[n - 1 - i] 
    }
    return res 
}

func main() {
    // Example 1:
    // Input: nums1 = [5,3,4,2], nums2 = [4,2,2,5]
    // Output: 40
    // Explanation: We can rearrange nums1 to become [3,5,4,2]. The product sum of [3,5,4,2] and [4,2,2,5] is 3*4 + 5*2 + 4*2 + 2*5 = 40.
    fmt.Println(minProductSum([]int{5,3,4,2}, []int{4,2,2,5})) // 40
    // Example 2:
    // Input: nums1 = [2,1,4,5,7], nums2 = [3,2,4,8,6]
    // Output: 65
    // Explanation: We can rearrange nums1 to become [5,7,4,1,2]. The product sum of [5,7,4,1,2] and [3,2,4,8,6] is 5*3 + 7*2 + 4*4 + 1*8 + 2*6 = 65.
    fmt.Println(minProductSum([]int{2,1,4,5,7}, []int{3,2,4,8,6})) // 65

    fmt.Println(minProductSum([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 165
    fmt.Println(minProductSum([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 165
    fmt.Println(minProductSum([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 165
    fmt.Println(minProductSum([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 165

    fmt.Println(minProductSum1([]int{5,3,4,2}, []int{4,2,2,5})) // 40
    fmt.Println(minProductSum1([]int{2,1,4,5,7}, []int{3,2,4,8,6})) // 65
    fmt.Println(minProductSum1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 165
    fmt.Println(minProductSum1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 165
    fmt.Println(minProductSum1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 165
    fmt.Println(minProductSum1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 165
}