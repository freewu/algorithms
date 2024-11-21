package main

// 2557. Maximum Number of Integers to Choose From a Range II
// You are given an integer array banned and two integers n and maxSum. 
// You are choosing some number of integers following the below rules:
//     1. The chosen integers have to be in the range [1, n].
//     2. Each integer can be chosen at most once.
//     3. The chosen integers should not be in the array banned.
//     4. The sum of the chosen integers should not exceed maxSum.

// Return the maximum number of integers you can choose following the mentioned rules.

// Constraints:
//     1 <= banned.length <= 10^5
//     1 <= banned[i] <= n <= 10^9
//     1 <= maxSum <= 10^15

import "fmt"
import "sort"

func maxCount(banned []int, n int, maxSum int64) int {
    banned = append(banned, []int{0, n + 1}...)
    sort.Ints(banned)
    res, ban := 0, []int{}
    for i, v := range banned {
        if i > 0 && v == banned[i - 1] { continue }
        ban = append(ban, v)
    }
    for k := 1; k < len(ban); k++ {
        i, j := ban[k-1], ban[k]
        left, right := 0, j - i - 1
        for left < right {
            mid := (left + right + 1) >> 1
            if int64((i + 1 + i + mid) * mid / 2) <= maxSum {
                left = mid
            } else {
                right = mid - 1
            }
        }
        res += left
        maxSum -= int64((i + 1 + i + left) * left / 2)
        if maxSum <= 0 { break }
    }
    return res
}

func maxCount1(banned []int, n int, maxSum int64) int {
    // Sort and remove duplicates from banned numbers
    sort.Ints(banned)
    removeDuplicates := func(arr []int) []int {
        if len(arr) <= 1 { return arr }
        unique := 0
        for i := 1; i < len(arr); i++ {
            if arr[i] != arr[unique] {
                unique++
                arr[unique] = arr[i]
            }
        }
        return arr[:unique + 1]
    }
    banned = removeDuplicates(banned)
    prefix := make([]int, len(banned) + 1) // Calculate prefix sums for efficient sum calculations
    for i := 0; i < len(banned); i++ {
        prefix[i + 1] = prefix[i] + banned[i]
    }
    binarySearch := func(arr []int, target int) int { // finds the index of the first element greater than target
        left, right := 0, len(arr)
        for left < right {
            mid := left + (right - left) / 2
            if arr[mid] > target {
                right = mid
            } else {
                left = mid + 1
            }
        }
        return left
    }
    isValidCount := func(count int) bool {
        totalSum := (int64(count) * int64(count + 1)) / 2 // Calculate sum of first 'count' numbers using arithmetic sequence formula
        bannedSum := int64(prefix[binarySearch(banned, count)]) // Subtract sum of banned numbers up to 'count'
        return totalSum - bannedSum <= maxSum
    }
    left, right := 1, n + 1
    for left < right { // binary search for the maximum valid count
        mid := left + (right - left) / 2  // Changed to prevent potential overflow
        if isValidCount(mid) {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left - 1 - binarySearch(banned, left - 1) // Subtract the count of banned numbers up to the found value
}

func main() {
    // Example 1:
    // Input: banned = [1,4,6], n = 6, maxSum = 4
    // Output: 1
    // Explanation: You can choose the integer 3.
    // 3 is in the range [1, 6], and do not appear in banned. The sum of the chosen integers is 3, which does not exceed maxSum.
    fmt.Println(maxCount([]int{1,4,6}, 6, 4)) // 1
    // Example 2:
    // Input: banned = [4,3,5,6], n = 7, maxSum = 18
    // Output: 3
    // Explanation: You can choose the integers 1, 2, and 7.
    // All these integers are in the range [1, 7], all do not appear in banned, and their sum is 18, which does not exceed maxSum.
    fmt.Println(maxCount([]int{4,3,5,6}, 7, 18)) // 3

    fmt.Println(maxCount1([]int{1,4,6}, 6, 4)) // 1
    fmt.Println(maxCount1([]int{4,3,5,6}, 7, 18)) // 3
}