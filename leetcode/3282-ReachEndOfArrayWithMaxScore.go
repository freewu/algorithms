package main

// 3282. Reach End of Array With Max Score
// You are given an integer array nums of length n.

// Your goal is to start at index 0 and reach index n - 1. 
// You can only jump to indices greater than your current index.

// The score for a jump from index i to index j is calculated as (j - i) * nums[i].

// Return the maximum possible total score by the time you reach the last index.

// Example 1:
// Input: nums = [1,3,1,5]
// Output: 7
// Explanation:
// First, jump to index 1 and then jump to the last index. The final score is 1 * 1 + 2 * 3 = 7.

// Example 2:
// Input: nums = [4,3,1,3,2]
// Output: 16
// Explanation:
// Jump directly to the last index. The final score is 4 * 4 = 16.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func findMaximumScore(nums []int) int64 {
    res, n := 0, len(nums)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if nums[j] > nums[i] {
                res += (nums[i] * (j - i))
                i = j - 1
                break
            }
            if j == n - 1 {
                res += (nums[i] * (j - i))
                i = j
            }
        }
    }
    return int64(res)
}

func findMaximumScore1(nums []int) int64 {
    queue := []int{0} //index asc <= ; 2 4 1 
    res := int64(0)
    for i := 1; i < len(nums); i++ {
        if nums[queue[len(queue) - 1]] <= nums[i] {
            m := queue[len(queue) - 1]
            res += int64(nums[m] * (i - m))
            queue = queue[:len(queue) - 1]
            queue = append(queue, i)
        }
    }
    i := len(nums)-1 
    m := queue[len(queue) - 1]
    res += int64(nums[m] * (i-m))
    return res
}

func findMaximumScore2(nums []int) int64 {
    res, mx := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        res += mx
        mx = max(mx, v)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,3,1,5]
    // Output: 7
    // Explanation:
    // First, jump to index 1 and then jump to the last index. The final score is 1 * 1 + 2 * 3 = 7.
    fmt.Println(findMaximumScore([]int{1,3,1,5})) // 7
    // Example 2:
    // Input: nums = [4,3,1,3,2]
    // Output: 16
    // Explanation:
    // Jump directly to the last index. The final score is 4 * 4 = 16.
    fmt.Println(findMaximumScore([]int{4,3,1,3,2})) // 16

    fmt.Println(findMaximumScore([]int{1,2,3,4,5,6,7,8,9})) // 36
    fmt.Println(findMaximumScore([]int{9,8,7,6,5,4,3,2,1})) // 72

    fmt.Println(findMaximumScore1([]int{1,3,1,5})) // 7
    fmt.Println(findMaximumScore1([]int{4,3,1,3,2})) // 16
    fmt.Println(findMaximumScore1([]int{1,2,3,4,5,6,7,8,9})) // 36
    fmt.Println(findMaximumScore1([]int{9,8,7,6,5,4,3,2,1})) // 72

    fmt.Println(findMaximumScore2([]int{1,3,1,5})) // 7
    fmt.Println(findMaximumScore2([]int{4,3,1,3,2})) // 16
    fmt.Println(findMaximumScore2([]int{1,2,3,4,5,6,7,8,9})) // 36
    fmt.Println(findMaximumScore2([]int{9,8,7,6,5,4,3,2,1})) // 72
}