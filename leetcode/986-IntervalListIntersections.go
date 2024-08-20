package main

// 986. Interval List Intersections
// You are given two lists of closed intervals, firstList and secondList, where firstList[i] = [starti, endi] and secondList[j] = [startj, endj]. 
// Each list of intervals is pairwise disjoint and in sorted order.

// Return the intersection of these two interval lists.

// A closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.

// The intersection of two closed intervals is a set of real numbers that are either empty or represented as a closed interval. 
// For example, the intersection of [1, 3] and [2, 4] is [2, 3].

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/01/30/interval1.png" />
// Input: firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
// Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]

// Example 2:
// Input: firstList = [[1,3],[5,9]], secondList = []
// Output: []

// Constraints:
//     0 <= firstList.length, secondList.length <= 1000
//     firstList.length + secondList.length >= 1
//     0 <= starti < endi <= 10^9
//     endi < starti+1
//     0 <= startj < endj <= 10^9
//     endj < startj+1

import "fmt"

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    res := [][]int{}
    if len(firstList) == 0 || len(secondList) == 0 {
        return res
    }
    i, j, m, n, mx, mn := 0, 0, len(firstList), len(secondList), -1 << 31, 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i < m && j < n {
        mx = max(firstList[i][0], secondList[j][0])
        mn = min(firstList[i][1], secondList[j][1])
        if mx <= mn {
            res = append(res, []int{ mx, mn })
        }
        if firstList[i][1] < secondList[j][1] {
            i++
        } else {
            j++
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/01/30/interval1.png" />
    // Input: firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
    // Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
    fmt.Println(intervalIntersection([][]int{{0,2},{5,10},{13,23},{24,25}}, [][]int{{1,5},{8,12},{15,24},{25,26}})) // [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
    // Example 2:
    // Input: firstList = [[1,3],[5,9]], secondList = []
    // Output: []
    fmt.Println(intervalIntersection([][]int{{1,3},{5,9}}, [][]int{})) // []
}