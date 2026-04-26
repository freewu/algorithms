package main

// 3909. Compare Sums of Bitonic Parts
// You are given a bitonic array nums of length n.

// Split the array into two parts:
//     1. Ascending part: from index 0 to the peak element (inclusive).
//     2. Descending part: from the peak element to index n - 1 (inclusive).

// The peak element belongs to both parts.

// Return:
//     1. 0 if the sum of the ascending part is greater.
//     2. 1 if the sum of the descending part is greater.
//     3. -1 if both sums are equal.

// Notes:
//     1. A bitonic array is an array that is strictly increasing up to a single peak element and then strictly decreasing.
//     2. An array is said to be strictly increasing if each element is strictly greater than its previous one (if exists).
//     3. An array is said to be strictly decreasing if each element is strictly smaller than its previous one (if exists).

// Example 1:
// Input: nums = [1,3,2,1]
// Output: 1
// Explanation:
// Peak element is nums[1] = 3
// Ascending part = [1, 3], sum is 1 + 3 = 4
// Descending part = [3, 2, 1], sum is 3 + 2 + 1 = 6
// Since the descending part has a larger sum, return 1.

// Example 2:
// Input: nums = [2,4,5,2]
// Output: 0
// Explanation:
// Peak element is nums[2] = 5
// Ascending part = [2, 4, 5], sum is 2 + 4 + 5 = 11
// Descending part = [5, 2], sum is 5 + 2 = 7
// Since the ascending part has a larger sum, return 0.

// Example 3:
// Input: nums = [1,2,4,3]
// Output: -1
// Explanation:
// Peak element is nums[2] = 4
// Ascending part = [1, 2, 4], sum is 1 + 2 + 4 = 7
// Descending part = [4, 3], sum is 4 + 3 = 7
// Since both parts have equal sums, return -1.

// Constraints:
//     3 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     nums is a bitonic array.

import "fmt"

func compareBitonicSums(nums []int) int {
    peak, n := 0, len(nums)
    for i := 1; i < n; i++ {
        if nums[i] > nums[peak] {
            peak = i
        }
    }
    ascSum, descSum := 0, 0
    for i := 0; i <= peak; i++ {
        ascSum += nums[i]
    }
    for i := peak; i < n; i++ {
        descSum += nums[i]
    }
    if ascSum > descSum { // 0 if the sum of the ascending part is greater.
        return 0
    } else if descSum > ascSum { // 1 if the sum of the descending part is greater.
        return 1
    }
    return -1 // -1 if both sums are equal.
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2,1]
    // Output: 1
    // Explanation:
    // Peak element is nums[1] = 3
    // Ascending part = [1, 3], sum is 1 + 3 = 4
    // Descending part = [3, 2, 1], sum is 3 + 2 + 1 = 6
    // Since the descending part has a larger sum, return 1.
    fmt.Println(compareBitonicSums([]int{1,3,2,1})) // 1
    // Example 2:
    // Input: nums = [2,4,5,2]
    // Output: 0
    // Explanation:
    // Peak element is nums[2] = 5
    // Ascending part = [2, 4, 5], sum is 2 + 4 + 5 = 11
    // Descending part = [5, 2], sum is 5 + 2 = 7
    // Since the ascending part has a larger sum, return 0.
    fmt.Println(compareBitonicSums([]int{2,4,5,2})) // 0
    // Example 3:
    // Input: nums = [1,2,4,3]
    // Output: -1
    // Explanation:
    // Peak element is nums[2] = 4
    // Ascending part = [1, 2, 4], sum is 1 + 2 + 4 = 7
    // Descending part = [4, 3], sum is 4 + 3 = 7
    // Since both parts have equal sums, return -1.
    fmt.Println(compareBitonicSums([]int{1,2,4,3})) // -1

    fmt.Println(compareBitonicSums([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(compareBitonicSums([]int{9,8,7,6,5,4,3,2,1})) // 1
}