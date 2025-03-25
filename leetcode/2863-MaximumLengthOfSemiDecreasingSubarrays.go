package main

// 2863. Maximum Length of Semi-Decreasing Subarrays
// You are given an integer array nums.

// Return the length of the longest semi-decreasing subarray of nums, and 0 if there are no such subarrays.
//     A subarray is a contiguous non-empty sequence of elements within an array.
//     A non-empty array is semi-decreasing if its first element is strictly greater than its last element.

// Example 1:
// Input: nums = [7,6,5,4,3,2,1,6,10,11]
// Output: 8
// Explanation: Take the subarray [7,6,5,4,3,2,1,6].
// The first element is 7 and the last one is 6 so the condition is met.
// Hence, the answer would be the length of the subarray or 8.
// It can be shown that there aren't any subarrays with the given condition with a length greater than 8.

// Example 2:
// Input: nums = [57,55,50,60,61,58,63,59,64,60,63]
// Output: 6
// Explanation: Take the subarray [61,58,63,59,64,60].
// The first element is 61 and the last one is 60 so the condition is met.
// Hence, the answer would be the length of the subarray or 6.
// It can be shown that there aren't any subarrays with the given condition with a length greater than 6.

// Example 3:
// Input: nums = [1,2,3,4]
// Output: 0
// Explanation: Since there are no semi-decreasing subarrays in the given array, the answer is 0.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"
import "sort"

func maxSubarrayLength(nums []int) int {
    mp := map[int][]int{}
    for i, v := range nums {
        mp[v] = append(mp[v], i)
    }
    keys := []int{}
    for k := range mp {
        keys = append(keys, k)
    }
    sort.Slice(keys, func(i, j int) bool { 
        return keys[i] > keys[j] 
    })
    res, k := 0, 1 << 30
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range keys {
        indexs := mp[v]
        res, k = max(res, indexs[len(indexs) - 1] - k + 1), min(k, indexs[0])
    }
    return res
}

func maxSubarrayLength1(nums []int) int {
    res, n, stack := 0, len(nums), []int{}
    for i := n - 1; i >= 0; i-- {
        if len(stack) == 0 || nums[stack[len(stack) - 1]] > nums[i] {
            stack = append(stack, i)
        }
    }
    for i := 0; i < n; i++ {
        for len(stack) != 0 && nums[stack[len(stack )- 1]] < nums[i] {
            res = max(res, stack[len(stack) - 1] - i + 1)
            stack = stack[:len(stack) - 1]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [7,6,5,4,3,2,1,6,10,11]
    // Output: 8
    // Explanation: Take the subarray [7,6,5,4,3,2,1,6].
    // The first element is 7 and the last one is 6 so the condition is met.
    // Hence, the answer would be the length of the subarray or 8.
    // It can be shown that there aren't any subarrays with the given condition with a length greater than 8.
    fmt.Println(maxSubarrayLength([]int{7,6,5,4,3,2,1,6,10,11})) // 8
    // Example 2:
    // Input: nums = [57,55,50,60,61,58,63,59,64,60,63]
    // Output: 6
    // Explanation: Take the subarray [61,58,63,59,64,60].
    // The first element is 61 and the last one is 60 so the condition is met.
    // Hence, the answer would be the length of the subarray or 6.
    // It can be shown that there aren't any subarrays with the given condition with a length greater than 6.
    fmt.Println(maxSubarrayLength([]int{57,55,50,60,61,58,63,59,64,60,63})) // 6
    // Example 3:
    // Input: nums = [1,2,3,4]
    // Output: 0
    // Explanation: Since there are no semi-decreasing subarrays in the given array, the answer is 0.
    fmt.Println(maxSubarrayLength([]int{1,2,3,4})) // 0

    fmt.Println(maxSubarrayLength([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(maxSubarrayLength([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(maxSubarrayLength1([]int{7,6,5,4,3,2,1,6,10,11})) // 8
    fmt.Println(maxSubarrayLength1([]int{57,55,50,60,61,58,63,59,64,60,63})) // 6
    fmt.Println(maxSubarrayLength1([]int{1,2,3,4})) // 0
    fmt.Println(maxSubarrayLength1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(maxSubarrayLength1([]int{9,8,7,6,5,4,3,2,1})) // 9
}