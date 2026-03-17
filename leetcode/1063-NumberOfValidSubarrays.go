package main

// 1063. Number of Valid Subarrays
// Given an integer array nums, return the number of non-empty subarrays with the leftmost element of the subarray not larger than other elements in the subarray.
// A subarray is a contiguous part of an array.

// Example 1:
// Input: nums = [1,4,2,5,3]
// Output: 11
// Explanation: There are 11 valid subarrays: [1],[4],[2],[5],[3],[1,4],[2,5],[1,4,2],[2,5,3],[1,4,2,5],[1,4,2,5,3].

// Example 2:
// Input: nums = [3,2,1]
// Output: 3
// Explanation: The 3 valid subarrays are: [3],[2],[1].

// Example 3:
// Input: nums = [2,2,2]
// Output: 6
// Explanation: There are 6 valid subarrays: [2],[2],[2],[2,2],[2,2],[2,2,2].

// Constraints:
//     1 <= nums.length <= 5 * 10^4
//     0 <= nums[i] <= 10^5

import "fmt"

func validSubarrays(nums []int) int {
    n, res := len(nums), 0
    nums = append(nums, -1)
    queue := []int{ n }
    for i := n-1; i >= 0; i-- {
        for nums[i] <= nums[queue[len(queue) - 1]] { // 子数组最左边的元素不大于子数组中的其他元素 
            queue = queue[:len(queue) - 1]
        }
        res += queue[len(queue) - 1] - i
        queue = append(queue, i) // push
    }
    return res
}

func validSubarrays1(nums []int) int {
    res, stack := 0, []int{-1}
    for _, v := range nums {
        c := stack[len(stack) - 1]
        if v >= c {
            stack = append(stack, v)
            res += len(stack) - 1            
        } else {
            for v < c && len(stack) > 0 {
                stack = stack[:len(stack) - 1]
                c = stack[len(stack) - 1]
            }
            stack = append(stack, v)
            res += len(stack) - 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,4,2,5,3]
    // Output: 11
    // Explanation: There are 11 valid subarrays: [1],[4],[2],[5],[3],[1,4],[2,5],[1,4,2],[2,5,3],[1,4,2,5],[1,4,2,5,3].
    fmt.Println(validSubarrays([]int{1,4,2,5,3})) // 11
    // Example 2:
    // Input: nums = [3,2,1]
    // Output: 3
    // Explanation: The 3 valid subarrays are: [3],[2],[1].
    fmt.Println(validSubarrays([]int{3,2,1})) // 3
    // Example 3:
    // Input: nums = [2,2,2]
    // Output: 6
    // Explanation: There are 6 valid subarrays: [2],[2],[2],[2,2],[2,2],[2,2,2].
    fmt.Println(validSubarrays([]int{2,2,2})) // 6

    fmt.Println(validSubarrays([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(validSubarrays([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(validSubarrays1([]int{1,4,2,5,3})) // 11
    fmt.Println(validSubarrays1([]int{3,2,1})) // 3
    fmt.Println(validSubarrays1([]int{2,2,2})) // 6
    fmt.Println(validSubarrays1([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(validSubarrays1([]int{9,8,7,6,5,4,3,2,1})) // 9
}