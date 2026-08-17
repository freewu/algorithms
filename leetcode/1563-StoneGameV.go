package main

// 1563. Stone Game V
// There are several stones arranged in a row, and each stone has an associated value which is an integer given in the array stoneValue.

// In each round of the game, Alice divides the row into two non-empty rows (i.e. left row and right row), 
// then Bob calculates the value of each row which is the sum of the values of all the stones in this row. 
// Bob throws away the row which has the maximum value, and Alice's score increases by the value of the remaining row. 
// If the value of the two rows are equal, Bob lets Alice decide which row will be thrown away. 
// The next round starts with the remaining row.

// The game ends when there is only one stone remaining. Alice's is initially zero.

// Return the maximum score that Alice can obtain.

// Example 1:
// Input: stoneValue = [6,2,3,4,5,5]
// Output: 18
// Explanation: In the first round, Alice divides the row to [6,2,3], [4,5,5]. The left row has the value 11 and the right row has value 14. Bob throws away the right row and Alice's score is now 11.
// In the second round Alice divides the row to [6], [2,3]. This time Bob throws away the left row and Alice's score becomes 16 (11 + 5).
// The last round Alice has only one choice to divide the row which is [2], [3]. Bob throws away the right row and Alice's score is now 18 (16 + 2). The game ends because only one stone is remaining in the row.

// Example 2:
// Input: stoneValue = [7,7,7,7,7,7,7]
// Output: 28

// Example 3:
// Input: stoneValue = [4]
// Output: 0

// Constraints:
//     1 <= stoneValue.length <= 500
//     1 <= stoneValue[i] <= 10^6

import "fmt"

func stoneGameV(stoneValue []int) int {
    type Key struct {
        r, l int
    }
    prefix := []int{ 0 }
    curr := 0
    for _, s := range stoneValue {
        curr += s
        prefix = append( prefix, curr )
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    memo := make(map[Key]int)
    var dp func(lb int, rb int) int
    dp = func(lb int, rb int) int {
        if rb == lb { return 0 }
        currKey := Key{lb, rb}
        if val, ok := memo[currKey]; ok { return val }
        res, curr, sl, sr := 0, 0, 0, 0
        for i := lb; i < rb; i++ {
            sl = prefix[i]-prefix[lb-1]
            sr = prefix[rb]-prefix[i]
            if sl < sr {
                curr = sl + dp(lb, i)
            } else if sl > sr {
                curr = sr + dp(i+1, rb)
            } else {
                curr = max(sl + dp(lb, i), sr + dp(i+1, rb))
            }
            res = max(res, curr)
        }
        memo[currKey] = res
        return res
    }
    return dp(1, len(prefix) - 1)
}

func stoneGameV1(stoneValue []int) int {
    n := len(stoneValue)
    prefix, dp := make([]int, n + 1), make([][]int, n)
    for i, v := range stoneValue {
        prefix[i + 1] = prefix[i] + v
    }
    for i := range dp {
        dp[i] = make([]int, n)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i == j { return 0 }
        if dp[i][j] != 0 { return dp[i][j] }
        res, a := 0, 0
        for k := i; k < j; k++ {
            a += stoneValue[k]
            b := prefix[j+1] - prefix[i] - a
            if a < b {
                if res >= a * 2 { continue }
                res = max(res, a + dfs(i, k))
            } else if a > b {
                if res >= b * 2 { break }
                res = max(res, b + dfs(k + 1, j))
            } else {
                res = max(res, max(a + dfs(i, k), b + dfs(k + 1, j)))
            }
        }
        dp[i][j] = res
        return res
    }
    return dfs(0, n - 1)
}

var psum [501]uint32
var pf, sf [502][502]uint32

func stoneGameV2(stones []int) int {
    n := len(stones)
    if n == 1 {
        return 0
    } else if n == 2 {
        return min(stones[0], stones[1])
    }
    for i := 1; i <= n; i++ {
        s := uint32(stones[i-1])
        psum[i] = psum[i-1] + s
        pf[i][i], sf[i][i] = s, s
        pf[i][i-1], sf[i+1][i] = 0, 0
    }
    res := uint32(0)
    for i := n - 1; i >= 1; i-- {
        k := i
        for j := i + 1; j <= n; j++ {
            t := psum[i-1] + psum[j]
            for psum[k]*2 < t {
                k++
            }
            var off int
            if psum[k]*2 == t {
                off++
            }
            res = max(pf[i][k-1+off], sf[k+1][j])
            next := psum[j] - psum[i-1] + res
            pf[i][j] = max(pf[i][j-1], next)
            sf[i][j] = max(sf[i+1][j], next)
        }
    }
    return int(res)
}

func main() {
    // Example 1:
    // Input: stoneValue = [6,2,3,4,5,5]
    // Output: 18
    // Explanation: In the first round, Alice divides the row to [6,2,3], [4,5,5]. The left row has the value 11 and the right row has value 14. Bob throws away the right row and Alice's score is now 11.
    // In the second round Alice divides the row to [6], [2,3]. This time Bob throws away the left row and Alice's score becomes 16 (11 + 5).
    // The last round Alice has only one choice to divide the row which is [2], [3]. Bob throws away the right row and Alice's score is now 18 (16 + 2). The game ends because only one stone is remaining in the row.
    fmt.Println(stoneGameV([]int{6,2,3,4,5,5})) // 18
    // Example 2:
    // Input: stoneValue = [7,7,7,7,7,7,7]
    // Output: 28
    fmt.Println(stoneGameV([]int{7,7,7,7,7,7,7})) // 28
    // Example 3:
    // Input: stoneValue = [4]
    // Output: 0
    fmt.Println(stoneGameV([]int{4})) // 0

    fmt.Println(stoneGameV([]int{1,2,3,4,5,6,7,8,9})) // 35
    fmt.Println(stoneGameV([]int{9,8,7,6,5,4,3,2,1})) // 35

    fmt.Println(stoneGameV1([]int{6,2,3,4,5,5})) // 18
    fmt.Println(stoneGameV1([]int{7,7,7,7,7,7,7})) // 28
    fmt.Println(stoneGameV1([]int{4})) // 0
    fmt.Println(stoneGameV1([]int{1,2,3,4,5,6,7,8,9})) // 35
    fmt.Println(stoneGameV1([]int{9,8,7,6,5,4,3,2,1})) // 35

    fmt.Println(stoneGameV2([]int{6,2,3,4,5,5})) // 18
    fmt.Println(stoneGameV2([]int{7,7,7,7,7,7,7})) // 28
    fmt.Println(stoneGameV2([]int{4})) // 0
    fmt.Println(stoneGameV2([]int{1,2,3,4,5,6,7,8,9})) // 35
    fmt.Println(stoneGameV2([]int{9,8,7,6,5,4,3,2,1})) // 35
}