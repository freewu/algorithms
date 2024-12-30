package main

// 2342. Max Sum of a Pair With Equal Sum of Digits
// You are given a 0-indexed array nums consisting of positive integers. 
// You can choose two indices i and j, such that i != j, and the sum of digits of the number nums[i] is equal to that of nums[j].

// Return the maximum value of nums[i] + nums[j] that you can obtain over all possible indices i and j that satisfy the conditions.

// Example 1:
// Input: nums = [18,43,36,13,7]
// Output: 54
// Explanation: The pairs (i, j) that satisfy the conditions are:
// - (0, 2), both numbers have a sum of digits equal to 9, and their sum is 18 + 36 = 54.
// - (1, 4), both numbers have a sum of digits equal to 7, and their sum is 43 + 7 = 50.
// So the maximum sum that we can obtain is 54.

// Example 2:
// Input: nums = [10,12,19,14]
// Output: -1
// Explanation: There are no two numbers that satisfy the conditions, so we return -1.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"
import "sort"

func maximumSum(nums []int) int {
    res, mp := -1, make(map[int][]int)
    calcSum := func (n int) int {
        sum := 0
        for n != 0 {
            sum += n % 10
            n /= 10
        }
        return sum
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        sum := calcSum(v)
        arr, ok := mp[sum]
        if !ok {
            mp[sum] = []int{ v }
        } else {
            arr = append(arr, v) // Add the current number to the arr
            sort.Sort(sort.Reverse(sort.IntSlice(arr))) // Sort the arr to have the largest numbers at the beginning
            if len(arr) > 2 { // Keep only the two largest numbers if there are more than two
                arr = arr[:2]
            }
            mp[sum] = arr // Save back to the mp
            if len(arr) > 1 { // Also, update the result if we have at least two numbers
                res = max(res, arr[0] + arr[1])
            }
        }
    }
    return res
}

func maximumSum1(nums []int) int {
    res, mp := -1, [82]int{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        val, sum := v, 0
        for val > 0 {
            sum += val % 10 
            val /= 10
        }
        if mp[sum] > 0 {
            res = max(res, v + mp[sum])
        } 
        mp[sum] = max(mp[sum], v)
    } 
    return res
}

func main() {
    // Example 1:
    // Input: nums = [18,43,36,13,7]
    // Output: 54
    // Explanation: The pairs (i, j) that satisfy the conditions are:
    // - (0, 2), both numbers have a sum of digits equal to 9, and their sum is 18 + 36 = 54.
    // - (1, 4), both numbers have a sum of digits equal to 7, and their sum is 43 + 7 = 50.
    // So the maximum sum that we can obtain is 54.
    fmt.Println(maximumSum([]int{18,43,36,13,7})) // 54
    // Example 2:
    // Input: nums = [10,12,19,14]
    // Output: -1
    // Explanation: There are no two numbers that satisfy the conditions, so we return -1.
    fmt.Println(maximumSum([]int{10,12,19,14})) // -1

    fmt.Println(maximumSum1([]int{18,43,36,13,7})) // 54
    fmt.Println(maximumSum1([]int{10,12,19,14})) // -1
}