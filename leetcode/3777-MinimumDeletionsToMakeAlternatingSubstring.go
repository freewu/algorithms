package main

// 3777. Minimum Deletions to Make Alternating Substring
// You are given a string s of length n consisting only of the characters 'A' and 'B'.

// You are also given a 2D integer array queries of length q, where each queries[i] is one of the following:
//     1. [1, j]: Flip the character at index j of s i.e. 
//               'A' changes to 'B' (and vice versa). 
//               This operation mutates s and affects subsequent queries.
//     2. [2, l, r]: Compute the minimum number of character deletions required to make the substring s[l..r] alternating. 
//                   This operation does not modify s; the length of s remains n.

// A substring is alternating if no two adjacent characters are equal. 
// A substring of length 1 is always alternating.

// Return an integer array answer, where answer[i] is the result of the ith query of type [2, l, r].

// Example 1:
// Input: s = "ABA", queries = [[2,1,2],[1,1],[2,0,2]]
// Output: [0,2]
// Explanation:
// i | queries[i]  | j | l | r | s before query | s[l..r]  | Result                            | Answer
// 0 | [2, 1, 2]   | - | 1 | 2 | "ABA"          | "BA"     | Already alternating               | 0
// 1 | [1, 1]      | 1 | - | - | "ABA"          | -        | Flip s[1] from 'B' to 'A'         | -
// 2 | [2, 0, 2]   | - | 0 | 2 | "AAA"          | "AAA"    | Delete any two 'A's to get "A"    | 2
// Thus, the answer is [0, 2].

// Example 2:
// Input: s = "ABB", queries = [[2,0,2],[1,2],[2,0,2]]
// Output: [1,0]
// Explanation:
// i | queries[i]  | j | l | r | s before query    | s[l..r]   | Result                        | Answer
// 0 | [2, 0, 2]   | - | 0 | 2 | "ABB"             | "ABB"     | Delete one 'B' to get "AB"    | 1
// 1 | [1, 2]      | 2 | - | - | "ABB"             | -         | Flip s[2] from 'B' to 'A'     | -
// 2 | [2, 0, 2]   | - | 0 | 2 | "ABA"             | "ABA"     | Already alternating           | 0
// Thus, the answer is [1, 0].

// Example 3:
// Input: s = "BABA", queries = [[2,0,3],[1,1],[2,1,3]]
// Output: [0,1]
// Explanation:
// i | queries[i]  | j | l | r | s before query    | s[l..r]   | Result                        | Answer
// 0 | [2, 0, 3]   | - | 0 | 3 | "BABA"            | "BABA"    | Already alternating           | 0
// 1 | [1, 1]      | 1 | - | - | "BABA"            | -         | Flip s[1] from 'A' to 'B'     | -
// 2 | [2, 1, 3]   | - | 1 | 3 | "BBBA"            | "BBA"     | Delete one 'B' to get "BA"    | 1 
// Thus, the answer is [0, 1].

// Constraints:
//     1 <= n == s.length <= 10^5
//     s[i] is either 'A' or 'B'.
//     1 <= q == queries.length <= 10^5
//     queries[i].length == 2 or 3
//     queries[i] == [1, j] or,
//     queries[i] == [2, l, r]
//     0 <= j <= n - 1
//     0 <= l <= r <= n - 1

import "fmt"

// 参考: https://judge.yosupo.jp/submission/334041 树状数组
type FenwickInfo[T any] interface {
    add(T, T) T // 加法操作，用于合并两个值
    sub(T, T) T // 减法操作，用于计算差值
    e() T       // 单位元（如加法的0，乘法的1）
}

type Fenwick[T any, M FenwickInfo[T]] struct {
    t   []T          // 树状数组内部存储
    n   int          // 元素个数
    add func(T, T) T // 加法函数
    sub func(T, T) T // 减法函数
    e   func() T     // 单位元函数
}

func NewFenwick[T any, M FenwickInfo[T]](n int, m M) *Fenwick[T, M] {
    t := make([]T, n+1) // 下标从1开始，所以需要n+1
    for i := range t {
        t[i] = m.e() // 初始化所有元素为单位元
    }
    return &Fenwick[T, M]{t, n, m.add, m.sub, m.e}
}
// Add 单点更新，在位置i增加x
func (fen *Fenwick[T, M]) Add(i int, x T) {
    // i++ 转换为1-based索引
    // i += i & -i 找到下一个需要更新的节点
    for i++; i <= fen.n; i += i & -i {
        fen.t[i] = fen.add(fen.t[i], x)
    }
}
// Pre 查询前缀和 [0, i)
func (fen *Fenwick[T, M]) Pre(i int) T {
    add := fen.add
    r := fen.e()
    for ; i > 0; i &= i - 1 {
        r = add(r, fen.t[i])
    }
    return r
}
// Sum 查询区间和 [l, r)
func (fen *Fenwick[T, M]) Sum(l, r int) T {
    return fen.sub(fen.Pre(r), fen.Pre(l))
}
// Get 单点查询（需要额外实现减法操作）
func (fen *Fenwick[T, M]) Get(i int) T {
    return fen.Sum(i, i+1)
}
type RangeAddFenwick[T any, M FenwickInfo[T]] struct {
    ft *Fenwick[T, M]
}
func NewRangeAddFenwick[T any, M FenwickInfo[T]](n int, m M) *RangeAddFenwick[T, M] {
    return &RangeAddFenwick[T, M]{ft: NewFenwick(n, m)}
}
func (raf *RangeAddFenwick[T, M]) RangeAdd(l, r int, x T) {
    raf.ft.Add(l, x) // 在l处加x
    if r < raf.ft.n {
        raf.ft.Add(r, raf.ft.sub(raf.ft.e(), x)) // 在r处减x，实现差分
    }
}
func (raf *RangeAddFenwick[T, M]) PointQuery(i int) T {
    return raf.ft.Pre(i + 1)
}

type Int struct{}
func (Int) add(a, b int) int {
    return a + b // 根据需要修改
}
func (Int) sub(a, b int) int {
    return a - b //根据需求修改
}
func (Int) e() int {
    return 0 //单位元
}

func minDeletions(s string, queries [][]int) []int {
    n := len(s)
    arr := make([]int, n)
    for i, x := range s {
        if x == 'A' {
            arr[i] = 1
        } else {
            arr[i] = 0
        }
    }
    arr = append(arr, 0)
    f := NewFenwick(n + 1,Int{})
    pre := -1
    for i,v := range arr {
        if v == pre {
            f.Add(i + 1,1)
        }else{
            pre = v
        }
    }
    res := []int{}
    for _, q := range queries {
        op := q[0]
        if op == 1 {
            x := q[1] + 1
            if x == 1 {
                if arr[x-1] == arr[x] {
                    f.Add(x+1,-1)
                } else {
                    f.Add(x+1,1)
                }
            } else if x == n {
                if arr[x-1] == arr[x-2] {
                    f.Add(x,-1)
                } else {
                    f.Add(x,1)
                }
            } else {
                if arr[x-2] != arr[x] {
                    if arr[x-2] == arr[x-1] {
                        f.Add(x,-1)
                        f.Add(x+1,1)
                    } else {
                        f.Add(x,1)
                        f.Add(x+1,-1)
                    }
                } else {
                    if arr[x-2] == arr[x-1] {   
                        f.Add(x,-1)
                        f.Add(x+1,-1)
                    } else {
                        f.Add(x,1)
                        f.Add(x+1,1)
                    }
                }

            }
            arr[x - 1] ^= 1
        } else {
            l, r := q[1]+1, q[2]+1
            mi := f.Sum(l, r+1)
            if f.Sum(l, l+1) == 1 {     
                mi--
            }
            res = append(res, mi)
        }
    }
    return res
}


type Segment struct {
    mi, ma int
    head   int
    tail   int
    cnt    int
    left   *Segment
    right  *Segment
}

func NewSegment(s string, l, r int) *Segment {
    t := &Segment{
        mi:   l,
        ma:   r,
        head: 0,
        tail: 0,
        cnt:  0,
    }
    if s[l] == 'A' {
        t.head = 1
    }
    if s[r] == 'A' {
        t.tail = 1
    }
    if l != r {
        m := l + (r-l)/2
        t.left = NewSegment(s, l, m)
        t.right = NewSegment(s, m+1, r)
        t.cnt = t.left.cnt + t.right.cnt
        if t.left.tail == t.right.head {
            t.cnt++
        }
    }
    return t
}

type SegmentInfo struct {
    head int
    tail int
    cnt  int
}

func (t *Segment) update(pos, x int) {
    if t.mi <= pos && pos <= t.ma {
        if t.mi == pos && pos == t.ma {
            t.head = x
            t.tail = x
            return
        }
        t.left.update(pos, x)
        t.right.update(pos, x)
        t.head = t.left.head
        t.tail = t.right.tail
        t.cnt = t.left.cnt + t.right.cnt
        if t.left.tail == t.right.head {
            t.cnt++
        }
    }
}

func (t *Segment) query(l, r int) SegmentInfo {
    if l <= t.mi && t.ma <= r {
        return SegmentInfo{t.head, t.tail, t.cnt}
    }
    if l > t.ma || r < t.mi {
        return SegmentInfo{-1, -1, 0}
    }
    le := t.left.query(l, r)
    ri := t.right.query(l, r)
    if le.head == -1 {
        return ri
    }
    if ri.head == -1 {
        return le
    }
    now := le.cnt + ri.cnt
    if le.tail == ri.head {
        now++
    }
    return SegmentInfo{le.head, ri.tail, now}
}

func minDeletions1(s string, queries [][]int) []int {
    n := len(s)
    seg := NewSegment(s, 0, n-1)
    res := make([]int, 0, len(queries))
    sb := []byte(s)
    for _, q := range queries {
        op := q[0]
        if op == 1 {
            idx := q[1]
            if sb[idx] == 'A' {
                sb[idx] = 'B'
                seg.update(idx, 0)
            } else {
                sb[idx] = 'A'
                seg.update(idx, 1)
            }
        } else {
            l, r := q[1], q[2]
            val := seg.query(l, r)
            res = append(res, val.cnt)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "ABA", queries = [[2,1,2],[1,1],[2,0,2]]
    // Output: [0,2]
    // Explanation:
    // i | queries[i]  | j | l | r | s before query | s[l..r]  | Result                            | Answer
    // 0 | [2, 1, 2]   | - | 1 | 2 | "ABA"          | "BA"     | Already alternating               | 0
    // 1 | [1, 1]      | 1 | - | - | "ABA"          | -        | Flip s[1] from 'B' to 'A'         | -
    // 2 | [2, 0, 2]   | - | 0 | 2 | "AAA"          | "AAA"    | Delete any two 'A's to get "A"    | 2
    // Thus, the answer is [0, 2].
    fmt.Println(minDeletions("ABA", [][]int{{2,1,2},{1,1},{2,0,2}})) // [0, 2]
    // Example 2:
    // Input: s = "ABB", queries = [[2,0,2],[1,2],[2,0,2]]
    // Output: [1,0]
    // Explanation:
    // i | queries[i]  | j | l | r | s before query    | s[l..r]   | Result                        | Answer
    // 0 | [2, 0, 2]   | - | 0 | 2 | "ABB"             | "ABB"     | Delete one 'B' to get "AB"    | 1
    // 1 | [1, 2]      | 2 | - | - | "ABB"             | -         | Flip s[2] from 'B' to 'A'     | -
    // 2 | [2, 0, 2]   | - | 0 | 2 | "ABA"             | "ABA"     | Already alternating           | 0
    // Thus, the answer is [1, 0].
    fmt.Println(minDeletions("ABB", [][]int{{2,0,2},{1,2},{2,0,2}})) // [1, 0]
    // Example 3:
    // Input: s = "BABA", queries = [[2,0,3],[1,1],[2,1,3]]
    // Output: [0,1]
    // Explanation:
    // i | queries[i]  | j | l | r | s before query    | s[l..r]   | Result                        | Answer
    // 0 | [2, 0, 3]   | - | 0 | 3 | "BABA"            | "BABA"    | Already alternating           | 0
    // 1 | [1, 1]      | 1 | - | - | "BABA"            | -         | Flip s[1] from 'A' to 'B'     | -
    // 2 | [2, 1, 3]   | - | 1 | 3 | "BBBA"            | "BBA"     | Delete one 'B' to get "BA"    | 1 
    // Thus, the answer is [0, 1].
    fmt.Println(minDeletions("BABA", [][]int{{2,0,3},{1,1},{2,1,3}})) // [0, 1]

    fmt.Println(minDeletions("AAAA", [][]int{{2,0,3},{1,1},{2,1,3}})) // [3 1]
    fmt.Println(minDeletions("BBBB", [][]int{{2,0,3},{1,1},{2,1,3}})) // [3 1]
    fmt.Println(minDeletions("ABAB", [][]int{{2,0,3},{1,1},{2,1,3}})) // [0, 1]
    fmt.Println(minDeletions("AABB", [][]int{{2,0,3},{1,1},{2,1,3}})) // [2 2]
    fmt.Println(minDeletions("BABA", [][]int{{2,0,3},{1,1},{2,1,3}})) // [0, 1]

    fmt.Println(minDeletions1("ABA", [][]int{{2,1,2},{1,1},{2,0,2}})) // [0, 2]
    fmt.Println(minDeletions1("ABB", [][]int{{2,0,2},{1,2},{2,0,2}})) // [1, 0]
    fmt.Println(minDeletions1("BABA", [][]int{{2,0,3},{1,1},{2,1,3}})) // [0, 1]
    fmt.Println(minDeletions1("AAAA", [][]int{{2,0,3},{1,1},{2,1,3}})) // [3 1]
    fmt.Println(minDeletions1("BBBB", [][]int{{2,0,3},{1,1},{2,1,3}})) // [3 1]
    fmt.Println(minDeletions1("ABAB", [][]int{{2,0,3},{1,1},{2,1,3}})) // [0, 1]
    fmt.Println(minDeletions1("AABB", [][]int{{2,0,3},{1,1},{2,1,3}})) // [2 2]
    fmt.Println(minDeletions1("BABA", [][]int{{2,0,3},{1,1},{2,1,3}})) // [0, 1]
}