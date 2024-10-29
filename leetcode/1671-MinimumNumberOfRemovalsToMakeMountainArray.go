package main

// 1671. Minimum Number of Removals to Make Mountain Array
// You may recall that an array arr is a mountain array if and only if:
//     arr.length >= 3
//     There exists some index i (0-indexed) with 0 < i < arr.length - 1 such that:
//         arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
//         arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

// Given an integer array nums​​​, return the minimum number of elements to remove to make nums​​​ a mountain array.

// Example 1:
// Input: nums = [1,3,1]
// Output: 0
// Explanation: The array itself is a mountain array so we do not need to remove any elements.

// Example 2:
// Input: nums = [2,1,1,5,6,2,3,1]
// Output: 3
// Explanation: One solution is to remove the elements at indices 0, 1, and 5, making the array nums = [1,5,6,3,1].

// Constraints:
//     3 <= nums.length <= 1000
//     1 <= nums[i] <= 10^9
//     It is guaranteed that you can make a mountain array out of nums.

import "fmt"

func minimumMountainRemovals(nums []int) int {
    n := len(nums)
    arr := make([]int, n)
    copy(arr, nums)
    left, right := 0, n - 1
    for left < right { // 翻转数组
        arr[left], arr[right] = arr[right], arr[left]
        left++
        right--
    }
    helper := func(nums []int) []int {
        dp, res := []int{}, make([]int, len(nums)) 
        for i, v := range nums {
            if len(dp) == 0 || v > dp[len(dp)-1] {
                dp = append(dp, v)
            } else {
                l, r := 0, len(dp) - 1
                for l <= r {
                    mid := (l + r) >> 1
                    if dp[mid] >= v {
                        r = mid - 1
                    } else if dp[mid] < v {
                        l = mid + 1
                    }
                }
                dp[l] = v
            }
            res[i] = len(dp)
        }
        return res
    }
    mx, arr1, arr2 := 0, helper(nums), helper(arr)
    for i := range nums {
        if arr1[i] >= 2 && arr2[n - i - 1] >= 2 && arr1[i] + arr2[n - i - 1] - 1 > mx {
            mx = arr1[i] + arr2[n - i - 1] - 1
        }
    }
    return n - mx
}

func main() {
    // Example 1:
    // Input: nums = [1,3,1]
    // Output: 0
    // Explanation: The array itself is a mountain array so we do not need to remove any elements.
    fmt.Println(minimumMountainRemovals([]int{1,3,1})) // 0
    // Example 2:
    // Input: nums = [2,1,1,5,6,2,3,1]
    // Output: 3
    // Explanation: One solution is to remove the elements at indices 0, 1, and 5, making the array nums = [1,5,6,3,1].
    fmt.Println(minimumMountainRemovals([]int{2,1,1,5,6,2,3,1})) // 3
}