package main

// 2874. Maximum Value of an Ordered Triplet II
// You are given a 0-indexed integer array nums.

// Return the maximum value over all triplets of indices (i, j, k) such that i < j < k. 
// If all such triplets have a negative value, return 0.

// The value of a triplet of indices (i, j, k) is equal to (nums[i] - nums[j]) * nums[k].

// Example 1:
// Input: nums = [12,6,1,2,7]
// Output: 77
// Explanation: The value of the triplet (0, 2, 4) is (nums[0] - nums[2]) * nums[4] = 77.
// It can be shown that there are no ordered triplets of indices with a value greater than 77. 

// Example 2:
// Input: nums = [1,10,3,4,19]
// Output: 133
// Explanation: The value of the triplet (1, 2, 4) is (nums[1] - nums[2]) * nums[4] = 133.
// It can be shown that there are no ordered triplets of indices with a value greater than 133.

// Example 3:
// Input: nums = [1,2,3]
// Output: 0
// Explanation: The only ordered triplet of indices (0, 1, 2) has a negative value of (nums[0] - nums[1]) * nums[2] = -3. Hence, the answer would be 0.

// Constraints:
//     3 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^6

import "fmt"

func maximumTripletValue(nums []int) int64 {
    prefixMax, suffixMax := make([]int, len(nums)), make([]int, len(nums))
    res, preMax, sufMax := int64(0), 0, 0
    for i, j := 0, len(nums)-1; i < len(nums) && j >= 0; i, j = i+1, j-1 {
        if preMax < nums[i] { preMax = nums[i] }
        if sufMax < nums[j] { sufMax = nums[j] }
        prefixMax[i], suffixMax[j] = preMax, sufMax
    }
    for i := 1; i < len(nums)-1; i++ {
        v := int64(prefixMax[i-1]-nums[i]) * int64(suffixMax[i+1])
        if res <  v {
            res = v
        }
    }
    return res
}

func maximumTripletValue1(nums []int) int64 {
    res, n := 0, len(nums)
    pre, suf := make([]int, n), make([]int, n) // 前缀最大值和后缀最大值
    pre[0], suf[n-1] = nums[0], nums[n-1]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        pre[i] = max(pre[i-1], nums[i])
    }
    for i := n - 2; i >= 0; i-- {
        suf[i] = max(suf[i+1], nums[i])
    }
    for i := 1; i < n-1; i++ {
        res = max(res, (pre[i-1]-nums[i]) * suf[i+1])
    }
    return int64(res)
}

func maximumTripletValue2(nums []int) int64 {
    res, imax, dmax := 0, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        res = max(res, dmax * v)
        dmax = max(dmax, imax - v)
        imax = max(imax, v)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [12,6,1,2,7]
    // Output: 77
    // Explanation: The value of the triplet (0, 2, 4) is (nums[0] - nums[2]) * nums[4] = 77.
    // It can be shown that there are no ordered triplets of indices with a value greater than 77. 
    fmt.Println(maximumTripletValue([]int{12,6,1,2,7})) // 77
    // Example 2:
    // Input: nums = [1,10,3,4,19]
    // Output: 133
    // Explanation: The value of the triplet (1, 2, 4) is (nums[1] - nums[2]) * nums[4] = 133.
    // It can be shown that there are no ordered triplets of indices with a value greater than 133.
    fmt.Println(maximumTripletValue([]int{1,10,3,4,19})) // 133
    // Example 3:
    // Input: nums = [1,2,3]
    // Output: 0
    // Explanation: The only ordered triplet of indices (0, 1, 2) has a negative value of (nums[0] - nums[1]) * nums[2] = -3. Hence, the answer would be 0.
    fmt.Println(maximumTripletValue([]int{1,2,3})) // 0

    fmt.Println(maximumTripletValue([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(maximumTripletValue([]int{9,8,7,6,5,4,3,2,1})) // 16

    fmt.Println(maximumTripletValue1([]int{12,6,1,2,7})) // 77
    fmt.Println(maximumTripletValue1([]int{1,10,3,4,19})) // 133
    fmt.Println(maximumTripletValue1([]int{1,2,3})) // 0
    fmt.Println(maximumTripletValue1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(maximumTripletValue1([]int{9,8,7,6,5,4,3,2,1})) // 16

    fmt.Println(maximumTripletValue2([]int{12,6,1,2,7})) // 77
    fmt.Println(maximumTripletValue2([]int{1,10,3,4,19})) // 133
    fmt.Println(maximumTripletValue2([]int{1,2,3})) // 0
    fmt.Println(maximumTripletValue2([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(maximumTripletValue2([]int{9,8,7,6,5,4,3,2,1})) // 16
}