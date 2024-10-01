package main

// 1403. Minimum Subsequence in Non-Increasing Order
// Given the array nums, obtain a subsequence of the array whose sum of elements is strictly greater than 
// the sum of the non included elements in such subsequence. 

// If there are multiple solutions, 
// return the subsequence with minimum size and if there still exist multiple solutions, 
// return the subsequence with the maximum total sum of all its elements. 
// A subsequence of an array can be obtained by erasing some (possibly zero) elements from the array. 

// Note that the solution with the given constraints is guaranteed to be unique. 
// Also return the answer sorted in non-increasing order.

// Example 1:
// Input: nums = [4,3,10,9,8]
// Output: [10,9] 
// Explanation: The subsequences [10,9] and [10,8] are minimal such that the sum of their elements is strictly greater than the sum of elements not included. However, the subsequence [10,9] has the maximum total sum of its elements. 

// Example 2:
// Input: nums = [4,4,7,6,7]
// Output: [7,7,6] 
// Explanation: The subsequence [7,7] has the sum of its elements equal to 14 which is not strictly greater than the sum of elements not included (14 = 4 + 4 + 6). Therefore, the subsequence [7,6,7] is the minimal satisfying the conditions. Note the subsequence has to be returned in non-increasing order.  

// Constraints:
//     1 <= nums.length <= 500
//     1 <= nums[i] <= 100

import "fmt"
import "sort"

func minSubsequence(nums []int) []int {
    sum, t, res := 0, 0, []int{}
    for _, v := range nums {
        sum += v
    }
    sort.Slice(nums, func(i, j int) bool { 
        return nums[i] > nums[j] 
    })
    for _, v := range nums {
        res = append(res, v)
        t += v
        if t > sum - t { break }
    }
    return res
}

func minSubsequence1(nums []int) []int {
    sort.Ints(nums)
    sum := 0
    for _, i := range nums {
        sum += i
    }
    check, res := 0, make([]int, 0, len(nums)/2)
    for i := len(nums) - 1; i >= 0; i-- {
        v := nums[i]
        check += v
        res = append(res, v)
        if check > sum / 2 {
            break
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,3,10,9,8]
    // Output: [10,9] 
    // Explanation: The subsequences [10,9] and [10,8] are minimal such that the sum of their elements is strictly greater than the sum of elements not included. However, the subsequence [10,9] has the maximum total sum of its elements. 
    fmt.Println(minSubsequence([]int{4,3,10,9,8})) // [10,9] 
    // Example 2:
    // Input: nums = [4,4,7,6,7]
    // Output: [7,7,6] 
    // Explanation: The subsequence [7,7] has the sum of its elements equal to 14 which is not strictly greater than the sum of elements not included (14 = 4 + 4 + 6). Therefore, the subsequence [7,6,7] is the minimal satisfying the conditions. Note the subsequence has to be returned in non-increasing order.  
    fmt.Println(minSubsequence([]int{4,4,7,6,7})) // [7,7,6] 

    fmt.Println(minSubsequence1([]int{4,3,10,9,8})) // [10,9] 
    fmt.Println(minSubsequence1([]int{4,4,7,6,7})) // [7,7,6] 
}