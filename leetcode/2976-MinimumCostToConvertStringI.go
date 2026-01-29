package main

// 2976. Minimum Cost to Convert String I
// You are given two 0-indexed strings source and target, 
// both of length n and consisting of lowercase English letters. 
// You are also given two 0-indexed character arrays original and changed, and an integer array cost, 
// where cost[i] represents the cost of changing the character original[i] to the character changed[i].

// You start with the string source. 
// In one operation, you can pick a character x from the string 
// and change it to the character y at a cost of z if there exists any index j such 
// that cost[j] == z, original[j] == x, and changed[j] == y.

// Return the minimum cost to convert the string source to the string target using any number of operations. 
// If it is impossible to convert source to target, return -1.

// Note that there may exist indices i, j such that original[j] == original[i] and changed[j] == changed[i].

// Example 1:
// Input: source = "abcd", target = "acbe", original = ["a","b","c","c","e","d"], changed = ["b","c","b","e","b","e"], cost = [2,5,5,1,2,20]
// Output: 28
// Explanation: To convert the string "abcd" to string "acbe":
// - Change value at index 1 from 'b' to 'c' at a cost of 5.
// - Change value at index 2 from 'c' to 'e' at a cost of 1.
// - Change value at index 2 from 'e' to 'b' at a cost of 2.
// - Change value at index 3 from 'd' to 'e' at a cost of 20.
// The total cost incurred is 5 + 1 + 2 + 20 = 28.
// It can be shown that this is the minimum possible cost.

// Example 2:
// Input: source = "aaaa", target = "bbbb", original = ["a","c"], changed = ["c","b"], cost = [1,2]
// Output: 12
// Explanation: To change the character 'a' to 'b' change the character 'a' to 'c' at a cost of 1, followed by changing the character 'c' to 'b' at a cost of 2, for a total cost of 1 + 2 = 3. To change all occurrences of 'a' to 'b', a total cost of 3 * 4 = 12 is incurred.

// Example 3:
// Input: source = "abcd", target = "abce", original = ["a"], changed = ["e"], cost = [10000]
// Output: -1
// Explanation: It is impossible to convert source to target because the value at index 3 cannot be changed from 'd' to 'e'.

// Constraints:
//     1 <= source.length == target.length <= 10^5
//     source, target consist of lowercase English letters.
//     1 <= cost.length == original.length == changed.length <= 2000
//     original[i], changed[i] are lowercase English letters.
//     1 <= cost[i] <= 10^6
//     original[i] != changed[i]

import "fmt"

func minimumCost(source, target string, original, changed []byte, cost []int) int64 {
    res, dis, inf := 0, [26][26]int{}, int(1e13)
    for i := range dis {
        for j := range dis[i] {
            if j != i {
                dis[i][j] = inf
            }
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, c := range cost {
        x, y := original[i] - 'a', changed[i] - 'a'
        dis[x][y] = min(dis[x][y], c)
    }
    for k := range dis {
        for i := range dis {
            for j := range dis {
                dis[i][j] = min(dis[i][j], dis[i][k] + dis[k][j])
            }
        }
    }
    for i, b := range source {
        res += dis[b-'a'][target[i]-'a']
    }
    if res >= inf {
        return -1
    }
    return int64(res)
}

// 
func minimumCost1(source string, target string, original []byte, changed []byte, cost []int) int64 {
    kind, inf := 26, (1 << 31) / 3
    min := func (x, y int) int { if x < y { return x; }; return y; }
    floyd := func() [][]int {
        g := make([][]int, kind)
        for i := 0; i < kind; i++ {
            g[i] = make([]int, kind)
            for j := 0; j < kind; j++ {
                if i != j {
                    g[i][j] = inf
                }
            }
        }
        for i, u := range original {
            u -= 'a'
            v, wt := changed[i]-'a', cost[i]
            g[u][v] = min(g[u][v], wt) // 有重边的情况下取花费最小的
        }
        f := g
        for k := 0; k < kind; k++ {
            for i := 0; i < kind; i++ {
                if f[i][k] == inf { // 剪枝,无法利用中间顶点优化i->j的路径花费
                    continue
                }
                for j := 0; j < kind; j++ {
                    f[i][j] = min(f[i][j], f[i][k]+f[k][j])
                }
            }
        }
        return f
    }
    f := floyd()
    res := int64(0)
    for i, s := range source {
        t := target[i]
        s -= 'a'
        t -= 'a'
        if uint8(s) == t {
            continue
        }
        if wt := f[s][t]; wt == inf {
            return -1
        } else {
            res += int64(wt)
        }
    }
    return res
}

func minimumCost2(source string, target string, original []byte, changed []byte, cost []int) int64 {
    res, n, inf := int64(0), 26, int64(1 << 61)
    dist := make([][]int64, n)
    for i := 0; i < n; i++ {
        dist[i] = make([]int64, n)
        for j := 0; j < n; j++ {
            if i == j {
                dist[i][j] = 0
            } else {
                dist[i][j] = inf
            }
        }
    }
    for i := range original {
        u, v, c := original[i] - 'a', changed[i] - 'a', int64(cost[i])
        if c < dist[u][v] {
            dist[u][v] = c
        }
    }
    for k := 0; k < n; k++ {
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                if dist[i][k]+dist[k][j] < dist[i][j] {
                    dist[i][j] = dist[i][k] + dist[k][j]
                }
            }
        }
    }
    for i := 0; i < len(source); i++ {
        u, v := source[i] - 'a', target[i] - 'a'
        if dist[u][v] >= inf { return -1 }
        res += dist[u][v]
    }
    return res
}

func main() {
    // Example 1:
    // Input: source = "abcd", target = "acbe", original = ["a","b","c","c","e","d"], changed = ["b","c","b","e","b","e"], cost = [2,5,5,1,2,20]
    // Output: 28
    // Explanation: To convert the string "abcd" to string "acbe":
    // - Change value at index 1 from 'b' to 'c' at a cost of 5.
    // - Change value at index 2 from 'c' to 'e' at a cost of 1.
    // - Change value at index 2 from 'e' to 'b' at a cost of 2.
    // - Change value at index 3 from 'd' to 'e' at a cost of 20.
    // The total cost incurred is 5 + 1 + 2 + 20 = 28.
    // It can be shown that this is the minimum possible cost.
    fmt.Println(minimumCost("abcd","acbe", []byte{'a','b','c','c','e','d'}, []byte{'b','c','b','e','b','e'}, []int{2,5,5,1,2,20})) // 28
    // Example 2:
    // Input: source = "aaaa", target = "bbbb", original = ["a","c"], changed = ["c","b"], cost = [1,2]
    // Output: 12
    // Explanation: To change the character 'a' to 'b' change the character 'a' to 'c' at a cost of 1, followed by changing the character 'c' to 'b' at a cost of 2, for a total cost of 1 + 2 = 3. To change all occurrences of 'a' to 'b', a total cost of 3 * 4 = 12 is incurred.
    fmt.Println(minimumCost("aaaa","bbbb", []byte{'a','c'}, []byte{'c','b'}, []int{1,2})) // 12
    // Example 3:
    // Input: source = "abcd", target = "abce", original = ["a"], changed = ["e"], cost = [10000]
    // Output: -1
    // Explanation: It is impossible to convert source to target because the value at index 3 cannot be changed from 'd' to 'e'.
    fmt.Println(minimumCost("abcd","abce", []byte{'a'}, []byte{'e'}, []int{10000})) // -1

    fmt.Println(minimumCost1("abcd","acbe", []byte{'a','b','c','c','e','d'}, []byte{'b','c','b','e','b','e'}, []int{2,5,5,1,2,20})) // 28
    fmt.Println(minimumCost1("aaaa","bbbb", []byte{'a','c'}, []byte{'c','b'}, []int{1,2})) // 12
    fmt.Println(minimumCost1("abcd","abce", []byte{'a'}, []byte{'e'}, []int{10000})) // -1

    fmt.Println(minimumCost2("abcd","acbe", []byte{'a','b','c','c','e','d'}, []byte{'b','c','b','e','b','e'}, []int{2,5,5,1,2,20})) // 28
    fmt.Println(minimumCost2("aaaa","bbbb", []byte{'a','c'}, []byte{'c','b'}, []int{1,2})) // 12
    fmt.Println(minimumCost2("abcd","abce", []byte{'a'}, []byte{'e'}, []int{10000})) // -1
}