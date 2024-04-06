package main

// 1283. Find the Smallest Divisor Given a Threshold
// Given an array of integers nums and an integer threshold, we will choose a positive integer divisor, divide all the array by it, and sum the division's result. 
// Find the smallest divisor such that the result mentioned above is less than or equal to threshold.
// Each result of the division is rounded to the nearest integer greater than or equal to that element. 
//     (For example: 7/3 = 3 and 10/2 = 5).

// The test cases are generated so that there will be an answer.

// Example 1:
// Input: nums = [1,2,5,9], threshold = 6
// Output: 5
// Explanation: We can get a sum to 17 (1+2+5+9) if the divisor is 1. 
// If the divisor is 4 we can get a sum of 7 (1+1+2+3) and if the divisor is 5 the sum will be 5 (1+1+1+2). 

// Example 2:
// Input: nums = [44,22,33,11,1], threshold = 5
// Output: 44
 
// Constraints:
//     1 <= nums.length <= 5 * 10^4
//     1 <= nums[i] <= 10^6
//     nums.length <= threshold <= 10^6

import "fmt"
import "math"

func smallestDivisor(nums []int, threshold int) int {
    // 得到最大值 边界就是 [1, max(nums)]
    left, right := 1, 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, num := range nums {
        right = max(right, num)
    }
    // 判断是否小于阈值
    check := func (nums []int, threshold, mid int) bool {
        sum := 0.0
        // 你需要选择一个正整数作为除数，然后将数组里每个数都除以它，并对除法结果求和
        for _, num := range nums {
            sum += math.Ceil(float64(num) / float64(mid))
        }
        return int(sum) <= threshold
    }
    for left <= right {
        mid := left + (right - left) >> 1
        if check(nums, threshold, mid) {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return left
}

func main() {
    // Explanation: We can get a sum to 17 (1+2+5+9) if the divisor is 1. 
    // If the divisor is 4 we can get a sum of 7 (1+1+2+3) and if the divisor is 5 the sum will be 5 (1+1+1+2). 
    fmt.Println(smallestDivisor([]int{1,2,5,9}, 6)) // 5 
    fmt.Println(smallestDivisor([]int{44,22,33,11,1}, 5)) // 44
    fmt.Println(smallestDivisor([]int{19}, 5)) // 4
}