package main

// 920. Number of Music Playlists
// Your music player contains n different songs. 
// You want to listen to goal songs (not necessarily different) during your trip. 
// To avoid boredom, you will create a playlist so that:
//     Every song is played at least once.
//     A song can only be played again only if k other songs have been played.

// Given n, goal, and k, return the number of possible playlists that you can create. 
// Since the answer can be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 3, goal = 3, k = 1
// Output: 6
// Explanation: There are 6 possible playlists: [1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], and [3, 2, 1].

// Example 2:
// Input: n = 2, goal = 3, k = 0
// Output: 6
// Explanation: There are 6 possible playlists: [1, 1, 2], [1, 2, 1], [2, 1, 1], [2, 2, 1], [2, 1, 2], and [1, 2, 2].

// Example 3:
// Input: n = 2, goal = 3, k = 1
// Output: 2
// Explanation: There are 2 possible playlists: [1, 2, 1] and [2, 1, 2].
 
// Constraints:
//     0 <= k < n <= goal <= 100

import "fmt"

func numMusicPlaylists(n int, goal int, k int) int {
    dp, mod := make([][]int, goal + 1), 1000000007 // 定义 dp[i][j] 代表播放列表里面有 i 首歌，其中包含 j 首不同的歌曲
    for i := 0; i < goal + 1; i++ {
        dp[i] = make([]int, n+1)
    }
    dp[0][0] = 1
    for i := 1; i <= goal; i++ {
        for j := 1; j <= n; j++ {
            dp[i][j] = (dp[i-1][j-1] * (n- (j - 1))) % mod
            if j > k {
                dp[i][j] = (dp[i][j] + (dp[i-1][j] * (j - k) ) % mod) % mod
            }
        }
    }
    return dp[goal][n]
}

func numMusicPlaylists1(n int, goal int, k int) int {
    dp, mod := make([][]int, goal + 1), int(1e9 + 7)
    for i := 0; i <= goal; i++ {
        dp[i] = make([]int, n + 1)
    }
    dp[0][0] = 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= goal; i++ {
        for j := 1; j <= n; j++ {
            dp[i][j] += dp[i-1][j-1] * (n - j + 1)
            dp[i][j] += dp[i-1][j] * max(j - k, 0)
            dp[i][j] = dp[i][j] % mod
        }
    }
    return dp[goal][n]
}

func numMusicPlaylists2(n int, goal int, k int) int {
    dp, mod := make([][]int, goal+1), 1_000_000_007
    for i := range dp {
        dp[i]= make([]int, n + 1)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    var dfs func(total,unique int) int 
    dfs = func(total,unique int) int {
        if total == goal {
            if unique == n {
                return 1
            }
            return 0
        }
        if dp[total][unique] > -1 {
            return dp[total][unique]
        }
        res := 0
        if unique < n {
            res = res % mod + (n-unique) * dfs(total+1, unique+1) % mod
        }
        
        if unique > k {
            res = res % mod + (unique - k) * dfs(total+1, unique) % mod
        }
        dp[total][unique] = res % mod
        return res
    }
    return dfs(0, 0) % mod
}

func main() {
    // Example 1:
    // Input: n = 3, goal = 3, k = 1
    // Output: 6
    // Explanation: There are 6 possible playlists: [1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], and [3, 2, 1].
    fmt.Println(numMusicPlaylists(3,3,1)) // 6
    // Example 2:
    // Input: n = 2, goal = 3, k = 0
    // Output: 6
    // Explanation: There are 6 possible playlists: [1, 1, 2], [1, 2, 1], [2, 1, 1], [2, 2, 1], [2, 1, 2], and [1, 2, 2].
    fmt.Println(numMusicPlaylists(2,3,0)) // 6
    // Example 3:
    // Input: n = 2, goal = 3, k = 1
    // Output: 2
    // Explanation: There are 2 possible playlists: [1, 2, 1] and [2, 1, 2].
    fmt.Println(numMusicPlaylists(2,3,1)) // 2

    fmt.Println(numMusicPlaylists1(3,3,1)) // 6
    fmt.Println(numMusicPlaylists1(2,3,0)) // 6
    fmt.Println(numMusicPlaylists1(2,3,1)) // 2

    fmt.Println(numMusicPlaylists2(3,3,1)) // 6
    fmt.Println(numMusicPlaylists2(2,3,0)) // 6
    fmt.Println(numMusicPlaylists2(2,3,1)) // 2
}