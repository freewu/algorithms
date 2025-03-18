package main

// 3480. Maximize Subarrays After Removing One Conflicting Pair
// You are given an integer n which represents an array nums containing the numbers from 1 to n in order. 
// Additionally, you are given a 2D array conflictingPairs, where conflictingPairs[i] = [a, b] indicates that a and b form a conflicting pair.

// Remove exactly one element from conflictingPairs. 
// Afterward, count the number of non-empty subarrays of nums which do not contain both a and b for any remaining conflicting pair [a, b].

// Return the maximum number of subarrays possible after removing exactly one conflicting pair.

// Example 1:
// Input: n = 4, conflictingPairs = [[2,3],[1,4]]
// Output: 9
// Explanation:
// Remove [2, 3] from conflictingPairs. Now, conflictingPairs = [[1, 4]].
// There are 9 subarrays in nums where [1, 4] do not appear together. They are [1], [2], [3], [4], [1, 2], [2, 3], [3, 4], [1, 2, 3] and [2, 3, 4].
// The maximum number of subarrays we can achieve after removing one element from conflictingPairs is 9.

// Example 2:
// Input: n = 5, conflictingPairs = [[1,2],[2,5],[3,5]]
// Output: 12
// Explanation:
// Remove [1, 2] from conflictingPairs. Now, conflictingPairs = [[2, 5], [3, 5]].
// There are 12 subarrays in nums where [2, 5] and [3, 5] do not appear together.
// The maximum number of subarrays we can achieve after removing one element from conflictingPairs is 12.

// Constraints:
//     2 <= n <= 10^5
//     1 <= conflictingPairs.length <= 2 * n
//     conflictingPairs[i].length == 2
//     1 <= conflictingPairs[i][j] <= n
//     conflictingPairs[i][0] != conflictingPairs[i][1]

import "fmt"
import "slices"

func maxSubarrays(n int, conflictingPairs [][]int) int64 {
    // Ensure conflictingPairs[i][0] is always the smaller value
    for i := range conflictingPairs {
        if conflictingPairs[i][0] > conflictingPairs[i][1] {
            conflictingPairs[i][0], conflictingPairs[i][1] = conflictingPairs[i][1], conflictingPairs[i][0]
        }
    }
    // Sort conflictingPairs based on the first element of each pair
    slices.SortFunc(conflictingPairs, func(a, b []int) int { 
        return b[0] - a[0]
    })
    rightBoundary, conflictIndex, largestCorrection, validSubarrays, curPenalty, maxPenalty := n + 1, 0, 0, 0, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // Traverse from n down to 1
    for b := n; b > 0; b-- {
        // Process conflicting pairs affecting `b`
        for ; conflictIndex < len(conflictingPairs) && conflictingPairs[conflictIndex][0] >= b; conflictIndex++ {
            if rightBoundary > conflictingPairs[conflictIndex][1] {
                // Found a new correction point
                largestCorrection = rightBoundary - conflictingPairs[conflictIndex][1]
                rightBoundary, curPenalty = conflictingPairs[conflictIndex][1], 0
            } else {
                // Minimize the correction penalty
                largestCorrection = min(largestCorrection, conflictingPairs[conflictIndex][1] - rightBoundary)
            }
        }
        validSubarrays += (rightBoundary - b)
        curPenalty += (largestCorrection)
        maxPenalty = max(maxPenalty, curPenalty)
    }
    return int64(validSubarrays + maxPenalty)
}

func maxSubarrays1(n int, conflictingPairs [][]int) int64 {
    groups := make([][2]int, n + 1)
    for i := range groups {
        groups[i] = [2]int{n + 1, n + 1}
    }
    for _, p := range conflictingPairs {
        a, b := p[0], p[1]
        if a > b {
            a, b = b, a
        }
        p := &groups[a]
        if b < p[0] {
            p[0], p[1] = b, p[0]
        } else if b < p[1] {
            p[1] = b
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, extra, maxExtra, b0, b1 := 0, 0, 0, n + 1, n + 1
    for a := n; a > 0; a-- {
        preB0 := b0
        for _, b := range groups[a] {
            if b < b0 {
                b0, b1 = b, b0
            } else if b < b1 {
                b1 = b
            }
        }
        res += (b0 - a)
        if b0 != preB0 {
            extra = 0
        }
        extra += b1 - b0
        maxExtra = max(maxExtra, extra)
    }
    return int64(res + maxExtra)
}

func main() {
    // Example 1:
    // Input: n = 4, conflictingPairs = [[2,3],[1,4]]
    // Output: 9
    // Explanation:
    // Remove [2, 3] from conflictingPairs. Now, conflictingPairs = [[1, 4]].
    // There are 9 subarrays in nums where [1, 4] do not appear together. They are [1], [2], [3], [4], [1, 2], [2, 3], [3, 4], [1, 2, 3] and [2, 3, 4].
    // The maximum number of subarrays we can achieve after removing one element from conflictingPairs is 9.
    fmt.Println(maxSubarrays(4, [][]int{{2,3},{1,4}})) // 9
    // Example 2:
    // Input: n = 5, conflictingPairs = [[1,2],[2,5],[3,5]]
    // Output: 12
    // Explanation:
    // Remove [1, 2] from conflictingPairs. Now, conflictingPairs = [[2, 5], [3, 5]].
    // There are 12 subarrays in nums where [2, 5] and [3, 5] do not appear together.
    // The maximum number of subarrays we can achieve after removing one element from conflictingPairs is 12.
    fmt.Println(maxSubarrays(5, [][]int{{1,2},{2,5},{3,5}})) // 12

    fmt.Println(maxSubarrays1(4, [][]int{{2,3},{1,4}})) // 9
    fmt.Println(maxSubarrays1(5, [][]int{{1,2},{2,5},{3,5}})) // 12
}