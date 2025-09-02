package main

// 3670. Maximum Product of Two Integers With No Common Bits
// You are given an integer array nums.

// Your task is to find two distinct indices i and j such that the product nums[i] * nums[j] is maximized, 
// and the binary representations of nums[i] and nums[j] do not share any common set bits.

// Return the maximum possible product of such a pair. 
// If no such pair exists, return 0.

// Example 1:
// Input: nums = [1,2,3,4,5,6,7]
// Output: 12
// Explanation:
// The best pair is 3 (011) and 4 (100). They share no set bits and 3 * 4 = 12.

// Example 2:
// Input: nums = [5,6,4]
// Output: 0
// Explanation:
// Every pair of numbers has at least one common set bit. Hence, the answer is 0.

// Example 3:
// Input: nums = [64,8,32]
// Output: 2048
// Explanation:
// No pair of numbers share a common bit, so the answer is the product of the two maximum elements, 64 and 32 (64 * 32 = 2048).

// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^6

import "fmt"
import "math/bits"
import "slices"

type Tree struct {
    Left *Tree
    Right *Tree
    Num int
    NumRest int
}

func (t *Tree) Calc(arr []int, cnt int) (res int) {
    if cnt == 0 { return t.NumRest }
    if len(arr) == 0 {
        return t.Num
    }
    if arr[0] == 1 && t.Left == nil {
        return 0
    } else if arr[0] == 1 {
        return t.Left.Calc(arr[1:], cnt-1)
    }
    if t.Left != nil {
        res = max(res, t.Left.Calc(arr[1:], cnt))
    } 
    if t.Right != nil {
        res = max(res, t.Right.Calc(arr[1:], cnt))
    }
    return res
}

func (t *Tree) Insert(arr []int, num int) *Tree {
    if len(arr) == 0 {
        t.Num = max(t.Num, num)
        return t
    }
    if arr[0] == 1 && t.Right == nil {
        t.Right = &Tree{}
        t.Right = t.Right.Insert(arr[1:], num)
    } else if arr[0] == 1 {
        t.Right = t.Right.Insert(arr[1:], num)
    } else if t.Left == nil {
        t.Left = &Tree{}
        t.Left.NumRest = max(t.Left.NumRest, num)
        t.Left = t.Left.Insert(arr[1:], num)
    } else {
        t.Left.NumRest = max(t.Left.NumRest, num)
        t.Left = t.Left.Insert(arr[1:], num)
    }
    return t
}

func maxProduct(nums []int) (res int64) {
    bits := make(map[int][]int)
    for i, n := range nums {
        bit := make([]int, 0)
        for n > 0 {
            bit = append(bit, n&1)
            n >>= 1
        }
        bits[nums[i]] = bit
    }
    t := &Tree{}
    for k := range bits {
        temp := make([]int, 20)
        for i := 0; i < len(bits[k]); i++ {
            temp[len(temp)-1-i] = bits[k][i]
        }
        bits[k] = temp
    }
    arr, visited := make([]bool, 1_000_001), make([]bool, 1_000_001)
    for _, n := range nums {
        if arr[n] { continue }
        arr[n] = true
        t = t.Insert(bits[n], n)
        if visited[n]  {continue }
        visited[n] = true
        count := 0
        for i := 0; i < len(bits[n]); i++ {
            if bits[n][i] == 1 { count++ }
        }
        temp := t.Calc(bits[n], count)
        res = max(res, int64(temp) * int64(n))
    }
    return res
}

func maxProduct1(nums []int) int64 {
    w := bits.Len(uint(slices.Max(nums)))
    res, u := 0, 1 << w
    dp := make([]int, u)
    for _, v := range nums {
        dp[v] = v
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < w; i++ {
        for s := 3; s < u; s++ {
            s |= 1 << i // 快速跳到第 i 位是 1 的 s
            dp[s] = max(dp[s], dp[s^1<<i])
        }
    }
    for _, v := range nums {
        res = max(res, v * dp[u-1^v])
    }
    return int64(res)
}

func maxProduct2(nums []int) int64 {
    if len(nums) < 2 {
        return 0
    }
    // Determine how many bits are actually needed
    res, bits, mx := 0, 0, 0
    for _, v := range nums {
        if v > mx {
            mx = v
        }
    }
    for (1 << bits) <= mx {
        bits++
    }
    if bits == 0 {
        bits = 1
    }
    size := 1 << bits
    dp := make([]int, size)
    // For each exact mask, keep the largest value with that mask
    for _, v := range nums {
        if v > dp[v] {
            dp[v] = v
        }
    }
    // SOS DP: for every mask store the maximum value among all of its submasks
    for bit := 0; bit < bits; bit++ {
        bitMask := 1 << bit
        for mask := 0; mask < size; mask++ {
            if mask&bitMask != 0 {
                sub := mask ^ bitMask
                if dp[sub] > dp[mask] {
                    dp[mask] = dp[sub]
                }
            }
        }
    }
    allMask := size - 1
    for _, v := range nums {
        best := dp[allMask^v]
        if best > 0 {
            product := v * best
            if product > res {
                res = product
            }
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5,6,7]
    // Output: 12
    // Explanation:
    // The best pair is 3 (011) and 4 (100). They share no set bits and 3 * 4 = 12.
    fmt.Println(maxProduct([]int{1,2,3,4,5,6,7})) // 12
    // Example 2:
    // Input: nums = [5,6,4]
    // Output: 0
    // Explanation:
    // Every pair of numbers has at least one common set bit. Hence, the answer is 0.
    fmt.Println(maxProduct([]int{5,6,4})) // 0
    // Example 3:
    // Input: nums = [64,8,32]
    // Output: 2048
    // Explanation:
    // No pair of numbers share a common bit, so the answer is the product of the two maximum elements, 64 and 32 (64 * 32 = 2048).
    fmt.Println(maxProduct([]int{64,8,32})) // 2048

    fmt.Println(maxProduct([]int{1,2,3,4,5,6,7,8,9})) // 56
    fmt.Println(maxProduct([]int{9,8,7,6,5,4,3,2,1})) // 56

    fmt.Println(maxProduct1([]int{1,2,3,4,5,6,7})) // 12
    fmt.Println(maxProduct1([]int{5,6,4})) // 0
    fmt.Println(maxProduct1([]int{64,8,32})) // 2048
    fmt.Println(maxProduct1([]int{1,2,3,4,5,6,7,8,9})) // 56
    fmt.Println(maxProduct1([]int{9,8,7,6,5,4,3,2,1})) // 56

    fmt.Println(maxProduct2([]int{1,2,3,4,5,6,7})) // 12
    fmt.Println(maxProduct2([]int{5,6,4})) // 0
    fmt.Println(maxProduct2([]int{64,8,32})) // 2048
    fmt.Println(maxProduct2([]int{1,2,3,4,5,6,7,8,9})) // 56
    fmt.Println(maxProduct2([]int{9,8,7,6,5,4,3,2,1})) // 56
}