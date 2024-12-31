package main

// 2358. Maximum Number of Groups Entering a Competition
// You are given a positive integer array grades which represents the grades of students in a university. 
// You would like to enter all these students into a competition in ordered non-empty groups, such that the ordering meets the following conditions:
//     1. The sum of the grades of students in the ith group is less than the sum of the grades of students in the (i + 1)th group, 
//        for all groups (except the last).
//     2. The total number of students in the ith group is less than the total number of students in the (i + 1)th group, 
//        for all groups (except the last).

// Return the maximum number of groups that can be formed.

// Example 1:
// Input: grades = [10,6,12,7,3,5]
// Output: 3
// Explanation: The following is a possible way to form 3 groups of students:
// - 1st group has the students with grades = [12]. Sum of grades: 12. Student count: 1
// - 2nd group has the students with grades = [6,7]. Sum of grades: 6 + 7 = 13. Student count: 2
// - 3rd group has the students with grades = [10,3,5]. Sum of grades: 10 + 3 + 5 = 18. Student count: 3
// It can be shown that it is not possible to form more than 3 groups.

// Example 2:
// Input: grades = [8,8]
// Output: 1
// Explanation: We can only form 1 group, since forming 2 groups would lead to an equal number of students in both groups.

// Constraints:
//     1 <= grades.length <= 10^5
//     1 <= grades[i] <= 10^5

import "fmt"
import "math"

func maximumGroups(grades []int) int {
    return int((math.Sqrt(float64(1 + 8 * len(grades))) - 1.0) / 2.0)
}

func main() {
    // Example 1:
    // Input: grades = [10,6,12,7,3,5]
    // Output: 3
    // Explanation: The following is a possible way to form 3 groups of students:
    // - 1st group has the students with grades = [12]. Sum of grades: 12. Student count: 1
    // - 2nd group has the students with grades = [6,7]. Sum of grades: 6 + 7 = 13. Student count: 2
    // - 3rd group has the students with grades = [10,3,5]. Sum of grades: 10 + 3 + 5 = 18. Student count: 3
    // It can be shown that it is not possible to form more than 3 groups.
    fmt.Println(maximumGroups([]int{10,6,12,7,3,5})) // 3
    // Example 2:
    // Input: grades = [8,8]
    // Output: 1
    // Explanation: We can only form 1 group, since forming 2 groups would lead to an equal number of students in both groups.
    fmt.Println(maximumGroups([]int{8,8})) // 1
}