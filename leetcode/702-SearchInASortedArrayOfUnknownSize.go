package main

// 702. Search in a Sorted Array of Unknown Size
// This is an interactive problem.
// You have a sorted array of unique elements and an unknown size. 
// You do not have an access to the array but you can use the ArrayReader interface to access it. 
// You can call ArrayReader.get(i) that:
//     returns the value at the ith index (0-indexed) of the secret array (i.e., secret[i]), or
//     returns 2^31 - 1 if the i is out of the boundary of the array.

// You are also given an integer target.
// Return the index k of the hidden array where secret[k] == target or return -1 otherwise.
// You must write an algorithm with O(log n) runtime complexity.

// Example 1:
// Input: secret = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in secret and its index is 4.

// Example 2:
// Input: secret = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in secret so return -1.

// Constraints:
//     1 <= secret.length <= 10^4
//     -10^4 <= secret[i], target <= 10^4
//     secret is sorted in a strictly increasing order.

import "fmt"

type ArrayReader struct {
    data []int
}

func Constructor(data []int) ArrayReader {
    return ArrayReader{data}
}

// returns the value at the ith index (0-indexed) of the secret array (i.e., secret[i]), or
// returns 2^31 - 1 if the i is out of the boundary of the array.
func (this *ArrayReader) get(index int) int {
    if index > len(this.data) -1 {
        return  1 << 31-1
    }
    return this.data[index]
}

/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

func search(reader ArrayReader, target int) int {
    left, right := 0, 10000 // 1 <= secret.length <= 10^4
    for left <= right {
        mid := (left + right) >> 1
        // 如果超出边界 或 大于目标值 <- 向左靠拢
        if v := reader.get(mid); v == 1 << 31-1 || v > target {
            right = mid - 1
        } else if v < target { // 小于目标值 -> 向右靠拢
            left = mid + 1
        } else { // 找到目标值
            return mid
        }
    }
    return -1
}

func main() {
    // Explanation: 9 exists in secret and its index is 4.
    fmt.Println(search(Constructor([]int{-1,0,3,5,9,12}), 9)) // 4

    // Explanation: 2 does not exist in secret so return -1.
    fmt.Println(search(Constructor([]int{-1,0,3,5,9,12}), 2)) // -1
}