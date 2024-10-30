package main

// 3327. Check if DFS Strings Are Palindromes
// You are given a tree rooted at node 0, consisting of n nodes numbered from 0 to n - 1. 
// The tree is represented by an array parent of size n, where parent[i] is the parent of node i. 
// Since node 0 is the root, parent[0] == -1.

// You are also given a string s of length n, where s[i] is the character assigned to node i.

// Consider an empty string dfsStr, and define a recursive function dfs(int x) that takes a node x as a parameter and performs the following steps in order:

//     Iterate over each child y of x in increasing order of their numbers, and call dfs(y).
//     Add the character s[x] to the end of the string dfsStr.

// Note that dfsStr is shared across all recursive calls of dfs.

// You need to find a boolean array answer of size n, where for each index i from 0 to n - 1, you do the following:
//     Empty the string dfsStr and call dfs(i).
//     If the resulting string dfsStr is a palindrome, then set answer[i] to true. Otherwise, set answer[i] to false.

// Return the array answer.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/09/01/tree1drawio.png" />
// Input: parent = [-1,0,0,1,1,2], s = "aababa"
// Output: [true,true,false,true,true,true]
// Explanation:
// Calling dfs(0) results in the string dfsStr = "abaaba", which is a palindrome.
// Calling dfs(1) results in the string dfsStr = "aba", which is a palindrome.
// Calling dfs(2) results in the string dfsStr = "ab", which is not a palindrome.
// Calling dfs(3) results in the string dfsStr = "a", which is a palindrome.
// Calling dfs(4) results in the string dfsStr = "b", which is a palindrome.
// Calling dfs(5) results in the string dfsStr = "a", which is a palindrome.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2024/09/01/tree2drawio-1.png" />
// Input: parent = [-1,0,0,0,0], s = "aabcb"
// Output: [true,true,true,true,true]
// Explanation:
// Every call on dfs(x) results in a palindrome string.

// Constraints:
//     n == parent.length == s.length
//     1 <= n <= 10^5
//     0 <= parent[i] <= n - 1 for all i >= 1.
//     parent[0] == -1
//     parent represents a valid tree.
//     s consists only of lowercase English letters.

import "fmt"
import "sort"

func findAnswer(parent []int, s string) []bool {
    root, n, mod := -1, len(parent), 1_000_000_007
    res, forwardHash, backwardHash, lengthArr := make([]bool, n), make([]int, n), make([]int, n), make([]int, n)
    flynorpexel := make([][]int, n)
    for i := 0; i < n; i++ {
        if parent[i] == -1 {
            root = i;
        } else {
            flynorpexel[parent[i]] = append(flynorpexel[parent[i]], i)
        }
    }
    for i := 0; i < n; i++ { // Sort each node's children in ascending order of their node numbers
        sort.Ints(flynorpexel[i])
    }
    power := make([]int, n + 1) // Precompute the powers of BASE
    power[0] = 1
    for i := 1; i <= n; i++ {
        power[i] = (power[i - 1] * 127) % mod
    }
    var dfs func(node int)
    dfs = func(node int) {
        forward, backward, curr := 0, 0, 0
        for _, v := range  flynorpexel[node] { // Traverse all child nodes in ascending order and perform DFS
            dfs(v)
            forward = (forward * power[lengthArr[v]] + forwardHash[v]) % mod // Merge the forward hash of child nodes
            backward = (backward + backwardHash[v] * power[curr]) % mod // Merge the backward hash of child nodes
            curr += lengthArr[v] // Update the current length
        }
        charVal := int(s[node] - 'a') + 1; // Add the current node's character
        forward = (forward * power[1] + charVal) % mod // Update the forward hash
        backward = (backward + charVal * power[curr]) % mod // Update the backward hash
        curr++
        forwardHash[node], backwardHash[node], lengthArr[node] = forward, backward, curr
        res[node] = (forwardHash[node] == backwardHash[node]) // Determine if the current dfsStr is a palindrome
    }
    dfs(root) // Start DFS traversal and hash calculation
    return res
}

func findAnswer1(parent []int, s string) []bool {
    n := len(parent)
    g := make([][]int, n)
    for i := 1; i < n; i++ {
        p := parent[i]
        g[p] = append(g[p], i) // 由于 i 是递增的，所以 g[p] 必然是有序的，下面无需排序
    }
    dfsStr := make([]byte, n) // dfsStr 是后序遍历整棵树得到的字符串
    nodes := make([]struct{ begin, end int }, n) // nodes[i] 表示子树 i 的后序遍历的开始时间戳和结束时间戳+1（左闭右开区间）
    time := 0
    var dfs func(int)
    dfs = func(x int) {
        nodes[x].begin = time
        for _, y := range g[x] {
            dfs(y)
        }
        dfsStr[time] = s[x] // 后序遍历
        time++
        nodes[x].end = time
    }
    dfs(0)
    // Manacher 模板
    // 将 dfsStr 改造为 t，这样就不需要讨论 n 的奇偶性，因为新串 t 的每个回文子串都是奇回文串（都有回文中心）
    // dfsStr 和 t 的下标转换关系：
    // (dfsStr_i+1)*2 = ti
    // ti/2-1 = dfsStr_i
    // ti 为偶数，对应奇回文串（从 2 开始）
    // ti 为奇数，对应偶回文串（从 3 开始）
    t := append(make([]byte, 0, n * 2 + 3), '^')
    for _, c := range dfsStr {
        t = append(t, '#', c)
    }
    t = append(t, '#', '$')

    // 定义一个奇回文串的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余字符串的长度
    // halfLen[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子串的回文半径
    // 即 [i-halfLen[i]+1,i+halfLen[i]-1] 是 t 上的一个回文子串
    halfLen := make([]int, len(t)-2)
    halfLen[1] = 1
    // boxR 表示当前右边界下标最大的回文子串的右边界下标+1
    // boxM 为该回文子串的中心位置，二者的关系为 r=mid+halfLen[mid]
    boxM, boxR := 0, 0
    for i := 2; i < len(halfLen); i++ { // 循环的起止位置对应着原串的首尾字符
        hl := 1
        if i < boxR {
            // 记 i 关于 boxM 的对称位置 i'=boxM*2-i
            // 若以 i' 为中心的最长回文子串范围超出了以 boxM 为中心的回文串的范围（即 i+halfLen[i'] >= boxR）
            // 则 halfLen[i] 应先初始化为已知的回文半径 boxR-i，然后再继续暴力匹配
            // 否则 halfLen[i] 与 halfLen[i'] 相等
            hl = min(halfLen[boxM*2-i], boxR-i)
        }
        // 暴力扩展
        // 算法的复杂度取决于这部分执行的次数
        // 由于扩展之后 boxR 必然会更新（右移），且扩展的的次数就是 boxR 右移的次数
        // 因此算法的复杂度 = O(len(t)) = O(n)
        for t[i-hl] == t[i+hl] {
            hl++
            boxM, boxR = i, i+hl
        }
        halfLen[i] = hl
    }
    // t 中回文子串的长度为 hl*2-1
    // 由于其中 # 的数量总是比字母的数量多 1
    // 因此其在 dfsStr 中对应的回文子串的长度为 hl-1
    // 这一结论可用在 isPalindrome 中

    // 判断左闭右开区间 [l,r) 是否为回文串  0<=l<r<=n
    // 根据下标转换关系得到 dfsStr 的 [l,r) 子串在 t 中对应的回文中心下标为 l+r+1
    // 需要满足 halfLen[l+r+1]-1 >= r-l，即 halfLen[l+r+1] > r-l
    isPalindrome := func(l, r int) bool { return halfLen[l+r+1] > r-l }
    res := make([]bool, n)
    for i, p := range nodes {
        res[i] = isPalindrome(p.begin, p.end)
    }
    return res
}

func findAnswer2(parent []int, s string) []bool {
    n, mod, base := len(parent), 1_000_000_007, 29
    res, g := make([]bool, n), make([][]int, n)
    for i, v := range parent {
        if v != -1 {
            g[v] = append(g[v], i)
        }
    }
    var dfs func (u int) (int, int, int)
    dfs = func(u int) (int, int, int) {
        h, r, b := 0, 0, 1
        for _, v := range g[u] {
            vh, vr, vb := dfs(v)
            h = (h * vb + vh) % mod 
            r = (vr * b + r) % mod
            b = (b * vb) % mod 
        }
        h = (h * base + int(s[u])) % mod
        r = (int(s[u]) * b + r) % mod
        b = (b * base) % mod
        res[u] = (h == r)
        return h, r, b
    }
    dfs(0)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/09/01/tree1drawio.png" />
    // Input: parent = [-1,0,0,1,1,2], s = "aababa"
    // Output: [true,true,false,true,true,true]
    // Explanation:
    // Calling dfs(0) results in the string dfsStr = "abaaba", which is a palindrome.
    // Calling dfs(1) results in the string dfsStr = "aba", which is a palindrome.
    // Calling dfs(2) results in the string dfsStr = "ab", which is not a palindrome.
    // Calling dfs(3) results in the string dfsStr = "a", which is a palindrome.
    // Calling dfs(4) results in the string dfsStr = "b", which is a palindrome.
    // Calling dfs(5) results in the string dfsStr = "a", which is a palindrome.
    fmt.Println(findAnswer([]int{-1,0,0,1,1,2}, "aababa")) // [true,true,false,true,true,true]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2024/09/01/tree2drawio-1.png" />
    // Input: parent = [-1,0,0,0,0], s = "aabcb"
    // Output: [true,true,true,true,true]
    // Explanation:
    // Every call on dfs(x) results in a palindrome string.
    fmt.Println(findAnswer([]int{-1,0,0,0,0}, "aabcb")) // [true,true,true,true,true]

    fmt.Println(findAnswer1([]int{-1,0,0,1,1,2}, "aababa")) // [true,true,false,true,true,true]
    fmt.Println(findAnswer1([]int{-1,0,0,0,0}, "aabcb")) // [true,true,true,true,true]

    fmt.Println(findAnswer2([]int{-1,0,0,1,1,2}, "aababa")) // [true,true,false,true,true,true]
    fmt.Println(findAnswer2([]int{-1,0,0,0,0}, "aabcb")) // [true,true,true,true,true]
}