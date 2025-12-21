package main

// 3784. Minimum Deletion Cost to Make All Characters Equal
// You are given a string s of length n and an integer array cost of the same length, where cost[i] is the cost to delete the ith character of s.

// You may delete any number of characters from s (possibly none), such that the resulting string is non-empty and consists of equal characters.

// Return an integer denoting the minimum total deletion cost required.

// Example 1:
// Input: s = "aabaac", cost = [1,2,3,4,1,10]
// Output: 11
// Explanation:
// Deleting the characters at indices 0, 1, 2, 3, 4 results in the string "c", which consists of equal characters, and the total cost is cost[0] + cost[1] + cost[2] + cost[3] + cost[4] = 1 + 2 + 3 + 4 + 1 = 11.

// Example 2:
// Input: s = "abc", cost = [10,5,8]
// Output: 13
// Explanation:
// Deleting the characters at indices 1 and 2 results in the string "a", which consists of equal characters, and the total cost is cost[1] + cost[2] = 5 + 8 = 13.

// Example 3:
// Input: s = "zzzzz", cost = [67,67,67,67,67]
// Output: 0
// Explanation:
// All characters in s are equal, so the deletion cost is 0.

// Constraints:
//     n == s.length == cost.length
//     1 <= n <= 10^5
//     1 <= cost[i] <= 10^9
//     s consists of lowercase English letters.

import "fmt"

// time: O(n), space: O(1)
func minCost(s string, cost []int) int64 {
    n := 'z' - 'a' + 1
    costs := make([]int64, n)
    sum, mx := int64(0), int64(0)
    for i, v := range cost {
        char := s[i] - 'a'
        costs[char] += int64(v)
        sum += int64(v)
        mx = max(mx, costs[char])
    }
    return sum - mx
}

func main() {
    // Example 1:
    // Input: s = "aabaac", cost = [1,2,3,4,1,10]
    // Output: 11
    // Explanation:
    // Deleting the characters at indices 0, 1, 2, 3, 4 results in the string "c", which consists of equal characters, and the total cost is cost[0] + cost[1] + cost[2] + cost[3] + cost[4] = 1 + 2 + 3 + 4 + 1 = 11.
    fmt.Println(minCost("aabaac", []int{1,2,3,4,1,10})) // 11
    // Example 2:
    // Input: s = "abc", cost = [10,5,8]
    // Output: 13
    // Explanation:
    // Deleting the characters at indices 1 and 2 results in the string "a", which consists of equal characters, and the total cost is cost[1] + cost[2] = 5 + 8 = 13.
    fmt.Println(minCost("abc", []int{10,5,8})) // 13
    // Example 3:
    // Input: s = "zzzzz", cost = [67,67,67,67,67]
    // Output: 0
    // Explanation:
    // All characters in s are equal, so the deletion cost is 0.
    fmt.Println(minCost("zzzzz", []int{67,67,67,67,67})) // 0

    fmt.Println(minCost("bluefrog", []int{1,2,3,4,5,6,7,8})) // 28
    fmt.Println(minCost("bluefrog", []int{8,7,6,5,4,3,2,1})) // 28
    fmt.Println(minCost("leetcode", []int{1,2,3,4,5,6,7,8})) // 23
    fmt.Println(minCost("leetcode", []int{8,7,6,5,4,3,2,1})) // 22

    // fmt.Println(minCost1("aabaac", []int{1,2,3,4,1,10})) // 11
    // fmt.Println(minCost1("abc", []int{10,5,8})) // 13
    // fmt.Println(minCost1("zzzzz", []int{67,67,67,67,67})) // 0
    // fmt.Println(minCost1("bluefrog", []int{1,2,3,4,5,6,7,8})) // 28
    // fmt.Println(minCost1("bluefrog", []int{8,7,6,5,4,3,2,1})) // 28
    // fmt.Println(minCost1("leetcode", []int{1,2,3,4,5,6,7,8})) // 23
    // fmt.Println(minCost1("leetcode", []int{8,7,6,5,4,3,2,1})) // 22
}