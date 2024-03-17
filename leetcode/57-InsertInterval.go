package main

// 57. Insert Interval
// You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi]
// represent the start and the end of the ith interval and intervals is sorted in ascending order by starti.
// You are also given an interval newInterval = [start, end] that represents the start and end of another interval.
// Insert newInterval into intervals such that intervals is still sorted in ascending order by starti
// and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).
// Return intervals after the insertion.

// Constraints:
//     0 <= intervals.length <= 10000
//     intervals[i].length == 2
//     0 <= starti <= endi <= 100000
//     intervals is sorted by starti in ascending order.
//     newInterval.length == 2
//     0 <= start <= end <= 100000

// Example 1:
//     Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
//     Output: [[1,5],[6,9]]

// Example 2:
//     Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//     Output: [[1,2],[3,10],[12,16]]
//     Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
    max := func (a int, b int) int { if a > b { return a; }; return b; }
    min := func (a int, b int) int { if a > b { return b; }; return a; }
    res := make([][]int, 0)
    if len(intervals) == 0 {
        res = append(res, newInterval)
        return res
    }
    curIndex := 0
    // 先找到要插入的位置 一直到 intervals[curIndex] 的结束区间 > newInterval开始区间
    for curIndex < len(intervals) && intervals[curIndex][1] < newInterval[0] {
        res = append(res, intervals[curIndex]) // 都可加入到结果集中
        curIndex++
    }
    //  intervals[curIndex][0] <= newInterval[1]
    for curIndex < len(intervals) && intervals[curIndex][0] <= newInterval[1] {
        newInterval = []int{ min(newInterval[0], intervals[curIndex][0]),  max(newInterval[1], intervals[curIndex][1])}
        curIndex++
    }
    res = append(res, newInterval)
    // 把剩余的区间加入到结果数组
    for curIndex < len(intervals) {
        res = append(res, intervals[curIndex])
        curIndex++
    }
    return res
}

func insert1(intervals [][]int, newInterval []int) [][]int {
    max := func (a int, b int) int { if a > b { return a; }; return b; }
    min := func (a int, b int) int { if a > b { return b; }; return a; }

    var res [][]int
    left, right := newInterval[0], newInterval[1]
    merged := false
    for _, interval := range intervals {
        if interval[0] > right {
            // 在插入区间的右侧且无交集
            if !merged {
                res = append(res, []int{left, right})
                merged = true
            }
            res = append(res, interval)
        } else if interval[1] < left {
            // 在插入区间的左侧且无交集
            res = append(res, interval)
        } else {
            // 与插入区间有交集，计算它们的并集
            left = min(left, interval[0])
            right = max(right, interval[1])
        }
    }
    // 没有合并处理 说明需要加在 res 里面
    if !merged {
        res = append(res, []int{left, right})
    }
    return res
}

// best solution
// time: O(n) space: O(n)
func insert2(intervals [][]int, newInterval []int) [][]int {
    //isOverlap := func (a, b []int) bool { return a[0] <= b[1] && b[0] <= a[1]; }
    merge := func (a, b []int) []int { 
        max := func (a int, b int) int { if a > b { return a; }; return b; }
        min := func (a int, b int) int { if a > b { return b; }; return a; }
        return []int{ min(a[0], b[0]), max(a[1], b[1]) } 
    }
    res := make([][]int, 0, len(intervals))
    for i := 0; i < len(intervals); i++ {
        itv := intervals[i]
        // newInterval is smaller than all the other intervals right side
        if newInterval[1] < itv[0] {
            res = append(res, newInterval)
            return append(res, intervals[i:]...)
        }
        if itv[1] < newInterval[0] {
            res = append(res, itv)
        } else {
            newInterval = merge(newInterval, itv) // overlap
        }
    }
    return append(res, newInterval)
}

func main() {
    fmt.Printf("insert([][]int{[]int{1,3},[]int{6,9}},[]int{2,5}) = %v\n",insert([][]int{[]int{1,3},[]int{6,9}},[]int{2,5}) ) //  [[1,5],[6,9]]
    fmt.Printf("insert([][]int{[]int{1,2},[]int{3,10},[]int{6,7},[]int{8,10},[]int{12,16}},[]int{4,8}) = %v\n",insert([][]int{[]int{1,2},[]int{3,10},[]int{6,7},[]int{8,10},[]int{12,16}},[]int{4,8}) ) //   [[1,2],[3,10],[12,16]]
    fmt.Println(insert1([][]int{[]int{1,3},[]int{6,9}},[]int{2,5}) ) // [[1,5],[6,9]]
    fmt.Println(insert1([][]int{[]int{1,2},[]int{3,10},[]int{6,7},[]int{8,10},[]int{12,16}},[]int{4,8}) ) // [[1,2],[3,10],[12,16]]
    fmt.Println(insert2([][]int{[]int{1,3},[]int{6,9}},[]int{2,5}) ) // [[1,5],[6,9]]
    fmt.Println(insert2([][]int{[]int{1,2},[]int{3,10},[]int{6,7},[]int{8,10},[]int{12,16}},[]int{4,8}) ) // [[1,2],[3,10],[12,16]]
}
