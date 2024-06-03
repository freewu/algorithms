package main

// 1272. Remove Interval
// A set of real numbers can be represented as the union of several disjoint intervals,
// where each interval is in the form [a, b). A real number x is in the set if one of its intervals [a, b) contains x (i.e. a <= x < b).

// You are given a sorted list of disjoint intervals intervals representing a set of real numbers as described above, 
// where intervals[i] = [ai, bi] represents the interval [ai, bi). You are also given another interval toBeRemoved.

// Return the set of real numbers with the interval toBeRemoved removed from intervals. 
// In other words, return the set of real numbers such that every x in the set is in intervals but not in toBeRemoved. 
// Your answer should be a sorted list of disjoint intervals as described above.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/24/removeintervalex1.png" />
// Input: intervals = [[0,2],[3,4],[5,7]], toBeRemoved = [1,6]
// Output: [[0,1],[6,7]]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/24/removeintervalex2.png" />
// Input: intervals = [[0,5]], toBeRemoved = [2,3]
// Output: [[0,2],[3,5]]

// Example 3:
// Input: intervals = [[-5,-4],[-3,-2],[1,2],[3,5],[8,9]], toBeRemoved = [-1,4]
// Output: [[-5,-4],[-3,-2],[4,5],[8,9]]
 
// Constraints:
//     1 <= intervals.length <= 10^4
//     -10^9 <= ai < bi <= 10^9

import "fmt"
import "sort"

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
    n, res := len(intervals), [][]int{}
    id := sort.Search(n, func(i int) bool {
        return toBeRemoved[0] <= intervals[i][1]
    })
    res = append(res, intervals[:id]...)
    if id < n && toBeRemoved[0] > intervals[id][0] {
        res = append(res, []int{intervals[id][0], toBeRemoved[0]})
    }
    id = sort.Search(n, func(i int) bool {
        return toBeRemoved[1] < intervals[i][1]
    })
    if id < n {
        res = append(res, []int{max(intervals[id][0], toBeRemoved[1]), intervals[id][1]})
    }
    if id+1 < n {
        res = append(res, intervals[id+1:]...)
    }
    return res
}

func removeInterval1(intervals [][]int, toBeRemoved []int) [][]int {
    if len(toBeRemoved) == 0 {
        return intervals
    }
    if len(intervals) == 1 && toBeRemoved[0] == intervals[0][0] && toBeRemoved[1] == intervals[0][1] {
        return nil
    }
    mn, mx, res:= toBeRemoved[0], toBeRemoved[1], make([][]int, 0, len(intervals))
    for _, v := range intervals {
        if v[1] <= mn {
            res = append(res, v)
        } else if v[1] >= mn && v[0] <= mn && v[1] <= mx {
            res = append(res, []int{v[0], mn})
        } else if v[0] <= mn && v[1] > mx {
            if v[0] != mn {
                res = append(res, []int{v[0], mn})
            }
            res = append(res, []int{mx, v[1]})
        } else if v[0] >= mn && v[0] < mx && v[1] >= mx {
            res = append(res, []int{mx, v[1]})
        } else if v[0] >= mx {
            res = append(res, v)
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/12/24/removeintervalex1.png" />
    // Input: intervals = [[0,2],[3,4],[5,7]], toBeRemoved = [1,6]
    // Output: [[0,1],[6,7]]
    fmt.Println(removeInterval([][]int{{0,2},{3,4},{5,7}},[]int{1,6})) // [[0,1],[6,7]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/12/24/removeintervalex2.png" />
    // Input: intervals = [[0,5]], toBeRemoved = [2,3]
    // Output: [[0,2],[3,5]]
    fmt.Println(removeInterval([][]int{{0,5}},[]int{2,3})) // [[0,2],[3,5]]
    // Example 3:
    // Input: intervals = [[-5,-4],[-3,-2],[1,2],[3,5],[8,9]], toBeRemoved = [-1,4]
    // Output: [[-5,-4],[-3,-2],[4,5],[8,9]]
    fmt.Println(removeInterval([][]int{{-5,-4},{-3,-1},{1,2},{3,5},{8,9}},[]int{-1,4})) //  [[-5,-4],[-3,-2],[4,5],[8,9]]

    fmt.Println(removeInterval1([][]int{{0,2},{3,4},{5,7}},[]int{1,6})) // [[0,1],[6,7]]
    fmt.Println(removeInterval1([][]int{{0,5}},[]int{2,3})) // [[0,2],[3,5]]
    fmt.Println(removeInterval1([][]int{{-5,-4},{-3,-1},{1,2},{3,5},{8,9}},[]int{-1,4})) //  [[-5,-4],[-3,-2],[4,5],[8,9]]
}