package main

// 724. Find Pivot Index
// Given an array of integers nums, calculate the pivot index of this array.
// The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.
// If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. 
// This also applies to the right edge of the array.
// Return the leftmost pivot index. If no such index exists, return -1.
 
// Example 1:
// Input: nums = [1,7,3,6,5,6]
// Output: 3
// Explanation:
// The pivot index is 3.
// Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
// Right sum = nums[4] + nums[5] = 5 + 6 = 11

// Example 2:
// Input: nums = [1,2,3]
// Output: -1
// Explanation:
// There is no index that satisfies the conditions in the problem statement.

// Example 3:
// Input: nums = [2,1,-1]
// Output: 0
// Explanation:
// The pivot index is 0.
// Left sum = 0 (no elements to the left of index 0)
// Right sum = nums[1] + nums[2] = 1 + -1 = 0

// Constraints:
//     1 <= nums.length <= 10^4
//     -1000 <= nums[i] <= 1000
 
import "fmt"

func pivotIndex(nums []int) int {
    rSum, lSum := 0, 0
    for _, v := range nums[1:] { // <--- 累加所有值(除 0 的位置)
        rSum += v
    }
    if rSum == 0 {
        return 0
    }
    for i := 1; i < len(nums); i++ {
        lSum += nums[i-1]
        rSum -= nums[i]
        if lSum == rSum {
            return i
        }
    }
    return -1
}

func pivotIndex1(nums []int) int {
    right, left := 0, 0
    for _, num := range nums {
        right += num
    }   
    for i := 0; i < len(nums); i++ {
        right -= nums[i]
        if left == right {
            return i
        }
        left += nums[i]
    }
    return -1
}

func pivotIndex2(nums []int) int {
    leftPrefixSum, rightPrefixSum := make([]int, len(nums)), make([]int, len(nums))
    leftPrefixSum[0] = nums[0]
    rightPrefixSum[len(nums) - 1] = nums[len(nums) - 1]
    for i := 1; i < len(nums); i++ {
        leftPrefixSum[i] = nums[i] + leftPrefixSum[i - 1]
        endIdx := len(nums) - i - 1 
        rightPrefixSum[endIdx] = nums[endIdx] + rightPrefixSum[endIdx + 1]
    }
    for i := 0; i < len(nums); i++ {
        if leftPrefixSum[i] - nums[i] == rightPrefixSum[i] - nums[i] {
            return i
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [1,7,3,6,5,6]
    // Output: 3
    // Explanation:
    // The pivot index is 3.
    // Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
    // Right sum = nums[4] + nums[5] = 5 + 6 = 11
    fmt.Println(pivotIndex([]int{1,7,3,6,5,6})) // 3
    // Example 2:
    // Input: nums = [1,2,3]
    // Output: -1
    // Explanation:
    // There is no index that satisfies the conditions in the problem statement.
    fmt.Println(pivotIndex([]int{1,2,3})) // -1
    // Example 3:
    // Input: nums = [2,1,-1]
    // Output: 0
    // Explanation:
    // The pivot index is 0.
    // Left sum = 0 (no elements to the left of index 0)
    // Right sum = nums[1] + nums[2] = 1 + -1 = 0
    fmt.Println(pivotIndex([]int{2,1,-1})) // 0

    fmt.Println(pivotIndex1([]int{1,7,3,6,5,6})) // 3
    fmt.Println(pivotIndex1([]int{1,2,3})) // -1
    fmt.Println(pivotIndex1([]int{2,1,-1})) // 0

    
    fmt.Println(pivotIndex2([]int{1,7,3,6,5,6})) // 3
    fmt.Println(pivotIndex2([]int{1,2,3})) // -1
    fmt.Println(pivotIndex2([]int{2,1,-1})) // 0
}