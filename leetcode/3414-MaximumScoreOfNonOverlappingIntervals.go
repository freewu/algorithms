package main

// 3414. Maximum Score of Non-overlapping Intervals
// You are given a 2D integer array intervals, where intervals[i] = [li, ri, weighti]. 
// Interval i starts at position li and ends at ri, and has a weight of weighti. 
// You can choose up to 4 non-overlapping intervals. 
// The score of the chosen intervals is defined as the total sum of their weights.

// Return the lexicographically smallest array of at most 4 indices from intervals with maximum score, representing your choice of non-overlapping intervals.

// Two intervals are said to be non-overlapping if they do not share any points. 
// In particular, intervals sharing a left or right boundary are considered overlapping.

// Example 1:
// Input: intervals = [[1,3,2],[4,5,2],[1,5,5],[6,9,3],[6,7,1],[8,9,1]]
// Output: [2,3]
// Explanation:
// You can choose the intervals with indices 2, and 3 with respective weights of 5, and 3.

// Example 2:
// Input: intervals = [[5,8,1],[6,7,7],[4,7,3],[9,10,6],[7,8,2],[11,14,3],[3,5,5]]
// Output: [1,3,5,6]
// Explanation:
// You can choose the intervals with indices 1, 3, 5, and 6 with respective weights of 7, 6, 3, and 5.

// Constraints:
//     1 <= intevals.length <= 5 * 10^4
//     intervals[i].length == 3
//     intervals[i] = [li, ri, weighti]
//     1 <= li <= ri <= 10^9
//     1 <= weighti <= 10^9

import "fmt"
import "sort"
import "slices"

func maximumWeight(intervals [][]int) []int {
    type Tuple struct{ left, right, weight, index int }
    n := len(intervals)
    arr := make([]Tuple, n)
    for i, v := range intervals {
        arr[i] = Tuple{ v[0], v[1], v[2], i }
    }
    slices.SortFunc(arr, func(a, b Tuple) int { 
        return a.right - b.right 
    })
    type Pair struct {
        sum int
        id  []int
    }
    dp := make([][5]Pair, n + 1)
    for i, t := range arr {
        k := sort.Search(i, func(k int) bool { 
            return arr[k].right >= t.left 
        })
        for j := 1; j < 5; j++ {
            s1 := dp[i][j].sum
            // 为什么是 dp[k] 不是 dp[k+1]：上面算的是 >= t.left，-1 后得到 < t.left，但由于还要 +1，抵消了
            s2 := dp[k][j - 1].sum + t.weight
            if s1 > s2 {
                dp[i + 1][j] = dp[i][j]
                continue
            }
            newId := slices.Clone(dp[k][j-1].id)
            newId = append(newId, t.index)
            slices.Sort(newId)
            if s1 == s2 && slices.Compare(dp[i][j].id, newId) < 0 {
                newId = dp[i][j].id
            }
            dp[i + 1][j] = Pair{s2, newId}
        }
    }
    return dp[n][4].id
}

func maximumWeight1(intervals [][]int) []int {
    type interval struct {
        l, r, w int64
        idx     int
    }
    n := len(intervals)
    if n == 0 {
        return []int{}
    }
    // Store intervals with original index and sort by end
    iv := make([]interval, n)
    for i, v := range intervals {
        iv[i] = interval{int64(v[0]), int64(v[1]), int64(v[2]), i}
    }
    // Sort by r ascending, tie-break by l ascending (not strictly necessary for correctness, but stable)
    // ensures consistent sorting for building p array
    // intervals that end earlier come first
    sort.Slice(iv, func(a, b int) bool {
        if iv[a].r < iv[b].r {
            return true
        } else if iv[a].r > iv[b].r {
            return false
        }
        return iv[a].l < iv[b].l
    })
    // Build array of ends for binary search
    ends := make([]int64, n)
    for i := 0; i < n; i++ {
        ends[i] = iv[i].r
    }
    // p[i] = largest j < i with ends[j] < iv[i].l (strictly less)
    p := make([]int, n)
    for i := 0; i < n; i++ {
        l := iv[i].l
        left, right, val := 0, i - 1, -1
        for left <= right {
            mid := (left + right) >> 1
            if ends[mid] < l {
                val, left = mid,  mid + 1
            } else {
                right = mid - 1
            }
        }
        p[i] = val
    }
    // We'll store DP in 1D flattened array: dp[i][k] -> dp[i*5 + k]
    // dpVal stores sum + up to 4 indices in ascending order
    type dpVal struct {
        sum    int64
        arr    [4]int
        length int
    }
    // Compare function: a is "better" if:
    // 1) a.sum > b.sum, or
    // 2) a.sum == b.sum but a.arr is lex-smaller than b.arr
    isBetter := func(a, b dpVal) bool {
        if a.sum > b.sum {
            return true
        }
        if a.sum < b.sum {
            return false
        }
        // sums tie -> compare lex
        la, lb := a.length, b.length
        for i := 0; i < la && i < lb; i++ {
            if a.arr[i] < b.arr[i] {
                return true
            } else if a.arr[i] > b.arr[i] {
                return false
            }
        }
        return la < lb
    }
    // Combine dpVal with an additional interval index (insert in ascending order)
    // iArg is the index in 'iv' whose original idx we want to insert.
    combine := func(base dpVal, iArg int) dpVal {
        var c dpVal
        c.sum = base.sum + iv[iArg].w
        c.length = base.length + 1
        // merge iv[iArg].idx into base.arr in sorted order
        x := iv[iArg].idx
        bLen := base.length
        pos := 0
        for pos < bLen && base.arr[pos] < x {
            c.arr[pos] = base.arr[pos]
            pos++
        }
        c.arr[pos] = x
        for j := pos; j < bLen; j++ {
            c.arr[j+1] = base.arr[j]
        }
        return c
    }
    // dp array, dimensions (n+1) x (5)
    dp := make([]dpVal, (n+1)*5)
    // Initialize dp[0][k] and dp[i][0]
    for k := 0; k < 5; k++ {
        dp[k] = dpVal{0, [4]int{}, 0}
    }
    for i := 0; i <= n; i++ {
        dp[i*5] = dpVal{0, [4]int{}, 0}
    }
    idx := func(i, k int) int {
        return i*5 + k
    }
    // Fill dp
    for i := 1; i <= n; i++ {
        for k := 1; k <= 4; k++ {
            // Not take current interval
            best := dp[idx(i-1, k)]
            // Take current interval if possible
            prev := p[i-1]
            var cand dpVal
            if k-1 >= 0 {
                if prev >= 0 {
                    cand = combine(dp[idx(prev+1, k-1)], i-1)
                } else {
                    // means no previous interval to worry about
                    cand = combine(dp[idx(0, k-1)], i-1)
                }
                // pick better between best and cand
                if isBetter(cand, best) {
                    best = cand
                }
            }
            dp[idx(i, k)] = best
        }
    }
    // Find best among dp[n][1], dp[n][2], dp[n][3], dp[n][4]
    tmp := dp[idx(n, 1)]
    for k := 2; k <= 4; k++ {
        if isBetter(dp[idx(n, k)], tmp) {
            tmp = dp[idx(n, k)]
        }
    }
    // Build answer from tmp.arr
    res := make([]int, tmp.length)
    for i := 0; i < tmp.length; i++ {
        res[i] = tmp.arr[i]
    }
    return res
}

func main() {
    // Example 1:
    // Input: intervals = [[1,3,2],[4,5,2],[1,5,5],[6,9,3],[6,7,1],[8,9,1]]
    // Output: [2,3]
    // Explanation:
    // You can choose the intervals with indices 2, and 3 with respective weights of 5, and 3.
    fmt.Println(maximumWeight([][]int{{1,3,2},{4,5,2},{1,5,5},{6,9,3},{6,7,1},{8,9,1}})) // [2,3]
    // Example 2:
    // Input: intervals = [[5,8,1],[6,7,7],[4,7,3],[9,10,6],[7,8,2],[11,14,3],[3,5,5]]
    // Output: [1,3,5,6]
    // Explanation:
    // You can choose the intervals with indices 1, 3, 5, and 6 with respective weights of 7, 6, 3, and 5.
    fmt.Println(maximumWeight([][]int{{5,8,1},{6,7,7},{4,7,3},{9,10,6},{7,8,2},{11,14,3},{3,5,5}})) // [1,3,5,6]

    fmt.Println(maximumWeight([][]int{{1,3,5}, {2,4,6},{3,5,7},{4,6,8}})) // [0,3]
    fmt.Println(maximumWeight([][]int{{1,1,1000000000},{1,1,1000000000},{1,1,1000000000},{1,1,1000000000}})) // [0]

    fmt.Println(maximumWeight1([][]int{{1,3,2},{4,5,2},{1,5,5},{6,9,3},{6,7,1},{8,9,1}})) // [2,3]
    fmt.Println(maximumWeight1([][]int{{5,8,1},{6,7,7},{4,7,3},{9,10,6},{7,8,2},{11,14,3},{3,5,5}})) // [1,3,5,6]
    fmt.Println(maximumWeight1([][]int{{1,3,5}, {2,4,6},{3,5,7},{4,6,8}})) // [0,3]
    fmt.Println(maximumWeight1([][]int{{1,1,1000000000},{1,1,1000000000},{1,1,1000000000},{1,1,1000000000}})) // [0]
}