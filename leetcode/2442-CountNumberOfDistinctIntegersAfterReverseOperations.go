package main

// 2442. Count Number of Distinct Integers After Reverse Operations
// You are given an array nums consisting of positive integers.

// You have to take each integer in the array, reverse its digits, and add it to the end of the array. 
// You should apply this operation to the original integers in nums.

// Return the number of distinct integers in the final array.

// Example 1:
// Input: nums = [1,13,10,12,31]
// Output: 6
// Explanation: After including the reverse of each number, the resulting array is [1,13,10,12,31,1,31,1,21,13].
// The reversed integers that were added to the end of the array are underlined. Note that for the integer 10, after reversing it, it becomes 01 which is just 1.
// The number of distinct integers in this array is 6 (The numbers 1, 10, 12, 13, 21, and 31).

// Example 2:
// Input: nums = [2,2,2]
// Output: 1
// Explanation: After including the reverse of each number, the resulting array is [2,2,2,2,2,2].
// The number of distinct integers in this array is 1 (The number 2).

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^6

import "fmt"

func countDistinctIntegers(nums []int) int {
    set := make(map[int]bool)
    reverse := func(n int) int {
        res := 0
        for ; n > 0; n = n / 10 {
            res = res * 10 + n % 10
        }
        return res
    }
    for _, v := range nums {
        set[v] = true
        set[reverse(v)] = true
    }
    return len(set)
}

func countDistinctIntegers1(nums []int) int {
    set := make(map[int]struct{})
    reverse := func (n int) int{
        res := 0
        for n > 0 {
            res = res * 10 + n % 10
            n /= 10
        }
        return res
    }
    for _, v := range nums {
        set[v] = struct{}{}
        set[reverse(v)] = struct{}{}
    }
    return len(set)
}

func main() {
    // Example 1:
    // Input: nums = [1,13,10,12,31]
    // Output: 6
    // Explanation: After including the reverse of each number, the resulting array is [1,13,10,12,31,1,31,1,21,13].
    // The reversed integers that were added to the end of the array are underlined. Note that for the integer 10, after reversing it, it becomes 01 which is just 1.
    // The number of distinct integers in this array is 6 (The numbers 1, 10, 12, 13, 21, and 31).
    fmt.Println(countDistinctIntegers([]int{1,13,10,12,31})) // 6
    // Example 2:
    // Input: nums = [2,2,2]
    // Output: 1
    // Explanation: After including the reverse of each number, the resulting array is [2,2,2,2,2,2].
    // The number of distinct integers in this array is 1 (The number 2).
    fmt.Println(countDistinctIntegers([]int{2,2,2})) // 1

    fmt.Println(countDistinctIntegers([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(countDistinctIntegers([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(countDistinctIntegers1([]int{1,13,10,12,31})) // 6
    fmt.Println(countDistinctIntegers1([]int{2,2,2})) // 1
    fmt.Println(countDistinctIntegers1([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(countDistinctIntegers1([]int{9,8,7,6,5,4,3,2,1})) // 9
}