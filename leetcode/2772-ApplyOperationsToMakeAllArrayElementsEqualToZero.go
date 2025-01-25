package main

// 2772. Apply Operations to Make All Array Elements Equal to Zero
// You are given a 0-indexed integer array nums and a positive integer k.

// You can apply the following operation on the array any number of times:
//     1. Choose any subarray of size k from the array and decrease all its elements by 1.

// Return true if you can make all the array elements equal to 0, or false otherwise.

// A subarray is a contiguous non-empty part of an array.

// Example 1:
// Input: nums = [2,2,3,1,1,0], k = 3
// Output: true
// Explanation: We can do the following operations:
// - Choose the subarray [2,2,3]. The resulting array will be nums = [1,1,2,1,1,0].
// - Choose the subarray [2,1,1]. The resulting array will be nums = [1,1,1,0,0,0].
// - Choose the subarray [1,1,1]. The resulting array will be nums = [0,0,0,0,0,0].

// Example 2:
// Input: nums = [1,3,1,1], k = 2
// Output: false
// Explanation: It is not possible to make all the array elements equal to 0.

// Constraints:
//     1 <= k <= nums.length <= 10^5
//     0 <= nums[i] <= 10^6

import "fmt"

func checkArray(nums []int, k int) bool {
    n, dec := len(nums), 0
    memo := make([]int, k) 
    for i := 0; i < n; i++ {
        if i >= k {
            dec -= memo[i % k]
        }
        rem := nums[i] - dec
        if i <= n - k {
            if rem < 0 { return false }
            memo[i % k] = rem 
            dec += rem 
        } else {
            if rem != 0 {
                return false 
            }
        }
    }
    return true 
}

func checkArray1(nums []int, k int) bool {
    n := len(nums)
    diff := make([]int, n + 1)
    for i := 1; i < n; i++ {
        diff[i] = nums[i] - nums[i-1]
    }
    diff[0], diff[n] = nums[0], -nums[n - 1]
    for i := 0; i <= n; i++ {
        if diff[i] < 0 { return false }
        if i > n - k { continue }
        if diff[i] > 0 {
            diff[i + k] += diff[i]
            diff[i] = 0
        }
    }
    return true
}

func checkArray2(nums []int, k int) bool {
    cur := 0
    for i := 0;i< len(nums); i++ {
        if nums[i] < cur { return false }
        nums[i] = nums[i] - cur
        cur += nums[i]
        if i >= k - 1 {
            cur -= nums[i - k + 1]
        }
    }
    if cur == 0 { return true }
    return false
}

func main() {
    // Example 1:
    // Input: nums = [2,2,3,1,1,0], k = 3
    // Output: true
    // Explanation: We can do the following operations:
    // - Choose the subarray [2,2,3]. The resulting array will be nums = [1,1,2,1,1,0].
    // - Choose the subarray [2,1,1]. The resulting array will be nums = [1,1,1,0,0,0].
    // - Choose the subarray [1,1,1]. The resulting array will be nums = [0,0,0,0,0,0].
    fmt.Println(checkArray([]int{2,2,3,1,1,0}, 3)) // true
    // Example 2:
    // Input: nums = [1,3,1,1], k = 2
    // Output: false
    // Explanation: It is not possible to make all the array elements equal to 0.
    fmt.Println(checkArray([]int{1,3,1,1}, 2)) // false

    fmt.Println(checkArray([]int{1,2,3,4,5,6,7,8,9}, 2)) // false
    fmt.Println(checkArray([]int{9,8,7,6,5,4,3,2,1}, 2)) // false

    fmt.Println(checkArray1([]int{2,2,3,1,1,0}, 3)) // true
    fmt.Println(checkArray1([]int{1,3,1,1}, 2)) // false
    fmt.Println(checkArray1([]int{1,2,3,4,5,6,7,8,9}, 2)) // false
    fmt.Println(checkArray1([]int{9,8,7,6,5,4,3,2,1}, 2)) // false

    fmt.Println(checkArray2([]int{2,2,3,1,1,0}, 3)) // true
    fmt.Println(checkArray2([]int{1,3,1,1}, 2)) // false
    fmt.Println(checkArray2([]int{1,2,3,4,5,6,7,8,9}, 2)) // false
    fmt.Println(checkArray2([]int{9,8,7,6,5,4,3,2,1}, 2)) // false
}