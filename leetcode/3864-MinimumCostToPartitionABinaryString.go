package main

// 3864. Minimum Cost to Partition a Binary String
// You are given a binary string s and two integers encCost and flatCost.

// For each index i, s[i] = '1' indicates that the ith element is sensitive, and s[i] = '0' indicates that it is not.

// The string must be partitioned into segments. Initially, the entire string forms a single segment.

// For a segment of length L containing X sensitive elements:

//     If X = 0, the cost is flatCost.
//     If X > 0, the cost is L * X * encCost.

// If a segment has even length, you may split it into two contiguous segments of equal length and the cost of this split is the sum of costs of the resulting segments.

// Return an integer denoting the minimum possible total cost over all valid partitions.

// Example 1:
// Input: s = "1010", encCost = 2, flatCost = 1
// Output: 6
// Explanation:
// The entire string s = "1010" has length 4 and contains 2 sensitive elements, giving a cost of 4 * 2 * 2 = 16.
// Since the length is even, it can be split into "10" and "10". Each segment has length 2 and contains 1 sensitive element, so each costs 2 * 1 * 2 = 4, giving a total of 8.
// Splitting both segments into four single-character segments yields the segments "1", "0", "1", and "0". A segment containing "1" has length 1 and exactly one sensitive element, giving a cost of 1 * 1 * 2 = 2, while a segment containing "0" has no sensitive elements and therefore costs flatCost = 1.
// ​​​​​​​The total cost is thus 2 + 1 + 2 + 1 = 6, which is the minimum possible total cost.

// Example 2:
// Input: s = "1010", encCost = 3, flatCost = 10
// Output: 12
// Explanation:
// The entire string s = "1010" has length 4 and contains 2 sensitive elements, giving a cost of 4 * 2 * 3 = 24.
// Since the length is even, it can be split into two segments "10" and "10".
// Each segment has length 2 and contains one sensitive element, so each costs 2 * 1 * 3 = 6, giving a total of 12, which is the minimum possible total cost.

// Example 3:
// Input: s = "00", encCost = 1, flatCost = 2
// Output: 2
// Explanation:
// The string s = "00" has length 2 and contains no sensitive elements, so storing it as a single segment costs flatCost = 2, which is the minimum possible total cost.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists only of '0' and '1'.
//     1 <= encCost, flatCost <= 10^5

import "fmt"

func minCost(s string, encCost, flatCost int) int64 {
    n := len(s)
    sum := make([]int, n + 1)
    for i, v := range s {
        sum[i+1] = sum[i] + int(v - '0')
    }
    // 计算子串 [l, r) 的最小费用，注意区间是左闭右开，方便计算
    var dfs func(l, r int) int 
    dfs = func(l, r int) int {
        // 不拆分
        res := flatCost
        if x := sum[r] - sum[l]; x > 0 {
            res = (r - l) * x * encCost
        }
        // 拆分
        if (r-l)%2 == 0 {
            m := (l + r) / 2
            res = min(res, dfs(l, m)+dfs(m, r))
        }
        return res
    }
    return int64(dfs(0, n))
}

func minCost1(s string, encCost int, flatCost int) int64 {
    mp := map[string]int64{}
    var rec func(s string ) int64
    rec = func(s string) int64 {
        if _, exists := mp[s]; exists {
            return mp[s]
        }
        n, cost, numSensitive := len(s), int64(0), 0 
        for _, r := range s {
            if r == '1' {
                numSensitive++
            }
        }
        if numSensitive == 0 {
            cost = int64(flatCost) 
        } else {
            cost = int64(n) * int64(encCost) * int64(numSensitive)
        }
        if n %2 == 0 {
            cost = min(cost, rec(s[0:n/2])+rec(s[n/2:]))
        }
        mp[s] = cost 
        return cost
    }
    return rec(s)
}

func main() {
    // Example 1:
    // Input: s = "1010", encCost = 2, flatCost = 1
    // Output: 6
    // Explanation:
    // The entire string s = "1010" has length 4 and contains 2 sensitive elements, giving a cost of 4 * 2 * 2 = 16.
    // Since the length is even, it can be split into "10" and "10". Each segment has length 2 and contains 1 sensitive element, so each costs 2 * 1 * 2 = 4, giving a total of 8.
    // Splitting both segments into four single-character segments yields the segments "1", "0", "1", and "0". A segment containing "1" has length 1 and exactly one sensitive element, giving a cost of 1 * 1 * 2 = 2, while a segment containing "0" has no sensitive elements and therefore costs flatCost = 1.
    // ​​​​​​​The total cost is thus 2 + 1 + 2 + 1 = 6, which is the minimum possible total cost.
    fmt.Println(minCost("1010", 2, 1)) // 6
    // Example 2:
    // Input: s = "1010", encCost = 3, flatCost = 10
    // Output: 12
    // Explanation:
    // The entire string s = "1010" has length 4 and contains 2 sensitive elements, giving a cost of 4 * 2 * 3 = 24.
    // Since the length is even, it can be split into two segments "10" and "10".
    // Each segment has length 2 and contains one sensitive element, so each costs 2 * 1 * 3 = 6, giving a total of 12, which is the minimum possible total cost.
    fmt.Println(minCost("1010", 3, 10)) // 12   
    // Example 3:
    // Input: s = "00", encCost = 1, flatCost = 2
    // Output: 2
    // Explanation:
    // The string s = "00" has length 2 and contains no sensitive elements, so storing it as a single segment costs flatCost = 2, which is the minimum possible total cost.
    fmt.Println(minCost("00", 1, 2)) // 2

    fmt.Println(minCost("0000000000", 2, 1)) // 1
    fmt.Println(minCost("1111111111", 2, 1)) // 100
    fmt.Println(minCost("0000011111", 2, 1)) // 51
    fmt.Println(minCost("1111100000", 2, 1)) // 51
    fmt.Println(minCost("0101010101", 2, 1)) // 50
    fmt.Println(minCost("1010101010", 2, 1)) // 50

    fmt.Println(minCost1("1010", 2, 1)) // 6
    fmt.Println(minCost1("1010", 3, 10)) // 12   
    fmt.Println(minCost1("00", 1, 2)) // 2
    fmt.Println(minCost1("0000000000", 2, 1)) // 1
    fmt.Println(minCost1("1111111111", 2, 1)) // 100
    fmt.Println(minCost1("0000011111", 2, 1)) // 51
    fmt.Println(minCost1("1111100000", 2, 1)) // 51
    fmt.Println(minCost1("0101010101", 2, 1)) // 50
    fmt.Println(minCost1("1010101010", 2, 1)) // 50
}