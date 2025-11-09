package main

// 3736. Minimum Moves to Equal Array Elements III
// You are given an integer array nums.

// In one move, you may increase the value of any single element nums[i] by 1.

// Return the minimum total number of moves required so that all elements in nums become equal.

// Example 1:
// Input: nums = [2,1,3]
// Output: 3
// Explanation:
// To make all elements equal:
// Increase nums[0] = 2 by 1 to make it 3.
// Increase nums[1] = 1 by 1 to make it 2.
// Increase nums[1] = 2 by 1 to make it 3.
// Now, all elements of nums are equal to 3. The minimum total moves is 3.

// Example 2:
// Input: nums = [4,4,5]
// Output: 2
// Explanation:
// To make all elements equal:
// Increase nums[0] = 4 by 1 to make it 5.
// Increase nums[1] = 4 by 1 to make it 5.
// Now, all elements of nums are equal to 5. The minimum total moves is 2.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"

func minMoves(nums []int) int {
    sum, n, mx := 0, len(nums), nums[0]
    for _, v := range nums {
        sum += v
        if v > mx {
            mx = v
        }
    }
    return n * mx - sum
}

func main() {
    // Example 1:
    // Input: nums = [2,1,3]
    // Output: 3
    // Explanation:
    // To make all elements equal:
    // Increase nums[0] = 2 by 1 to make it 3.
    // Increase nums[1] = 1 by 1 to make it 2.
    // Increase nums[1] = 2 by 1 to make it 3.
    // Now, all elements of nums are equal to 3. The minimum total moves is 3.
    fmt.Println(minMoves([]int{2,1,3})) // 3
    // Example 2:
    // Input: nums = [4,4,5]
    // Output: 2
    // Explanation:
    // To make all elements equal:
    // Increase nums[0] = 4 by 1 to make it 5.
    // Increase nums[1] = 4 by 1 to make it 5.
    // Now, all elements of nums are equal to 5. The minimum total moves is 2.
    fmt.Println(minMoves([]int{4,4,5})) // 2

    fmt.Println(minMoves([]int{1,2,3,4,5,6,7,8,9})) // 36
    fmt.Println(minMoves([]int{9,8,7,6,5,4,3,2,1})) // 36
}