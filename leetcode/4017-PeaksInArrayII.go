package main

// 4017. Peaks in Array II
// You are given an integer array nums of length n and a 2D integer array queries.

// A subarray nums[i..j] is called a peak subarray if:
//     1. Its length is at least 3.
//     2. There exists an index k such that i < k < j and:
//         2.1 nums[k] > nums[k - 1]
//         2.2 nums[k] > nums[k + 1]

// You have to process queries of two types:
//     1. [1, li, ri]: Calculate the number of peak subarrays fully contained within nums[li..ri].
//     2. [2, indexi, vali]: Update nums[indexi] to vali. 
//        This update applies to all subsequent queries.

// Return an array answer, where answer[i] is the answer to the ith query of type 1 in the order they appear.

// Example 1:
// Input: nums = [1,3,2,4], queries = [[1,0,3],[2,1,1],[1,0,3]]
// Output: [2,0]
// Explanation:​​​​​​​
// Query [1, 0, 3]:
// [1, 3, 2]: choose k = 1. Then nums[k] = 3, nums[k - 1] = 1, and nums[k + 1] = 2. Since 3 > 1 and 3 > 2, this is a peak subarray.
// [1, 3, 2, 4]: choose k = 1. Then nums[k] = 3, nums[k - 1] = 1, and nums[k + 1] = 2. Since 3 > 1 and 3 > 2, this is a peak subarray.
// Query [2, 1, 1]: Update nums[1] to 1. The array becomes [1, 1, 2, 4].
// Query [1, 0, 3]: There are no peak subarrays now.
// Thus, answer = [2, 0].

// Example 2:
// Input: nums = [9,8,9,8], queries = [[1,1,3],[2,2,1],[1,0,2]]
// Output: [1,0]
// Explanation:
// Query [1, 1, 3]:
// nums[1..3] = [8, 9, 8]: choose k = 2. Then nums[k] = 9, nums[k - 1] = 8, and nums[k + 1] = 8. Since 9 > 8 and 9 > 8, this is a peak subarray.
// Query [2, 2, 1]: Update nums[2] to 1. The array becomes [9, 8, 1, 8].
// Query [1, 0, 2]: There are no peak subarrays.
// Thus, answer = [1, 0].

// Example 3:
// Input: nums = [3,6,2,7,1], queries = [[1,1,3],[2,3,0],[1,0,4]]
// Output: [0,3]
// Explanation:
// Query [1, 1, 3]: The only subarray of length at least 3 is [6, 2, 7]. Its only possible peak index is k = 2, but nums[2] = 2 is less than both nums[1] = 6 and nums[3] = 7, so it is not a peak subarray.
// Query [2, 3, 0]: Update nums[3] to 0. The array becomes [3, 6, 2, 0, 1].
// Query [1, 0, 4]:
// [3, 6, 2]: choose k = 1. Then nums[k] = 6, nums[k - 1] = 3, and nums[k + 1] = 2. Since 6 > 3 and 6 > 2, this is a peak subarray.
// [3, 6, 2, 0]: choose k = 1. Then nums[k] = 6, nums[k - 1] = 3, and nums[k + 1] = 2. Since 6 > 3 and 6 > 2, this is a peak subarray.
// [3, 6, 2, 0, 1]: choose k = 1. Then nums[k] = 6, nums[k - 1] = 3, and nums[k + 1] = 2. Since 6 > 3 and 6 > 2, this is a peak subarray.
// Thus, answer = [0, 3].
 
// Constraints:
//     3 <= n == nums.length <= 10^5
//     0 <= nums[i] <= 10^5
//     1 <= queries.length <= 10^5
//     queries[i] = [1, li, ri] or queries[i] = [2, indexi, vali]
//     0 <= li < ri <= n - 1
//     0 <= indexi <= n - 1
//     0 <= vali <= 10^5

import "fmt"
import "math/bits"

type Data struct {
    count, prefix, suffix, length int
    hasPeak            bool
}

type SegmentTree []Data

func NewSegmentTree(a []int) SegmentTree {
    n := len(a)
    t := make(SegmentTree, 2 << bits.Len(uint(n-1)))
    t.build(a, 1, 0, n-1)
    return t
}

func (seg SegmentTree) mergeData(a, b Data) Data {
    count := a.count + b.count + a.length*b.length - a.suffix * b.prefix
    prefix := a.prefix
    if !a.hasPeak {
        prefix += b.prefix
    }
    suffix := b.suffix
    if !b.hasPeak {
        suffix += a.suffix
    }
    return Data{count, prefix, suffix, a.length + b.length, a.hasPeak || b.hasPeak }
}

func (t SegmentTree) maintain(node int) {
    t[node] = t.mergeData(t[node * 2], t[node * 2 + 1])
}

func (t SegmentTree) build(a []int, node, l, r int) {
    if l == r { // 叶子
        hasPeak := 0 < l && l < len(a)-1 && a[l-1] < a[l] && a[l] > a[l+1]
        t[node] = Data{0, 1, 1, 1, hasPeak} // 初始化叶节点的值
        return
    }
    m := (l + r) >> 1
    t.build(a, node*2, l, m)     // 初始化左子树
    t.build(a, node*2+1, m+1, r) // 初始化右子树
    t.maintain(node)
}

func (t SegmentTree) update(node, l, r, i int) {
    if l == r { // 叶子（到达目标）
        t[node].hasPeak = !t[node].hasPeak
        return
    }
    m := (l + r) >> 1
    if i <= m { // i 在左子树
        t.update(node * 2, l, m, i)
    } else { // i 在右子树
        t.update(node * 2 + 1, m + 1, r, i)
    }
    t.maintain(node)
}

func (t SegmentTree) query(node, l, r, ql, qr int) Data {
    if ql <= l && r <= qr { // 当前子树完全在 [ql, qr] 内
        return t[node]
    }
    m := (l + r) >> 1
    if qr <= m { // [ql, qr] 与右子树无交集，仅需递归左子树
        return t.query(node * 2, l, m, ql, qr)
    }
    if ql > m { // [ql, qr] 与左子树无交集，仅需递归右子树
        return t.query(node*2+1, m+1, r, ql, qr)
    }
    // [ql, qr] 与左右子树均有交集，分别递归，然后合并结果
    return t.mergeData(t.query(node*2, l, m, ql, qr), t.query(node*2+1, m+1, r, ql, qr))
}

func countOfPeaks(nums []int, queries [][]int) []int64 {
    res, n := make([]int64, 0), len(nums)
    t := NewSegmentTree(nums)
    for _, q := range queries {
        if q[0] == 1 {
            res = append(res, int64(t.query(1, 0, n-1, q[1], q[2]).count))      
            continue
        }
        i, v := q[1], q[2]
        if i > 1 {
            oldHas := nums[i-2] < nums[i-1] && nums[i-1] > nums[i]
            newHas := nums[i-2] < nums[i-1] && nums[i-1] > v
            if newHas != oldHas {
                t.update(1, 0, n-1, i-1)
            }
        }
        if 0 < i && i < n-1 {
            oldHas := nums[i-1] < nums[i] && nums[i] > nums[i+1]
            newHas := nums[i-1] < v && v > nums[i+1]
            if newHas != oldHas {
                t.update(1, 0, n-1, i)
            }
        }
        if i < n-2 {
            oldHas := nums[i] < nums[i+1] && nums[i+1] > nums[i+2]
            newHas := v < nums[i+1] && nums[i+1] > nums[i+2]
            if newHas != oldHas {
                t.update(1, 0, n-1, i+1)
            }
        }
        nums[i] = v
    }
    return res
}

func countOfPeaks1(nums[]int,queries[][]int) []int64 {
    res, n := make([]int64, 0), len(nums)
    cnt, sa, sb := make([]int64,n + 1), make([]int64,n + 1), make([]int64,n + 1)
    update := func(t[]int64,i int,v int64) {
        for j := i + 1; j <= n; j += j&-j {
            t[j] += v
        }
    }
    query := func(t[]int64,i int)int64{
        var s int64
        for j := i + 1; j > 0; j -= j&-j {
            s += t[j]
        }
        return s
    }
    kth := func(k int64) int {
        if k <= 0 || k > query(cnt,n-1) {
            return -1
        }
        pos := 0
        for pw := 1 << 17; pw > 0; pw >>= 1 {
            if pos + pw <= n && cnt[pos + pw] < k {
                pos += pw
                k -= cnt[pos]
            }
        }
        return pos
    }
    succ := func(x int) int {
        if x < 0 {
            x = 0
        }
        return kth(query(cnt,x - 1) + 1)
    }
    pred := func(x int) int {
        return kth(query(cnt,x - 1))    
    }
    pp, isPk := make([]int,n), make([]bool,n)
    peak := func(k int) bool {
        return k >= 1 && k <= n-2 && nums[k] > nums[k-1] && nums[k] > nums[k + 1]
    }
    addPeak := func(k int) {
        p, nx := pred(k), succ(k + 1)
        pp[k] = p
        a := int64(k - p)
        update(cnt, k, 1)
        update(sa, k, a)
        update(sb, k, int64(k) * a)
        if nx != -1 {
            d := int64(pp[nx] - k)
            update(sa, nx, d)
            update(sb, nx, int64(nx) * d)
            pp[nx] = k
        }
        isPk[k] = true
    }
    removePeak := func(k int) {
        p, nx := pp[k], succ(k + 1)
        a := int64(k - p)
        update(cnt, k, -1)
        update(sa, k, -a)
        update(sb, k, -int64(k) * a)
        if nx != -1 {
            d := int64(k - p)
            update(sa, nx, d)
            update(sb, nx, int64(nx) * d)
            pp[nx] = p
        }
        isPk[k] = false
    }
    for k := 1; k <= n - 2; k++ {
        if peak(k) {
            addPeak(k)
        }
    }
    for _, q := range queries {
        if q[0] == 2 {
            nums[q[1]] = q[2]
            for k := q[1] - 1; k <= q[1] + 1; k++ {
                if k < 1 || k > n - 2 {
                    continue
                }
                now := peak(k)
                if now && !isPk[k] {
                    addPeak(k)
                } else if !now && isPk[k] {
                    removePeak(k)
                }
            }
        } else {
            l, r := q[1], q[2]
            val, lo, hi := int64(0),l + 1, r - 1
            if lo <= hi {
                Sa := query(sa, hi) - query(sa,lo - 1)
                Sb := query(sb, hi) - query(sb,lo - 1)
                val = int64(r) * Sa - Sb
                if k1 := succ(lo); k1 != -1 && k1 <= hi {
                    val += int64(pp[k1] - l) * int64(r - k1)
                }
            }
            res = append(res, val)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2,4], queries = [[1,0,3],[2,1,1],[1,0,3]]
    // Output: [2,0]
    // Explanation:​​​​​​​
    // Query [1, 0, 3]:
    // [1, 3, 2]: choose k = 1. Then nums[k] = 3, nums[k - 1] = 1, and nums[k + 1] = 2. Since 3 > 1 and 3 > 2, this is a peak subarray.
    // [1, 3, 2, 4]: choose k = 1. Then nums[k] = 3, nums[k - 1] = 1, and nums[k + 1] = 2. Since 3 > 1 and 3 > 2, this is a peak subarray.
    // Query [2, 1, 1]: Update nums[1] to 1. The array becomes [1, 1, 2, 4].
    // Query [1, 0, 3]: There are no peak subarrays now.
    // Thus, answer = [2, 0].
    fmt.Println(countOfPeaks([]int{1,3,2,4}, [][]int{{1,0,3},{2,1,1},{1,0,3}})) // [2,0]
    // Example 2:
    // Input: nums = [9,8,9,8], queries = [[1,1,3],[2,2,1],[1,0,2]]
    // Output: [1,0]
    // Explanation:
    // Query [1, 1, 3]:
    // nums[1..3] = [8, 9, 8]: choose k = 2. Then nums[k] = 9, nums[k - 1] = 8, and nums[k + 1] = 8. Since 9 > 8 and 9 > 8, this is a peak subarray.
    // Query [2, 2, 1]: Update nums[2] to 1. The array becomes [9, 8, 1, 8].
    // Query [1, 0, 2]: There are no peak subarrays.
    // Thus, answer = [1, 0].
    fmt.Println(countOfPeaks([]int{9,8,9,8}, [][]int{{1,1,3},{2,2,1},{1,0,2}})) // [1,0]
    // Example 3:
    // Input: nums = [3,6,2,7,1], queries = [[1,1,3],[2,3,0],[1,0,4]]
    // Output: [0,3]
    // Explanation:
    // Query [1, 1, 3]: The only subarray of length at least 3 is [6, 2, 7]. Its only possible peak index is k = 2, but nums[2] = 2 is less than both nums[1] = 6 and nums[3] = 7, so it is not a peak subarray.
    // Query [2, 3, 0]: Update nums[3] to 0. The array becomes [3, 6, 2, 0, 1].
    // Query [1, 0, 4]:
    // [3, 6, 2]: choose k = 1. Then nums[k] = 6, nums[k - 1] = 3, and nums[k + 1] = 2. Since 6 > 3 and 6 > 2, this is a peak subarray.
    // [3, 6, 2, 0]: choose k = 1. Then nums[k] = 6, nums[k - 1] = 3, and nums[k + 1] = 2. Since 6 > 3 and 6 > 2, this is a peak subarray.
    // [3, 6, 2, 0, 1]: choose k = 1. Then nums[k] = 6, nums[k - 1] = 3, and nums[k + 1] = 2. Since 6 > 3 and 6 > 2, this is a peak subarray.
    // Thus, answer = [0, 3].
    fmt.Println(countOfPeaks([]int{3,6,2,7,1}, [][]int{{1,1,3},{2,3,0},{1,0,4}})) // [0,3]

    fmt.Println(countOfPeaks([]int{1,2,3,4,5,6,7,8,9}, [][]int{{1,1,3},{2,3,0},{1,0,4}})) // [0,4]
    fmt.Println(countOfPeaks([]int{9,8,7,6,5,4,3,2,1}, [][]int{{1,1,3},{2,3,0},{1,0,4}})) // [0,0]

    fmt.Println(countOfPeaks1([]int{1,3,2,4}, [][]int{{1,0,3},{2,1,1},{1,0,3}})) // [2,0]
    fmt.Println(countOfPeaks1([]int{9,8,9,8}, [][]int{{1,1,3},{2,2,1},{1,0,2}})) // [1,0]
    fmt.Println(countOfPeaks1([]int{3,6,2,7,1}, [][]int{{1,1,3},{2,3,0},{1,0,4}})) // [0,3]
    fmt.Println(countOfPeaks1([]int{1,2,3,4,5,6,7,8,9}, [][]int{{1,1,3},{2,3,0},{1,0,4}})) // [0,4]
    fmt.Println(countOfPeaks1([]int{9,8,7,6,5,4,3,2,1}, [][]int{{1,1,3},{2,3,0},{1,0,4}})) // [0,0]
}