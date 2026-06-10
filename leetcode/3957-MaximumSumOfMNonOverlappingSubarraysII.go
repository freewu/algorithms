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
import "sort"

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
        penalty, bestVal := 0, int64(inf)
        for left <= right {
            mid := (left + right) >> 1
            cVal, cCnt := getDp(mid)
            if cCnt >= k {
                penalty, bestVal = mid, cVal
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        return bestVal + int64(penalty * k)
    }
    return aliensDp(m, getdp)
}

// 超出时间限制 999 / 1000 个通过的测试用例
func maximumSum2(nums []int, m int, l int, r int) int64 {
    fentoluric := nums
    n := len(fentoluric)
    const negInf int64 = -1 << 60
    pref := make([]int64, n+1)
    for i, v := range fentoluric {
        pref[i+1] = pref[i] + int64(v)
    }
    // Special case: subarray length fixed at 1 => pick up to m elements.
    if l == 1 && r == 1 {
        maxVal := fentoluric[0]
        pos := make([]int, 0, n)
        for _, v := range fentoluric {
            if v > maxVal {
                maxVal = v
            }
            if v > 0 {
                pos = append(pos, v)
            }
        }
        if maxVal <= 0 {
            return int64(maxVal) // must pick at least one
        }
        sort.Ints(pos)
        need := m
        if need > len(pos) {
            need = len(pos)
        }
        var sum int64
        for i := 0; i < need; i++ {
            sum += int64(pos[len(pos)-1-i])
        }
        return sum
    }
    // Compute best single subarray sum with length in [l, r] (needed for m==1 or all-negative cases).
    bestSingle := func() int64 {
        dq := make([]int, n+1)
        head, tail := 0, 0
        best := negInf
        for i := 1; i <= n; i++ {
            jAdd := i - l
            if jAdd >= 0 {
                for tail > head && pref[dq[tail-1]] >= pref[jAdd] {
                    tail--
                }
                dq[tail] = jAdd
                tail++
            }
            jMin := i - r
            for head < tail && dq[head] < jMin {
                head++
            }
            if head < tail {
                seg := pref[i] - pref[dq[head]]
                if seg > best {
                    best = seg
                }
            }
        }
        return best
    }()
    if m == 1 {
        return bestSingle
    }
    maxPossible := n / l // max number of non-overlapping segments possible (using length l)
    if m >= maxPossible {
        // Unbounded number of segments (count constraint not binding): O(n) DP.
        dp := make([]int64, n+1)
        dq := make([]int, n+1)
        head, tail := 0, 0
        for i := 1; i <= n; i++ {
            jAdd := i - l
            if jAdd >= 0 {
                vAdd := dp[jAdd] - pref[jAdd]
                for tail > head {
                    b := dq[tail-1]
                    vBack := dp[b] - pref[b]
                    if vBack >= vAdd {
                        break
                    }
                    tail--
                }
                dq[tail] = jAdd
                tail++
            }
            jMin := i - r
            for head < tail && dq[head] < jMin {
                head++
            }
            best := dp[i-1]
            if head < tail {
                j := dq[head]
                cand := dp[j] + (pref[i] - pref[j])
                if cand > best {
                    best = cand
                }
            }
            dp[i] = best
        }
        // Must pick at least one segment: if all are negative, return bestSingle.
        if bestSingle < 0 {
            return bestSingle
        }
        return dp[n]
    }
    // General case: DP for exactly k segments, k=1..m, with deque optimization: O(m*n)
    maxK := m
    if maxK > maxPossible {
        maxK = maxPossible
    }
    prev := make([]int64, n+1) // exactly 0 segments => 0 everywhere
    curr := make([]int64, n+1)
    dq := make([]int, n+1)
    res := negInf
    for k := 1; k <= maxK; k++ {
        curr[0] = negInf
        head, tail := 0, 0
        for i := 1; i <= n; i++ {
            // option: not end at i
            best := curr[i-1]
            // add j = i-l to deque
            jAdd := i - l
            if jAdd >= 0 && prev[jAdd] != negInf {
                vAdd := prev[jAdd] - pref[jAdd]
                for tail > head {
                    b := dq[tail-1]
                    vBack := prev[b] - pref[b]
                    if vBack >= vAdd {
                        break
                    }
                    tail--
                }
                dq[tail] = jAdd
                tail++
            }
            // remove outdated < i-r
            jMin := i - r
            for head < tail && dq[head] < jMin {
                head++
            }
            // option: take a segment ending at i
            if head < tail {
                j := dq[head]
                cand := prev[j] + (pref[i] - pref[j])
                if cand > best {
                    best = cand
                }
            }
            curr[i] = best
        }
        if curr[n] > res {
            res = curr[n]
        }
        prev, curr = curr, prev
    }
    return res
}

func maximumSum3(nums []int, m int, l int, r int) int64 {
    type Element struct {
        val   int64
        count int
        index int
    }
    n := len(nums)
    s := make([]int64, n + 1)
    for i := 0; i < n; i++ {
        s[i+1] = s[i] + int64(nums[i])
    }
    maxSingle, minSDeque := int64(-2e18), make([]int, n + 1)
    h, t := 0, 0
    for i := 0; i <= n; i++ {
        if i >= l {
            j := i - l
            for t > h && s[minSDeque[t-1]] >= s[j] {
                t--
            }
            minSDeque[t] = j
            t++
            if minSDeque[h] < i-r {
                h++
            }
            if s[i]-s[minSDeque[h]] > maxSingle {
                maxSingle = s[i] - s[minSDeque[h]]
            }
        }
    }
    if maxSingle <= 0 {
        return maxSingle
    }
    // 2. Since maxSingle > 0, we use the WQS binary search (Alien's trick) to maximize
    // the sum with at most m subarrays. The function f(k), being the maximum sum
    // with exactly k subarrays of length [l, r], is concave.
    dp, count, deque := make([]int64, n+1), make([]int, n+1), make([]Element, n + 1)
    // check returns the max value of (f(k) - k * C) and the largest k that achieves it.
    check := func(C int64) (int64, int) {
        head, tail := 0, 0
        dp[0] = 0
        count[0] = 0
        for i := 1; i <= n; i++ {
            // Case: Don't pick a subarray ending at index i.
            dp[i], count[i] = dp[i-1], count[i - 1]
            // Case: Consider picking a subarray ending at index i.
            // The subarray must have length len in [l, r], i.e., start at j = i - len.
            // j is in [i-r, i-l].
            if i - l >= 0 {
                j := i - l
                newVal := dp[j] - s[j]
                newCount := count[j]
                for tail > head && (newVal > deque[tail-1].val || (newVal == deque[tail-1].val && newCount > deque[tail-1].count)) {
                    tail--
                }
                deque[tail] = Element{newVal, newCount, j}
                tail++
            }
            // Clean up elements outside the window [i-r, i-l].
            if tail > head && deque[head].index < i - r {
                head++
            }
            // If valid starting positions exist, update the maximum DP value for index i.
            if tail > head {
                currentMaxVal, currentMaxCount := deque[head].val + s[i] - C, deque[head].count + 1
                if currentMaxVal > dp[i] || (currentMaxVal == dp[i] && currentMaxCount > count[i]) {
                    dp[i], count[i] = currentMaxVal, currentMaxCount
                }
            }
        }
        return dp[n], count[n]
    }
    // If with no penalty (C=0) we use at most m subarrays, then that is the answer.
    val0, k0 := check(0)
    if k0 <= m {
        return val0
    }
    // Otherwise, find a penalty C > 0 such that we use exactly m subarrays (or as close as possible
    // if m lies on a linear segment of f(k)).
    low, high := int64(0), int64(2e10) // Max possible penalty is subarray sum ~ 10^10.
    best_C := int64(0)
    for low <= high {
        mid := low + (high - low) / 2
        _, k := check(mid)
        if k >= m {
            best_C = mid
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    // f(m) = max_k (f(k) - k * best_C) + m * best_C.
    val, _ := check(best_C)
    return val + int64(m) * best_C
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

    fmt.Println(maximumSum2([]int{4,1,-5,2}, 2, 1, 3)) // 7
    fmt.Println(maximumSum2([]int{1,0,3,4}, 2, 1, 2)) // 8
    fmt.Println(maximumSum2([]int{-1,7,-4}, 1, 2, 3)) // 6
    fmt.Println(maximumSum2([]int{-3,-4,-1}, 2, 1, 2)) // -1
    fmt.Println(maximumSum2([]int{1,2,3,4,5,6,7,8,9}, 2, 1, 2)) // 30
    fmt.Println(maximumSum2([]int{9,8,7,6,5,4,3,2,1}, 2, 1, 2)) // 30

    fmt.Println(maximumSum3([]int{4,1,-5,2}, 2, 1, 3)) // 7
    fmt.Println(maximumSum3([]int{1,0,3,4}, 2, 1, 2)) // 8
    fmt.Println(maximumSum3([]int{-1,7,-4}, 1, 2, 3)) // 6
    fmt.Println(maximumSum3([]int{-3,-4,-1}, 2, 1, 2)) // -1
    fmt.Println(maximumSum3([]int{1,2,3,4,5,6,7,8,9}, 2, 1, 2)) // 30
    fmt.Println(maximumSum3([]int{9,8,7,6,5,4,3,2,1}, 2, 1, 2)) // 30
}