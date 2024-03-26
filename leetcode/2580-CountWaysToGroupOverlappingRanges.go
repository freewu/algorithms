package main

// 2580. Count Ways to Group Overlapping Ranges
// You are given a 2D integer array ranges where ranges[i] = [starti, endi] denotes 
// that all integers between starti and endi (both inclusive) are contained in the ith range.

// You are to split ranges into two (possibly empty) groups such that:
//     Each range belongs to exactly one group.
//     Any two overlapping ranges must belong to the same group.

// Two ranges are said to be overlapping if there exists at least one integer that is present in both ranges.
//     For example, [1, 3] and [2, 5] are overlapping because 2 and 3 occur in both ranges.

// Return the total number of ways to split ranges into two groups. Since the answer may be very large, 
// return it modulo 10^9 + 7.

 
// Example 1:
// Input: ranges = [[6,10],[5,15]]
// Output: 2
// Explanation: 
// The two ranges are overlapping, so they must be in the same group.
// Thus, there are two possible ways:
// - Put both the ranges together in group 1.
// - Put both the ranges together in group 2.

// Example 2:
// Input: ranges = [[1,3],[10,20],[2,5],[4,8]]
// Output: 4
// Explanation: 
// Ranges [1,3], and [2,5] are overlapping. So, they must be in the same group.
// Again, ranges [2,5] and [4,8] are also overlapping. So, they must also be in the same group. 
// Thus, there are four possible ways to group them:
// - All the ranges in group 1.
// - All the ranges in group 2.
// - Ranges [1,3], [2,5], and [4,8] in group 1 and [10,20] in group 2.
// - Ranges [1,3], [2,5], and [4,8] in group 2 and [10,20] in group 1.

// Constraints:
//     1 <= ranges.length <= 10^5
//     ranges[i].length == 2
//     0 <= starti <= endi <= 10^9

import "fmt"
import "sort"

func countWays(ranges [][]int) int {
    // 先从小到大排序
    sort.Slice(ranges, func(i, j int) bool {
        if ranges[i][0] == ranges[j][0] {
            return ranges[i][1] < ranges[j][1]
        }
        return ranges[i][0] < ranges[j][0]
    })
    prevEnd := -1
    distinct := 0
    for _, v := range ranges {
        // 每个区间只属于一个组。
        // 两个有 交集 的区间必须在 同一个 组内。
        // 如果两个区间有至少 一个 公共整数，那么这两个区间是 有交集 的
        // 存在重叠
        if v[0] > prevEnd {
            distinct++
            prevEnd = v[1]
        } else {
            if v[1] > prevEnd {
                prevEnd = v[1]
            }
        }
    }
    ways := 1
    for i := 0; i < distinct; i++ {
        ways *= 2
        ways %= 1000000007 // 由于答案可能很大，将它对 10^9 + 7 取余 后返回
    }
    return ways
}

func main() {
    // The two ranges are overlapping, so they must be in the same group.
    // Thus, there are two possible ways:
    // - Put both the ranges together in group 1.
    // - Put both the ranges together in group 2.
    fmt.Println(
        countWays(
            [][]int{
                []int{6,10},
                []int{5,15},
            },
        ),
    ) // 2

    // Ranges [1,3], and [2,5] are overlapping. So, they must be in the same group.
    // Again, ranges [2,5] and [4,8] are also overlapping. So, they must also be in the same group. 
    // Thus, there are four possible ways to group them:
    // - All the ranges in group 1.
    // - All the ranges in group 2.
    // - Ranges [1,3], [2,5], and [4,8] in group 1 and [10,20] in group 2.
    // - Ranges [1,3], [2,5], and [4,8] in group 2 and [10,20] in group 1.
    fmt.Println(
        countWays(
            [][]int{
                []int{1,3},
                []int{10,20},
                []int{2,5},
                []int{4,8},
            },
        ),
    ) // 4
}