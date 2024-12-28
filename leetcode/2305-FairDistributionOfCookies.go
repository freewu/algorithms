package main

// 2305. Fair Distribution of Cookies
// You are given an integer array cookies, where cookies[i] denotes the number of cookies in the ith bag. 
// You are also given an integer k that denotes the number of children to distribute all the bags of cookies to. 
// All the cookies in the same bag must go to the same child and cannot be split up.

// The unfairness of a distribution is defined as the maximum total cookies obtained by a single child in the distribution.

// Return the minimum unfairness of all distributions.

// Example 1:
// Input: cookies = [8,15,10,20,8], k = 2
// Output: 31
// Explanation: One optimal distribution is [8,15,8] and [10,20]
// - The 1st child receives [8,15,8] which has a total of 8 + 15 + 8 = 31 cookies.
// - The 2nd child receives [10,20] which has a total of 10 + 20 = 30 cookies.
// The unfairness of the distribution is max(31,30) = 31.
// It can be shown that there is no distribution with an unfairness less than 31.

// Example 2:
// Input: cookies = [6,1,3,2,2,4,1,2], k = 3
// Output: 7
// Explanation: One optimal distribution is [6,1], [3,2,2], and [4,1,2]
// - The 1st child receives [6,1] which has a total of 6 + 1 = 7 cookies.
// - The 2nd child receives [3,2,2] which has a total of 3 + 2 + 2 = 7 cookies.
// - The 3rd child receives [4,1,2] which has a total of 4 + 1 + 2 = 7 cookies.
// The unfairness of the distribution is max(7,7,7) = 7.
// It can be shown that there is no distribution with an unfairness less than 7.

// Constraints:
//     2 <= cookies.length <= 8
//     1 <= cookies[i] <= 10^5
//     2 <= k <= cookies.length

import "fmt"
import "sort"

// dfs  backtracking
func distributeCookies(cookies []int, k int) int {
    sort.Slice(cookies, func(i, j int) bool {
        return cookies[i] > cookies[j]
    })
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    children := make([]int, k)
    var dfs func(index int, unfairness int) int
    dfs = func(index int, unfairness int) int {
        if index < 0 { return unfairness }
        res := 1 << 31 - 1
        for i := 0; i < k; i++ {
            if max(children[i] + cookies[index], unfairness) >= res { continue }
            children[i] += cookies[index]
            res = min(res, dfs(index-1, max(children[i], unfairness)))
            children[i] -= cookies[index]
        }
        return res
    }
    return dfs(len(cookies) - 1, 0)
}

func distributeCookies1(cookies []int, k int) int {
    n := len(cookies)
    facts, values := make([][]int, 1 << n), make([]int, 1 << n)
    for i, v := range cookies {
        values[1<<i] = v
        for j := 0; j < 1 << i; j++ {
            values[1 << i|j] = values[j] + v
        }
    }
    for i := range facts {
        facts[i] = make([]int, k)
        for j := range facts[i] {
            facts[i][j] = -1
        }
    }
    var dfs func(mask int, i int) int
    dfs = func(mask int, i int) int {
        if mask == 0 { return 0 }
        if i == k { return 1 << 31 }
        if facts[mask][i] != -1 { return facts[mask][i] }
        cur := 1 << 31
        for sub := mask; sub > 0; sub = (sub - 1) & mask {
            c := values[sub]
            cur = min(cur, max(c, dfs(mask ^ sub, i + 1)))
        }
        facts[mask][i] = cur
        return cur
    }
    return dfs(1 << n - 1, 0)
}

func main() {
    // Example 1:
    // Input: cookies = [8,15,10,20,8], k = 2
    // Output: 31
    // Explanation: One optimal distribution is [8,15,8] and [10,20]
    // - The 1st child receives [8,15,8] which has a total of 8 + 15 + 8 = 31 cookies.
    // - The 2nd child receives [10,20] which has a total of 10 + 20 = 30 cookies.
    // The unfairness of the distribution is max(31,30) = 31.
    // It can be shown that there is no distribution with an unfairness less than 31.
    fmt.Println(distributeCookies([]int{8,15,10,20,8}, 2)) // 31
    // Example 2:
    // Input: cookies = [6,1,3,2,2,4,1,2], k = 3
    // Output: 7
    // Explanation: One optimal distribution is [6,1], [3,2,2], and [4,1,2]
    // - The 1st child receives [6,1] which has a total of 6 + 1 = 7 cookies.
    // - The 2nd child receives [3,2,2] which has a total of 3 + 2 + 2 = 7 cookies.
    // - The 3rd child receives [4,1,2] which has a total of 4 + 1 + 2 = 7 cookies.
    // The unfairness of the distribution is max(7,7,7) = 7.
    // It can be shown that there is no distribution with an unfairness less than 7.
    fmt.Println(distributeCookies([]int{6,1,3,2,2,4,1,2}, 3)) // 7

    fmt.Println(distributeCookies1([]int{8,15,10,20,8}, 2)) // 31
    fmt.Println(distributeCookies1([]int{6,1,3,2,2,4,1,2}, 3)) // 7
}