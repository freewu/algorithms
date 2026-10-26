package main

// 4087. Maximum Subarray Sum of Fixed Length After at Most K Swaps
// You are given an integer array nums and two integers length and k.

// You may perform at most k swap operations on nums.

// In one swap operation, you may choose two distinct indices i and j and swap nums[i] and nums[j].

// Return the maximum possible sum of a subarray of length exactly length after performing the swaps.

// Example 1:
// Input: nums = [1,-1,0,2], length = 2, k = 1
// Output: 3
// Explanation:
// - Swap the elements at indices 1 and 3, resulting in [1, 2, 0, -1].
// - The subarray [1, 2] has length 2 and sum 3.
// - No subarray of length 2 can have a larger sum, since 1 and 2 are the two largest elements.

// Example 2:
// Input: nums = [4,3,2,4], length = 3, k = 2
// Output: 11
// Explanation:
// - Swap the elements at indices 2 and 3, resulting in [4, 3, 4, 2].
// - The subarray [4, 3, 4] has length 3 and sum 11.
// - This is the sum of the three largest elements, so no larger answer is possible. Only one swap is needed.

// Example 3:
// Input: nums = [-1,-2], length = 2, k = 0
// Output: -3
// Explanation:
// No swaps are allowed. The only subarray of length 2 is [-1, -2], whose sum is -3.

// Constraints:
//     - 1 <= nums.length <= 300
//     - -10^{9} <= nums[i] <= 10^{9}
//     - 1 <= length <= nums.length
//     - 0 <= k <= nums.length

import "fmt"
import "sort"
import "container/heap"

// 贪心
func maxSum(nums []int, length int, k int) int64 {
    res, n := int64(-1 << 61), len(nums)
    for start := 0; start <= n - length; start++ {
        inside, outside := make([]int, 0, length), make([]int, 0, n - length)
        for i := 0; i < n; i++ {
            if i >= start && i < start+length {
                inside = append(inside, nums[i])
            } else {
                outside = append(outside, nums[i])
            }
        }
        sort.Ints(inside)
        sort.Ints(outside)
        for i, j := 0, len(outside)-1; i < j; i, j = i+1, j-1 {
            outside[i], outside[j] = outside[j], outside[i]
        }
        // Base sum of the window
        curr := int64(0)
        for _, v := range inside {
            curr += int64(v)
        }
        // Greedily swap up to k times
        swaps := min(k, min(length, len(outside)))
        for i := 0; i < swaps; i++ {
            if outside[i] > inside[i] {
                curr += int64(outside[i] - inside[i])
            } else {
                break // no further improvement possible
            }
        }
        if curr > res {
            res = curr
        }
    }
    return res
}

// HeapItem 对应 python [val, -x]
type HeapItem struct {
    qVal int
    negX int
}

type Tre struct {
    sm  int64
    h   Heap
    lis map[int]bool
    len int
    jd  int
}

func NewTre(x, y int) *Tre {
    t := &Tre{
        sm:  0,
        h:   make(Heap, 0),
        lis: make(map[int]bool),
        len: x,
        jd:  y,
    }
    heap.Init(&t.h)
    return t
}

// cr 清理无效元素，对应 python cr()
func (t *Tre) cr() {
    for t.h.Len() > 0 {
        top := t.h[0]
        realX := -top.negX
        if t.lis[realX] {
            break
        }
        heap.Pop(&t.h)
    }
}

// add 返回 (p, q) 和原版Python完全一致
func (t *Tre) add(x, y int) (int, int) {
    t.cr()
    var p, q int = -1, -1
    if t.len > 0 {
        t.len--
    } else {
        top := t.h[0]
        if top.qVal*t.jd > y {
            return x, y
        }
        popped := heap.Pop(&t.h).(HeapItem)
        q = popped.qVal * t.jd
        p = -popped.negX
        t.cr()
        delete(t.lis, p)
        t.sm -= int64(q)
    }
    pushQ := y * t.jd
    heap.Push(&t.h, HeapItem{qVal: pushQ, negX: -x})
    t.lis[x] = true
    t.sm += int64(y)
    return p, q
}

func (t *Tre) dlt(x, y int) {
    delete(t.lis, x)
    t.sm -= int64(y)
    t.len++
    t.cr()
}

// gdt 返回 a,b （下标，值），原版Python，不做空保护
func (t *Tre) gdt() (int, int) {
    t.cr()
    top := t.h[0]
    b := top.qVal * t.jd
    a := -top.negX
    t.dlt(a, b)
    return a, b
}

type Heap []HeapItem

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].qVal < h[j].qVal }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(HeapItem))
}

func (h *Heap) Pop() interface{} {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

func maxSum1(nums []int, length int, k int) int64 {
    doadd := func(A, B *Tre, x, y int) (int, int) {
        a, b := A.add(x, y)
        if a >= 0 {
            B.add(a, b)
        }
        return a, b
    }
    dodlt := func(A, B *Tre, x, y int) (int, int) {
        if A.lis[x] {
            A.dlt(x, y)
            a, b := B.gdt()
            A.add(a, b)
            return a, b
        } else {
            B.dlt(x, y)
            return -1, -1
        }
    }
    n := len(nums)
    if k == 0 {
        var sm int64 = 0
        for i := 0; i < length; i++ {
            sm += int64(nums[i])
        }
        res := sm
        for i := length; i < n; i++ {
            sm += int64(nums[i]) - int64(nums[i-length])
            res = max(res, sm)
        }
        return res
    }
    if length <= k+1 {
        tmp := make([]int, n)
        copy(tmp, nums)
        // 排序取最大 length 个
        for i := 0; i < n; i++ {
            for j := i + 1; j < n; j++ {
                if tmp[i] > tmp[j] {
                    tmp[i], tmp[j] = tmp[j], tmp[i]
                }
            }
        }
        var sum int64 = 0
        for i := n - length; i < n; i++ {
            sum += int64(tmp[i])
        }
        return sum
    }
    if n == length {
        var sum int64 = 0
        for _, v := range nums {
            sum += int64(v)
        }
        return sum
    }
    A := NewTre(length-k, 1)
    B := NewTre(k, -1)
    C := NewTre(k, 1)
    D := NewTre(n, -1)
    for x, y := range nums {
        doadd(C, D, x, y)
    }
    res := int64(-1 << 61)
    for x, y := range nums {
        if x >= length {
            nx := x - length
            ny := nums[nx]
            a, b := dodlt(A, B, nx, ny)
            if a >= 0 {
                dodlt(C, D, a, b)
                doadd(C, D, nx, ny)
            }
        }
        a, b := doadd(A, B, x, y)
        if a != x {
            dodlt(C, D, x, y)
            if a >= 0 {
                doadd(C, D, a, b)
            }
        }
        if x >= length-1 {
            res = max(res, A.sm+C.sm)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,-1,0,2], length = 2, k = 1
    // Output: 3
    // Explanation:
    // - Swap the elements at indices 1 and 3, resulting in [1, 2, 0, -1].
    // - The subarray [1, 2] has length 2 and sum 3.
    // - No subarray of length 2 can have a larger sum, since 1 and 2 are the two largest elements.
    fmt.Println(maxSum([]int{1, -1, 0, 2}, 2, 1)) // 3
    // Example 2:
    // Input: nums = [4,3,2,4], length = 3, k = 2
    // Output: 11
    // Explanation:
    // - Swap the elements at indices 2 and 3, resulting in [4, 3, 4, 2].
    // - The subarray [4, 3, 4] has length 3 and sum 11.
    // - This is the sum of the three largest elements, so no larger answer is possible. Only one swap is needed.
    fmt.Println(maxSum([]int{4, 3, 2, 4}, 3, 2)) // 11
    // Example 3:
    // Input: nums = [-1,-2], length = 2, k = 0
    // Output: -3
    // Explanation:
    // No swaps are allowed. The only subarray of length 2 is [-1, -2], whose sum is -3.
    fmt.Println(maxSum([]int{-1, -2}, 2, 0)) // -3

    fmt.Println(maxSum([]int{46,29,-44}, 1, 2)) // 46
    fmt.Println(maxSum([]int{1,2,3,4,5,6,7,8,9}, 2, 1)) // 17
    fmt.Println(maxSum([]int{9,8,7,6,5,4,3,2,1}, 2, 1)) // 17
    fmt.Println(maxSum([]int{-24,44,25,38}, 3, 1)) // 107
    fmt.Println(maxSum([]int{18,-6,2,-17,-1}, 4, 1)) // 13
    fmt.Println(maxSum([]int{-36,-13,5,-30,8,-50}, 3, 1)) // 0

    fmt.Println(maxSum1([]int{1, -1, 0, 2}, 2, 1)) // 3
    fmt.Println(maxSum1([]int{4, 3, 2, 4}, 3, 2)) // 11
    fmt.Println(maxSum1([]int{-1, -2}, 2, 0)) // -3
    fmt.Println(maxSum1([]int{46,29,-44}, 1, 2)) // 46
    fmt.Println(maxSum1([]int{1,2,3,4,5,6,7,8,9}, 2, 1)) // 17
    fmt.Println(maxSum1([]int{9,8,7,6,5,4,3,2,1}, 2, 1)) // 17
    fmt.Println(maxSum1([]int{-24,44,25,38}, 3, 1)) // 107
    fmt.Println(maxSum1([]int{18,-6,2,-17,-1}, 4, 1)) // 13
    fmt.Println(maxSum1([]int{-36,-13,5,-30,8,-50}, 3, 1)) // 0
}