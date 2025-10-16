package main

// 2598. Smallest Missing Non-negative Integer After Operations
// You are given a 0-indexed integer array nums and an integer value.

// In one operation, you can add or subtract value from any element of nums.
//     For example, if nums = [1,2,3] and value = 2, you can choose to subtract value from nums[0] to make nums = [-1,2,3].

// The MEX (minimum excluded) of an array is the smallest missing non-negative integer in it.
//     For example, the MEX of [-1,2,3] is 0 while the MEX of [1,0,3] is 2.

// Return the maximum MEX of nums after applying the mentioned operation any number of times.

// Example 1:
// Input: nums = [1,-10,7,13,6,8], value = 5
// Output: 4
// Explanation: One can achieve this result by applying the following operations:
// - Add value to nums[1] twice to make nums = [1,0,7,13,6,8]
// - Subtract value from nums[2] once to make nums = [1,0,2,13,6,8]
// - Subtract value from nums[3] twice to make nums = [1,0,2,3,6,8]
// The MEX of nums is 4. It can be shown that 4 is the maximum MEX we can achieve.

// Example 2:
// Input: nums = [1,-10,7,13,6,8], value = 7
// Output: 2
// Explanation: One can achieve this result by applying the following operation:
// - subtract value from nums[2] once to make nums = [1,-10,0,13,6,8]
// The MEX of nums is 2. It can be shown that 2 is the maximum MEX we can achieve.

// Constraints:
//     1 <= nums.length, value <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"

func findSmallestInteger(nums []int, value int) int {
    arr := make([]int, value)
    for i := 0 ; i < len(nums); i++ {
        v := nums[i] % value
        if v < 0 {
            v += value
        }
        arr[v]++
    }
    res, mn := 0, 1 << 31
    for i := 0 ; i < len(arr) ; i++ {
        if arr[i] < mn {
            mn, res = arr[i], value * arr[i] + i
        }
    }
    return res
}

func findSmallestInteger1(nums []int, value int) int {
    count := make([]int, value)
    for _, v := range nums {
        count[(v % value + value) % value]++
    }
    for i := 0; ; i++ {
        if count[i % value] == 0 { return i }
        count[i % value]--
    }
    return -1
}

func findSmallestInteger2(nums []int, value int) int {
    count := make([]int, value)
    for _, v := range nums {
        count[(v%value + value) % value]++
    }
    res := len(nums)
    for i, v := range count {
        res = min(res, v * value + i)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,-10,7,13,6,8], value = 5
    // Output: 4
    // Explanation: One can achieve this result by applying the following operations:
    // - Add value to nums[1] twice to make nums = [1,0,7,13,6,8]
    // - Subtract value from nums[2] once to make nums = [1,0,2,13,6,8]
    // - Subtract value from nums[3] twice to make nums = [1,0,2,3,6,8]
    // The MEX of nums is 4. It can be shown that 4 is the maximum MEX we can achieve.
    fmt.Println(findSmallestInteger([]int{1,-10,7,13,6,8}, 5)) // 4
    // Example 2:
    // Input: nums = [1,-10,7,13,6,8], value = 7
    // Output: 2
    // Explanation: One can achieve this result by applying the following operation:
    // - subtract value from nums[2] once to make nums = [1,-10,0,13,6,8]
    // The MEX of nums is 2. It can be shown that 2 is the maximum MEX we can achieve.
    fmt.Println(findSmallestInteger([]int{1,-10,7,13,6,8}, 7)) // 2

    fmt.Println(findSmallestInteger([]int{1,2,3,4,5,6,7,8,9}, 7)) // 7
    fmt.Println(findSmallestInteger2([]int{9,8,7,6,5,4,3,2,1}, 7)) // 7

    fmt.Println(findSmallestInteger1([]int{1,-10,7,13,6,8}, 5)) // 4
    fmt.Println(findSmallestInteger1([]int{1,-10,7,13,6,8}, 7)) // 2
    fmt.Println(findSmallestInteger1([]int{1,2,3,4,5,6,7,8,9}, 7)) // 7
    fmt.Println(findSmallestInteger2([]int{9,8,7,6,5,4,3,2,1}, 7)) // 7

    fmt.Println(findSmallestInteger2([]int{1,-10,7,13,6,8}, 5)) // 4
    fmt.Println(findSmallestInteger2([]int{1,-10,7,13,6,8}, 7)) // 2
    fmt.Println(findSmallestInteger2([]int{1,2,3,4,5,6,7,8,9}, 7)) // 7
    fmt.Println(findSmallestInteger2([]int{9,8,7,6,5,4,3,2,1}, 7)) // 7
}