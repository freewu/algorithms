package main

// 4088. Difference Between Peaks and Valleys
// You are given an integer array nums of length n.

// An index i is called a peak if all of the following conditions hold:
//     1. 0 < i < n - 1
//     2. nums[i] > nums[i - 1]
//     3. nums[i] > nums[i + 1]

// An index i is called a valley if all of the following conditions hold:
//     1. 0 < i < n - 1
//     2. nums[i] < nums[i - 1]
//     3. nums[i] < nums[i + 1]

// Return an integer denoting the difference between the sum of all peak values and the sum of all valley values.

// Example 1:
// Input: nums = [1,5,2,4,1,3,0]
// Output: 9
// Explanation:
// - The peak values are nums[1] = 5, nums[3] = 4, and nums[5] = 3.
// - The valley values are nums[2] = 2 and nums[4] = 1.
// - The difference between the sum of the peak values and the sum of the valley values is (5 + 4 + 3) - (2 + 1) = 9. 
//   Hence, the answer is 9.

// Example 2:
// Input: nums = [-5,-6,-1,-2,5,3,2]
// Output: 12
// Explanation:
// - The peak values are nums[2] = -1 and nums[4] = 5.
// - The valley values are nums[1] = -6 and nums[3] = -2.
// - The difference between the sum of the peak values and the sum of the valley values is (-1 + 5) - (-6 + (-2)) = 12. 
//   Hence, the answer is 12.

// Example 3:
// Input: nums = [6,5,2,5,5,3,6]
// Output: -5
// Explanation:
// - There are no peak values in the array.
// - The valley values are nums[2] = 2 and nums[5] = 3.
// - The difference between the sum of the peak values and the sum of the valley values is (0) - (2 + 3) = -5. 
//   Hence, the answer is -5.

// Constraints:
//     - 3 <= n == nums.length <= 100
//     - -100 <= nums[i] <= 100

import "fmt"

func peakValleyDifference(nums []int) int {
    peak, valley := 0, 0
    isPeak := func(left, center, right int) bool {
        return center > left && center > right
    }
    isValley := func(left, center, right int) bool {
        return center < left && center < right
    }
    for i := 1; i < len(nums) - 1; i++ {
        left, center, right := nums[i-1], nums[i], nums[i+1]
        if isPeak(left, center, right) {
            peak += center
        } else if isValley(left, center, right) {
            valley += center
        }
    }
    return peak - valley
}

func main() {
    // Example 1:
    // Input: nums = [1,5,2,4,1,3,0]
    // Output: 9
    // Explanation:
    // - The peak values are nums[1] = 5, nums[3] = 4, and nums[5] = 3.
    // - The valley values are nums[2] = 2 and nums[4] = 1.
    // - The difference between the sum of the peak values and the sum of the valley values is (5 + 4 + 3) - (2 + 1) = 9. 
    //   Hence, the answer is 9.
    fmt.Println(peakValleyDifference([]int{1,5,2,4,1,3,0})) // 9
    // Example 2:
    // Input: nums = [-5,-6,-1,-2,5,3,2]
    // Output: 12
    // Explanation:
    // - The peak values are nums[2] = -1 and nums[4] = 5.
    // - The valley values are nums[1] = -6 and nums[3] = -2.
    // - The difference between the sum of the peak values and the sum of the valley values is (-1 + 5) - (-6 + (-2)) = 12. 
    //   Hence, the answer is 12.
    fmt.Println(peakValleyDifference([]int{-5,-6,-1,-2,5,3,2})) // 12
    // Example 3:
    // Input: nums = [6,5,2,5,5,3,6]
    // Output: -5
    // Explanation:
    // - There are no peak values in the array.
    // - The valley values are nums[2] = 2 and nums[5] = 3.
    // - The difference between the sum of the peak values and the sum of the valley values is (0) - (2 + 3) = -5. 
    //   Hence, the answer is -5.
    fmt.Println(peakValleyDifference([]int{6,5,2,5,5,3,6})) // -5

    fmt.Println(peakValleyDifference([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(peakValleyDifference([]int{9,8,7,6,5,4,3,2,1})) // 0
}
