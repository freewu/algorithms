package main

// 378. Kth Smallest Element in a Sorted Matrix
// Given an n x n matrix where each of the rows and columns is sorted in ascending order, 
// return the kth smallest element in the matrix.

// Note that it is the kth smallest element in the sorted order, not the kth distinct element.
// You must find a solution with a memory complexity better than O(n2).

// Example 1:
// Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
// Output: 13
// Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13

// Example 2:
// Input: matrix = [[-5]], k = 1
// Output: -5
 
// Constraints:
//     n == matrix.length == matrix[i].length
//     1 <= n <= 300
//     -10^9 <= matrix[i][j] <= 10^9
//     All the rows and columns of matrix are guaranteed to be sorted in non-decreasing order.
//     1 <= k <= n2

// Follow up:
//     Could you solve the problem with a constant memory (i.e., O(1) memory complexity)?
//     Could you solve the problem in O(n) time complexity? The solution may be too advanced for an interview but you may find reading this paper fun.

import "fmt"
import "sort"

// 二分
func kthSmallest(matrix [][]int, k int) int {
    i, j := matrix[0][0], matrix[len(matrix) - 1][len(matrix[0]) - 1]
    smallerCount := func(matrix [][]int, k int) int {
        res := 0
        for _, v := range matrix {
            res = res + sort.Search(len(v), func(i int) bool {
                return v[i] > k
            })
        }
        return res
    }
    for i < j {
        mid := i + (j - i) / 2
        if smallerCount(matrix, mid) < k {
            i = mid + 1
        } else {
            j = mid
        }
    }
    return i
}

func kthSmallest1(matrix [][]int, k int) int {
    n := len(matrix)
    left, right := matrix[0][0], matrix[n-1][n-1]
    check := func (matrix [][]int, mid, k, n int) bool {
        i, j, res := n - 1, 0, 0
        for i >= 0 && j < n {
            if matrix[i][j] <= mid {
                res += i + 1
                j++
            } else {
                i--
            }
        }
        return res >= k
    }
    for left < right {
        mid := left + (right - left) / 2
        if check(matrix, mid, k, n) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func main() {
    // Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13
    fmt.Println(kthSmallest([][]int{{1,5,9},{10,11,13},{12,13,15}},8)) // 13
    fmt.Println(kthSmallest([][]int{{-5}},1)) // -5

    fmt.Println(kthSmallest1([][]int{{1,5,9},{10,11,13},{12,13,15}},8)) // 13
    fmt.Println(kthSmallest1([][]int{{-5}},1)) // -5
}