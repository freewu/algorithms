package main

// 3266. Final Array State After K Multiplication Operations II
// You are given an integer array nums, an integer k, and an integer multiplier.

// You need to perform k operations on nums. In each operation:
//     1. Find the minimum value x in nums. If there are multiple occurrences of the minimum value, 
//        select the one that appears first.
//     2. Replace the selected minimum value x with x * multiplier.

// After the k operations, apply modulo 10^9 + 7 to every value in nums.

// Return an integer array denoting the final state of nums after performing all k operations and then applying the modulo.

// Example 1:
// Input: nums = [2,1,3,5,6], k = 5, multiplier = 2
// Output: [8,4,6,5,6]
// Explanation:
// Operation	Result
// After operation 1	[2, 2, 3, 5, 6]
// After operation 2	[4, 2, 3, 5, 6]
// After operation 3	[4, 4, 3, 5, 6]
// After operation 4	[4, 4, 6, 5, 6]
// After operation 5	[8, 4, 6, 5, 6]
// After applying modulo	[8, 4, 6, 5, 6]

// Example 2:
// Input: nums = [100000,2000], k = 2, multiplier = 1000000
// Output: [999999307,999999993]
// Explanation:
// Operation	Result
// After operation 1	[100000, 2000000000]
// After operation 2	[100000000000, 2000000000]
// After applying modulo	[999999307, 999999993]

// Constraints:
//     1 <= nums.length <= 10^4
//     1 <= nums[i] <= 10^9
//     1 <= k <= 10^9
//     1 <= multiplier <= 10^6

import "fmt"
import "math/bits"
import "slices"
import "sort"
import "container/heap"

func getFinalState(nums []int, k int, multiplier int) []int {
    if multiplier == 1 {  return nums }
    n, mx, mod := len(nums), 0, 1_000_000_007
    h := make(MinHeap, n)
    pow := func(x, n int) int {
        res := 1
        for ; n > 0; n /= 2 {
            if n%2 > 0 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    less := func(a, b Pair) bool { return a.x < b.x || a.x == b.x && a.i < b.i }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range nums {
        mx = max(mx, v)
        h[i] = Pair{v, i}
    }
    clone := slices.Clone(h)
    // 打表，计算出最小的 e 满足 multiplier^e >= 2^i
    mxLen := bits.Len(uint(mx))
    type ep struct{ e, powM int }
    ePowM := make([]ep, 0, mxLen)
    for pow2, powM, e := 1, 1, 0; pow2 <= mx; pow2 <<= 1 {
        if powM < pow2 { // 由于 multiplier >= 2，这里只需写 if 而不是 for
            powM *= multiplier
            e++
        }
        ePowM = append(ePowM, ep{e, powM})
    }
    // 把每个数都操作到 >= mx
    left := k
    for i := range h {
        x := h[i].x
        p := ePowM[mxLen-bits.Len(uint(x))]
        e, powM := p.e, p.powM
        if powM/multiplier*x >= mx { // 多操作了一次
            powM /= multiplier
            e--
        } else if x*powM < mx { // 少操作了一次
            powM *= multiplier
            e++
        }
        left -= e
        if left < 0 {
            break
        }
        h[i].x *= powM
    }
    if left < 0 {
        // 暴力模拟
        h = clone
        heap.Init(&h)
        for ; k > 0; k-- {
            h[0].x *= multiplier
            heap.Fix(&h, 0)
        }
        sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
        for _, p := range h {
            nums[p.i] = p.x % mod
        }
        return nums
    }
    // 剩余的操作可以直接用公式计算
    k = left
    pow1 := pow(multiplier, k/n)
    pow2 := pow1 * multiplier % mod
    sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
    for i, p := range h {
        pw := pow1
        if i < k%n {
            pw = pow2
        }
        nums[p.i] = p.x % mod * pw % mod
    }
    return nums
}

type Pair struct{ x, i int }
type MinHeap []Pair
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].x < h[j].x || h[i].x == h[j].x && h[i].i < h[j].i  }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v interface{}) { *h = append(*h, v.(Pair)) }
func (h *MinHeap) Pop() interface{} {
    a := *h
    res := a[len(a) - 1]
    *h = a[:len(a) - 1]
    return res
}

func getFinalState1(nums []int, k int, multiplier int) []int {
    n, mod := len(nums), 1_000_000_007
    pow := func (a, b int) int {
        res := 1
        for ; b > 0; b >>= 1 {
            if (b & 1) == 1 {
                res = (res * a) % mod
            }
            a = ( a * a ) % mod
        }
        return res
    }
    switch {
        case multiplier == 1:
        return nums
        case n == 1:
        return []int{int((nums[0] * pow(multiplier, k)) % mod) }
    }
    idx := make([]int, n)
    for i := range idx {
        idx[i] = i
    }
    sort.Slice(idx, func (i, j int) bool { 
        return nums[idx[i]] < nums[idx[j]] || nums[idx[i]] == nums[idx[j]] && idx[i] < idx[j] 
    })
    var group_idx, group_pow []int
    sum := []int{0}
    b, p := 1, -1
    for i, j := range idx {
        num := nums[j]
        if num >= b {
            group_idx = append(group_idx, i)
            s := 0
            for num >= b {
                b *= multiplier
                p++
                s++
            }
            sum = append(sum, sum[len(sum) - 1] + s * i)
            group_pow = append(group_pow, p)
        }
    }
    group_idx = append(group_idx, n)
    sum = sum[1:]
    g := sort.Search(len(sum), func (i int) bool { return sum[i] > k }) - 1
    group_size := group_idx[g + 1]
    p = (k - sum[g]) / group_size
    res := make([]int, n)
    for u := 0; u <= g; u++ {
        b := pow(multiplier, p + group_pow[g] - group_pow[u])
        for _, j := range idx[group_idx[u]:group_idx[u + 1]] {
            res[j] = int((nums[j] * b) % mod)
        }
    }
    for _, j := range idx[group_size:] {
        res[j] = nums[j]
    }
    idx = idx[:group_size]
    p = (k - sum[g]) % group_size
    if p > 0 {
        if g > 0 {
            new_nums := make([]int, n)
            b = 1
            for u, q := g, group_pow[g]; u >= 0; u-- {
                for q > group_pow[u] {
                    b *= multiplier
                    q--
                }
                for _, j := range idx[group_idx[u]:group_idx[u + 1]] {
                    new_nums[j] = nums[j] * b
                }
            }
            sort.Slice(idx, func (i, j int) bool { 
                return new_nums[idx[i]] < new_nums[idx[j]] || new_nums[idx[i]] == new_nums[idx[j]] && idx[i] < idx[j] 
            })
        }
        for _, j := range idx[:p] {
            res[j] = int((res[j] * multiplier) % mod)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,1,3,5,6], k = 5, multiplier = 2
    // Output: [8,4,6,5,6]
    // Explanation:
    // Operation	Result
    // After operation 1	[2, 2, 3, 5, 6]
    // After operation 2	[4, 2, 3, 5, 6]
    // After operation 3	[4, 4, 3, 5, 6]
    // After operation 4	[4, 4, 6, 5, 6]
    // After operation 5	[8, 4, 6, 5, 6]
    // After applying modulo	[8, 4, 6, 5, 6]
    fmt.Println(getFinalState([]int{2,1,3,5,6}, 5, 2)) // [8,4,6,5,6]
    // Example 2:
    // Input: nums = [100000,2000], k = 2, multiplier = 1000000
    // Output: [999999307,999999993]
    // Explanation:
    // Operation	Result
    // After operation 1	[100000, 2000000000]
    // After operation 2	[100000000000, 2000000000]
    // After applying modulo	[999999307, 999999993]
    fmt.Println(getFinalState([]int{100000,2000}, 2, 1000000)) // [999999307,999999993]

    fmt.Println(getFinalState1([]int{2,1,3,5,6}, 5, 2)) // [8,4,6,5,6]
    fmt.Println(getFinalState1([]int{100000,2000}, 2, 1000000)) // [999999307,999999993]
}