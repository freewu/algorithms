package main

// 851. Loud and Rich
// There is a group of n people labeled from 0 to n - 1 where each person has a different amount of money and a different level of quietness.

// You are given an array richer where richer[i] = [ai, bi] indicates that ai has more money than bi and an integer array quiet where quiet[i] is the quietness of the ith person. 
// All the given data in richer are logically correct (i.e., the data will not lead you to a situation where x is richer than y and y is richer than x at the same time).

// Return an integer array answer where answer[x] = y if y is the least quiet person (that is, the person y with the smallest value of quiet[y]) among all people who definitely have equal to or more money than the person x.

// Example 1:
// Input: richer = [[1,0],[2,1],[3,1],[3,7],[4,3],[5,3],[6,3]], quiet = [3,2,5,4,6,1,7,0]
// Output: [5,5,2,5,4,5,6,7]
// Explanation: 
// answer[0] = 5.
// Person 5 has more money than 3, which has more money than 1, which has more money than 0.
// The only person who is quieter (has lower quiet[x]) is person 7, but it is not clear if they have more money than person 0.
// answer[7] = 7.
// Among all people that definitely have equal to or more money than person 7 (which could be persons 3, 4, 5, 6, or 7), the person who is the quietest (has lower quiet[x]) is person 7.
// The other answers can be filled out with similar reasoning.

// Example 2:
// Input: richer = [], quiet = [0]
// Output: [0]

// Constraints:
//     n == quiet.length
//     1 <= n <= 500
//     0 <= quiet[i] < n
//     All the values of quiet are unique.
//     0 <= richer.length <= n * (n - 1) / 2
//     0 <= ai, bi < n
//     ai != bi
//     All the pairs of richer are unique.
//     The observations in richer are all logically consistent.

import "fmt"

// dfs
func loudAndRich(richer [][]int, quiet []int) []int {
    list := make([][]int, len(quiet))
    for _, v := range richer {
        list[v[1]] = append(list[v[1]], v[0])
    }
    dict :=  make(map[int]int)
    var dfs func(n int) int
    dfs = func(n int) int {
        if v, ok := dict[n]; ok {
            return v
        }
        res := n
        for _, v := range list[n] {
            t := dfs(v)
            if quiet[t] < quiet[res] {
                res = t
            }
        }
        dict[n] = res
        return res
    }
    res := make([]int, len(quiet))
    for i := 0; i < len(res); i++  {
        res[i] = dfs(i)
    }
    return res
}

// bfs
func loudAndRich1(richer [][]int, quiet []int) []int {
    n := len(quiet)
    d, g := make([]int, n), make([][]int, n)
    for _, r := range richer {
        x, y := r[0], r[1]
        g[x] = append(g[x], y)
        d[y]++
    }
    res, queue := make([]int, n), []int{}
    for i, c := range d {
        if c == 0 {
            queue = append(queue, i)
        }
        res[i] = i
    }
    for len(queue) > 0 {
        x := queue[0]
        queue = queue[1:] // pop
        for _, y := range g[x] {
            if quiet[res[x]] < quiet[res[y]] {
                res[y] = res[x]
            }
            if d[y]--; d[y] == 0 {
                queue = append(queue, y)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: richer = [[1,0],[2,1],[3,1],[3,7],[4,3],[5,3],[6,3]], quiet = [3,2,5,4,6,1,7,0]
    // Output: [5,5,2,5,4,5,6,7]
    // Explanation: 
    // answer[0] = 5.
    // Person 5 has more money than 3, which has more money than 1, which has more money than 0.
    // The only person who is quieter (has lower quiet[x]) is person 7, but it is not clear if they have more money than person 0.
    // answer[7] = 7.
    // Among all people that definitely have equal to or more money than person 7 (which could be persons 3, 4, 5, 6, or 7), the person who is the quietest (has lower quiet[x]) is person 7.
    // The other answers can be filled out with similar reasoning.
    fmt.Println(loudAndRich([][]int{{1,0},{2,1},{3,1},{3,7},{4,3},{5,3},{6,3}}, []int{3,2,5,4,6,1,7,0})) // [5,5,2,5,4,5,6,7]
    // Example 2:
    // Input: richer = [], quiet = [0]
    // Output: [0]
    fmt.Println(loudAndRich([][]int{}, []int{0})) // [0]

    fmt.Println(loudAndRich1([][]int{{1,0},{2,1},{3,1},{3,7},{4,3},{5,3},{6,3}}, []int{3,2,5,4,6,1,7,0})) // [5,5,2,5,4,5,6,7]
    fmt.Println(loudAndRich1([][]int{}, []int{0})) // [0]
}