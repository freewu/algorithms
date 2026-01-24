package main

// 1877. Minimize Maximum Pair Sum in Array
// The pair sum of a pair (a,b) is equal to a + b. 
// The maximum pair sum is the largest pair sum in a list of pairs.
//     For example, if we have pairs (1,5), (2,3), and (4,4), 
//     the maximum pair sum would be max(1+5, 2+3, 4+4) = max(6, 5, 8) = 8.

// Given an array nums of even length n, pair up the elements of nums into n / 2 pairs such that:
//     Each element of nums is in exactly one pair, and
//     The maximum pair sum is minimized.

// Return the minimized maximum pair sum after optimally pairing up the elements.

// Example 1:
// Input: nums = [3,5,2,3]
// Output: 7
// Explanation: The elements can be paired up into pairs (3,3) and (5,2).
// The maximum pair sum is max(3+3, 5+2) = max(6, 7) = 7.

// Example 2:
// Input: nums = [3,5,4,2,4,6]
// Output: 8
// Explanation: The elements can be paired up into pairs (3,5), (4,4), and (6,2).
// The maximum pair sum is max(3+5, 4+4, 6+2) = max(8, 8, 8) = 8.

// Constraints:
//     n == nums.length
//     2 <= n <= 10^5
//     n is even.
//     1 <= nums[i] <= 10^5

import "fmt"
import "sort"

func minPairSum(nums []int) int {
    sort.Ints(nums)
    res, n := -1 << 31, len(nums)
    for i, j := 0, n - 1; i < n / 2; i, j = i + 1, j - 1 {
         if nums[i] + nums[j] > res {
            res = nums[i] + nums[j]
         }
    }
    return res
}
 
func minPairSum1(nums []int) int {
    sort.Ints(nums)
    res, n := -1 << 31, len(nums)
    for i := 0; i < n / 2; i++ {
        if nums[i] + nums[n - i - 1] > res {
            res = nums[i] + nums[n - i - 1]
        }
    }
    return res
}

func minPairSum2(nums []int) int {
    sort.Ints(nums)
    j := len(nums) - 1
    res := nums[0] + nums[j]
    j--
    for i := 1; i < len(nums) / 2; i++{
        if res < nums[i] + nums[j] {
            res = nums[i] + nums[j]
        }
        j--
    }
    return res
}

func minPairSum3(nums []int) int {
    sort.Ints(nums)
    n, res := len(nums), 0
    for i, v := range nums[:n / 2] {
        res = max(res, v + nums[n - 1 - i])
    }
    return res
}

func minPairSum4(nums []int) int {
    sort.Ints(nums)
    res, n := 0, len(nums)
    for i := 0; i < n / 2; i++ {
        sum := nums[i] + nums[n - 1 - i]
        if sum > res {
            res = sum
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,5,2,3]
    // Output: 7
    // Explanation: The elements can be paired up into pairs (3,3) and (5,2).
    // The maximum pair sum is max(3+3, 5+2) = max(6, 7) = 7.
    fmt.Println(minPairSum([]int{3,5,2,3})) // 7
    // Example 2:
    // Input: nums = [3,5,4,2,4,6]
    // Output: 8
    // Explanation: The elements can be paired up into pairs (3,5), (4,4), and (6,2).
    // The maximum pair sum is max(3+5, 4+4, 6+2) = max(8, 8, 8) = 8.
    fmt.Println(minPairSum([]int{3,5,4,2,4,6})) // 8

    fmt.Println(minPairSum([]int{1,2,3,4,5,6,7,8,9})) // 10
    fmt.Println(minPairSum([]int{9,8,7,6,5,4,3,2,1})) // 10

    fmt.Println(minPairSum1([]int{3,5,2,3})) // 7
    fmt.Println(minPairSum1([]int{3,5,4,2,4,6})) // 8
    fmt.Println(minPairSum1([]int{1,2,3,4,5,6,7,8,9})) // 10
    fmt.Println(minPairSum1([]int{9,8,7,6,5,4,3,2,1})) // 10

    fmt.Println(minPairSum2([]int{3,5,2,3})) // 7
    fmt.Println(minPairSum2([]int{3,5,4,2,4,6})) // 8
    fmt.Println(minPairSum2([]int{1,2,3,4,5,6,7,8,9})) // 10
    fmt.Println(minPairSum2([]int{9,8,7,6,5,4,3,2,1})) // 10

    fmt.Println(minPairSum3([]int{3,5,2,3})) // 7
    fmt.Println(minPairSum3([]int{3,5,4,2,4,6})) // 8
    fmt.Println(minPairSum3([]int{1,2,3,4,5,6,7,8,9})) // 10
    fmt.Println(minPairSum3([]int{9,8,7,6,5,4,3,2,1})) // 10

    fmt.Println(minPairSum4([]int{3,5,2,3})) // 7
    fmt.Println(minPairSum4([]int{3,5,4,2,4,6})) // 8
    fmt.Println(minPairSum4([]int{1,2,3,4,5,6,7,8,9})) // 10
    fmt.Println(minPairSum4([]int{9,8,7,6,5,4,3,2,1})) // 10
}