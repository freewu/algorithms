package main

// 4001. Aggregate Two Time Series
// You are given two 2D integer arrays series1 and series2.

// Each element in both series is of the form [timestamp, value], where:
//     1. timestamp is an integer representing the time.
//     2. value is an integer representing the value at that timestamp.

// Each array is sorted in strictly increasing order of timestamp.

// For any timestamp not present in a series, its value is taken from the next available timestamp in the same series if one exists. 
// Otherwise, its value is considered 0.

// The aggregated series is formed by summing the corresponding values from both series at every timestamp that appears in either series.

// Return the aggregated series as a 2D integer array of [timestamp, summedValue] pairs, sorted in strictly increasing order of timestamp.

// Example 1:
// Input: series1 = [[1,3],[4,1]], series2 = [[2,2],[5,2]]
// Output: [[1,5],[2,3],[4,3],[5,2]]
// Explanation:
// Timestamp | series1 | series2 | summedValue
// 1         | 3       | 2       | 5
// 2         | 1       | 2       | 3
// 4         | 1       | 2       | 3
// 5         | 0       | 2       | 2
// Thus, the aggregated series is [[1, 5], [2, 3], [4, 3], [5, 2]].

// Example 2:
// Input: series1 = [[1,5],[3,1]], series2 = [[2,2]]
// Output: [[1,7],[2,3],[3,1]]
// Explanation:
// Timestamp | series1 | series2 | summedValue
// 1         | 5       | 2       | 7
// 2         | 1       | 2       | 3
// 3         | 1       | 0       | 1
// Thus, the aggregated series is [[1, 7], [2, 3], [3, 1]].

// Example 3:
// Input: series1 = [[1,5]], series2 = [[1000000000,2]]
// Output: [[1,7],[1000000000,2]]
// Explanation:
// At timestamp 1, the next available value in series2 is 2 at timestamp 1000000000. 
// At timestamp 1000000000, there is no later timestamp in series1, so its value is 0. 
// Only timestamps that appear in at least one of the two series are included.

// Constraints:
//     1 <= series1.length, series2.length <= 10^5
//     series1[i].length == series2[i].length == 2
//     1 <= series1[i][0], series2[i][0] <= 10^9
//     1 <= series1[i][1], series2[i][1] <= 10^9
//     Each series is sorted in strictly increasing order of timestamp.

import "fmt"

var dp [200_000][]int

func init() {
    for i := range dp {
        dp[i] = make([]int, 2)
    }
}

func aggregateTimeSeries(series1 [][]int, series2 [][]int) [][]int {
    n1, n2, r := len(series1), len(series2), 0
    series1 = append(series1, []int{1 << 30, 0})
    series2 = append(series2, []int{1 << 30, 0})
    for i, j := 0, 0; i < n1 || j < n2; r++ {
        t := min(series1[i][0], series2[j][0])
        dp[r][0] = t
        dp[r][1] = series1[i][1] + series2[j][1]
        if series1[i][0] == t {
            i++
        }
        if series2[j][0] == t {
            j++
        }
    }
    return dp[:r]
}

func aggregateTimeSeries1(series1 [][]int, series2 [][]int) [][]int {
    // 既然最后都放在一起，最开始就聚合并且从末尾开始遍历，如果时间戳没有值则使用之前的值。
    // 初始值为0。
    // 过程：
    // 1.从末尾开始访问，根据timestamp的值判断。
    // 2.选择当前timestamp大的放入结果中，由于严格递增所以直接比较大小访问即可。
    // 2.1.如果两者相同，则累加两者之和。更新两者当前值，用于后续更新。
    // 2.2.存在某个timestamp比较大，则添加timestamp的值和另一个当前值；更新存在的值。
    // 3.最后还有剩下的则全部添加，添加方式直接加入即可。
    res := make([][]int, 0, len(series1)+len(series2))
    i, j := len(series1)-1, len(series2)-1
    s1v, s2v := 0, 0
    for i >= 0 && j >= 0 {
        if series1[i][0] == series2[j][0] {
            s1v = series1[i][1]
            s2v = series2[j][1]
            res = append(res, []int{series1[i][0], s1v + s2v})
            i--
            j--
        } else if series1[i][0] > series2[j][0] {
            s1v = series1[i][1]
            res = append(res, []int{series1[i][0], s1v + s2v})
            i--
        } else {
            s2v = series2[j][1]
            res = append(res, []int{series2[j][0], s1v + s2v})
            j--
        }
    }
    for i >= 0 {
        s1v = series1[i][1]
        res = append(res, []int{series1[i][0], s1v + s2v})
        i--
    }
    for j >= 0 {
        s2v = series2[j][1]
        res = append(res, []int{series2[j][0], s1v + s2v})
        j--
    }
    for i := 0; i < len(res)/2; i++ {
        res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
    }
    return res
}

func main() {
    // Example 1:
    // Input: series1 = [[1,3],[4,1]], series2 = [[2,2],[5,2]]
    // Output: [[1,5],[2,3],[4,3],[5,2]]
    // Explanation:
    // Timestamp | series1 | series2 | summedValue
    // 1         | 3       | 2       | 5
    // 2         | 1       | 2       | 3
    // 4         | 1       | 2       | 3
    // 5         | 0       | 2       | 2
    // Thus, the aggregated series is [[1, 5], [2, 3], [4, 3], [5, 2]].
    fmt.Println(aggregateTimeSeries([][]int{{1,3},{4,1}}, [][]int{{2,2},{5,2}})) // [[1,5],[2,3],[4,3],[5,2]]
    // Example 2:
    // Input: series1 = [[1,5],[3,1]], series2 = [[2,2]]
    // Output: [[1,7],[2,3],[3,1]]
    // Explanation:
    // Timestamp | series1 | series2 | summedValue
    // 1         | 5       | 2       | 7
    // 2         | 1       | 2       | 3
    // 3         | 1       | 0       | 1
    // Thus, the aggregated series is [[1, 7], [2, 3], [3, 1]].
    fmt.Println(aggregateTimeSeries([][]int{{1,5},{3,1}}, [][]int{{2,2}})) // [[1,7],[2,3],[3,1]]
    // Example 3:
    // Input: series1 = [[1,5]], series2 = [[1000000000,2]]
    // Output: [[1,7],[1000000000,2]]
    // Explanation:
    // At timestamp 1, the next available value in series2 is 2 at timestamp 1000000000. 
    // At timestamp 1000000000, there is no later timestamp in series1, so its value is 0. 
    // Only timestamps that appear in at least one of the two series are included.
    fmt.Println(aggregateTimeSeries([][]int{{1,5}}, [][]int{{1000000000,2}})) // [[1,7],[1000000000,2]]

    fmt.Println(aggregateTimeSeries1([][]int{{1,3},{4,1}}, [][]int{{2,2},{5,2}})) // [[1,5],[2,3],[4,3],[5,2]]
    fmt.Println(aggregateTimeSeries1([][]int{{1,5},{3,1}}, [][]int{{2,2}})) // [[1,7],[2,3],[3,1]]
    fmt.Println(aggregateTimeSeries1([][]int{{1,5}}, [][]int{{1000000000,2}})) // [[1,7],[1000000000,2]]
}