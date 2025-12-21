package main

// 3779. Minimum Number of Operations to Have Distinct Elements
// You are given an integer array nums.

// In one operation, you remove the first three elements of the current array. 
// If there are fewer than three elements remaining, all remaining elements are removed.

// Repeat this operation until the array is empty or contains no duplicate values.

// Return an integer denoting the number of operations required.

// Example 1:
// Input: nums = [3,8,3,6,5,8]
// Output: 1
// Explanation:
// In the first operation, we remove the first three elements. 
// The remaining elements [6, 5, 8] are all distinct, so we stop. 
// Only one operation is needed.

// Example 2:
// Input: nums = [2,2]
// Output: 1
// Explanation:
// After one operation, the array becomes empty, which meets the stopping condition.

// Example 3:
// Input: nums = [4,3,5,1,2]
// Output: 0
// Explanation:
// All elements in the array are distinct, therefore no operations are needed.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

// time: O(n), space: O(max(nums) - min(nums))
func minOperations(nums []int) int {
    mn, mx := nums[0], nums[0]
    for _, v := range nums { // 找到最大和最小的元素
        mn = min(mn, v)
        mx = max(mx, v)
    }
    freq := make([]int, mx - mn + 1)
    res, count := 0, 0
    for _, v := range nums {
        v -= mn
        freq[v]++ // 统计每个元素的出现次数
        if freq[v] == 2 { // 如果出现次数为2，说明有重复元素
            count++
        }
    }
    for i, l := 0, len(nums); count != 0 && i < l; {
        for j := 3; i < l && j != 0; j, i = j-1, i+1 {
            v := nums[i] - mn
            freq[v]--
            if freq[v] == 1 {
                count--
            }
        }
        res++
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,8,3,6,5,8]
    // Output: 1
    // Explanation:
    // In the first operation, we remove the first three elements. 
    // The remaining elements [6, 5, 8] are all distinct, so we stop. 
    // Only one operation is needed.
    fmt.Println(minOperations([]int{3,8,3,6,5,8})) // 1
    // Example 2:
    // Input: nums = [2,2]
    // Output: 1
    // Explanation:
    // After one operation, the array becomes empty, which meets the stopping condition.
    fmt.Println(minOperations([]int{2,2})) // 1
    // Example 3:
    // Input: nums = [4,3,5,1,2]
    // Output: 0
    // Explanation:
    // All elements in the array are distinct, therefore no operations are needed.
    fmt.Println(minOperations([]int{4,3,5,1,2})) // 0

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 0
}