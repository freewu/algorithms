package main

// 3802. Number of Ways to Paint Sheets
// You are given an integer n representing the number of sheets.

// You are also given an integer array limit of size m, where limit[i] is the maximum number of sheets that can be painted using color i.

// You must paint all n sheets under the following conditions:
//     1. Exactly two distinct colors are used.
//     2. Each color must cover a single contiguous segment of sheets.
//     3. The number of sheets painted with color i cannot exceed limit[i].

// Return an integer denoting the number of distinct ways to paint all sheets. Since the answer may be large, return it modulo 10^9 + 7.

// Note: Two ways differ if at least one sheet is painted with a different color.

// Example 1:
// Input: n = 4, limit = [3,1,2]
// Output: 6
// Explanation:​​​​​​​
// For each ordered pair (i, j), where color i is used for the first segment and color j for the second segment (i != j), a split of x and 4 - x is valid if 1 <= x <= limit[i] and 1 <= 4 - x <= limit[j].
// Valid pairs and counts are:
// (0, 1): x = 3
// (0, 2): x = 2, 3
// (1, 0): x = 1
// (2, 0): x = 1, 2
// Therefore, there are 6 valid ways in total.

// Example 2:
// Input: n = 3, limit = [1,2]
// Output: 2
// Explanation:
// For each ordered pair (i, j), where color i is used for the first segment and color j for the second segment (i != j), a split of x and 3 - x is valid if 1 <= x <= limit[i] and 1 <= 3 - x <= limit[j].
// Valid pairs and counts are:
// (0, 1): x = 1
// (1, 0): x = 2
// Hence, there are 2 valid ways in total.

// Example 3:
// Input: n = 3, limit = [2,2]
// Output: 4
// Explanation:
// For each ordered pair (i, j), where color i is used for the first segment and color j for the second segment (i != j), a split of x and 3 - x is valid if 1 <= x <= limit[i] and 1 <= 3 - x <= limit[j].
// Valid pairs and counts are:
// (0, 1): x = 1, 2
// (1, 0): x = 1, 2
// Therefore, there are 4 valid ways in total.

// Constraints:
//     2 <= n <= 10^9
//     2 <= m == limit.length <= 10^5
//     1 <= limit[i] <= 10^9

import "fmt"
import "sort"

func numberOfWays(n int, limit []int) int {
    res, sum, mod := 0, 0, 1_000_000_007
    for i := range limit {
        limit[i] = min(limit[i], n - 1)
        sum += limit[i]
    }
    sort.Ints(limit)
    i, j := 0, len(limit) - 1
    for i < j {
        if limit[i]+ limit[j] < n {
            sum -= limit[i]
            i++
        } else {
            sum -= limit[j]
            res = (res + sum - (n - limit[j] - 1) * (j - i)) % mod
            j--
        }
    }
    return (res * 2 % mod + mod) % mod // 保证结果非负
}

func main() {
    // Example 1:
    // Input: n = 4, limit = [3,1,2]
    // Output: 6
    // Explanation:​​​​​​​
    // For each ordered pair (i, j), where color i is used for the first segment and color j for the second segment (i != j), a split of x and 4 - x is valid if 1 <= x <= limit[i] and 1 <= 4 - x <= limit[j].
    // Valid pairs and counts are:
    // (0, 1): x = 3
    // (0, 2): x = 2, 3
    // (1, 0): x = 1
    // (2, 0): x = 1, 2
    // Therefore, there are 6 valid ways in total.
    fmt.Println(numberOfWays(4, []int{3,1,2})) // 6
    // Example 2:
    // Input: n = 3, limit = [1,2]
    // Output: 2
    // Explanation:
    // For each ordered pair (i, j), where color i is used for the first segment and color j for the second segment (i != j), a split of x and 3 - x is valid if 1 <= x <= limit[i] and 1 <= 3 - x <= limit[j].
    // Valid pairs and counts are:
    // (0, 1): x = 1
    // (1, 0): x = 2
    // Hence, there are 2 valid ways in total.
    fmt.Println(numberOfWays(3, []int{1,2})) // 2
    // Example 3:
    // Input: n = 3, limit = [2,2]
    // Output: 4
    // Explanation:
    // For each ordered pair (i, j), where color i is used for the first segment and color j for the second segment (i != j), a split of x and 3 - x is valid if 1 <= x <= limit[i] and 1 <= 3 - x <= limit[j].
    // Valid pairs and counts are:
    // (0, 1): x = 1, 2
    // (1, 0): x = 1, 2
    // Therefore, there are 4 valid ways in total.
    fmt.Println(numberOfWays(3, []int{2,2})) // 4

    fmt.Println(numberOfWays(4, []int{1,2,3,4,5,6,7,8,9})) // 168
    fmt.Println(numberOfWays(4, []int{9,8,7,6,5,4,3,2,1})) // 168
}