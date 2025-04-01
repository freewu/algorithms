package main

// 3505. Minimum Operations to Make Elements Within K Subarrays Equal
// You are given an integer array nums and two integers, x and k. 
// You can perform the following operation any number of times (including zero):
//     Increase or decrease any element of nums by 1.

// Return the minimum number of operations needed to have at least k non-overlapping subarrays of size exactly x in nums, where all elements within each subarray are equal.

// Example 1:
// Input: nums = [5,-2,1,3,7,3,6,4,-1], x = 3, k = 2
// Output: 8
// Explanation:
// Use 3 operations to add 3 to nums[1] and use 2 operations to subtract 2 from nums[3]. The resulting array is [5, 1, 1, 1, 7, 3, 6, 4, -1].
// Use 1 operation to add 1 to nums[5] and use 2 operations to subtract 2 from nums[6]. The resulting array is [5, 1, 1, 1, 7, 4, 4, 4, -1].
// Now, all elements within each subarray [1, 1, 1] (from indices 1 to 3) and [4, 4, 4] (from indices 5 to 7) are equal. Since 8 total operations were used, 8 is the output.

// Example 2:
// Input: nums = [9,-2,-2,-2,1,5], x = 2, k = 2
// Output: 3
// Explanation:
// Use 3 operations to subtract 3 from nums[4]. The resulting array is [9, -2, -2, -2, -2, 5].
// Now, all elements within each subarray [-2, -2] (from indices 1 to 2) and [-2, -2] (from indices 3 to 4) are equal. Since 3 operations were used, 3 is the output.

// Constraints:
//     2 <= nums.length <= 10^5
//     -10^6 <= nums[i] <= 10^6
//     2 <= x <= nums.length
//     1 <= k <= 15
//     2 <= k * x <= nums.length

import "fmt"
import "sort"
import "container/heap"

func minOperations(nums []int, x, k int) int64 {
    n := len(nums)
    dis := medianSlidingWindow(nums, x)
    f, g := make([]int, n + 1), make([]int, n + 1) // 滚动数组
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= k; i++ {
        g[i*x-1] = 1 << 61
        for j := i * x; j <= n-(k-i)*x; j++ {
            g[j] = min(g[j-1], f[j-x]+dis[j-x])
        }
        f, g = g, f
    }
    return int64(f[n])
}

// 480. 滑动窗口中位数（有改动）
// 返回 nums 的所有长为 k 的子数组的（到子数组中位数的）距离和
func medianSlidingWindow(nums []int, k int) []int {
    ans := make([]int, len(nums)-k+1)
    left := newLazyHeap()  // 最大堆（元素取反）
    right := newLazyHeap() // 最小堆
    for i, in := range nums {
        // 1. 进入窗口
        if left.size == right.size {
            left.push(-right.pushPop(in))
        } else {
            right.push(-left.pushPop(-in))
        }
        l := i + 1 - k
        if l < 0 { // 窗口大小不足 k
            continue
        }
        // 2. 计算答案
        v := -left.top()
        s1 := v*left.size + left.sum // sum 取反
        s2 := right.sum - v*right.size
        ans[l] = s1 + s2
        // 3. 离开窗口
        out := nums[l]
        if out <= -left.top() {
            left.remove(-out)
            if left.size < right.size {
                left.push(-right.pop()) // 平衡两个堆的大小
            }
        } else {
            right.remove(out)
            if left.size > right.size+1 {
                right.push(-left.pop()) // 平衡两个堆的大小
            }
        }
    }
    return ans
}

func newLazyHeap() *lazyHeap {
    return &lazyHeap{removeCnt: map[int]int{}}
}

// 懒删除堆
type lazyHeap struct {
    sort.IntSlice
    removeCnt map[int]int // 每个元素剩余需要删除的次数
    size      int         // 实际大小
    sum       int         // 堆中元素总和
}

// 必须实现的两个接口
func (h *lazyHeap) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *lazyHeap) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

// 删除
func (h *lazyHeap) remove(v int) {
    h.removeCnt[v]++ // 懒删除
    h.size--
    h.sum -= v
}

// 正式执行删除操作
func (h *lazyHeap) applyRemove() {
    for h.removeCnt[h.IntSlice[0]] > 0 {
        h.removeCnt[h.IntSlice[0]]--
        heap.Pop(h)
    }
}

// 查看堆顶
func (h *lazyHeap) top() int {
    h.applyRemove()
    return h.IntSlice[0]
}

// 出堆
func (h *lazyHeap) pop() int {
    h.applyRemove()
    h.size--
    h.sum -= h.IntSlice[0]
    return heap.Pop(h).(int)
}

// 入堆
func (h *lazyHeap) push(v int) {
    if h.removeCnt[v] > 0 {
        h.removeCnt[v]-- // 抵消之前的删除
    } else {
        heap.Push(h, v)
    }
    h.size++
    h.sum += v
}

// push(v) 然后 pop()
func (h *lazyHeap) pushPop(v int) int {
    if h.size > 0 && v > h.top() { // 最小堆，v 比堆顶大就替换堆顶
        h.sum += v - h.IntSlice[0]
        v, h.IntSlice[0] = h.IntSlice[0], v
        heap.Fix(h, 0)
    }
    return v
}

func minOperations1(nums []int, x int, k int) int64 {
    n := len(nums)
    var quickSort func(a []int, l, r int) 
    quickSort = func(a []int, l, r int) {
        if l >= r { return }
        pivot := a[(l+r)>>1]
        i, j := l, r
        for i <= j {
            for a[i] < pivot {
                i++
            }
            for a[j] > pivot {
                j--
            }
            if i <= j {
                a[i], a[j] = a[j], a[i]
                i++
                j--
            }
        }
        if l < j {
            quickSort(a, l, j)
        }
        if i < r {
            quickSort(a, i, r)
        }
    }
    // Step 1: Coordinate Compression
    allVals := make([]int, n)
    copy(allVals, nums)
    quickSort(allVals, 0, n-1)
    uniqueVals := allVals[:1]
    for i := 1; i < n; i++ {
        if allVals[i] != allVals[i-1] {
            uniqueVals = append(uniqueVals, allVals[i])
        }
    }
    compressMap := make(map[int]int, len(uniqueVals))
    for i, v := range uniqueVals {
        compressMap[v] = i + 1
    }
    sizeFenwicks := len(uniqueVals) + 1
    freqFenwicks := make([]int64, sizeFenwicks)
    sumFenwicks := make([]int64, sizeFenwicks)

    updateFreq := func(pos int, delta int64) {
        for pos < sizeFenwicks {
            freqFenwicks[pos] += delta
            pos += pos & -pos
        }
    }
    updateSum := func(pos int, delta int64) {
        for pos < sizeFenwicks {
            sumFenwicks[pos] += delta
            pos += pos & -pos
        }
    }
    getFreq := func(pos int) int64 {
        var s int64
        for pos > 0 {
            s += freqFenwicks[pos]
            pos -= pos & -pos
        }
        return s
    }
    getSum := func(pos int) int64 {
        var s int64
        for pos > 0 {
            s += sumFenwicks[pos]
            pos -= pos & -pos
        }
        return s
    }
    insertVal := func(v int) {
        cpos := compressMap[v]
        updateFreq(cpos, 1)
        updateSum(cpos, int64(v))
    }
    removeVal := func(v int) {
        cpos := compressMap[v]
        updateFreq(cpos, -1)
        updateSum(cpos, -int64(v))
    }
    findMedianPos := func(windowSize int) int {
        need := (windowSize + 1) >> 1
        left, right := 1, sizeFenwicks-1
        for left < right {
            mid := (left + right) >> 1
            if getFreq(mid) >= int64(need) {
                right = mid
            } else {
                left = mid + 1
            }
        }
        return left
    }
    getCost := func() int64 {
        countTotal := getFreq(sizeFenwicks - 1)
        if countTotal == 0 {
            return 0
        }
        medianPos := findMedianPos(int(countTotal))
        medianVal := uniqueVals[medianPos-1]
        leftFreq := getFreq(medianPos)
        leftSum := getSum(medianPos)
        allSum := getSum(sizeFenwicks - 1)
        allFreq := getFreq(sizeFenwicks - 1)
        rightFreq := allFreq - leftFreq
        rightSum := allSum - leftSum
        costLeft := int64(medianVal)*leftFreq - leftSum
        costRight := rightSum - int64(medianVal)*rightFreq
        return costLeft + costRight
    }
    costArr := make([]int64, n-x+1)
    leftPtr := 0
    for rightPtr := 0; rightPtr < n; rightPtr++ {
        insertVal(nums[rightPtr])
        if rightPtr-leftPtr+1 > x {
            removeVal(nums[leftPtr])
            leftPtr++
        }
        if rightPtr-leftPtr+1 == x {
            costArr[leftPtr] = getCost()
        }
    }
    INF := int64(1) << 60
    dp := make([][]int64, k+1)
    for r := 0; r <= k; r++ {
        dp[r] = make([]int64, n+1)
        for i := 0; i <= n; i++ {
            dp[r][i] = INF
        }
    }
    dp[0][0] = 0
    for i := 1; i <= n; i++ {
        dp[0][i] = 0
    }
    for r := 1; r <= k; r++ {
        for i := 1; i <= n; i++ {
            dp[r][i] = dp[r][i-1]
            if i >= x {
                cand := dp[r-1][i-x] + costArr[i-x]
                if cand < dp[r][i] {
                    dp[r][i] = cand
                }
            }
        }
    }
    return dp[k][n]
}

func main() {
    // Example 1:
    // Input: nums = [5,-2,1,3,7,3,6,4,-1], x = 3, k = 2
    // Output: 8
    // Explanation:
    // Use 3 operations to add 3 to nums[1] and use 2 operations to subtract 2 from nums[3]. The resulting array is [5, 1, 1, 1, 7, 3, 6, 4, -1].
    // Use 1 operation to add 1 to nums[5] and use 2 operations to subtract 2 from nums[6]. The resulting array is [5, 1, 1, 1, 7, 4, 4, 4, -1].
    // Now, all elements within each subarray [1, 1, 1] (from indices 1 to 3) and [4, 4, 4] (from indices 5 to 7) are equal. Since 8 total operations were used, 8 is the output.
    fmt.Println(minOperations([]int{5,-2,1,3,7,3,6,4,-1}, 3, 2)) // 8
    // Example 2:
    // Input: nums = [9,-2,-2,-2,1,5], x = 2, k = 2
    // Output: 3
    // Explanation:
    // Use 3 operations to subtract 3 from nums[4]. The resulting array is [9, -2, -2, -2, -2, 5].
    // Now, all elements within each subarray [-2, -2] (from indices 1 to 2) and [-2, -2] (from indices 3 to 4) are equal. Since 3 operations were used, 3 is the output.
    fmt.Println(minOperations([]int{9,-2,-2,-2,1,5}, 2, 2)) // 3

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9}, 2, 2)) // 2
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1}, 2, 2)) // 2

    fmt.Println(minOperations1([]int{5,-2,1,3,7,3,6,4,-1}, 3, 2)) // 8
    fmt.Println(minOperations1([]int{9,-2,-2,-2,1,5}, 2, 2)) // 3
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9}, 2, 2)) // 2
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1}, 2, 2)) // 2
}