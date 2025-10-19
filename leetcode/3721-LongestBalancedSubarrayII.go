package main

// 3721. Longest Balanced Subarray II
// You are given an integer array nums.

// A subarray is called balanced if the number of distinct even numbers in the subarray is equal to the number of distinct odd numbers.

// Return the length of the longest balanced subarray.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [2,5,4,3]
// Output: 4
// Explanation:
// The longest balanced subarray is [2, 5, 4, 3].
// It has 2 distinct even numbers [2, 4] and 2 distinct odd numbers [5, 3]. Thus, the answer is 4.

// Example 2:
// Input: nums = [3,2,2,5,4]
// Output: 5
// Explanation:
// The longest balanced subarray is [3, 2, 2, 5, 4].
// It has 2 distinct even numbers [2, 4] and 2 distinct odd numbers [3, 5]. Thus, the answer is 5.

// Example 3:
// Input: nums = [1,2,3,2]
// Output: 3
// Explanation:
// The longest balanced subarray is [2, 3, 2].
// It has 1 distinct even number [2] and 1 distinct odd number [3]. Thus, the answer is 3.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"
import "math/bits"

type Pair struct{ mn, mx int }
type lazySegment []struct {
    l, r int
    Pair
    todo int
}

func merge(l, r Pair) Pair {
	return Pair{min(l.mn, r.mn), max(l.mx, r.mx)}
}

func (t lazySegment) apply(o int, f int) {
    cur := &t[o]
    cur.mn += f
    cur.mx += f
    cur.todo += f
}

func (t lazySegment) maintain(o int) {
    t[o].Pair = merge(t[o<<1].Pair, t[o<<1|1].Pair)
}

func (t lazySegment) spread(o int) {
    f := t[o].todo
    if f == 0 { return }
    t.apply(o<<1, f)
    t.apply(o<<1|1, f)
    t[o].todo = 0
}

func (t lazySegment) build(o, l, r int) {
    t[o].l, t[o].r = l, r
    if l == r {return }
    m := (l + r) >> 1
    t.build(o<<1, l, m)
    t.build(o<<1|1, m + 1, r)
}

func (t lazySegment) update(o, l, r int, f int) {
    if l <= t[o].l && t[o].r <= r {
        t.apply(o, f)
        return
    }
    t.spread(o)
    m := (t[o].l + t[o].r) >> 1
    if l <= m {
        t.update(o<<1, l, r, f)
    }
    if m < r {
        t.update(o<<1|1, l, r, f)
    }
    t.maintain(o)
}

// 查询 [l,r] 内第一个等于 target 的元素下标
func (t lazySegment) findFirst(o, l, r, target int) int {
    if t[o].l > r || t[o].r < l || target < t[o].Pair.mn || target > t[o].Pair.mx {
        return -1
    }
    if t[o].l == t[o].r {  return t[o].l }
    t.spread(o)
    index := t.findFirst(o<<1, l, r, target)
    if index < 0 {
        // 去右子树找
        index = t.findFirst(o<<1|1, l, r, target)
    }
    return index
}

func longestBalanced(nums []int) int {
    res, sum, n := 0, 0, len(nums)
    t := make(lazySegment, 2 << bits.Len(uint(n)))
    t.build(1, 0, n)
    last := map[int]int{} // nums 的元素上一次出现的位置
    for i := 1; i <= n; i++ {
        x := nums[i-1]
        v := x % 2 * 2 - 1
        if j := last[x]; j == 0 { // 首次遇到 x
            sum += v
            t.update(1, i, n, v) // sum[i:] 增加 v
        } else { // 再次遇到 x
            t.update(1, j, i-1, -v) // 撤销之前对 sum[j:i] 的增加
        }
        last[x] = i
        j := t.findFirst(1, 0, i-1, sum)
        if j >= 0 {
            res = max(res, i - j)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,5,4,3]
    // Output: 4
    // Explanation:
    // The longest balanced subarray is [2, 5, 4, 3].
    // It has 2 distinct even numbers [2, 4] and 2 distinct odd numbers [5, 3]. Thus, the answer is 4.
    fmt.Println(longestBalanced([]int{2,5,4,3})) // 4
    // Example 2:
    // Input: nums = [3,2,2,5,4]
    // Output: 5
    // Explanation:
    // The longest balanced subarray is [3, 2, 2, 5, 4].
    // It has 2 distinct even numbers [2, 4] and 2 distinct odd numbers [3, 5]. Thus, the answer is 5.
    fmt.Println(longestBalanced([]int{3,2,2,5,4})) // 5
    // Example 3:
    // Input: nums = [1,2,3,2]
    // Output: 3
    // Explanation:
    // The longest balanced subarray is [2, 3, 2].
    // It has 1 distinct even number [2] and 1 distinct odd number [3]. Thus, the answer is 3.
    fmt.Println(longestBalanced([]int{1,2,3,2})) // 3

    fmt.Println(longestBalanced([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(longestBalanced([]int{9,8,7,6,5,4,3,2,1})) // 8
}