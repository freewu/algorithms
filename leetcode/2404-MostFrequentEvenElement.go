package main

// 2404. Most Frequent Even Element
// Given an integer array nums, return the most frequent even element.

// If there is a tie, return the smallest one. If there is no such element, return -1.

// Example 1:
// Input: nums = [0,1,2,2,4,4,1]
// Output: 2
// Explanation:
// The even elements are 0, 2, and 4. Of these, 2 and 4 appear the most.
// We return the smallest one, which is 2.

// Example 2:
// Input: nums = [4,4,4,9,2,4]
// Output: 4
// Explanation: 4 is the even element appears the most.

// Example 3:
// Input: nums = [29,47,21,41,13,37,25,7]
// Output: -1
// Explanation: There is no even element.

// Constraints:
//     1 <= nums.length <= 2000
//     0 <= nums[i] <= 10^5

import "fmt"

func mostFrequentEven(nums []int) int {
    res, mx := -1, -1<<31
    freq := make(map[int]int)
    for _, v := range nums {
        if v % 2 == 0 {
            freq[v]++
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for n, c := range freq {
        if c >= mx {
            if c == mx {
                res = min(res, n)
            } else {
                res = n
            }
            mx = c
        }
    }
    return res
}

func mostFrequentEven1(nums []int) int {
    mp := make(map[int]int)
    for _, v := range nums {
        if v % 2 == 0 {
            mp[v]++
        }
    }
    res, count := -1, 0 // 记录频数最大，值更小的偶数,  对应的频数
    for k, v := range mp { 
        if v > count {// 存在更加频繁的偶数
            res, count = k, v
        }
        if v == count && k < res { // 频数相同，值更小
            res = k
        }
    }
    return res
}

func mostFrequentEven2(nums []int) int {
    freq := make(map[int]int)
    hasEven := false
    for _, v := range nums {
        if v % 2 == 0 {
            freq[v]++
            hasEven = true
        }
    }
    if !hasEven { return -1 }
    mx, res := 0, 1 << 31
    for k, v := range freq {
        if v > mx || (v == mx && k < res) {
            mx, res = v, k
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,2,2,4,4,1]
    // Output: 2
    // Explanation:
    // The even elements are 0, 2, and 4. Of these, 2 and 4 appear the most.
    // We return the smallest one, which is 2.
    fmt.Println(mostFrequentEven([]int{0,1,2,2,4,4,1})) // 2
    // Example 2:
    // Input: nums = [4,4,4,9,2,4]
    // Output: 4
    // Explanation: 4 is the even element appears the most.
    fmt.Println(mostFrequentEven([]int{4,4,4,9,2,4})) // 4
    // Example 3:
    // Input: nums = [29,47,21,41,13,37,25,7]
    // Output: -1
    // Explanation: There is no even element.
    fmt.Println(mostFrequentEven([]int{29,47,21,41,13,37,25,7})) // -1

    fmt.Println(mostFrequentEven1([]int{0,1,2,2,4,4,1})) // 2
    fmt.Println(mostFrequentEven1([]int{4,4,4,9,2,4})) // 4
    fmt.Println(mostFrequentEven1([]int{29,47,21,41,13,37,25,7})) // -1

    fmt.Println(mostFrequentEven2([]int{0,1,2,2,4,4,1})) // 2
    fmt.Println(mostFrequentEven2([]int{4,4,4,9,2,4})) // 4
    fmt.Println(mostFrequentEven2([]int{29,47,21,41,13,37,25,7})) // -1
}