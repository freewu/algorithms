package main

// 1871. Jump Game VII
// You are given a 0-indexed binary string s and two integers minJump and maxJump. 
// In the beginning, you are standing at index 0, which is equal to '0'. 
// You can move from index i to index j if the following conditions are fulfilled:
//     i + minJump <= j <= min(i + maxJump, s.length - 1), and
//     s[j] == '0'.

// Return true if you can reach index s.length - 1 in s, or false otherwise.

// Example 1:
// Input: s = "011010", minJump = 2, maxJump = 3
// Output: true
// Explanation:
// In the first step, move from index 0 to index 3. 
// In the second step, move from index 3 to index 5.

// Example 2:
// Input: s = "01101110", minJump = 2, maxJump = 3
// Output: false

// Constraints:
//     2 <= s.length <= 10^5
//     s[i] is either '0' or '1'.
//     s[0] == '0'
//     1 <= minJump <= maxJump < s.length

import "fmt"

// bfs
func canReach(s string, minJump int, maxJump int) bool {
    r, farthest := len(s) - 1, 0
    queue := []int{ farthest }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for len(queue) > 0 {
        l := queue[0]
        queue = queue[1:] // pop
        from, to := max(l + minJump, farthest + 1),  min(l + maxJump, r)
        for i := from; i <= to; i++ {
            if s[i] == '0' {
                if i == r { return true }
                queue = append(queue, i)
            }
        }
        farthest = l + maxJump
    }
    return false
}

func canReach1(s string, minJump, maxJump int) bool {
    n := len(s)
    sum := make([]int, n + 1)
    sum[1] = 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        sum[i + 1] = sum[i]
        if i >= minJump && s[i] == '0' && sum[i - minJump + 1] > sum[max(0, i - maxJump)] {
            sum[i + 1]++
        }
    }
    return sum[n] > sum[n - 1]
}

func canReach2(s string, minJump, maxJump int) bool {
    n, count := len(s), 0
    dp := make([]bool, n)
    dp[0] = true
    for i := minJump; i < n; i++ {
        if dp[i - minJump] {
            count++
        }
        if i > maxJump && dp[i - maxJump - 1] {
            count--
        }
        if s[i] == '0' && count > 0 {
            dp[i] = true
        }
    }
    return dp[n-1]
}

func main() {
    // Example 1:
    // Input: s = "011010", minJump = 2, maxJump = 3
    // Output: true
    // Explanation:
    // In the first step, move from index 0 to index 3. 
    // In the second step, move from index 3 to index 5.
    fmt.Println(canReach("011010", 2, 3)) // true
    // Example 2:
    // Input: s = "01101110", minJump = 2, maxJump = 3
    // Output: false
    fmt.Println(canReach("01101110", 2, 3)) // false

    fmt.Println(canReach1("011010", 2, 3)) // true
    fmt.Println(canReach1("01101110", 2, 3)) // false

    fmt.Println(canReach2("011010", 2, 3)) // true
    fmt.Println(canReach2("01101110", 2, 3)) // false
}