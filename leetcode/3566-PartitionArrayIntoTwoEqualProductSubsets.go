package main

// 3566. Partition Array into Two Equal Product Subsets
// You are given an integer array nums containing distinct positive integers and an integer target.

// Determine if you can partition nums into two non-empty disjoint subsets, with each element belonging to exactly one subset, 
// such that the product of the elements in each subset is equal to target.

// Return true if such a partition exists and false otherwise.

// A subset of an array is a selection of elements of the array.

// Example 1:
// Input: nums = [3,1,6,8,4], target = 24
// Output: true
// Explanation: The subsets [3, 8] and [1, 6, 4] each have a product of 24. Hence, the output is true.

// Example 2:
// Input: nums = [2,5,3,7], target = 15
// Output: false
// Explanation: There is no way to partition nums into two non-empty disjoint subsets such that both subsets have a product of 15. Hence, the output is false.

// Constraints:
//     3 <= nums.length <= 12
//     1 <= target <= 10^15
//     1 <= nums[i] <= 100
//     All elements of nums are distinct.

import "fmt"

func checkEqualPartitions(nums []int, target int64) bool {
    mul := int64(1) // Calculate the total product of all numbers in nums
    for _, v := range nums {
        mul *= int64(v)
    }
    // If the total product is not equal to target squared, return false
    // This ensures that the product can be split into two equal parts
    if mul != target * target {
        return false
    }
    // Recursive function to explore possible subsets
    var helper func(index int, curr int64) bool
    helper = func(index int, curr int64) bool {
        // If we find a subset whose product equals target, return true
        if curr == target { return true }
        // If current product exceeds target, stop searching this path
        if curr > target { return false }
        // If all elements are used and we haven't found a valid subset, return false
        if index == len(nums) {
            return false
        }
        // Try including the current number in the subset
        // Try excluding the current number from the subset
        // Return true if either approach finds a valid partition
        return helper(index+1, curr * int64(nums[index])) ||  helper(index + 1, curr)
    }
    // Start the recursion with an initial product of 1
    return helper(0, 1)
}

func main() {
    // Example 1:
    // Input: nums = [3,1,6,8,4], target = 24
    // Output: true
    // Explanation: The subsets [3, 8] and [1, 6, 4] each have a product of 24. Hence, the output is true.
    fmt.Println(checkEqualPartitions([]int{3,1,6,8,4}, 24)) // true
    // Example 2:
    // Input: nums = [2,5,3,7], target = 15
    // Output: false
    // Explanation: There is no way to partition nums into two non-empty disjoint subsets such that both subsets have a product of 15. Hence, the output is false.
    fmt.Println(checkEqualPartitions([]int{2,5,3,7}, 15)) // false

    fmt.Println(checkEqualPartitions([]int{1,2,3,4,5,6,7,8,9}, 15)) // false
    fmt.Println(checkEqualPartitions([]int{9,8,7,6,5,4,3,2,1}, 15)) // false
}