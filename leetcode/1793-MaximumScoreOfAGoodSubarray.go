package main

// 1793. Maximum Score of a Good Subarray
// You are given an array of integers nums (0-indexed) and an integer k.
// The score of a subarray (i, j) is defined as min(nums[i], nums[i+1], ..., nums[j]) * (j - i + 1). 
// A good subarray is a subarray where i <= k <= j.

// Return the maximum possible score of a good subarray.

// Example 1:
// Input: nums = [1,4,3,7,4,5], k = 3
// Output: 15
// Explanation: The optimal subarray is (1, 5) with a score of min(4,3,7,4,5) * (5-1+1) = 3 * 5 = 15. 

// Example 2:
// Input: nums = [5,5,4,5,4,1,1,1], k = 0
// Output: 20
// Explanation: The optimal subarray is (0, 4) with a score of min(5,5,4,5,4) * (4-0+1) = 4 * 5 = 20.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 2 * 10^4
//     0 <= k < nums.length

import "fmt"

// 2 Pointers O(n) O(1)
func maximumScore(nums []int, k int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x < y { return y; }; return x; }
    l := len(nums)
    // iterate to the left to update the minimum value at each index
    val := nums[k]
    for i := k - 1; i >= 0; i-- {
        val = min(val, nums[i])
        nums[i] = val
    }
    // iterate to the right to update the minimum value at each index
    val = nums[k]
    for i := k + 1; i < l; i++ {
        val = min(val, nums[i])
        nums[i] = val
    }
    // start with 2 pointers at opposite ends of nums
    res, left, right := 0, 0, l - 1
    for left <= right {
        res = max(res, min(nums[left], nums[right]) * (right - left + 1))
        // first to check if either pointer is at k
        if left == k { // if it is at k then we must move the other pointer inwards
            right--
        } else if right == k {
            left++
        } else if nums[left] < nums[right] { // if neither index is at k move the pointer that is smaller inwards
            left++
        } else {
            right--
        }
    }
    return res
}

func main() {
    // The optimal subarray is (1, 5) with a score of min(4,3,7,4,5) * (5-1+1) = 3 * 5 = 15. 
    fmt.Println(maximumScore([]int{1,4,3,7,4,5} ,3)) // 15
    // The optimal subarray is (0, 4) with a score of min(5,5,4,5,4) * (4-0+1) = 4 * 5 = 20.
    fmt.Println(maximumScore([]int{5,5,4,5,4,1,1,1} ,0)) // 20
}