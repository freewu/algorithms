package main

// 3841. Palindromic Path Queries in a Tree
// You are given an undirected tree with n nodes labeled 0 to n - 1. 
// This is represented by a 2D array edges of length n - 1, where edges[i] = [ui, vi] indicates an undirected edge between nodes ui and vi.

// You are also given a string s of length n consisting of lowercase English letters, where s[i] represents the character assigned to node i.

// You are also given a string array queries, where each queries[i] is either:
//     1. "update ui c": Change the character at node ui to c. Formally, update s[ui] = c.
//     2. "query ui vi": Determine whether the string formed by the characters on the unique path from ui to vi (inclusive) can be rearranged into a palindrome.

// Return a boolean array answer, where answer[j] is true if the jth query of type "query ui vi"​​​​​​​ can be rearranged into a palindrome, and false otherwise.

// Example 1:
// Input: n = 3, edges = [[0,1],[1,2]], s = "aac", queries = ["query 0 2","update 1 b","query 0 2"]
// Output: [true,false]
// Explanation:
// "query 0 2": Path 0 → 1 → 2 gives "aac", which can be rearranged to form "aca", a palindrome. Thus, answer[0] = true.
// "update 1 b": Update node 1 to 'b', now s = "abc".
// "query 0 2": Path characters are "abc", which cannot be rearranged to form a palindrome. Thus, answer[1] = false.
// Thus, answer = [true, false].

// Example 2:
// Input: n = 4, edges = [[0,1],[0,2],[0,3]], s = "abca", queries = ["query 1 2","update 0 b","query 2 3","update 3 a","query 1 3"]
// Output: [false,false,true]
// Explanation:
// "query 1 2": Path 1 → 0 → 2 gives "bac", which cannot be rearranged to form a palindrome. Thus, answer[0] = false.
// "update 0 b": Update node 0 to 'b', now s = "bbca".
// "query 2 3": Path 2 → 0 → 3 gives "cba", which cannot be rearranged to form a palindrome. Thus, answer[1] = false.
// "update 3 a": Update node 3 to 'a', s = "bbca".
// "query 1 3": Path 1 → 0 → 3 gives "bba", which can be rearranged to form "bab", a palindrome. Thus, answer[2] = true.
// Thus, answer = [false, false, true].

// Constraints:
//     1 <= n == s.length <= 5 * 10^4
//     edges.length == n - 1
//     edges[i] = [ui, vi]
//     0 <= ui, vi <= n - 1
//     s consists of lowercase English letters.
//     The input is generated such that edges represents a valid tree.
//     1 <= queries.length <= 5 * 10^4​​​​​​​
//         queries[i] = "update ui c" or
//         queries[i] = "query ui vi"
//         0 <= ui, vi <= n - 1
//         c is a lowercase English letter.

import "fmt"
import "strings"
import "math/bits"
import "strconv"

type Fenwick []int

func NewFenwickTree(n int) Fenwick {
    return make(Fenwick, n + 1) // 使用下标 1 到 n
}

// a[i] ^= val
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f Fenwick) update(i int, val int) {
    for ; i < len(f); i += i & -i {
        f[i] ^= val
    }
}

// 计算前缀异或和 a[1] ^ ... ^ a[i]
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f Fenwick) pre(i int) (res int) {
    for ; i > 0; i &= i - 1 {
        res ^= f[i]
    }
    return
}

func palindromePath(n int, edges [][]int, s string, queries []string) []bool {
    graph := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        graph[x] = append(graph[x], y)
        graph[y] = append(graph[y], x)
    }
    res, mx := make([]bool,0), bits.Len(uint(n))
    pa, dep := make([][16]int, n), make([]int, n)
    timeIn ,timeOut := make([]int, n), make([]int, n) // DFS 时间戳
    clock := 0
    pathXorFromRoot := make([]int, n) // 从根开始的路径中的字母奇偶性的集合
    pathXorFromRoot[0] = 1 << (s[0] - 'a')
    var dfs func(int, int)
    dfs = func(x, p int) {
        pa[x][0] = p
        clock++
        timeIn[x] = clock
        for _, y := range graph[x] {
            if y != p {
                dep[y] = dep[x] + 1
                pathXorFromRoot[y] = pathXorFromRoot[x] ^ 1<<(s[y]-'a')
                dfs(y, x)
            }
        }
        timeOut[x] = clock
    }
    dfs(0, -1)
    for i := range mx - 1 {
        for x := range pa {
            p := pa[x][i]
            if p != -1 {
                pa[x][i+1] = pa[p][i]
            } else {
                pa[x][i+1] = -1
            }
        }
    }
    uptoDep := func(x, d int) int {
        for k := uint32(dep[x] - d); k > 0; k &= k - 1 {
            x = pa[x][bits.TrailingZeros32(k)]
        }
        return x
    }
    // 返回 x 和 y 的最近公共祖先
    getLCA := func(x, y int) int {
        if dep[x] > dep[y] {
            x, y = y, x
        }
        y = uptoDep(y, dep[x]) // 使 y 和 x 在同一深度
        if y == x { return x  }
        for i := mx - 1; i >= 0; i-- {
            px, py := pa[x][i], pa[y][i]
            if px != py {
                x, y = px, py // 同时往上跳 2^i 步
            }
        }
        return pa[x][0]
    }
    t := []byte(s)
    f := NewFenwickTree(n) // 注意树状数组是异或运算
    for _, q := range queries {
        if q[0] == 'u' {
            x, _ := strconv.Atoi(q[7 : len(q)-2])
            c := q[len(q)-1]
            val := 1<<(t[x]-'a') ^ 1<<(c-'a') // 擦除旧的，换上新的
            t[x] = c
            // 子树 x 全部异或 val，转换成对区间 [timeIn[x], timeOut[x]] 的差分更新
            f.update(timeIn[x], val)
            f.update(timeOut[x]+1, val)
        } else {
            q = q[6:]
            i := strings.IndexByte(q, ' ')
            x, _ := strconv.Atoi(q[:i])
            y, _ := strconv.Atoi(q[i+1:])
            lca := getLCA(x, y)
            val := pathXorFromRoot[x] ^ pathXorFromRoot[y] ^ f.pre(timeIn[x]) ^ f.pre(timeIn[y]) ^ 1<<(t[lca]-'a')
            res = append(res, val & (val - 1) == 0) // 至多一个字母的出现次数是奇数
        }
    }
    return res
}

var adj, deg, sizes, heavy, chain, queue [50_000]uint16
var bit, masks [50_001]uint32

func palindromePath1(n int, edges [][]int, s string, queries []string) []bool {
    sentinel := uint16(n)
    clear(adj[:n])
    clear(deg[:n])
    deg[0] = 2
    for _, edge := range edges {
        u, v := uint16(edge[0]), uint16(edge[1])
        adj[u] ^= v
        adj[v] ^= u
        deg[u]++
        deg[v]++
    }
    q := 0
    for u := range uint16(n) {
        sizes[u] = 1
        heavy[u] = sentinel
        if deg[u] == 1 {
            queue[q] = u
            q++
        }
    }
    for i := 0; i < q; i++ {
        u := queue[i]
        v := adj[u]
        size := sizes[u]
        sizes[v] += size
        if child := heavy[v]; child == sentinel || size > sizes[child] {
            heavy[v] = u
        }
        adj[v] ^= u
        deg[v]--
        if deg[v] == 1 {
            queue[q] = v
            q++
        }
    }
    queue[q] = 0
    q++
    idxs := deg[:n]
    clear(idxs)
    var pos uint16
    for i := q - 1; i >= 0; i-- {
        u := queue[i]
        if idxs[u] != 0 {
            continue
        }
        i := u
        for u != sentinel {
            pos++
            chain[u] = i
            idxs[u] = pos
            masks[pos] = 1 << (s[u] - 'a')
            u = heavy[u]
        }
    }
    copy(bit[:n+1], masks[:n+1])
    for i := 1; i <= n; i++ {
        if j := i + (i & -i); j <= n {
            bit[j] ^= bit[i]
        }
    }
    res := make([]bool, 0, len(queries))
    for _, query := range queries {
        update := query[0] == 'u'
        ptr := 6
        if update {
            ptr = 7
        }
        var u uint16
        for ; query[ptr] != ' '; ptr++ {
            u = u*10 + uint16(query[ptr]-'0')
        }
        ptr++
        if update {
            mask := uint32(1) << (query[ptr] - 'a')
            pos := int(idxs[u])
            delta := mask ^ masks[pos]
            if delta != 0 {
                masks[pos] = mask
                for i := pos; i <= n; i += i & -i {
                    bit[i] ^= delta
                }
            }
            continue
        }
        var v uint16
        for ; ptr < len(query); ptr++ {
            v = v*10 + uint16(query[ptr]-'0')
        }
        var path uint32
        for chain[u] != chain[v] {
            if idxs[chain[u]] < idxs[chain[v]] {
                u, v = v, u
            }
            for right := idxs[u]; right > 0; right &= right - 1 {
                path ^= bit[right]
            }
            for left := idxs[chain[u]] - 1; left > 0; left &= left - 1 {
                path ^= bit[left]
            }
            u = adj[chain[u]]
        }
        if idxs[u] > idxs[v] {
            u, v = v, u
        }
        for right := idxs[v]; right > 0; right &= right - 1 {
            path ^= bit[right]
        }
        for left := idxs[u] - 1; left > 0; left &= left - 1 {
            path ^= bit[left]
        }
        res = append(res, path & (path - 1) == 0)
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, edges = [[0,1],[1,2]], s = "aac", queries = ["query 0 2","update 1 b","query 0 2"]
    // Output: [true,false]
    // Explanation:
    // "query 0 2": Path 0 → 1 → 2 gives "aac", which can be rearranged to form "aca", a palindrome. Thus, answer[0] = true.
    // "update 1 b": Update node 1 to 'b', now s = "abc".
    // "query 0 2": Path characters are "abc", which cannot be rearranged to form a palindrome. Thus, answer[1] = false.
    // Thus, answer = [true, false].
    fmt.Println(palindromePath(3, [][]int{{0,1},{1,2}}, "aac", []string{"query 0 2","update 1 b","query 0 2"})) // [true, false]
    // Example 2:
    // Input: n = 4, edges = [[0,1],[0,2],[0,3]], s = "abca", queries = ["query 1 2","update 0 b","query 2 3","update 3 a","query 1 3"]
    // Output: [false,false,true]
    // Explanation:
    // "query 1 2": Path 1 → 0 → 2 gives "bac", which cannot be rearranged to form a palindrome. Thus, answer[0] = false.
    // "update 0 b": Update node 0 to 'b', now s = "bbca".
    // "query 2 3": Path 2 → 0 → 3 gives "cba", which cannot be rearranged to form a palindrome. Thus, answer[1] = false.
    // "update 3 a": Update node 3 to 'a', s = "bbca".
    // "query 1 3": Path 1 → 0 → 3 gives "bba", which can be rearranged to form "bab", a palindrome. Thus, answer[2] = true.
    // Thus, answer = [false, false, true]. 
    fmt.Println(palindromePath(4, [][]int{{0,1},{0,2},{0,3}}, "abca", []string{"query 1 2","update 0 b","query 2 3","update 3 a","query 1 3"})) // [false, false, true]


    fmt.Println(palindromePath1(3, [][]int{{0,1},{1,2}}, "aac", []string{"query 0 2","update 1 b","query 0 2"})) // [true, false]
    fmt.Println(palindromePath1(4, [][]int{{0,1},{0,2},{0,3}}, "abca", []string{"query 1 2","update 0 b","query 2 3","update 3 a","query 1 3"})) // [false, false, true]
}