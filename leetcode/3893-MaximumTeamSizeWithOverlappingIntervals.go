package main

// 3893. Maximum Team Size with Overlapping Intervals
// You are given two integer arrays startTime and endTime of length n.
//     1. startTime[i] represents the start time of the ith employee.
//     2. endTime[i] represents the end time of the ith employee.

// Two employees i and j can interact if their time intervals overlap. 
// Two intervals are considered overlapping if they share at least one common time point.

// A team is valid if there exists at least one employee in the team who can interact with every other member of the team.

// Return an integer denoting the maximum possible size of such a team.

// Example 1:
// Input: startTime = [1,2,3], endTime = [4,5,6]
// Output: 3
// Explanation:
// For i = 0 with interval [1, 4].
// It overlaps with i = 1 having interval [2, 5] and i = 2 having interval [3, 6].
// Thus, index 0 can interact with all other indices, so the team size is 3.

// Example 2:
// Input: startTime = [2,5,8], endTime = [3,7,9]
// Output: 1
// Explanation:
// For i = 0, interval [2, 3] does not overlap with [5, 7] or [8, 9].
// For i = 1, interval [5, 7] does not overlap with [2, 3] or [8, 9].
// For i = 2, interval [8, 9] does not overlap with [2, 3] or [5, 7].
// Thus, no index can interact with others, so the maximum team size is 1.

// Example 3:
// Input: startTime = [3,4,6], endTime = [8,5,7]
// Output: 3
// Explanation:
// For i = 0 with interval [3, 8].
// It overlaps with i = 1 having interval [4, 5] and i = 2 having interval [6, 7].
// Thus, index 0 can interact with all other indices, so the team size is 3.

// Constraints:
//     1 <= n == startTime.length == endTime.length <= 10^5
//     0 <= startTime[i] <= endTime[i] <= 10^9

import "fmt"
import "sort"

func maximumTeamSize(startTime []int, endTime []int) int {
    res, n := 1, len(startTime)
    // 复制并排序开始时间数组
    sortedStarts := make([]int, n)
    copy(sortedStarts, startTime)
    sort.Ints(sortedStarts)
    // 复制并排序结束时间数组
    sortedEnds := make([]int, n)
    copy(sortedEnds, endTime)
    sort.Ints(sortedEnds)
    upperBound := func(arr []int, target int) int { // 返回排序数组中 <= target 的元素个数
        low, high := 0, len(arr)
        for low < high {
            mid := (low + high) / 2
            if arr[mid] <= target {
                low = mid + 1
            } else {
                high = mid
            }
        }
        return low
    }
    lowerBound := func(arr []int, target int) int { // 返回排序数组中 < target 的元素个数
        low, high := 0, len(arr)
        for low < high {
            mid := (low + high) / 2
            if arr[mid] < target {
                low = mid + 1
            } else {
                high = mid
            }
        }
        return low
    }
    // 遍历每一个区间作为中心区间
    for i := 0; i < n; i++ {
        s, e := startTime[i], endTime[i]
        // 二分查找：数组中 <= target 的元素个数（对应TS的upperBound）
        startsBeforeOrAtEnd := upperBound(sortedStarts, e)
        // 二分查找：数组中 < target 的元素个数（对应TS的lowerBound）
        endsBeforeStart := lowerBound(sortedEnds, s)
        // 计算当前重叠区间数量
        count := startsBeforeOrAtEnd - endsBeforeStart
        if count > res {
            res = count
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: startTime = [1,2,3], endTime = [4,5,6]
    // Output: 3
    // Explanation:
    // For i = 0 with interval [1, 4].
    // It overlaps with i = 1 having interval [2, 5] and i = 2 having interval [3, 6].
    // Thus, index 0 can interact with all other indices, so the team size is 3.
    fmt.Println(maximumTeamSize([]int{1,2,3}, []int{4,5,6})) // 3
    // Example 2:
    // Input: startTime = [2,5,8], endTime = [3,7,9]
    // Output: 1
    // Explanation:
    // For i = 0, interval [2, 3] does not overlap with [5, 7] or [8, 9].
    // For i = 1, interval [5, 7] does not overlap with [2, 3] or [8, 9].
    // For i = 2, interval [8, 9] does not overlap with [2, 3] or [5, 7].
    // Thus, no index can interact with others, so the maximum team size is 1.
    fmt.Println(maximumTeamSize([]int{2,5,8}, []int{3,7,9})) // 1
    // Example 3:
    // Input: startTime = [3,4,6], endTime = [8,5,7]
    // Output: 3
    // Explanation:
    // For i = 0 with interval [3, 8].
    // It overlaps with i = 1 having interval [4, 5] and i = 2 having interval [6, 7].
    // Thus, index 0 can interact with all other indices, so the team size is 3.  
    fmt.Println(maximumTeamSize([]int{3,4,6}, []int{8,5,7})) // 3

    fmt.Println(maximumTeamSize([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(maximumTeamSize([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 9
    fmt.Println(maximumTeamSize([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(maximumTeamSize([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 1
}