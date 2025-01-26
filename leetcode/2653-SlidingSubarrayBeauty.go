package main

// 2653. Sliding Subarray Beauty
// Given an integer array nums containing n integers, find the beauty of each subarray of size k.

// The beauty of a subarray is the xth smallest integer in the subarray if it is negative, or 0 if there are fewer than x negative integers.

// Return an integer array containing n - k + 1 integers, which denote the beauty of the subarrays in order from the first index in the array.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,-1,-3,-2,3], k = 3, x = 2
// Output: [-1,-2,-2]
// Explanation: There are 3 subarrays with size k = 3. 
// The first subarray is [1, -1, -3] and the 2nd smallest negative integer is -1. 
// The second subarray is [-1, -3, -2] and the 2nd smallest negative integer is -2. 
// The third subarray is [-3, -2, 3] and the 2nd smallest negative integer is -2.

// Example 2:
// Input: nums = [-1,-2,-3,-4,-5], k = 2, x = 2
// Output: [-1,-2,-3,-4]
// Explanation: There are 4 subarrays with size k = 2.
// For [-1, -2], the 2nd smallest negative integer is -1.
// For [-2, -3], the 2nd smallest negative integer is -2.
// For [-3, -4], the 2nd smallest negative integer is -3.
// For [-4, -5], the 2nd smallest negative integer is -4. 

// Example 3:
// Input: nums = [-3,1,2,-3,0,-3], k = 2, x = 1
// Output: [-3,0,-3,-3,-3]
// Explanation: There are 5 subarrays with size k = 2.
// For [-3, 1], the 1st smallest negative integer is -3.
// For [1, 2], there is no negative integer so the beauty is 0.
// For [2, -3], the 1st smallest negative integer is -3.
// For [-3, 0], the 1st smallest negative integer is -3.
// For [0, -3], the 1st smallest negative integer is -3.

// Constraints:
//     n == nums.length 
//     1 <= n <= 10^5
//     1 <= k <= n
//     1 <= x <= k 
//     -50 <= nums[i] <= 50 

import "fmt"

func getSubarrayBeauty(nums []int, k int, x int) []int {
    count := make(map[int]int, 101) // -50 -- 50 
    for i := 0; i < k - 1; i++ {
        count[nums[i]]++
    }
    res := make([]int, 0, len(nums) - k + 1)
    for i := k-1; i < len(nums); i++ {
        count[nums[i]]++
        if i - k >= 0 {
            count[nums[i-k]]--
        }
        negative := 0
        for i := -50; i <= 50; i++ {
            if i >= 0 {
                res = append(res, 0)
                break
            }
            negative += count[i]
            if negative >= x {
                res = append(res, i)
                break
            }
        }
    }
    return res
}

func getSubarrayBeauty1(nums []int, k int, x int) []int {
    n := len(nums)
    count := [101]int{}
    for _, v := range nums[:k] {
        count[v + 50]++
    }
    res := make([]int, n - k + 1)
    helper := func(v int) int {
        sum := 0
        for i := 0; i < 50; i++ {
            sum += count[i]
            if sum >= v {
                return i - 50
            }
        }
        return 0
    }
    res[0] = helper(x)
    for i, j := k, 1; i < n; i, j = i + 1, j + 1 {
        count[nums[i] + 50]++
        count[nums[i - k] + 50]--
        res[j] = helper(x)
    }
    return res
}

func getSubarrayBeauty2(nums []int, k int, x int) []int {
    if len(nums) == 0 || k <= 0 { return []int{} }
    half := 50
    res, count := make([]int, len(nums) - k + 1), make([]int, half * 2 + 1)
    for i := 0; i < k - 1; i++ {
        count[nums[i] + half]++
    }
    for i, j := range nums[k-1:] {
        count[j + half]++
        c := x
        for m, n := range count[:half] {
            c -= n
            if c <= 0  {
                res[i] = m - half
                break
            }
        }
        count[nums[i] + half]--
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,-1,-3,-2,3], k = 3, x = 2
    // Output: [-1,-2,-2]
    // Explanation: There are 3 subarrays with size k = 3. 
    // The first subarray is [1, -1, -3] and the 2nd smallest negative integer is -1. 
    // The second subarray is [-1, -3, -2] and the 2nd smallest negative integer is -2. 
    // The third subarray is [-3, -2, 3] and the 2nd smallest negative integer is -2.
    fmt.Println(getSubarrayBeauty([]int{1,-1,-3,-2,3}, 3, 2)) // [-1,-2,-2]
    // Example 2:
    // Input: nums = [-1,-2,-3,-4,-5], k = 2, x = 2
    // Output: [-1,-2,-3,-4]
    // Explanation: There are 4 subarrays with size k = 2.
    // For [-1, -2], the 2nd smallest negative integer is -1.
    // For [-2, -3], the 2nd smallest negative integer is -2.
    // For [-3, -4], the 2nd smallest negative integer is -3.
    // For [-4, -5], the 2nd smallest negative integer is -4. 
    fmt.Println(getSubarrayBeauty([]int{-1,-2,-3,-4,-5}, 2, 2)) // [-1,-2,-3,-4]
    // Example 3:
    // Input: nums = [-3,1,2,-3,0,-3], k = 2, x = 1
    // Output: [-3,0,-3,-3,-3]
    // Explanation: There are 5 subarrays with size k = 2.
    // For [-3, 1], the 1st smallest negative integer is -3.
    // For [1, 2], there is no negative integer so the beauty is 0.
    // For [2, -3], the 1st smallest negative integer is -3.
    // For [-3, 0], the 1st smallest negative integer is -3.
    // For [0, -3], the 1st smallest negative integer is -3.
    fmt.Println(getSubarrayBeauty([]int{-3,1,2,-3,0,-3}, 2, 1)) // [-3,0,-3,-3,-3]

    fmt.Println(getSubarrayBeauty([]int{1,2,3,4,5,6,7,8,9}, 2, 1)) // [0 0 0 0 0 0 0 0]
    fmt.Println(getSubarrayBeauty([]int{9,8,7,6,5,4,3,2,1}, 2, 1)) // [0 0 0 0 0 0 0 0]

    fmt.Println(getSubarrayBeauty1([]int{1,-1,-3,-2,3}, 3, 2)) // [-1,-2,-2]
    fmt.Println(getSubarrayBeauty1([]int{-1,-2,-3,-4,-5}, 2, 2)) // [-1,-2,-3,-4]
    fmt.Println(getSubarrayBeauty1([]int{-3,1,2,-3,0,-3}, 2, 1)) // [-3,0,-3,-3,-3]
    fmt.Println(getSubarrayBeauty1([]int{1,2,3,4,5,6,7,8,9}, 2, 1)) // [0 0 0 0 0 0 0 0]
    fmt.Println(getSubarrayBeauty1([]int{9,8,7,6,5,4,3,2,1}, 2, 1)) // [0 0 0 0 0 0 0 0]

    fmt.Println(getSubarrayBeauty2([]int{1,-1,-3,-2,3}, 3, 2)) // [-1,-2,-2]
    fmt.Println(getSubarrayBeauty2([]int{-1,-2,-3,-4,-5}, 2, 2)) // [-1,-2,-3,-4]
    fmt.Println(getSubarrayBeauty2([]int{-3,1,2,-3,0,-3}, 2, 1)) // [-3,0,-3,-3,-3]
    fmt.Println(getSubarrayBeauty2([]int{1,2,3,4,5,6,7,8,9}, 2, 1)) // [0 0 0 0 0 0 0 0]
    fmt.Println(getSubarrayBeauty2([]int{9,8,7,6,5,4,3,2,1}, 2, 1)) // [0 0 0 0 0 0 0 0]
}