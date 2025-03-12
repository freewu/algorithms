package main

// 3444. Minimum Increments for Target Multiples in an Array
// You are given two arrays, nums and target.

// In a single operation, you may increment any element of nums by 1.

// Return the minimum number of operations required so that each element in target has at least one multiple in nums.

// Example 1:
// Input: nums = [1,2,3], target = [4]
// Output: 1
// Explanation:
// The minimum number of operations required to satisfy the condition is 1.
// Increment 3 to 4 with just one operation, making 4 a multiple of itself.

// Example 2:
// Input: nums = [8,4], target = [10,5]
// Output: 2
// Explanation:
// The minimum number of operations required to satisfy the condition is 2.
// Increment 8 to 10 with 2 operations, making 10 a multiple of both 5 and 10.

// Example 3:
// Input: nums = [7,9,10], target = [7]
// Output: 0
// Explanation:
// Target 7 already has a multiple in nums, so no additional operations are needed.

// Constraints:
//     1 <= nums.length <= 5 * 10^4
//     1 <= target.length <= 4
//     target.length <= nums.length
//     1 <= nums[i], target[i] <= 10^4

import "fmt"
import "container/heap"
import "slices"

func minimumIncrements(nums []int, target []int) int {
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    lcm := func(x, y int) int { return x * y / gcd(x, y) }
    n, inf := len(target), 1 << 31
    arr, dp := make([]int, 1 << n), make([]int, 1 << n)
    for i := 1; i < (1 << n); i++ {
        val := 1
        for j := 0; j < n; j++ {
            if (i & (1 << j)) != 0 {
                val = lcm(val, target[j])
            }
        }
        arr[i] = val
    }
    for i := 0; i < (1 << n); i++ {
        dp[i] = inf
    }
    dp[0] = 0
    for _, x := range nums {
        newdp := make([]int, len(dp))
        copy(newdp, dp)
        for i := 1; i < (1 << n); i++ {
            r, cost := x % arr[i], 0
            if r != 0 {
                cost = arr[i] - r
            }
            for j := 0; j < (1 << n); j++ {
                if newdp[j | i] > dp[j] + cost {
                    newdp[j | i] = dp[j] + cost
                }
            }
        }
        dp = newdp
    }
    return dp[(1 << n) - 1]
}

type Pair struct{ op, i int }
type MaxHeap []Pair
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].op > h[j].op }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(v any)        { *h = append(*h, v.(Pair)) }
func (h *MaxHeap) Pop() (_ any)      { return }
func (h *MaxHeap) update(p Pair) {
    if p.op < (*h)[0].op {
        (*h)[0] = p
        heap.Fix(h, 0)
    }
}

func minimumIncrements1(nums []int, target []int) int {
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    lcm := func(x, y int) int { return x * y / gcd(x, y) }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    n := len(target)
    arr := make([]int, 1 << n)
    arr[0] = 1
    for i, t := range target {
        bit := 1 << i
        for mask, l := range arr[:bit] {
            arr[bit|mask] = lcm(t, l)
        }
    }
    mx := max(slices.Max(nums) * 2, slices.Max(target))
    mp := make(map[int]bool)
    for _, v := range arr[1:] {
        if v > mx { continue }
        h := MaxHeap{}
        for i, x := range nums {
            p := Pair{(v - x % v) % v, i}
            if len(h) < n {
                heap.Push(&h, p)
            } else {
                h.update(p)
            }
        }
        for _, p := range h {
            mp[p.i] = true
        }
    }
    f := make([]int, 1 << n)
    for j := 1; j < 1 << n; j++ {
        f[j] = 1 << 31
    }
    for i := range mp {
        x := nums[i]
        for j := 1 << n - 1; j > 0; j-- {
            for sub := j; sub > 0; sub = (sub - 1) & j {
                l := arr[sub]
                f[j] = min(f[j], f[j^sub]+(l - x % l) % l)
            }
        }
    }
    return f[1 << n - 1]
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3], target = [4]
    // Output: 1
    // Explanation:
    // The minimum number of operations required to satisfy the condition is 1.
    // Increment 3 to 4 with just one operation, making 4 a multiple of itself.
    fmt.Println(minimumIncrements([]int{1,2,3}, []int{4})) // 1
    // Example 2:
    // Input: nums = [8,4], target = [10,5]
    // Output: 2
    // Explanation:
    // The minimum number of operations required to satisfy the condition is 2.
    // Increment 8 to 10 with 2 operations, making 10 a multiple of both 5 and 10.
    fmt.Println(minimumIncrements([]int{8,4}, []int{10,5})) // 2
    // Example 3:
    // Input: nums = [7,9,10], target = [7]
    // Output: 0
    // Explanation:
    // Target 7 already has a multiple in nums, so no additional operations are needed.
    fmt.Println(minimumIncrements([]int{7,9,10}, []int{7})) // 0

    fmt.Println(minimumIncrements([]int{1,2,3,4,5,6,7,8,9}, []int{4})) // 0
    fmt.Println(minimumIncrements([]int{9,8,7,6,5,4,3,2,1}, []int{4})) // 0

    fmt.Println(minimumIncrements1([]int{1,2,3}, []int{4})) // 1
    fmt.Println(minimumIncrements1([]int{8,4}, []int{10,5})) // 2
    fmt.Println(minimumIncrements1([]int{7,9,10}, []int{7})) // 0
    fmt.Println(minimumIncrements1([]int{1,2,3,4,5,6,7,8,9}, []int{4})) // 0
    fmt.Println(minimumIncrements1([]int{9,8,7,6,5,4,3,2,1}, []int{4})) // 0
}