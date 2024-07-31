package main

// 757. Set Intersection Size At Least Two
// You are given a 2D integer array intervals where intervals[i] = [starti, endi] represents all the integers from starti to endi inclusively.

// A containing set is an array nums where each interval from intervals has at least two integers in nums.
//     For example, if intervals = [[1,3], [3,7], [8,9]], then [1,2,4,7,8,9] and [2,3,4,8,9] are containing sets.

// Return the minimum possible size of a containing set.

// Example 1:
// Input: intervals = [[1,3],[3,7],[8,9]]
// Output: 5
// Explanation: let nums = [2, 3, 4, 8, 9].
// It can be shown that there cannot be any containing array of size 4.

// Example 2:
// Input: intervals = [[1,3],[1,4],[2,5],[3,5]]
// Output: 3
// Explanation: let nums = [2, 3, 4].
// It can be shown that there cannot be any containing array of size 2.

// Example 3:
// Input: intervals = [[1,2],[2,3],[2,4],[4,5]]
// Output: 5
// Explanation: let nums = [1, 2, 3, 4, 5].
// It can be shown that there cannot be any containing array of size 4.
 
// Constraints:
//     1 <= intervals.length <= 3000
//     intervals[i].length == 2
//     0 <= starti < endi <= 10^8

import "fmt"
import "sort"

func intersectionSizeTwo(intervals [][]int) int {
    res, end1, end2, inf := 0, -1, -1, 1_000_000_000
    sort.Slice(intervals, func(i, j int) bool {
        // 先按集合再按集合右边缘升序排序, 在按集合左边缘降序，，这样保证i区间和i+1区间的交集一定是在i+1区间的前
        return intervals[i][1] * inf + intervals[j][0] < intervals[j][1] * inf + intervals[i][0]
    })
    for _, v := range intervals {
        if v[0] > end2 {
            end1, end2 = v[1] - 1, v[1]
            res += 2
        } else if v[0] > end1 {
            end1, end2 = end2, v[1]
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: intervals = [[1,3],[3,7],[8,9]]
    // Output: 5
    // Explanation: let nums = [2, 3, 4, 8, 9].
    // It can be shown that there cannot be any containing array of size 4.
    fmt.Println(intersectionSizeTwo([][]int{{1,3},{3,7},{8,9}})) // 5 [2, 3, 4, 8, 9]
    // Example 2:
    // Input: intervals = [[1,3],[1,4],[2,5],[3,5]]
    // Output: 3
    // Explanation: let nums = [2, 3, 4].
    // It can be shown that there cannot be any containing array of size 2.
    fmt.Println(intersectionSizeTwo([][]int{{1,3},{1,4},{2,5},{3,5}})) // 3 [2, 3, 4]
    // Example 3:
    // Input: intervals = [[1,2],[2,3],[2,4],[4,5]]
    // Output: 5
    // Explanation: let nums = [1, 2, 3, 4, 5].
    // It can be shown that there cannot be any containing array of size 4.
    fmt.Println(intersectionSizeTwo([][]int{{1,2},{2,3},{2,4},{4,5}})) // 5 [1, 2, 3, 4, 5]
}