package main

// 1943. Describe the Painting
// There is a long and thin painting that can be represented by a number line. 
// The painting was painted with multiple overlapping segments where each segment was painted with a unique color. 
// You are given a 2D integer array segments, where segments[i] = [starti, endi, colori] represents the half-closed segment [starti, endi) with colori as the color.

// The colors in the overlapping segments of the painting were mixed when it was painted. 
// When two or more colors mix, they form a new color that can be represented as a set of mixed colors.
//     For example, if colors 2, 4, and 6 are mixed, then the resulting mixed color is {2,4,6}.

// For the sake of simplicity, you should only output the sum of the elements in the set rather than the full set.

// You want to describe the painting with the minimum number of non-overlapping half-closed segments of these mixed colors. 
// These segments can be represented by the 2D array painting where painting[j] = [leftj, rightj, mixj] describes a half-closed segment [leftj, rightj) with the mixed color sum of mixj.

//     For example, the painting created with segments = [[1,4,5],[1,7,7]] can be described by painting = [[1,4,12],[4,7,7]] because:
//         [1,4) is colored {5,7} (with a sum of 12) from both the first and second segments.
//         [4,7) is colored {7} from only the second segment.

// Return the 2D array painting describing the finished painting (excluding any parts that are not painted). 
// You may return the segments in any order.

// A half-closed segment [a, b) is the section of the number line between points a and b including point a and not including point b.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/18/1.png" />
// Input: segments = [[1,4,5],[4,7,7],[1,7,9]]
// Output: [[1,4,14],[4,7,16]]
// Explanation: The painting can be described as follows:
// - [1,4) is colored {5,9} (with a sum of 14) from the first and third segments.
// - [4,7) is colored {7,9} (with a sum of 16) from the second and third segments.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/18/2.png" />
// Input: segments = [[1,7,9],[6,8,15],[8,10,7]]
// Output: [[1,6,9],[6,7,24],[7,8,15],[8,10,7]]
// Explanation: The painting can be described as follows:
// - [1,6) is colored 9 from the first segment.
// - [6,7) is colored {9,15} (with a sum of 24) from the first and second segments.
// - [7,8) is colored 15 from the second segment.
// - [8,10) is colored 7 from the third segment.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/07/04/c1.png" />
// Input: segments = [[1,4,5],[1,4,7],[4,7,1],[4,7,11]]
// Output: [[1,4,12],[4,7,12]]
// Explanation: The painting can be described as follows:
// - [1,4) is colored {5,7} (with a sum of 12) from the first and second segments.
// - [4,7) is colored {1,11} (with a sum of 12) from the third and fourth segments.
// Note that returning a single segment [1,7) is incorrect because the mixed color sets are different.

// Constraints:
//     1 <= segments.length <= 2 * 10^4
//     segments[i].length == 3
//     1 <= starti < endi <= 10^5
//     1 <= colori <= 10^9
//     Each colori is distinct.

import "fmt"

func splitPainting(segments [][]int) [][]int64 {
    n, MAX := len(segments), 100010
    res, sweep, change := [][]int64{},[100011]int64{}, [100011]bool{}
    for i :=0 ; i < n; i++ {
        start, end, val := segments[i][0], segments[i][1], segments[i][2]
        sweep[start] += int64(val)
        sweep[end] -= int64(val)
        change[start], change[end] = true, true
    }
    for i := 1 ; i < MAX; i++ {
        sweep[i] += sweep[i - 1]
    }
    for start := 1; start <= MAX; start++ {
        currentColor, end := sweep[start], start + 1
        for end <= MAX && !change[end] {
            end++
        }
        if currentColor > 0 {
            res = append(res, []int64{ int64(start), int64(end), currentColor })
        }
        if end == MAX { break }
        start = end - 1
    }
    return res
}

func splitPainting1(segments [][]int) [][]int64 {
    diff, diffIndex := make([]int, 100000 + 10), make([]int, 100000 + 10)
    for i := 0; i < len(segments) ; i ++ {
        start, end, val := segments[i][0], segments[i][1], segments[i][2]
        diff[start] += val
        diffIndex[start] += i
        diff[end] -= val
        diffIndex[end] -= i
    }
    prefix, prefixIndex, res := 0, 0, [][]int64{}
    for i := 1 ; i <= 100000; i ++ {
        prefix += diff[i]
        prefixIndex += diffIndex[i]
        if prefix == 0 { continue }
        match, matchIndex := prefix, prefixIndex
        for j := i + 1; ; j++ {
            prefix += diff[j]
            prefixIndex += diffIndex[j]
            if prefix == match && prefixIndex == matchIndex { continue } 
            prefix -= diff[j]
            prefixIndex -= diffIndex[j]
            res = append(res, []int64{int64(i), int64(j), int64(prefix)})
            i = j - 1
            break
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/18/1.png" />
    // Input: segments = [[1,4,5],[4,7,7],[1,7,9]]
    // Output: [[1,4,14],[4,7,16]]
    // Explanation: The painting can be described as follows:
    // - [1,4) is colored {5,9} (with a sum of 14) from the first and third segments.
    // - [4,7) is colored {7,9} (with a sum of 16) from the second and third segments.
    fmt.Println(splitPainting([][]int{{1,4,5},{4,7,7},{1,7,9}})) // [[1,4,14],[4,7,16]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/18/2.png" />
    // Input: segments = [[1,7,9],[6,8,15],[8,10,7]]
    // Output: [[1,6,9],[6,7,24],[7,8,15],[8,10,7]]
    // Explanation: The painting can be described as follows:
    // - [1,6) is colored 9 from the first segment.
    // - [6,7) is colored {9,15} (with a sum of 24) from the first and second segments.
    // - [7,8) is colored 15 from the second segment.
    // - [8,10) is colored 7 from the third segment.
    fmt.Println(splitPainting([][]int{{1,7,9},{6,8,15},{8,10,7}})) // [[1,6,9],[6,7,24],[7,8,15],[8,10,7]]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/07/04/c1.png" />
    // Input: segments = [[1,4,5],[1,4,7],[4,7,1],[4,7,11]]
    // Output: [[1,4,12],[4,7,12]]
    // Explanation: The painting can be described as follows:
    // - [1,4) is colored {5,7} (with a sum of 12) from the first and second segments.
    // - [4,7) is colored {1,11} (with a sum of 12) from the third and fourth segments.
    // Note that returning a single segment [1,7) is incorrect because the mixed color sets are different.
    fmt.Println(splitPainting([][]int{{1,4,5},{1,4,7},{4,7,1},{4,7,11}})) // [[1,4,12],[4,7,12]]

    fmt.Println(splitPainting1([][]int{{1,4,5},{4,7,7},{1,7,9}})) // [[1,4,14],[4,7,16]]
    fmt.Println(splitPainting1([][]int{{1,7,9},{6,8,15},{8,10,7}})) // [[1,6,9],[6,7,24],[7,8,15],[8,10,7]]
    fmt.Println(splitPainting1([][]int{{1,4,5},{1,4,7},{4,7,1},{4,7,11}})) // [[1,4,12],[4,7,12]]
}