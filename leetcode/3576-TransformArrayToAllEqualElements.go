package main

// 3576. Transform Array to All Equal Elements
// You are given an integer array nums of size n containing only 1 and -1, and an integer k.

// You can perform the following operation at most k times:

//     Choose an index i (0 <= i < n - 1), and multiply both nums[i] and nums[i + 1] by -1.

// Note that you can choose the same index i more than once in different operations.

// Return true if it is possible to make all elements of the array equal after at most k operations, and false otherwise.

// Example 1:
// Input: nums = [1,-1,1,-1,1], k = 3
// Output: true
// Explanation:
// We can make all elements in the array equal in 2 operations as follows:
// Choose index i = 1, and multiply both nums[1] and nums[2] by -1. Now nums = [1,1,-1,-1,1].
// Choose index i = 2, and multiply both nums[2] and nums[3] by -1. Now nums = [1,1,1,1,1].

// Example 2:
// Input: nums = [-1,-1,-1,1,1,1], k = 5
// Output: false
// Explanation:
// It is not possible to make all array elements equal in at most 5 operations.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     nums[i] is either -1 or 1.
//     1 <= k <= n

import "fmt"

func canMakeEqual(nums []int, k int) bool {
    convert := func(nums []int, k, expected int) bool {
        i := 0
        for ; i < len(nums)-1; i++ {
            if nums[i] == expected { continue }
            if k == 0 { return false }
            k--
            nums[i] = -nums[i]
            nums[i+1] = -nums[i+1]
        }
        return i == len(nums)-1 && nums[len(nums)-1] == expected
    }
    // make 1
    arr := make([]int, 0, len(nums))
    arr = append(arr, nums...)
    if convert(arr, k, 1) {
        return true
    }
    // make -1
    arr = make([]int, 0, len(nums))
    arr = append(arr, nums...)
    if convert(arr, k, -1) {
        return true
    }
    return false
}

func canMakeEqual1(nums []int, k int) bool {
    check := func(target int) bool {
        count, mul := k, 1
        for i, v := range nums {
            if v * mul == target {
                mul = 1
                continue
            }
            if count == 0 || i == len(nums) - 1 {
                return false
            }
            count--
            mul = -1
        }
        return true
    }
    return check(-1) || check(1)
}

func main() {
    // Example 1:
    // Input: nums = [1,-1,1,-1,1], k = 3
    // Output: true
    // Explanation:
    // We can make all elements in the array equal in 2 operations as follows:
    // Choose index i = 1, and multiply both nums[1] and nums[2] by -1. Now nums = [1,1,-1,-1,1].
    // Choose index i = 2, and multiply both nums[2] and nums[3] by -1. Now nums = [1,1,1,1,1].
    fmt.Println(canMakeEqual([]int{1,-1,1,-1,1}, 3)) // true
    // Example 2:
    // Input: nums = [-1,-1,-1,1,1,1], k = 5
    // Output: false
    // Explanation:
    // It is not possible to make all array elements equal in at most 5 operations.
    fmt.Println(canMakeEqual([]int{-1,-1,-1,1,1,1}, 5)) // false

    fmt.Println(canMakeEqual([]int{-1,-1,-1,-1,-1,-1,-1,-1,-1, -1}, 3)) // true
    fmt.Println(canMakeEqual([]int{1,1,1,1,1,1,1,1,1,1}, 3)) // true
    fmt.Println(canMakeEqual([]int{1,1,1,1,1,-1,-1,-1,-1,-1}, 3)) // false
    fmt.Println(canMakeEqual([]int{-1,-1,-1,-1,-1,-1,1,1,1,1,}, 3)) // true

    fmt.Println(canMakeEqual1([]int{1,-1,1,-1,1}, 3)) // true
    fmt.Println(canMakeEqual1([]int{-1,-1,-1,1,1,1}, 5)) // false
    fmt.Println(canMakeEqual1([]int{-1,-1,-1,-1,-1,-1,-1,-1,-1, -1}, 3)) // true
    fmt.Println(canMakeEqual1([]int{1,1,1,1,1,1,1,1,1,1}, 3)) // true
    fmt.Println(canMakeEqual1([]int{1,1,1,1,1,-1,-1,-1,-1,-1}, 3)) // false
    fmt.Println(canMakeEqual1([]int{-1,-1,-1,-1,-1,-1,1,1,1,1,}, 3)) // true
}