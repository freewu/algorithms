package main

// 4077. Minimize Inversions in a Prefix Tree
// You are given an array words of n unique strings containing only lowercase English letters from 'a' to 'j'. 
// You are also given an integer array target containing each index from 0 to n - 1 exactly once.

// Build a prefix tree from words:
//     1. The root represents the empty string "".
//     2. Each other node represents a different non-empty prefix of a word. 
//        The parent of a node is the prefix obtained by removing its last character.
//     3. The node representing the complete word words[i] is marked with index i. 
//        A marked node may also have children.

// You may choose the order of the children of each node independently. 
// You cannot change which node is a child's parent.

// Traverse the tree starting at the root. 
// When visiting a node, first record its index if it is marked. 
// Then visit its children in the chosen order, completing the traversal of each child before moving to the next. 
// This is a pre-order traversal.

// The recorded indices form an array traversal of length n. 
// Let pos[i] be the position of index i in target.

// An inversion is a pair of positions (a, b) such that 0 <= a < b < n and pos[traversal[a]] > pos[traversal[b]].

// Return the minimum possible number of inversions over all choices of child orderings.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2026/03/04/screenshot-2026-03-04-at-55332pm.png" style="width: 264px; height: 250px;">​​​​​​​​​​​​​​​​​​​​​
// Input: words = ["bad","bag","fig","fed"], target = [3,1,0,2]
// Output: 2
// Explanation:
// Visit the root's "f" child before its "b" child. 
// Within the "f" branch, visit "fed" before "fig". 
// Within the "b" branch, visit "bag" before "bad".
// This produces traversal = [3,2,1,0]:
// Traversal position | Word index | Word  | Position in target
//  0                 | 3          | "fed" | 0
//  1                 | 2          | "fig" | 3
//  2                 | 1          | "bag" | 1
//  3                 | 0          | "bad" | 2
// The positions in target are [0,3,1,2]. 
// The inversions are the pairs of traversal positions (1,2) and (1,3), since 3 > 1 and 3 > 2.
// No child ordering produces fewer than 2 inversions.

// Example 2:
// Input: words = ["ab","ac","ba","bc"], target = [0,3,1,2]
// Output: 1
// Explanation:
// Visit the root's "a" child before its "b" child. Within these branches, visit "ab" before "ac", and "bc" before "ba".
// This produces traversal = [0,1,3,2]:
// Traversal position | Word index | Word  | Position in target
//  0                 | 0          | "ab" | 0
//  1                 | 1          | "ac" | 2
//  2                 | 3          | "bc" | 1
//  3                 | 2          | "ba" | 3
// The positions in target are [0,2,1,3]. 
// The only inversion is the pair of traversal positions (1,2), since 2 > 1.
// No child ordering produces zero inversions.

// Constraints:
//     - 1 <= n == words.length <= 10^{4}
//     - 1 <= words[i].length <= 20
//     - words[i] contains only lowercase English letters from 'a' to 'j'​​​​​​​
//     - All words[i] are unique
//     - 0 <= target[i] <= n - 1
//     - target is a permutation of integers from 0 to n - 1

import "fmt"
import "slices"
import "sort"

func minInversions(words []string, target []int) int{
    res, n := 0, len(target)
    pos := make([]int, n)
    for i, v := range target {
        pos[v] = i
    }
    // 计算 a 在 b 左侧时，a 和 b 之间的逆序对个数
    // a 和 b 已经升序排序
    calc := func(a, b []int) int {
        res, n, m, i, j := 0, len(a), len(b), 0, 0
        for i < n && j < m {
            if a[i] > b[j] {
                // a 的子数组 [i, n-1] 中的数都大于 b[j] 
                res += n - i
                j++
            } else {
                i++
            }
        }
        return res
    }
    // 创建前缀树（字典树）
    type Node struct {
        son [10]*Node
        pos int
    }
    root := &Node{}
    for i, word := range words {
        curr := root
        for _, ch := range word {
            ch -= 'a'
            if curr.son[ch] == nil {
                curr.son[ch] = &Node{}
            }
            curr = curr.son[ch]
        }
        // 对于不是 word 的节点，pos 是 0
        curr.pos = pos[i] + 1 // +1，与 pos = 0 的节点区分开
    }
    // 返回 node 子树的 pos 列表（已从小到大排序）
    var dfs func(*Node) []int
    dfs = func(node *Node) []int {
        arr := [][]int{}
        for _, son := range node.son {
            if son != nil {
                arr = append(arr, dfs(son))
            }
        }
        // 预处理逆序对个数
        m := len(arr)
        inv := make([][]int, m)
        for i, p := range arr {
            inv[i] = make([]int, m)
            for j, q := range arr {
                if j != i {
                    inv[i][j] = calc(p, q)
                }
            }
        }
        u := 1 << m
        sumInv := make([][]int, m)
        for j := range sumInv {
            sum := make([]int, u)
            for i := range m {
                highBit := 1 << i
                for mask, s := range sum[:highBit] {
                    sum[highBit|mask] = s + inv[i][j]
                }
            }
            sumInv[j] = sum
        }
        f := make([]int, u)
        for i := 1; i < u; i++ {
            f[i] = 1 << 61
        }
        for s, fs := range f { // 之前选的下标集合是 s
            // 在 s 的右边添加 j
            for j := range m {
                if s >> j & 0x1 == 0 {
                    f[s|1<<j] = min(f[s|1 << j], fs + sumInv[j][s])
                }
            }
        }
        res += f[u-1]
        val := []int{}
        curPos := node.pos
        if curPos > 0 {
            val = append(val, curPos)
        }
        for _, sub := range arr {
            if curPos > 0 {
                // curPos 与子树之间的逆序对个数
                res += sort.SearchInts(sub, curPos)
            }
            val = append(val, sub...)
        }
        slices.Sort(val)
        return val
    }
    dfs(root)
    return res
}

func minInversions1(words []string, target []int) int {
    n := len(words)
    rank := make([]int, n)
    for i, v := range target {
        rank[v] = i
    }
    type Node struct {
        child [10]int32
        mark  int32
    }
    newNode := func() Node {
        return Node{
            child: [10]int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
            mark:  -1,
        }
    }
    totalChars := 0
    for _, w := range words {
        totalChars += len(w)
    }
    nodes := make([]Node, 1, totalChars + 1)
    nodes[0] = newNode()
    for i, word := range words {
        curr := int32(0)
        for _, ch := range word {
            c := ch - 'a'
            if nodes[curr].child[c] == -1 {
                nodes = append(nodes, newNode())
                nodes[curr].child[c] = int32(len(nodes) - 1)
            }
            curr = nodes[curr].child[c]
        }
        nodes[curr].mark = int32(i)
    }
    dp, bit := make([]int, 1 << 10), make([]int, 1 << 10)
    var rowSum [10][1 << 10]int
    for i := 0; i < 10; i++ {
        bit[1 << i] = i
    }
    countGreater := func(a, b []int) int {
        res, j := 0, 0
        for _, x := range a {
            for j < len(b) && b[j] < x {
                j++
            }
            res += j
        }
        return res
    }
    countLess := func(a []int, x int) int {
        low, high := 0, len(a)
        for low < high {
            mid := (low + high) / 2
            if a[mid] < x {
                low = mid + 1
            } else {
                high = mid
            }
        }
        return low
    }
    buildSortedList := func(childLists [][]int, ownRank, totalSize int) []int {
        var lists [11][]int
        var pos [11]int
        count := 0
        if ownRank != -1 {
            ownArr := [1]int{ ownRank }
            lists[0] = ownArr[:]
            count = 1
        }
        for _, l := range childLists {
            lists[count] = l
            count++
        }
        full := make([]int, 0, totalSize)
        for k := 0; k < totalSize; k++ {
            best := -1
            for i := 0; i < count; i++ {
                if pos[i] < len(lists[i]) {
                    if best == -1 || lists[i][pos[i]] < lists[best][pos[best]] {
                        best = i
                    }
                }
            }
            full = append(full, lists[best][pos[best]])
            pos[best]++
        }
        return full
    }
    computeCrossDP := func(lists [][]int) int {
        m := len(lists)
        var cost [10][10]int
        for i := 0; i < m; i++ {
            for j := i + 1; j < m; j++ {
                c := countGreater(lists[i], lists[j])
                cost[i][j] = c
                cost[j][i] = len(lists[i])*len(lists[j]) - c
            }
        }
        fullMask := (1 << m) - 1
        for i := 0; i < m; i++ {
            rowSum[i][0] = 0
            for mask := 1; mask <= fullMask; mask++ {
                index := mask & -mask
                j := bit[index]
                rowSum[i][mask] = rowSum[i][mask^index] + cost[j][i]
            }
        }
        dp[0] = 0
        for mask := 1; mask <= fullMask; mask++ {
            best := 1 << 30
            for i := 0; i < m; i++ {
                bit := 1 << i
                if mask&bit != 0 {
                    prev := mask ^ bit
                    cand := dp[prev] + rowSum[i][prev]
                    if cand < best {
                        best = cand
                    }
                }
            }
            dp[mask] = best
        }
        return dp[fullMask]
    }
    var dfs func(u int32) ([]int, int)
    dfs = func(u int32) ([]int, int) {
        node := &nodes[u]
        var childIdx [10]int32
        m := 0
        for _, c := range node.child {
            if c != -1 {
                childIdx[m] = c
                m++
            }
        }
        var childLists [10][]int
        var childCosts [10]int
        totalSize := 0
        for i := 0; i < m; i++ {
            list, cost := dfs(childIdx[i])
            childLists[i] = list
            childCosts[i] = cost
            totalSize += len(list)
        }
        totalCost := 0
        for i := 0; i < m; i++ {
            totalCost += childCosts[i]
        }
        ownRank := -1
        if node.mark != -1 {
            ownRank = rank[node.mark]
            less := 0
            for i := 0; i < m; i++ {
                less += countLess(childLists[i], ownRank)
            }
            totalCost += less
        }
        if m > 1 {
            totalCost += computeCrossDP(childLists[:m])
        }
        fullSize := totalSize
        if ownRank != -1 {
            fullSize++
        }
        full := buildSortedList(childLists[:m], ownRank, fullSize)
        return full, totalCost
    }
    _, res := dfs(0)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2026/03/04/screenshot-2026-03-04-at-55332pm.png" style="width: 264px; height: 250px;">​​​​​​​​​​​​​​​​​​​​​
    // Input: words = ["bad","bag","fig","fed"], target = [3,1,0,2]
    // Output: 2
    // Explanation:
    // Visit the root's "f" child before its "b" child. 
    // Within the "f" branch, visit "fed" before "fig". 
    // Within the "b" branch, visit "bag" before "bad".
    // This produces traversal = [3,2,1,0]:
    // Traversal position | Word index | Word  | Position in target
    //  0                 | 3          | "fed" | 0
    //  1                 | 2          | "fig" | 3
    //  2                 | 1          | "bag" | 1
    //  3                 | 0          | "bad" | 2
    // The positions in target are [0,3,1,2]. 
    // The inversions are the pairs of traversal positions (1,2) and (1,3), since 3 > 1 and 3 > 2.
    // No child ordering produces fewer than 2 inversions.
    fmt.Println(minInversions([]string{"bad","bag","fig","fed"}, []int{3,1,0,2})) // 2
    // Example 2:
    // Input: words = ["ab","ac","ba","bc"], target = [0,3,1,2]
    // Output: 1
    // Explanation:
    // Visit the root's "a" child before its "b" child. Within these branches, visit "ab" before "ac", and "bc" before "ba".
    // This produces traversal = [0,1,3,2]:
    // Traversal position | Word index | Word  | Position in target
    //  0                 | 0          | "ab" | 0
    //  1                 | 1          | "ac" | 2
    //  2                 | 3          | "bc" | 1
    //  3                 | 2          | "ba" | 3
    // The positions in target are [0,2,1,3]. 
    // The only inversion is the pair of traversal positions (1,2), since 2 > 1.
    // No child ordering produces zero inversions.
    fmt.Println(minInversions([]string{"ab","ac","ba","bc"}, []int{0,3,1,2})) // 1

    fmt.Println(minInversions1([]string{"bad","bag","fig","fed"}, []int{3,1,0,2})) // 2
    fmt.Println(minInversions1([]string{"ab","ac","ba","bc"}, []int{0,3,1,2})) // 1
}