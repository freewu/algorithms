package main

// 1923. Longest Common Subpath
// There is a country of n cities numbered from 0 to n - 1. 
// In this country, there is a road connecting every pair of cities.

// There are m friends numbered from 0 to m - 1 who are traveling through the country. 
// Each one of them will take a path consisting of some cities. 
// Each path is represented by an integer array that contains the visited cities in order. 
// The path may contain a city more than once, but the same city will not be listed consecutively.

// Given an integer n and a 2D integer array paths where paths[i] is an integer array representing the path of the ith friend, 
// return the length of the longest common subpath that is shared by every friend's path, 
// or 0 if there is no common subpath at all.

// A subpath of a path is a contiguous sequence of cities within that path.

// Example 1:
// Input: n = 5, paths = [[0,1,2,3,4], [2,3,4], [4,0,1,2,3]]
// Output: 2
// Explanation: The longest common subpath is [2,3].

// Example 2:
// Input: n = 3, paths = [[0],[1],[2]]
// Output: 0
// Explanation: There is no common subpath shared by the three paths.

// Example 3:
// Input: n = 5, paths = [[0,1,2,3,4], [4,3,2,1,0]]
// Output: 1
// Explanation: The possible longest common subpaths are [0], [1], [2], [3], and [4]. All have a length of 1.

// Constraints:
//     1 <= n <= 10^5
//     m == paths.length
//     2 <= m <= 10^5
//     sum(paths[i].length) <= 10^5
//     0 <= paths[i][j] < n
//     The same city is not listed multiple times consecutively in paths[i].

import "fmt"

func longestCommonSubpath(n int, paths [][]int) int {
    mn := len(paths[0])
    for _, path := range paths {
        if len(path) < mn { mn = len(path) }
    }
    equal := func(arr1, arr2 []int) bool {
        if len(arr1) != len(arr2) { return false }
        for i := range arr1 {
            if arr1[i] != arr2[i] { return false }
        }
        return true
    }
    hasCommon := func(paths [][]int, length, n int) bool {
        n++
        mod, pow := int(1e9 + 7), int64(1)
        for i := 0; i < length-1; i++ {
            pow = (pow * int64(n)) % int64(mod)
        }
        mp := make(map[int64][][]int)
        for p := 0; p < len(paths); p++ {
            hash := int64(0)
            for i := 0; i < len(paths[p]); i++ {
                if i >= length { hash = (hash - pow * int64(paths[p][i-length])) % int64(mod) }
                hash = ((hash * int64(n) + int64(paths[p][i])) % int64(mod) + int64(mod)) % int64(mod)
                if i >= length-1 { mp[hash] = append(mp[hash], []int{p, i}) }
            }
        }
        for _, endIdxs := range mp {
            if len(endIdxs) < len(paths) { continue }
            common, friends := []int{}, make(map[int]bool)
            for _, endIdx := range endIdxs {
                i, j := endIdx[0], endIdx[1]
                if _, ok := friends[i]; ok { continue }
                if len(common) == 0 {
                    common = append(common, paths[i][j-length+1:j+1]...)
                } else {
                    if !equal(common, paths[i][j-length+1:j+1]) {
                        common = []int{}
                        break
                    }
                }
                friends[i] = true
            }
            if len(friends) == len(paths) && common != nil { return true }
        }
        return false
    }
    helper := func(n, mn int) int {
        left, right := 0, mn
        for left < right {
            mid := left + (right - left + 1) / 2
            if hasCommon(paths, mid, n) {
                left = mid
            } else {
                right = mid - 1
            }
        }
        return left
    }
    return helper(n, mn)
}

// // 解答错误 80 / 81 
// func longestCommonSubpath1(n int, paths [][]int) int {
//     h, p := make([]int, 100010),  make([]int, 100010)
//     count, inner := make(map[int]int), make(map[int]int)
//     left, right := 0, 100010
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     for _, v := range paths { // 得到最短的
//         right = min(right, len(v))
//     }
//     get := func(l, r int) int { return h[r] - h[l - 1] * p[r - l + 1] }
//     check := func(mid int) bool {
//         count, inner = make(map[int]int), make(map[int]int) // cnt.clear(); inner.clear();
//         p[0] = 1
//         for j := 0; j < len(paths); j++ {
//             n := len(paths[j])
//             for i := 1; i <= n; i++ {
//                 p[i], h[i] = p[i - 1] * 133331, h[i - 1] * 133331 + paths[j][i - 1]
//             }
//             for i := mid; i <= n; i++ {
//                 val := get(i - mid + 1, i)
//                 _, ok := inner[val]
//                 if !ok || inner[val] != j {
//                     inner[val]= j
//                     count[val]++
//                 }
//             }
//         }
//         mx := 0
//         for _, v := range count {
//             mx = max(mx, v)
//         }
//         return mx == len(paths)
//     }
//     for left < right {
//         mid := (left + right + 1) >> 1
//         if (check(mid)) {
//             left = mid
//         } else {
//             right = mid - 1
//         }
//     }
//     return left
// }

func main() {
    // Example 1:
    // Input: n = 5, paths = [[0,1,2,3,4], [2,3,4], [4,0,1,2,3]]
    // Output: 2
    // Explanation: The longest common subpath is [2,3].
    fmt.Println(longestCommonSubpath(5, [][]int{{0,1,2,3,4}, {2,3,4}, {4,0,1,2,3}})) // 2
    // Example 2:
    // Input: n = 3, paths = [[0],[1],[2]]
    // Output: 0
    // Explanation: There is no common subpath shared by the three paths.
    fmt.Println(longestCommonSubpath(3, [][]int{{0},{1},{2}})) // 0
    // Example 3:
    // Input: n = 5, paths = [[0,1,2,3,4], [4,3,2,1,0]]
    // Output: 1
    // Explanation: The possible longest common subpaths are [0], [1], [2], [3], and [4]. All have a length of 1.
    fmt.Println(longestCommonSubpath(5, [][]int{{0,1,2,3,4}, {4,3,2,1,0}})) // 1

    // fmt.Println(longestCommonSubpath1(5, [][]int{{0,1,2,3,4}, {2,3,4}, {4,0,1,2,3}})) // 2
    // fmt.Println(longestCommonSubpath1(3, [][]int{{0},{1},{2}})) // 0
    // fmt.Println(longestCommonSubpath1(5, [][]int{{0,1,2,3,4}, {4,3,2,1,0}})) // 1
}