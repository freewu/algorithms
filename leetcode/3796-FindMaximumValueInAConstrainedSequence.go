package main

// 3796. Find Maximum Value in a Constrained Sequence
// You are given an integer n, a 2D integer array restrictions, and an integer array diff of length n - 1. 
// Your task is to construct a sequence of length n, denoted by a[0], a[1], ..., a[n - 1], such that it satisfies the following conditions:
//     1. a[0] is 0.
//     2. All elements in the sequence are non-negative.
//     3. For every index i (0 <= i <= n - 2), abs(a[i] - a[i + 1]) <= diff[i].
//     4. For each restrictions[i] = [idx, maxVal], the value at position idx in the sequence must not exceed maxVal (i.e., a[idx] <= maxVal).

// Your goal is to construct a valid sequence that maximizes the largest value within the sequence while satisfying all the above conditions.

// Return an integer denoting the largest value present in such an optimal sequence.

// Example 1:
// Input: n = 10, restrictions = [[3,1],[8,1]], diff = [2,2,3,1,4,5,1,1,2]
// Output: 6
// Explanation:
// The sequence a = [0, 2, 4, 1, 2, 6, 2, 1, 1, 3] satisfies the given constraints (a[3] <= 1 and a[8] <= 1).
// The maximum value in the sequence is 6.

// Example 2:
// Input: n = 8, restrictions = [[3,2]], diff = [3,5,2,4,2,3,1]
// Output: 12
// Explanation:
// The sequence a = [0, 3, 3, 2, 6, 8, 11, 12] satisfies the given constraints (a[3] <= 2).
// The maximum value in the sequence is 12.
 
// Constraints:
//     2 <= n <= 10^5
//     1 <= restrictions.length <= n - 1
//     restrictions[i].length == 2
//     restrictions[i] = [idx, maxVal]
//     1 <= idx < n
//     1 <= maxVal <= 10^6
//     diff.length == n - 1
//     1 <= diff[i] <= 10
//     The values of restrictions[i][0] are unique.

import "fmt"
import "slices"

func findMaxVal(n int, restrictions [][]int, diff []int) int {
    mx := make([]int, n)
    for i := range mx {
        mx[i] = 1 << 31
    }
    for _, v := range restrictions {
        mx[v[0]] = v[1]
    }
    arr := make([]int, n)
    for i, d := range diff {
        arr[i + 1] = min(arr[i] + d, mx[i + 1])
    }
    for i := n - 2; i > 0; i-- {
        arr[i] = min(arr[i], arr[i + 1] + diff[i])
    }
    return slices.Max(arr)
}

func findMaxVal1(n int, restrictions [][]int, diff []int) int {
    arr := make([]int, n)
    for _, r := range restrictions {
        arr[r[0]] = r[1]
    }
    for i := 1; i < n; i++ {
        t := arr[i - 1] + diff[i-1]
        if arr[i] > 0 {
            arr[i] = min(t, arr[i])
        } else {
            arr[i] = t
        }
    }
    for i := n - 2; i >= 0; i-- {
        arr[i] = min(arr[i], arr[i + 1] + diff[i])
    }
    return slices.Max(arr)
}

func main() {
    // Example 1:
    // Input: n = 10, restrictions = [[3,1],[8,1]], diff = [2,2,3,1,4,5,1,1,2]
    // Output: 6
    // Explanation:
    // The sequence a = [0, 2, 4, 1, 2, 6, 2, 1, 1, 3] satisfies the given constraints (a[3] <= 1 and a[8] <= 1).
    // The maximum value in the sequence is 6.
    fmt.Println(findMaxVal(10, [][]int{{3,1},{8,1}}, []int{2,2,3,1,4,5,1,1,2})) // 6
    // Example 2:
    // Input: n = 8, restrictions = [[3,2]], diff = [3,5,2,4,2,3,1]
    // Output: 12
    // Explanation:
    // The sequence a = [0, 3, 3, 2, 6, 8, 11, 12] satisfies the given constraints (a[3] <= 2).
    // The maximum value in the sequence is 12. 
    fmt.Println(findMaxVal(8, [][]int{{3,2}}, []int{3,5,2,4,2,3,1})) // 12

    fmt.Println(findMaxVal(10, [][]int{{3,2}}, []int{1,2,3,4,5,6,7,8,9})) // 41
    fmt.Println(findMaxVal(10, [][]int{{3,2}}, []int{9,8,7,6,5,4,3,2,1})) // 23

    fmt.Println(findMaxVal1(10, [][]int{{3,1},{8,1}}, []int{2,2,3,1,4,5,1,1,2})) // 6
    fmt.Println(findMaxVal1(8, [][]int{{3,2}}, []int{3,5,2,4,2,3,1})) // 12
    fmt.Println(findMaxVal1(10, [][]int{{3,2}}, []int{1,2,3,4,5,6,7,8,9})) // 41
    fmt.Println(findMaxVal1(10, [][]int{{3,2}}, []int{9,8,7,6,5,4,3,2,1})) // 23
}