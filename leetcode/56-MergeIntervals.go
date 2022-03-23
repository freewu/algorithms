package main

import (
	"fmt"
	"sort"
)

/**
56. Merge Intervals
Given an array of intervals where intervals[i] = [starti, endi],
merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

Constraints:
	1 <= intervals.length <= 10000
	intervals[i].length == 2
	0 <= starti <= endi <= 10000

Example 1:

	Input: [[1,3],[2,6],[8,10],[15,18]]
	Output: [[1,6],[8,10],[15,18]]
	Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

Example 2:

	Input: [[1,4],[4,5]]
	Output: [[1,5]]
	Explanation: Intervals [1,4] and [4,5] are considered overlapping.

解题思路:
	先按照区间起点进行排序。然后从区间起点小的开始扫描，依次合并每个有重叠的区间。
 */

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	quickSort(intervals, 0, len(intervals)-1) // 先把所有的区间按从小到大排好序
	res := make([][]int, 0)
	res = append(res, intervals[0]) // 把第一个区间先放入返回数组中
	curIndex := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[curIndex][1] { // 发现 后一个区间开始 比 当前区间结尾 要大 说明没有交集 不需要合并
			curIndex++
			res = append(res, intervals[i]) // 把后一个区间加入返回结果数组中
		} else {
			// 这里是做合并操作的
			// 如果下个区间结尾 和 当前结尾 哪个更大使用哪个
			res[curIndex][1] = max(intervals[i][1], res[curIndex][1])
		}
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

func partitionSort(a [][]int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if (a[j][0] < pivot[0]) || (a[j][0] == pivot[0] && a[j][1] < pivot[1]) {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
func quickSort(a [][]int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionSort(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

type sortSlice [][]int
func (l sortSlice) Less(i,j int) bool {
	return l[i][0] < l[j][0]
}
func (l sortSlice) Len() int {
	return len(l)
}
func (l sortSlice) Swap(i,j int)  {
	l[i],l[j] = l[j],l[i]
}

func mergeBest(intervals [][]int) [][]int {
	length := len(intervals)
	var result [][]int
	sort.Sort(sortSlice(intervals))
	tLeft, tRight := intervals[0][0], intervals[0][1]
	for i:=1;i<length;i++ {
		if tRight >= intervals[i][0] {
			if tLeft > intervals[i][0] {
				tLeft = intervals[i][0]
			}
			if tRight < intervals[i][1] {
				tRight = intervals[i][1]
			}
		} else {
			result = append(result, []int{tLeft,tRight})
			tLeft = intervals[i][0]
			tRight = intervals[i][1]
		}
	}
	result = append(result, []int{tLeft,tRight})
	return result
}

func main() {
	fmt.Printf("merge([][]int{[]int{1,3},[]int{2,6},[]int{8,10},[]int{15,18}}) = %v\n",merge([][]int{[]int{1,3},[]int{2,6},[]int{8,10},[]int{15,18}})) // [[1,6],[8,10],[15,18]]
	fmt.Printf("merge([][]int{[]int{1,4},[]int{4,5}}) = %v\n",merge([][]int{[]int{1,4},[]int{4,5}})) // [[1,5]]

	fmt.Printf("mergeBest([][]int{[]int{1,3},[]int{2,6},[]int{8,10},[]int{15,18}}) = %v\n",mergeBest([][]int{[]int{1,3},[]int{2,6},[]int{8,10},[]int{15,18}})) // [[1,6],[8,10],[15,18]]
	fmt.Printf("mergeBest([][]int{[]int{1,4},[]int{4,5}}) = %v\n",mergeBest([][]int{[]int{1,4},[]int{4,5}})) // [[1,5]]
}
