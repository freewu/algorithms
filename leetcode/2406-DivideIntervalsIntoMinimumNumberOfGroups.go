package main

// 2406. Divide Intervals Into Minimum Number of Groups
// You are given a 2D integer array intervals where intervals[i] = [lefti, righti] represents the inclusive interval [lefti, righti].

// You have to divide the intervals into one or more groups such that each interval is in exactly one group, 
// and no two intervals that are in the same group intersect each other.

// Return the minimum number of groups you need to make.

// Two intervals intersect if there is at least one common number between them. 
// For example, the intervals [1, 5] and [5, 8] intersect.

// Example 1:
// Input: intervals = [[5,10],[6,8],[1,5],[2,3],[1,10]]
// Output: 3
// Explanation: We can divide the intervals into the following groups:
// - Group 1: [1, 5], [6, 8].
// - Group 2: [2, 3], [5, 10].
// - Group 3: [1, 10].
// It can be proven that it is not possible to divide the intervals into fewer than 3 groups.

// Example 2:
// Input: intervals = [[1,3],[5,6],[8,10],[11,13]]
// Output: 1
// Explanation: None of the intervals overlap, so we can put all of them in one group.

// Constraints:
//     1 <= intervals.length <= 10^5
//     intervals[i].length == 2
//     1 <= lefti <= righti <= 10^6

import "fmt"
import "sort"

func minGroups(intervals [][]int) int {
    res, mp := 0, make(map[int]int)
    for _, v := range intervals {
        mp[v[0]]++
        mp[v[1] + 1]--
    }
    keys := make([]int, 0, len(mp))
    for k := range mp {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    for i := range keys {
        if i > 0 {
            mp[keys[i]] += mp[keys[i-1]]
        }
    }
    for _, v := range mp {
        if v > res {
            res = v
        }
    }
    return res
}

func minGroups1(intervals [][]int) int {
    res, sum, end := 0, 0, 0
    diff := make([]int, 1_000_005)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v:= range intervals {
        diff[v[0]]++
        diff[v[1]+1]--
        end = max(end, v[1] + 1)
    }
    for i := 0; i <= end; i++ {
        sum += diff[i]
        res = max(res, sum)
    }
    return res
}

func minGroups2(intervals [][]int) int {
    start, end := 1 << 31 , -1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range intervals { // Find the minimum and maximum value in the interval
        start = min(start, v[0])
        end = max(end, v[1])
    }
    pointToCount := make([]int, end + 2) // Initialize the list with all zeroes
    // Iterate over the intervals:
    //  - For each interval [start, end], increment pointToCount[start] to indicate the start of an interval
    //  - Decrement pointToCount[end + 1] to mark the point where the interval ends
    for _, v := range intervals {
        pointToCount[v[0]]++
        pointToCount[v[1] + 1]--
    }
    // Loop from rangeStart to rangeEnd and maintain a running sum of active intervals.
    // Track the maximum number of concurrent intervals as res.
    res, sum := 0, 0
    for i := start; i <= end; i++ {
        // Update currently active intervals
        sum += pointToCount[i]
        res = max(res, sum)
    }
    return res
}

func main() {
    // Example 1:
    // Input: intervals = [[5,10],[6,8],[1,5],[2,3],[1,10]]
    // Output: 3
    // Explanation: We can divide the intervals into the following groups:
    // - Group 1: [1, 5], [6, 8].
    // - Group 2: [2, 3], [5, 10].
    // - Group 3: [1, 10].
    // It can be proven that it is not possible to divide the intervals into fewer than 3 groups.
    fmt.Println(minGroups([][]int{{5,10},{6,8},{1,5},{2,3},{1,10}})) // 3
    // Example 2:
    // Input: intervals = [[1,3],[5,6],[8,10],[11,13]]
    // Output: 1
    // Explanation: None of the intervals overlap, so we can put all of them in one group.
    fmt.Println(minGroups([][]int{{1,3},{5,6},{8,10},{11,13}})) // 1

    fmt.Println(minGroups1([][]int{{5,10},{6,8},{1,5},{2,3},{1,10}})) // 3
    fmt.Println(minGroups1([][]int{{1,3},{5,6},{8,10},{11,13}})) // 1

    fmt.Println(minGroups2([][]int{{5,10},{6,8},{1,5},{2,3},{1,10}})) // 3
    fmt.Println(minGroups2([][]int{{1,3},{5,6},{8,10},{11,13}})) // 1
}