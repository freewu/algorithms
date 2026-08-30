package main

// 4038. Count Integers Appearing in a Single Block
// You are given an integer array nums.

// An integer x is special if all occurrences of x in nums appear in a single contiguous block.

// Return the number of distinct special integers in nums.

// Example 1:
// Input: nums = [1,2,2,1]
// Output: 1
// Explanation:
// 1 appears at indices 0 and 3, forming two separate blocks, so it is not special.
// 2 appears in a single contiguous block at indices [1, 2], so it is special.
// Therefore, there is one special integer.

// Example 2:
// Input: nums = [3,3,1,2,2,1]
// Output: 2
// Explanation:
// 3 appears in a single contiguous block at indices [0, 1], so it is special.
// 1 appears at indices 2 and 5, forming two separate blocks, so it is not special.
// 2 appears in a single contiguous block at indices [3, 4], so it is special.
// Therefore, there are two special integers.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"

func countSpecialIntegers(nums []int) int {
    res, arr, mp := 0, []int{}, map[int]bool{}
    for _, v := range nums {
        if !mp[v] {
            arr = append(arr, v)
            mp[v] = true
        }
    }
    for i := 0; i < len(arr); i++ {
        isSpecial, changedNow, appear := true, false, false
        x := arr[i] 
        for j := 0; j < len(nums); j++ {
            y := nums[j]
            if x == y {
                appear = true
            }
            if x != y && appear {
                changedNow = true
            }
            if changedNow && x == y {
                isSpecial = false
            }
        }
        if isSpecial {
            res++
        }
    }
    return res
}

func countSpecialIntegers1(nums[]int) int {
    res, f, l, c := 0, make(map[int]int), make(map[int]int), make(map[int]int)
    for i, v :=range nums {
        if _, ok := f[v]; !ok {
            f[v] = i
        }
        l[v] = i
        c[v]++
    }
    for v, n := range c {
        if l[v] - f[v] + 1 == n {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,1]
    // Output: 1
    // Explanation:
    // 1 appears at indices 0 and 3, forming two separate blocks, so it is not special.
    // 2 appears in a single contiguous block at indices [1, 2], so it is special.
    // Therefore, there is one special integer.
    fmt.Println(countSpecialIntegers([]int{1,2,2,1})) // 1
    // Example 2:
    // Input: nums = [3,3,1,2,2,1]
    // Output: 2
    // Explanation:
    // 3 appears in a single contiguous block at indices [0, 1], so it is special.
    // 1 appears at indices 2 and 5, forming two separate blocks, so it is not special.
    // 2 appears in a single contiguous block at indices [3, 4], so it is special.
    // Therefore, there are two special integers. 
    fmt.Println(countSpecialIntegers([]int{3,3,1,2,2,1})) // 2

    fmt.Println(countSpecialIntegers([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(countSpecialIntegers([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(countSpecialIntegers1([]int{1,2,2,1})) // 1
    fmt.Println(countSpecialIntegers1([]int{3,3,1,2,2,1})) // 2
    fmt.Println(countSpecialIntegers1([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(countSpecialIntegers1([]int{9,8,7,6,5,4,3,2,1})) // 9
}