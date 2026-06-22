package main

// 3962. Maximum Subarray Sum After at Most K Swaps
// You are given an integer array nums and an integer k.

// You are allowed to perform at most k swap operations on the array.

// In one swap operation, you may choose any two indices i and j and swap nums[i] and nums[j].

// Return an integer denoting the maximum possible subarray sum after performing the swaps.

// Example 1:
// Input: nums = [1,-1,0,2], k = 1
// Output: 3
// Explanation:
// We can swap on indices 1 and 3, resulting in the array [1, 2, 0, -1].
// The subarray [1, 2] has a sum of 3, which is the maximum possible subarray sum after at most k = 1​​​​​​​ swap.

// Example 2:
// Input: nums = [4,3,2,4], k = 2
// Output: 13
// Explanation:
// The maximum possible subarray sum after at most k = 2 swaps is the sum of the entire array, which is 13.

// Example 3:
// Input: nums = [-1,-2], k = 0
// Output: -1
// Explanation:
// k = 0 swaps are allowed.
// The possible subarrays are [-1], [-2], and [-1, -2], with sums -1, -2, and -3 respectively.
// Among these sums, the maximum is -1.

// Constraints:
//     1 <= nums.length <= 1500
//     -10^5 <= nums[i] <= 10^5
//     0 <= k <= nums.length

import "fmt"
import "math/bits"
import "slices"
import "sort"

var width int

type Pair struct{ count, sum int }
type Fenwick []Pair

// 添加 num 个 val，其中 val 离散化后的值为 i
// 如果 num < 0，表示减少 -num 个 val
func (t Fenwick) update(i, num, val int) {
    for ; i < len(t); i += i & -i {
        t[i].count += num
        t[i].sum += val
    }
}

// 返回第 k 小的数（k 从 1 开始）
func (t Fenwick) kth(k int, sorted []int) int {
    i := 0
    for b := 1 << (width - 1); b > 0; b >>= 1 {
        if j := i | b; j < len(t) && t[j].count < k {
            k -= t[j].count
            i = j
        }
    }
    return sorted[i]
}

// 返回前 k 小的数之和（k 从 1 开始）
func (t Fenwick) preSum(k int, sorted []int) int {
    res, i := 0, 0
    for b := 1 << (width - 1); b > 0; b >>= 1 {
        if j := i | b; j < len(t) && t[j].count < k {
            k -= t[j].count
            res += t[j].sum
            i = j
        }
    }
    res += sorted[i] * k // 加上等于第 k 小的数
    return res
}

func maxSum(nums []int, k int) int64 {
    // O(n) 特判：能否把正数都聚在一起
    allPosSum, allPosCount := 0, 0
    for _, v := range nums {
        if v > 0 {
            allPosSum += v
            allPosCount++
        }
    }
    if allPosCount == 0 { // 没有正数
        return int64(slices.Max(nums))
    }
    // 定长滑动窗口模板，窗口长度为 allPosCount
    count := 0
    for i, v := range nums {
        if v > 0 {
            count++
        }
        left := i - allPosCount + 1
        if left < 0 {
            continue
        }
        if count + k >= allPosCount { // 可以把正数都聚在一起   
            return int64(allPosSum)
        }
        if nums[left] > 0 {
            count--
        }
    }
    // 离散化
    res, n := -1 << 61,len(nums)
    sorted := slices.Clone(nums)
    slices.Sort(sorted)
    sorted = slices.Compact(sorted)
    m := len(sorted)
    width = bits.Len(uint(m))
    rank := make([]int, n) // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
    allPosTree := make(Fenwick, m + 1) // 包含所有正数的树状数组

    for i, v := range nums {
        rank[i] = sort.SearchInts(sorted, v) + 1
        if v > 0 {
            allPosTree.update(rank[i], 1, v)
        }
    }
    // 枚举子数组左端点
    for left := range nums {
        negTree := make(Fenwick, m + 1)   
        posTree := slices.Clone(allPosTree)
        posSum, posCount, negCount, subSum := allPosSum, allPosCount, 0, 0
        // 枚举子数组右端点
        for right := left; right < n; right++ {
            // v 从子数组外移到子数组内
            v := nums[right]
            rk := rank[right]
            subSum += v
            if v > 0 {
                posTree.update(rk, -1, -v)
                posSum -= v
                posCount--
            } else if v < 0 {
                negTree.update(rk, 1, v)
                negCount++
            }
            // 计算通过交换导致的元素和的增量
            delta := 0
            needSwap := min(negCount, posCount, k)
            if needSwap > 0 {
                inSum := negTree.preSum(needSwap, sorted)
                outSum := posSum - posTree.preSum(posCount - needSwap, sorted)
                delta = outSum - inSum
            }
            res = max(res, subSum + delta)
        }
    }
    return int64(res)
}

type Pair1 struct{ count, sum int }
type Fenwick1 struct {
    t       []Pair1
    sorted  []int
    width int
}

func newFenwickTree(sorted []int) Fenwick1 {
    n := len(sorted)
    return Fenwick1{
        t:       make([]Pair1, n+1),
        sorted:  sorted,
        width: 1 << (bits.Len(uint(n)) - 1),
    }
}

// 添加 num 个 val，其中 val 离散化后的值为 i（i 从 1 开始）
// 如果 num < 0，表示减少 -num 个 val
func (f *Fenwick1) update(i, num, val int) {
    for ; i < len(f.t); i += i & -i {
        f.t[i].count += num
        f.t[i].sum += val
    }
}

// 返回第 k 小的数（k 从 1 开始）
func (f *Fenwick1) kth(k int) int {
    i := 0
    for b := f.width; b > 0; b >>= 1 {
        if nxt := i | b; nxt < len(f.t) && f.t[nxt].count < k {
            k -= f.t[nxt].count
            i = nxt
        }
    }
    return f.sorted[i]
}

// 返回前 k 小的数之和（k 从 1 开始）
func (f *Fenwick1) preSum(k int) int {
    res, i := 0, 0
    for b := f.width; b > 0; b >>= 1 {
        if nxt := i | b; nxt < len(f.t) && f.t[nxt].count < k {
            k -= f.t[nxt].count
            res += f.t[nxt].sum
            i = nxt
        }
    }
    // 加上等于第 k 小的数
    res += f.sorted[i] * k
    return res
}

func maxSum1(nums []int, k int) int64 {
    // 特判：能否把正数都聚在一起
    allPosSum, allPosCount := 0, 0
    for _, x := range nums {
        if x > 0 {
            allPosSum += x
            allPosCount++
        }
    }
    if allPosCount == 0 { // 没有正数
        return int64(slices.Max(nums))
    }
    // 定长滑动窗口模板，窗口长度为 allPosCount
    count := 0
    for i, x := range nums {
        if x > 0 {
            count++
        }
        left := i - allPosCount + 1
        if left < 0 {
            continue
        }
        if count + k >= allPosCount { // 可以把正数都聚在一起
            return int64(allPosSum)
        }
        if nums[left] > 0 {
            count--
        }
    }

    // 离散化
    n := len(nums)
    sorted := slices.Clone(nums)
    slices.Sort(sorted)
    sorted = slices.Compact(sorted)
    rank := make([]int, n) // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
    allPosTree := newFenwickTree(sorted) // 包含所有正数的树状数组
    for i, x := range nums {
        rank[i] = sort.SearchInts(sorted, x) + 1
        if x > 0 {
            allPosTree.update(rank[i], 1, x)
        }
    }
    negTree := newFenwickTree(sorted)
    posTree := newFenwickTree(sorted)
    res := -1 << 61
    // 枚举子数组左端点
    for left := range nums {
        clear(negTree.t)
        copy(posTree.t, allPosTree.t)
        posSum, posCount := allPosSum, allPosCount
        negCount, subSum := 0, 0
        // 枚举子数组右端点
        for right := left; right < n; right++ {
            // x 从子数组外移到子数组内
            x := nums[right]
            rk := rank[right]
            subSum += x
            if x > 0 {
                posTree.update(rk, -1, -x)
                posSum -= x
                posCount--
            } else if x < 0 {
                negTree.update(rk, 1, x)
                negCount++
            }
            // 计算通过交换导致的元素和的增量
            delta := 0
            needSwap := min(negCount, posCount, k)
            if needSwap > 0 {
                inSum := negTree.preSum(needSwap)
                outSum := posSum - posTree.preSum(posCount - needSwap)
                delta = outSum - inSum
            }
            res = max(res, subSum +delta)
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,-1,0,2], k = 1
    // Output: 3
    // Explanation:
    // We can swap on indices 1 and 3, resulting in the array [1, 2, 0, -1].
    // The subarray [1, 2] has a sum of 3, which is the maximum possible subarray sum after at most k = 1​​​​​​​ swap.
    fmt.Println(maxSum([]int{1,-1,0,2}, 1)) // 3
    // Example 2:
    // Input: nums = [4,3,2,4], k = 2
    // Output: 13
    // Explanation:
    // The maximum possible subarray sum after at most k = 2 swaps is the sum of the entire array, which is 13.
    fmt.Println(maxSum([]int{1,-1,0,2}, 1)) // 3
    // Example 3:
    // Input: nums = [-1,-2], k = 0
    // Output: -1
    // Explanation:
    // k = 0 swaps are allowed.
    // The possible subarrays are [-1], [-2], and [-1, -2], with sums -1, -2, and -3 respectively.
    // Among these sums, the maximum is -1.
    fmt.Println(maxSum([]int{1,-1,0,2}, 1)) // 3

    fmt.Println(maxSum([]int{1,2,3,4,5,6,7,8,9}, 2)) // 45
    fmt.Println(maxSum([]int{9,8,7,6,5,4,3,2,1}, 2)) // 45

    fmt.Println(maxSum1([]int{1,-1,0,2}, 1)) // 3
    fmt.Println(maxSum1([]int{1,-1,0,2}, 1)) // 3
    fmt.Println(maxSum1([]int{1,-1,0,2}, 1)) // 3
    fmt.Println(maxSum1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 45
    fmt.Println(maxSum1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 45
}