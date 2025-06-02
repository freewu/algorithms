package main

// 3569. Maximize Count of Distinct Primes After Split
// You are given an integer array nums having length n and a 2D integer array queries where queries[i] = [idx, val].

// For each query:
//     1. Update nums[idx] = val.
//     2. Choose an integer k with 1 <= k < n to split the array into the non-empty prefix nums[0..k-1] 
//        and suffix nums[k..n-1] such that the sum of the counts of distinct prime values in each part is maximum.

// Note: The changes made to the array in one query persist into the next query.

// Return an array containing the result for each query, in the order they are given.

// Example 1:
// Input: nums = [2,1,3,1,2], queries = [[1,2],[3,3]]
// Output: [3,4]
// Explanation:
// Initially nums = [2, 1, 3, 1, 2].
// After 1st query, nums = [2, 2, 3, 1, 2]. Split nums into [2] and [2, 3, 1, 2]. [2] consists of 1 distinct prime and [2, 3, 1, 2] consists of 2 distinct primes. Hence, the answer for this query is 1 + 2 = 3.
// After 2nd query, nums = [2, 2, 3, 3, 2]. Split nums into [2, 2, 3] and [3, 2] with an answer of 2 + 2 = 4.
// The output is [3, 4].

// Example 2:
// Input: nums = [2,1,4], queries = [[0,1]]
// Output: [0]
// Explanation:
// Initially nums = [2, 1, 4].
// After 1st query, nums = [1, 1, 4]. There are no prime numbers in nums, hence the answer for this query is 0.
// The output is [0].

// Constraints:
//     2 <= n == nums.length <= 5 * 10^4
//     1 <= queries.length <= 5 * 10^4
//     1 <= nums[i] <= 10^5
//     0 <= queries[i][0] < nums.length
//     1 <= queries[i][1] <= 10^5

import "fmt"
// import "github.com/emirpasic/gods/v2/trees/redblacktree"

// const mx int = 1e5

// var np = [mx + 1]bool{true, true}

// func init() {
//     for i := 2; i <= mx; i++ {
//         if !np[i] {
//             for j := i * i; j <= mx; j += i {
//                 np[j] = true
//             }
//         }
//     }
// }

// type lazySeg []struct {
//     l, r int
//     mx   int
//     todo int
// }

// func mergeInfo(a, b int) int {
//     return max(a, b)
// }

// const todoInit = 0

// func mergeTodo(f, old int) int {
//     return f + old
// }

// func (t lazySeg) apply(o int, f int) {
//     cur := &t[o]
//     cur.mx += f
//     cur.todo = mergeTodo(f, cur.todo)
// }

// func (t lazySeg) maintain(o int) {
//     t[o].mx = mergeInfo(t[o<<1].mx, t[o<<1|1].mx)
// }

// func (t lazySeg) spread(o int) {
//     f := t[o].todo
//     if f == todoInit {
//         return
//     }
//     t.apply(o<<1, f)
//     t.apply(o<<1|1, f)
//     t[o].todo = todoInit
// }

// func (t lazySeg) build(a []int, o, l, r int) {
//     t[o].l, t[o].r = l, r
//     t[o].todo = todoInit
//     if l == r {
//         t[o].mx = a[l]
//         return
//     }
//     m := (l + r) >> 1
//     t.build(a, o<<1, l, m)
//     t.build(a, o<<1|1, m+1, r)
//     t.maintain(o)
// }

// func (t lazySeg) update(o, l, r int, f int) {
//     if l <= t[o].l && t[o].r <= r {
//         t.apply(o, f)
//         return
//     }
//     t.spread(o)
//     m := (t[o].l + t[o].r) >> 1
//     if l <= m {
//         t.update(o<<1, l, r, f)
//     }
//     if m < r {
//         t.update(o<<1|1, l, r, f)
//     }
//     t.maintain(o)
// }

// func (t lazySeg) query(o, l, r int) int {
//     if l <= t[o].l && t[o].r <= r {
//         return t[o].mx
//     }
//     t.spread(o)
//     m := (t[o].l + t[o].r) >> 1
//     if r <= m {
//         return t.query(o<<1, l, r)
//     }
//     if l > m {
//         return t.query(o<<1|1, l, r)
//     }
//     return mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
// }

// func newLazySegmentTreeWithArray(a []int) lazySeg {
//     n := len(a)
//     t := make(lazySeg, 2<<bits.Len(uint(n-1)))
//     t.build(a, 1, 0, n-1)
//     return t
// }

// func newLazySegmentTree(n int, initVal int) lazySeg {
//     a := make([]int, n)
//     for i := range a {
//         a[i] = initVal
//     }
//     return newLazySegmentTreeWithArray(a)
// }

// func maximumCount(nums []int, queries [][]int) []int {
//     res, n := []int{}, len(nums)
//     pos := map[int]*redblacktree.Tree[int, struct{}]{}
//     for i, x := range nums {
//         if np[x] { continue }
//         if _, ok := pos[x]; !ok {
//             pos[x] = redblacktree.New[int, struct{}]()
//         }
//         pos[x].Put(i, struct{}{})
//     }
//     t := newLazySegmentTree(n, 0)
//     for _, ps := range pos {
//         if ps.Size() > 1 {
//             t.update(1, ps.Left().Key, ps.Right().Key, 1)
//         }
//     }
//     update := func(ps *redblacktree.Tree[int, struct{}], i, delta int) {
//         l, r := ps.Left().Key, ps.Right().Key
//         if l == r {
//             t.update(1, min(l, i), max(r, i), delta)
//         } else if i < l {
//             t.update(1, i, l-1, delta)
//         } else if i > r {
//             t.update(1, r+1, i, delta)
//         }
//     }
//     for _, q := range queries {
//         i, v := q[0], q[1]
//         old := nums[i]
//         nums[i] = v
//         // 处理旧值
//         if !np[old] {
//             ps := pos[old]
//             ps.Remove(i)
//             if ps.Empty() {
//                 delete(pos, old)
//             } else {
//                 update(ps, i, -1)
//             }
//         }
//         // 处理新值
//         if !np[v] {
//             if ps, ok := pos[v]; !ok {
//                 pos[v] = redblacktree.New[int, struct{}]()
//             } else {
//                 update(ps, i, 1)
//             }
//             pos[v].Put(i, struct{}{})
//         }
//         // 整个数组的不同质数个数 + 切一刀的最大额外收益
//         res = append(res, len(pos) + t.query(1, 0, n - 1))
//     }
//     return res
// }

func maximumCount(nums []int, queries [][]int) []int {
    n := len(nums)
    const MAXV = 101010
    prime := make([]bool, MAXV)
    for i := 2; i < MAXV; i++ {
        prime[i] = true
    }
    for i := 2; i*i < MAXV; i++ {
        if prime[i] {
            for j := i * i; j < MAXV; j += i {
                prime[j] = false
            }
        }
    }
    type Seg struct {
        mi, ma, lazy, best     int
        left, right            *Seg
    }
    var newSeg func(l, r int) *Seg
    newSeg = func(l, r int) *Seg {
        s := &Seg{mi: l, ma: r}
        if l != r {
            m := l + (r-l)/2
            s.left = newSeg(l, m)
            s.right = newSeg(m+1, r)
        }
        return s
    }
    seg := newSeg(0, n-1)
    var pushDown func(s *Seg)
    pushDown = func(s *Seg) {
        if s.lazy != 0 {
            s.best += s.lazy
            if s.left != nil {
                s.left.lazy += s.lazy
                s.right.lazy += s.lazy
            }
            s.lazy = 0
        }
    }
    var segUpdate func(s *Seg, l, r, x int)
    segUpdate = func(s *Seg, l, r, x int) {
        pushDown(s)
        if l <= s.mi && s.ma <= r {
            s.lazy += x
            pushDown(s)
            return
        }
        if l > s.ma || r < s.mi {
            return
        }
        segUpdate(s.left, l, r, x)
        segUpdate(s.right, l, r, x)
        if s.left.best > s.right.best {
            s.best = s.left.best
        } else {
            s.best = s.right.best
        }
    }
    container := make(map[int]map[int]struct{})
    ask := func(v int) (l, r int) {
        m, ok := container[v]
        if !ok || len(m) == 0 {
            return -1, -1
        }
        l, r = n, -1
        for idx := range m {
            if idx < l {
                l = idx
            }
            if idx > r {
                r = idx
            }
        }
        return
    }
    push := func(idx int) {
        v := nums[idx]
        if !prime[v] {
            return
        }
        if container[v] == nil {
            container[v] = make(map[int]struct{})
        }
        container[v][idx] = struct{}{}
    }
    pop := func(idx int) {
        v := nums[idx]
        if !prime[v] {
            return
        }
        m := container[v]
        delete(m, idx)
        if len(m) == 0 {
            delete(container, v)
        }
    }
    update := func(l, r, op int) {
        if l != -1 {
            segUpdate(seg, l, n-1, op)
        }
        if r != -1 {
            segUpdate(seg, 0, r-1, op)
        }
    }
    for i := 0; i < n; i++ {
        push(i)
    }
    for v := range container {
        l, r := ask(v)
        update(l, r, 1)
    }
    res := make([]int, len(queries))
    for i, q := range queries {
        idx, val := q[0], q[1]
        if nums[idx] != val {
            if prime[nums[idx]] {
                l1, r1 := ask(nums[idx])
                update(l1, r1, -1)
                pop(idx)
                l2, r2 := ask(nums[idx])
                update(l2, r2, 1)
            }
            nums[idx] = val
            if prime[val] {
                l1, r1 := ask(val)
                update(l1, r1, -1)
                push(idx)
                l2, r2 := ask(val)
                update(l2, r2, 1)
            }
        }
        res[i] = seg.best
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,1,3,1,2], queries = [[1,2],[3,3]]
    // Output: [3,4]
    // Explanation:
    // Initially nums = [2, 1, 3, 1, 2].
    // After 1st query, nums = [2, 2, 3, 1, 2]. Split nums into [2] and [2, 3, 1, 2]. [2] consists of 1 distinct prime and [2, 3, 1, 2] consists of 2 distinct primes. Hence, the answer for this query is 1 + 2 = 3.
    // After 2nd query, nums = [2, 2, 3, 3, 2]. Split nums into [2, 2, 3] and [3, 2] with an answer of 2 + 2 = 4.
    // The output is [3, 4].
    fmt.Println(maximumCount([]int{2,1,3,1,2}, [][]int{{1,2},{3,3}})) // [3,4]
    // Example 2:
    // Input: nums = [2,1,4], queries = [[0,1]]
    // Output: [0]
    // Explanation:
    // Initially nums = [2, 1, 4].
    // After 1st query, nums = [1, 1, 4]. There are no prime numbers in nums, hence the answer for this query is 0.
    // The output is [0].
    fmt.Println(maximumCount([]int{2,1,4}, [][]int{{0,1}})) // [0]
}