package main

// 3843. First Element with Unique Frequency
// You are given an integer array nums.

// Return an integer denoting the first element (scanning from left to right) in nums whose frequency is unique. 
// That is, no other integer appears the same number of times in nums. 
// If there is no such element, return -1.

// Example 1:
// Input: nums = [20,10,30,30]
// Output: 30
// Explanation:
// 20 appears once.
// 10 appears once.
// 30 appears twice.
// The frequency of 30 is unique because no other integer appears exactly twice.

// Example 2:
// Input: nums = [20,20,10,30,30,30]
// Output: 20
// Explanation:
// 20 appears twice.
// 10 appears once.
// 30 appears 3 times.
// The frequency of 20, 10, and 30 are unique. The first element that has unique frequency is 20.

// Example 3:
// Input: nums = [10,10,20,20]
// Output: -1
// Explanation:
// 10 appears twice.
// 20 appears twice.
// No element has a unique frequency.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func firstUniqueFreq(nums []int) int {
    if len(nums) == 0 { return - 1 }
    mp := make(map[int]int)
    for _, v := range nums {
        mp[v]++
    }
    freq := make(map[int]int)
    for _, v := range mp {
        freq[v]++
    }
    for _, v := range nums {
        count := mp[v]
        if freq[count] == 1 {
            return v
        }
    }
    return -1
}

func firstUniqueFreq1(nums []int) int {
    mp := make(map[int]int)
    for _, v := range nums {
        mp[v]++
    }
    freq := make(map[int]int)
    for _, v := range mp {
        freq[v]++
    }
    for _, v := range nums {
        if freq[mp[v]] == 1 {
            return v
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [20,10,30,30]
    // Output: 30
    // Explanation:
    // 20 appears once.
    // 10 appears once.
    // 30 appears twice.
    // The frequency of 30 is unique because no other integer appears exactly twice.
    fmt.Println(firstUniqueFreq([]int{20,10,30,30})) // 30
    // Example 2:
    // Input: nums = [20,20,10,30,30,30]
    // Output: 20
    // Explanation:
    // 20 appears twice.
    // 10 appears once.
    // 30 appears 3 times.
    // The frequency of 20, 10, and 30 are unique. The first element that has unique frequency is 20.
    fmt.Println(firstUniqueFreq([]int{20,20,10,30,30,30})) // 20
    // Example 3:
    // Input: nums = [10,10,20,20]
    // Output: -1
    // Explanation:
    // 10 appears twice.
    // 20 appears twice.
    // No element has a unique frequency.
    fmt.Println(firstUniqueFreq([]int{10,10,20,20})) // -1

    fmt.Println(firstUniqueFreq([]int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(firstUniqueFreq([]int{9,8,7,6,5,4,3,2,1})) // -1

    fmt.Println(firstUniqueFreq1([]int{20,10,30,30})) // 30
    fmt.Println(firstUniqueFreq1([]int{20,20,10,30,30,30})) // 20
    fmt.Println(firstUniqueFreq1([]int{10,10,20,20})) // -1
    fmt.Println(firstUniqueFreq1([]int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(firstUniqueFreq1([]int{9,8,7,6,5,4,3,2,1})) // -1
}