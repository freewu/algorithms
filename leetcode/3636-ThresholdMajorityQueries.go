package main

// 3636. Threshold Majority Queries
// You are given an integer array nums of length n and an array queries, where queries[i] = [li, ri, thresholdi].

// Return an array of integers ans where ans[i] is equal to the element in the subarray nums[li...ri] that appears at least thresholdi times, selecting the element with the highest frequency (choosing the smallest in case of a tie), or -1 if no such element exists.

// Example 1:
// Input: nums = [1,1,2,2,1,1], queries = [[0,5,4],[0,3,3],[2,3,2]]
// Output: [1,-1,2]
// Explanation:
// Query	Sub-array	Threshold	Frequency table	Answer
// [0, 5, 4]	[1, 1, 2, 2, 1, 1]	4	1 → 4, 2 → 2	1
// [0, 3, 3]	[1, 1, 2, 2]	3	1 → 2, 2 → 2	-1
// [2, 3, 2]	[2, 2]	2	2 → 2	2

// Example 2:
// Input: nums = [3,2,3,2,3,2,3], queries = [[0,6,4],[1,5,2],[2,4,1],[3,3,1]]
// Output: [3,2,3,2]
// Explanation:
// Query	Sub-array	Threshold	Frequency table	Answer
// [0, 6, 4]	[3, 2, 3, 2, 3, 2, 3]	4	3 → 4, 2 → 3	3
// [1, 5, 2]	[2, 3, 2, 3, 2]	2	2 → 3, 3 → 2	2
// [2, 4, 1]	[3, 2, 3]	1	3 → 2, 2 → 1	3
// [3, 3, 1]	[2]	1	2 → 1	2
 
// Constraints:
//     1 <= nums.length == n <= 10^4
//     1 <= nums[i] <= 10^9
//     1 <= queries.length <= 5 * 10^4
//     queries[i] = [li, ri, thresholdi]
//     0 <= li <= ri < n
//     1 <= thresholdi <= ri - li + 1

import "fmt"
import "slices"
import "cmp"
import "sort"
import "math"

func subarrayMajority(nums []int, queries [][]int) []int {
    n, m := len(nums), len(queries)
    a := slices.Clone(nums)
    slices.Sort(a)
    a = slices.Compact(a)
    indexToValue := make([]int, n)
    for i, x := range nums {
        indexToValue[i] = sort.SearchInts(a, x)
    }
    cnt := make([]int, len(a)+1)
    maxCnt, minVal := 0, 0
    add := func(i int) {
        v := indexToValue[i]
        cnt[v]++
        c := cnt[v]
        x := nums[i]
        if c > maxCnt {
            maxCnt, minVal = c, x
        } else if c == maxCnt {
            minVal = min(minVal, x)
        }
    }
    res := make([]int, m)
    blockSize := int(math.Ceil(float64(n) / math.Sqrt(float64(m))))
    type query struct{ bid, l, r, threshold, qid int } // [l,r) 左闭右开
    qs := []query{}
    for i, q := range queries {
        l, r, threshold := q[0], q[1]+1, q[2] // 左闭右开
        // 大区间离线（保证 l 和 r 不在同一个块中）
        if r-l > blockSize {
            qs = append(qs, query{l / blockSize, l, r, threshold, i})
            continue
        }
        // 小区间暴力
        for j := l; j < r; j++ {
            add(j)
        }
        if maxCnt >= threshold {
            res[i] = minVal
        } else {
            res[i] = -1
        }
        // 重置数据
        for _, v := range indexToValue[l:r] {
            cnt[v]--
        }
        maxCnt = 0
    }
    slices.SortFunc(qs, func(a, b query) int { return cmp.Or(a.bid-b.bid, a.r-b.r) })
    r := 0
    for i, q := range qs {
        l0 := (q.bid + 1) * blockSize
        if i == 0 || q.bid > qs[i-1].bid { // 遍历到一个新的块
            r = l0 // 右端点移动的起点
            // 重置数据
            clear(cnt)
            maxCnt = 0
        }
        // 右端点从 r 移动到 q.r（q.r 不计入）
        for ; r < q.r; r++ {
            add(r)
        }
        tmpMaxCnt, tmpMinVal := maxCnt, minVal
        // 左端点从 l0 移动到 q.l（l0 不计入）
        for l := q.l; l < l0; l++ {
            add(l)
        }
        if maxCnt >= q.threshold {
            res[q.qid] = minVal
        } else {
            res[q.qid] = -1
        }
        // 回滚
        maxCnt, minVal = tmpMaxCnt, tmpMinVal
        for _, v := range indexToValue[q.l:l0] {
            cnt[v]--
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,1,2,2,1,1], queries = [[0,5,4],[0,3,3],[2,3,2]]
    // Output: [1,-1,2]
    // Explanation:
    // Query	Sub-array	Threshold	Frequency table	Answer
    // [0, 5, 4]	[1, 1, 2, 2, 1, 1]	4	1 → 4, 2 → 2	1
    // [0, 3, 3]	[1, 1, 2, 2]	3	1 → 2, 2 → 2	-1
    // [2, 3, 2]	[2, 2]	2	2 → 2	2
    fmt.Println(subarrayMajority([]int{1,1,2,2,1,1}, [][]int{{0,5,4},{0,3,3},{2,3,2}})) // [1,-1,2]
    // Example 2:
    // Input: nums = [3,2,3,2,3,2,3], queries = [[0,6,4],[1,5,2],[2,4,1],[3,3,1]]
    // Output: [3,2,3,2]
    // Explanation:
    // Query	Sub-array	Threshold	Frequency table	Answer
    // [0, 6, 4]	[3, 2, 3, 2, 3, 2, 3]	4	3 → 4, 2 → 3	3
    // [1, 5, 2]	[2, 3, 2, 3, 2]	2	2 → 3, 3 → 2	2
    // [2, 4, 1]	[3, 2, 3]	1	3 → 2, 2 → 1	3
    // [3, 3, 1]	[2]	1	2 → 1	2
    fmt.Println(subarrayMajority([]int{3,2,3,2,3,2,3}, [][]int{{0,6,4},{1,5,2},{2,4,1},{3,3,1}})) // [3,2,3,2]
}