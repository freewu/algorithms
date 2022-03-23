package main

import "fmt"

/**
57. Insert Interval
You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi]
represent the start and the end of the ith interval and intervals is sorted in ascending order by starti.
You are also given an interval newInterval = [start, end] that represents the start and end of another interval.
Insert newInterval into intervals such that intervals is still sorted in ascending order by starti
and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).
Return intervals after the insertion.

Constraints:

	0 <= intervals.length <= 10000
	intervals[i].length == 2
	0 <= starti <= endi <= 100000
	intervals is sorted by starti in ascending order.
	newInterval.length == 2
	0 <= start <= end <= 100000

Example 1:

	Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
	Output: [[1,5],[6,9]]

Example 2:

	Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
	Output: [[1,2],[3,10],[12,16]]
	Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].

 */

func insert(intervals [][]int, newInterval []int) [][]int {
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

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Printf("insert([][]int{[]int{1,3},[]int{6,9}},[]int{2,5}) = %v\n",insert([][]int{[]int{1,3},[]int{6,9}},[]int{2,5}) ) //  [[1,5],[6,9]]
	fmt.Printf("insert([][]int{[]int{1,2},[]int{3,10},[]int{6,7},[]int{8,10},[]int{12,16}},[]int{4,8}) = %v\n",insert([][]int{[]int{1,2},[]int{3,10},[]int{6,7},[]int{8,10},[]int{12,16}},[]int{4,8}) ) //   [[1,2],[3,10],[12,16]]
}
