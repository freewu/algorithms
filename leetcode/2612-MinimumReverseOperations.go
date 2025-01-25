package main

// 2612. Minimum Reverse Operations
// You are given an integer n and an integer p representing an array arr of length n where all elements are set to 0's, 
// except position p which is set to 1. 
// You are also given an integer array banned containing restricted positions. 
// Perform the following operation on arr:
//     Reverse a subarray with size k if the single 1 is not set to a position in banned.

// Return an integer array answer with n results 
// where the ith result is the minimum number of operations needed to bring the single 1 to position i in arr, or -1 if it is impossible.

// Example 1:
// Input: n = 4, p = 0, banned = [1,2], k = 4
// Output: [0,-1,-1,1]
// Explanation:
// Initially 1 is placed at position 0 so the number of operations we need for position 0 is 0.
// We can never place 1 on the banned positions, so the answer for positions 1 and 2 is -1.
// Perform the operation of size 4 to reverse the whole array.
// After a single operation 1 is at position 3 so the answer for position 3 is 1.

// Example 2:
// Input: n = 5, p = 0, banned = [2,4], k = 3
// Output: [0,-1,-1,-1,-1]
// Explanation
// Initially 1 is placed at position 0 so the number of operations we need for position 0 is 0.
// We cannot perform the operation on the subarray positions [0, 2] because position 2 is in banned.
// Because 1 cannot be set at position 2, it is impossible to set 1 at other positions in more operations.

// Example 3:
// Input: n = 4, p = 2, banned = [0,1,3], k = 1
// Output: [-1,-1,0,-1]
// Explanation:
// Perform operations of size 1 and 1 never changes its position.

// Constraints:
//     1 <= n <= 10^5
//     0 <= p <= n - 1
//     0 <= banned.length <= n - 1
//     0 <= banned[i] <= n - 1
//     1 <= k <= n 
//     banned[i] != p
//     all values in banned are unique 

import "fmt"
import "sort"

type UnionFind struct {
    fa []int
}

func NewUnionFind(n int) UnionFind {
    fa := make([]int, n)
    for i := range fa {
        fa[i] = i
    }
    return UnionFind{fa}
}

func (u *UnionFind) find(x int) int {
    if u.fa[x] != x {
        u.fa[x] = u.find(u.fa[x])
    }
    return u.fa[x]
}

func (u *UnionFind) merge(from, to int) {
    x, y := u.find(from), u.find(to)
    u.fa[x] = y
}

func minReverseOperations(n, p int, banned []int, k int) []int {
    ban := map[int]bool{p: true}
    for _, v := range banned {
        ban[v] = true
    }
    notBanned := [2][]int{}
    for i := 0; i < n; i++ {
        if !ban[i] {
            notBanned[i%2] = append(notBanned[i%2], i)
        }
    }
    notBanned[0] = append(notBanned[0], n)
    notBanned[1] = append(notBanned[1], n) // 哨兵
    ufs := [2]UnionFind{ NewUnionFind(len(notBanned[0])), NewUnionFind(len(notBanned[1])) }
    res := make([]int, n)
    for i := range res {
        res[i] = -1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    queue := []int{ p }
    for step := 0; len(queue) > 0; step++ {
        tmp := queue
        queue = nil
        for _, i := range tmp {
            res[i] = step
            // 从 mn 到 mx 的所有位置都可以翻转到
            mn, mx := max(i-k+1, k-i-1), min(i+k-1, n*2-k-i-1)
            a, u := notBanned[mn % 2], ufs[mn % 2]
            for j := u.find(sort.SearchInts(a, mn)); a[j] <= mx; j = u.find(j + 1) {
                queue = append(queue, a[j])
                u.merge(j, j + 1) // 删除 j
            }
        }
    }
    return res
}

func minReverseOperations1(n int, p int, banned []int, k int) []int {
    m := 1
    neven := n / 2 + n & 1
    for m < neven {
        m *= 2
    }
    segtree := [2][]bool{ make([]bool, m * 2), make([]bool, m * 2) }
    arr := [2][]int{ make([]int, neven), make([]int, n / 2)}
    for i := range arr {
        for j := range arr[i] {
            arr[i][j] = -2
        }
    }
    // mark marks the provided index with the provided value
    mark := func(segtree []bool, res []int, i, val int) {
        segtree[m+i] = true
        for k := (m + i) / 2; k >= 1; k /= 2 {
            segtree[k] = segtree[k*2] && segtree[k*2+1]
            if !segtree[k] { break }
        }
        if i < len(res) {
            res[i] = val
        }
    }
    for _, i := range banned {
        mark(segtree[i & 1], arr[i & 1], i / 2, -1)
    }
    // Mark out-of-bounds elements as seen
    for i := neven; i < m; i++ {
        mark(segtree[0], arr[0], i, -1)
    }
    for i := n / 2; i < m; i++ {
        mark(segtree[1], arr[1], i, -1)
    }
    type Pos struct { i, odd int }
    // Update marks the range
    // [lo,hi] is the current range in the segment tree
    // [qlo,qhi] is the range being updated by the caller
    // j is 0 if using even segtree, otherwise 1
    var update func(segtree []bool, res []int, next *[]Pos, i, lo, hi, qlo, qhi, val, odd int)
    update = func(segtree []bool, res []int, next *[]Pos, i, lo, hi, qlo, qhi, val, odd int) {
        if qhi < lo || qlo > hi { return }
        if lo >= qlo && hi <= qhi { // This range of the segment tree should be marked as "done"
            if segtree[i] { return } // If it is already done, return
            for j := lo; j <= hi; j++ { // Mark all nodes as done, adding unseen nodes to next iteration
                if res[j] != -2 { continue }
                res[j] = val
                mark(segtree, res, j, val)
                *next = append(*next, Pos{ j, odd })
            }
            return
        }
        // Split query into left/right
        mid := lo + (hi-lo)/2
        update(segtree, res, next, i*2, lo, mid, qlo, qhi, val, odd)
        update(segtree, res, next, i*2+1, mid+1, hi, qlo, qhi, val, odd)
    }
    reachableRange := func(n, i, k int) (int, int) {
        left, right := i - k + 1, i + k - 1
        if left < 0 {
            d := -left
            left += d * 2
        }
        if right >= n {
            d := right - n + 1
            right -= d * 2
        }
        return left, right
    }
    curr, next := []Pos{},  []Pos{}
    // Start by marking p as done
    // This will populate next with the first node.
    update(segtree[p&1], arr[p&1], &curr, 1, 0, m-1, p/2, p/2, 0, p&1)
    for steps := 1; len(curr) > 0; steps++ {
        next = next[:0]
        for _, x := range curr {
            ii := x.i*2 + x.odd
            l, r := reachableRange(n, ii, k)
            update(segtree[l&1], arr[l&1], &next, 1, 0, m-1, l/2, r/2, steps, l&1)
        }
        curr, next = next, curr
    }
    res := make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = arr[i&1][i/2]
    }
    for i := range res {
        if res[i] == -2 {
            res[i] = -1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, p = 0, banned = [1,2], k = 4
    // Output: [0,-1,-1,1]
    // Explanation:
    // Initially 1 is placed at position 0 so the number of operations we need for position 0 is 0.
    // We can never place 1 on the banned positions, so the answer for positions 1 and 2 is -1.
    // Perform the operation of size 4 to reverse the whole array.
    // After a single operation 1 is at position 3 so the answer for position 3 is 1.
    fmt.Println(minReverseOperations(4, 0, []int{1,2}, 4)) // [0,-1,-1,1]
    // Example 2:
    // Input: n = 5, p = 0, banned = [2,4], k = 3
    // Output: [0,-1,-1,-1,-1]
    // Explanation
    // Initially 1 is placed at position 0 so the number of operations we need for position 0 is 0.
    // We cannot perform the operation on the subarray positions [0, 2] because position 2 is in banned.
    // Because 1 cannot be set at position 2, it is impossible to set 1 at other positions in more operations.
    fmt.Println(minReverseOperations(5, 0, []int{2,4}, 3)) // [0,-1,-1,-1,-1]
    // Example 3:
    // Input: n = 4, p = 2, banned = [0,1,3], k = 1
    // Output: [-1,-1,0,-1]
    // Explanation:
    // Perform operations of size 1 and 1 never changes its position.
    fmt.Println(minReverseOperations(4, 2, []int{0,1,3}, 1)) // [-1,-1,0,-1]

    fmt.Println(minReverseOperations1(4, 0, []int{1,2}, 4)) // [0,-1,-1,1]
    fmt.Println(minReverseOperations1(5, 0, []int{2,4}, 3)) // [0,-1,-1,-1,-1]
    fmt.Println(minReverseOperations1(4, 2, []int{0,1,3}, 1)) // [-1,-1,0,-1]
}