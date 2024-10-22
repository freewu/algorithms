package main

// 1868. Product of Two Run-Length Encoded Arrays
// Run-length encoding is a compression algorithm that allows for an integer array nums with many segments of consecutive repeated numbers to be represented by a (generally smaller) 2D array encoded. 
// Each encoded[i] = [vali, freqi] describes the ith segment of repeated numbers in nums where vali is the value that is repeated freqi times.

//     For example, nums = [1,1,1,2,2,2,2,2] is represented by the run-length encoded array encoded = [[1,3],[2,5]]. 
//     Another way to read this is "three 1's followed by five 2's".

// The product of two run-length encoded arrays encoded1 and encoded2 can be calculated using the following steps:
//     1. Expand both encoded1 and encoded2 into the full arrays nums1 and nums2 respectively.
//     2. Create a new array prodNums of length nums1.length and set prodNums[i] = nums1[i] * nums2[i].
//     3. Compress prodNums into a run-length encoded array and return it.

// You are given two run-length encoded arrays encoded1 and encoded2 representing full arrays nums1 and nums2 respectively. 
// Both nums1 and nums2 have the same length. 
// Each encoded1[i] = [vali, freqi] describes the ith segment of nums1, and each encoded2[j] = [valj, freqj] describes the jth segment of nums2.

// Return the product of encoded1 and encoded2.

// Note: Compression should be done such that the run-length encoded array has the minimum possible length.

// Example 1:
// Input: encoded1 = [[1,3],[2,3]], encoded2 = [[6,3],[3,3]]
// Output: [[6,6]]
// Explanation: encoded1 expands to [1,1,1,2,2,2] and encoded2 expands to [6,6,6,3,3,3].
// prodNums = [6,6,6,6,6,6], which is compressed into the run-length encoded array [[6,6]].

// Example 2:
// Input: encoded1 = [[1,3],[2,1],[3,2]], encoded2 = [[2,3],[3,3]]
// Output: [[2,3],[6,1],[9,2]]
// Explanation: encoded1 expands to [1,1,1,2,3,3] and encoded2 expands to [2,2,2,3,3,3].
// prodNums = [2,2,2,6,9,9], which is compressed into the run-length encoded array [[2,3],[6,1],[9,2]].

// Constraints:
//     1 <= encoded1.length, encoded2.length <= 10^5
//     encoded1[i].length == 2
//     encoded2[j].length == 2
//     1 <= vali, freqi <= 10^4 for each encoded1[i].
//     1 <= valj, freqj <= 10^4 for each encoded2[j].
//     The full arrays that encoded1 and encoded2 represent are the same length.

import "fmt"

func findRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
    res := [][]int{}
    i, j, n1, n2 := 0, 0, len(encoded1), len(encoded2)
    for i < n1 && j < n2 {
        v1, f1 := encoded1[i][0], encoded1[i][1]
        v2, f2 := encoded2[j][0], encoded2[j][1]
        if f1 == f2 {
            res = append(res, []int{ v1 * v2, f2 })
            i++
            j++
        } else if f1 > f2 {
            res = append(res, []int{ v1 * v2, f2 })
            encoded1[i][1] -= f2
            j++
        } else {
            res = append(res, []int{ v1 * v2, f1 })
            encoded2[j][1] -= f1
            i++
        }
        // 处理被分割的情况
        n := len(res)
        if n > 1 && res[n-1][0] == res[n-2][0] {
            res[n-2][1] += res[n-1][1]
            res = res[:n-1]
        }
    }
    return res
}

func findRLEArray1(encoded1 [][]int, encoded2 [][]int) [][]int {
    res := [][]int{}
    i1, i2 := 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i1 < len(encoded1) && i2 < len(encoded2) {
        mn := min(encoded1[i1][1], encoded2[i2][1])
        encoded1[i1][1] -= mn
        encoded2[i2][1] -= mn
        dotProduct := encoded1[i1][0] * encoded2[i2][0]
        // clear invalid encode
        if encoded1[i1][1] == 0 { i1++ }
        if encoded2[i2][1] == 0 { i2++ }
        n := len(res)
        // compress result
        if n > 0 && res[n - 1][0] == dotProduct {
            res[n - 1][1] += mn
        } else {
            res = append(res, []int{ dotProduct, mn})
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: encoded1 = [[1,3],[2,3]], encoded2 = [[6,3],[3,3]]
    // Output: [[6,6]]
    // Explanation: encoded1 expands to [1,1,1,2,2,2] and encoded2 expands to [6,6,6,3,3,3].
    // prodNums = [6,6,6,6,6,6], which is compressed into the run-length encoded array [[6,6]].
    fmt.Println(findRLEArray([][]int{{1,3},{2,3}}, [][]int{{6,3},{3,3}})) // [[6,6]]
    // Example 2:
    // Input: encoded1 = [[1,3],[2,1],[3,2]], encoded2 = [[2,3],[3,3]]
    // Output: [[2,3],[6,1],[9,2]]
    // Explanation: encoded1 expands to [1,1,1,2,3,3] and encoded2 expands to [2,2,2,3,3,3].
    // prodNums = [2,2,2,6,9,9], which is compressed into the run-length encoded array [[2,3],[6,1],[9,2]].
    fmt.Println(findRLEArray([][]int{{1,3},{2,1},{3,2}}, [][]int{{2,3},{3,3}})) // [[2,3],[6,1],[9,2]]

    fmt.Println(findRLEArray1([][]int{{1,3},{2,3}}, [][]int{{6,3},{3,3}})) // [[6,6]]
    fmt.Println(findRLEArray1([][]int{{1,3},{2,1},{3,2}}, [][]int{{2,3},{3,3}})) // [[2,3],[6,1],[9,2]]
}