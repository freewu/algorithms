package main

// 3957. Maximum Sum of M Non-Overlapping Subarrays II
// You are given an integer array nums of length n, and three integers m, l, and r.

// Your task is to select at least one and at most m non-overlapping subarrays from nums such that:
//     1. Each selected subarray has a length between [l, r] (inclusive).
//     2. The total sum of all selected subarrays is maximized.

// Return the maximum total sum you can achieve.

// Example 1:
// Input: nums = [4,1,-5,2], m = 2, l = 1, r = 3
// Output: 7
// Explanation:
// One optimal strategy is to:
// Select the subarray [4, 1] with sum 4 + 1 = 5 and the subarray [2] with sum 2. Both subarrays have length between [l, r].
// The total sum of these subarrays is 5 + 2 = 7, which is the maximum achievable sum with at most m = 2 subarrays.

// Example 2:
// Input: nums = [1,0,3,4], m = 2, l = 1, r = 2
// Output: 8
// Explanation:
// One optimal strategy is to:
// Select the subarray [1] with sum 1 and the subarray [3, 4] with sum 3 + 4 = 7. Both subarrays have length between [l, r].
// The total sum of these subarrays is 1 + 7 = 8, which is the maximum achievable sum with at most m = 2 subarrays.

// Example 3:
// Input: nums = [-1,7,-4], m = 1, l = 2, r = 3
// Output: 6
// Explanation:
// Select the subarray [-1, 7] from nums which has length between [l, r].
// The total sum of this subarray is -1 + 7 = 6, which is the maximum achievable sum with at most m = 1 subarray.

// Example 4:
// Input: nums = [-3,-4,-1], m = 2, l = 1, r = 2
// Output: -1
// Explanation:
// All subarrays of nums have negative sums. The optimal strategy is to select the subarray [-1], which has length between [l, r].
// The total sum of this subarray is -1, which is the maximum achievable sum with at most m = 2 subarrays.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     -10^5 <= nums[i] <= 10^5​​​​​​​
//     1 <= m <= n
//     1 <= l <= r <= n

import "fmt"
import "container/list"

// Time Limit Exceeded 989 / 999 testcases passed
func maximumSum(nums []int, m, left, right int) int64 {
    res, n := -1 << 61, len(nums)
    prefix := make([]int, n + 1) // nums 的前缀和
    for i, v := range nums {
        prefix[i + 1] = prefix[i] + v
    }
    // f[i][j] 表示在前 j 个数（下标 0 到 j-1）中选出恰好 i 个子数组，所选元素之和的最大值
    f := make([]int, n + 1)
    for i := 1; i <= m; i++ {
        nf := make([]int, n + 1)
        for j := range nf {
            nf[j] = -1 << 61
        }
        q := []int{}
        // 前 i 个子数组至少占用了 i * left 个位置
        for j := i * left; j <= n; j++ {
            // 1. 入
            k := j - left
            v := f[k] - prefix[k]
            for len(q) > 0 && f[q[len(q)-1]] - prefix[q[len(q)-1]] <= v {
                q = q[:len(q)-1]
            }
            q = append(q, k)
            // 2. 更新
            // 不选 nums[j-1] vs 选一个以 j-1 结尾的子数组
            nf[j] = max(nf[j-1], f[q[0]] - prefix[q[0]] + prefix[j])
            // 3. 出，下一轮循环队首离开窗口
            if q[0] <= j-right {
                q = q[1:]
            }
        }
        // 枚举恰好选 i 个子数组
        f = nf
        res = max(res, f[n])
    }
    return int64(res)
}

func maximumSum1(nums []int, m, l, r int) int64 {
    const inf = -1 << 61
    n := len(nums)
    prefix := make([]int64, n + 1)
    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + int64(nums[i])
    }
    getdp := func(p int) (int64, int) {
        f, count := make([][2]int64, n + 1), make([]int, n + 1)
        for i := range f {
            f[i][0] = inf
            count[i] = 0
        }
        f[0][0] = 0
        count[0] = 0
        fg := func(x, y int) bool {
            cx := f[x][0] - prefix[x]
            cy := f[y][0] - prefix[y]
            if cy > cx {
                return true
            }
            if cy == cx && count[y] < count[x] {
                return true
            }
            return false
        }
        max2 := func(aVal int64, aCnt int, bVal int64, bCnt int) (int64, int) {
            if aVal > bVal {
                return aVal, aCnt
            }
            if bVal > aVal {
                return bVal, bCnt
            }
            if aCnt > bCnt {
                return aVal, aCnt
            }
            return bVal, bCnt
        }
        q := list.New()
        for i := 1; i <= n; i++ {
            for q.Len() > 0 && q.Front().Value.(int) < i-r {
                q.Remove(q.Front())
            }
            j := i - l
            if j >= 0 {
                for q.Len() > 0 && fg(q.Back().Value.(int), j) {
                    q.Remove(q.Back())
                }
                q.PushBack(j)
            }
            curVal, curCnt := f[i-1][0], count[i-1]
            if q.Len() > 0 {
                front := q.Front().Value.(int)
                newVal := f[front][0] - prefix[front] + prefix[i] - int64(p)
                newCnt := count[front] + 1
                curVal, curCnt = max2(curVal, curCnt, newVal, newCnt)
            }
            f[i][0], count[i] = curVal, curCnt
        }
        totalVal := f[n][0]
        totalCnt := count[n]
        if totalCnt == 0 {
            var mx int64 = inf
            q := list.New()
            for i := 1; i <= n; i++ {
                for q.Len() > 0 && q.Front().Value.(int) < i-r {
                    q.Remove(q.Front())
                }
                j := i - l
                if j >= 0 {
                    for q.Len() > 0 && prefix[j] < prefix[q.Back().Value.(int)] {
                        q.Remove(q.Back())
                    }
                    q.PushBack(j)
                }
                if q.Len() > 0 {
                    front := q.Front().Value.(int)
                    current := prefix[i] - prefix[front]
                    if current > mx {
                        mx = current
                    }
                }
            }
            return mx - int64(p), 1
        }
        return totalVal, totalCnt
    }
    aliensDp := func(k int, getDp func(int) (int64, int)) int64 {
        left := 0
        right := int(1e18)
        f1Val, f1Cnt := getDp(0)
        if f1Cnt >= 1 && f1Cnt <= k {
            return f1Val
        }
        penalty := 0
        var bestVal int64 = inf
        // bestCnt 只声明不使用，直接删除

        for left <= right {
            mid := (left + right) >> 1
            cVal, cCnt := getDp(mid)
            if cCnt >= k {
                penalty = mid
                bestVal = cVal
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        return bestVal + int64(penalty * k)
    }
    return aliensDp(m, getdp)
}

func main() {
    // Example 1:
    // Input: nums = [4,1,-5,2], m = 2, l = 1, r = 3
    // Output: 7
    // Explanation:
    // One optimal strategy is to:
    // Select the subarray [4, 1] with sum 4 + 1 = 5 and the subarray [2] with sum 2. Both subarrays have length between [l, r].
    // The total sum of these subarrays is 5 + 2 = 7, which is the maximum achievable sum with at most m = 2 subarrays.
    fmt.Println(maximumSum([]int{4,1,-5,2}, 2, 1, 3)) // 7
    // Example 2:
    // Input: nums = [1,0,3,4], m = 2, l = 1, r = 2
    // Output: 8
    // Explanation:
    // One optimal strategy is to:
    // Select the subarray [1] with sum 1 and the subarray [3, 4] with sum 3 + 4 = 7. Both subarrays have length between [l, r].
    // The total sum of these subarrays is 1 + 7 = 8, which is the maximum achievable sum with at most m = 2 subarrays.
    fmt.Println(maximumSum([]int{1,0,3,4}, 2, 1, 2)) // 8
    // Example 3:
    // Input: nums = [-1,7,-4], m = 1, l = 2, r = 3
    // Output: 6
    // Explanation:
    // Select the subarray [-1, 7] from nums which has length between [l, r].
    // The total sum of this subarray is -1 + 7 = 6, which is the maximum achievable sum with at most m = 1 subarray.
    fmt.Println(maximumSum([]int{-1,7,-4}, 1, 2, 3)) // 6
    // Example 4:
    // Input: nums = [-3,-4,-1], m = 2, l = 1, r = 2
    // Output: -1
    // Explanation:
    // All subarrays of nums have negative sums. The optimal strategy is to select the subarray [-1], which has length between [l, r].
    // The total sum of this subarray is -1, which is the maximum achievable sum with at most m = 2 subarrays.
    fmt.Println(maximumSum([]int{-3,-4,-1}, 2, 1, 2)) // -1

    fmt.Println(maximumSum([]int{1,2,3,4,5,6,7,8,9}, 2, 1, 2)) // 30
    fmt.Println(maximumSum([]int{9,8,7,6,5,4,3,2,1}, 2, 1, 2)) // 30

    fmt.Println(maximumSum1([]int{4,1,-5,2}, 2, 1, 3)) // 7
    fmt.Println(maximumSum1([]int{1,0,3,4}, 2, 1, 2)) // 8
    fmt.Println(maximumSum1([]int{-1,7,-4}, 1, 2, 3)) // 6
    fmt.Println(maximumSum1([]int{-3,-4,-1}, 2, 1, 2)) // -1
    fmt.Println(maximumSum1([]int{1,2,3,4,5,6,7,8,9}, 2, 1, 2)) // 30
    fmt.Println(maximumSum1([]int{9,8,7,6,5,4,3,2,1}, 2, 1, 2)) // 30
}