package main

// 1128. Number of Equivalent Domino Pairs
// Given a list of dominoes, dominoes[i] = [a, b] is equivalent to dominoes[j] = [c, d] if 
// and only if either (a == c and b == d), or (a == d and b == c) - that is, 
// one domino can be rotated to be equal to another domino.

// Return the number of pairs (i, j) for which 0 <= i < j < dominoes.length, and dominoes[i] is equivalent to dominoes[j].

// Example 1:
// Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
// Output: 1

// Example 2:
// Input: dominoes = [[1,2],[1,2],[1,1],[1,2],[2,2]]
// Output: 3

// Constraints:
//     1 <= dominoes.length <= 4 * 10^4
//     dominoes[i].length == 2
//     1 <= dominoes[i][j] <= 9

import "fmt"
import "sort"

func numEquivDominoPairs(dominoes [][]int) int {
    res, magic := 0, map[[2]int]int{}
    for i := range dominoes {
        sort.Ints(dominoes[i])
        magic[[2]int{dominoes[i][0], dominoes[i][1]}]++
    }
    helper := func(n int) int  {
        if n < 2 { return 0 }
        n *= n-1
        return n / 2
    }
    for _, v := range magic {
        res += helper(v)
    }
    return res
}

func numEquivDominoPairs1(dominoes [][]int) int {
    res, set := 0, make(map[int]int)
    for _, v := range dominoes {
        if v[0] > v[1] {
            set[v[1] * 10 + v[0]]++
        } else {
            set[v[0] * 10 + v[1]]++
        }
    }
    for _, v := range set {
        if v > 1 {
            res += v * (v - 1) / 2
        }
    }
    return res
}

func numEquivDominoPairs2(dominoes [][]int) int {
    res, mp := 0, make(map[[2]int]int)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, d := range dominoes {
        key := [2]int{min(d[0], d[1]), max(d[0], d[1])}
        res += mp[key]
        mp[key]++
    }
    return res
}

func main() {
    // Example 1:
    // Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
    // Output: 1
    fmt.Println(numEquivDominoPairs([][]int{{1,2},{2,1},{3,4},{5,6}})) // 1
    // Example 2:
    // Input: dominoes = [[1,2],[1,2],[1,1],[1,2],[2,2]]
    // Output: 3
    fmt.Println(numEquivDominoPairs([][]int{{1,2},{1,2},{1,1},{1,2},{2,2}})) // 3

    fmt.Println(numEquivDominoPairs1([][]int{{1,2},{2,1},{3,4},{5,6}})) // 1
    fmt.Println(numEquivDominoPairs1([][]int{{1,2},{1,2},{1,1},{1,2},{2,2}})) // 3

    fmt.Println(numEquivDominoPairs2([][]int{{1,2},{2,1},{3,4},{5,6}})) // 1
    fmt.Println(numEquivDominoPairs2([][]int{{1,2},{1,2},{1,1},{1,2},{2,2}})) // 3
}