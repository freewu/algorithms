package main

// 768. Max Chunks To Make Sorted II
// You are given an integer array arr.

// We split arr into some number of chunks (i.e., partitions), and individually sort each chunk. 
// After concatenating them, the result should equal the sorted array.

// Return the largest number of chunks we can make to sort the array.

// Example 1:
// Input: arr = [5,4,3,2,1]
// Output: 1
// Explanation:
// Splitting into two or more chunks will not return the required result.
// For example, splitting into [5, 4], [3, 2, 1] will result in [4, 5, 1, 2, 3], which isn't sorted.

// Example 2:
// Input: arr = [2,1,3,4,4]
// Output: 4
// Explanation:
// We can split into two chunks, such as [2, 1], [3, 4, 4].
// However, splitting into [2, 1], [3], [4], [4] is the highest number of chunks possible.

// Constraints:
//     1 <= arr.length <= 2000
//     0 <= arr[i] <= 10^8

import "fmt"

// Monotonic stack
func maxChunksToSorted(arr []int) int {
    res, n := 0, len(arr)
    if n <= 1 {
        return n
    }
    stack := []int{}
    for i, v := range arr {
        if len(stack) == 0 || v >= arr[stack[len(stack)-1]] {
            stack = append(stack, i)
        }
    }
    mn := arr[n-1]
    for i := n-1; i > 0; i-- {
        if arr[i] < mn {
            mn = arr[i]
        }
        if stack[len(stack)-1] == i {
            stack = stack[:len(stack)-1]
        }
        if mn >= arr[stack[len(stack)-1]] {
            res++
        }
    }
    return res + 1
}

func main() {
    // Example 1:
    // Input: arr = [5,4,3,2,1]
    // Output: 1
    // Explanation:
    // Splitting into two or more chunks will not return the required result.
    // For example, splitting into [5, 4], [3, 2, 1] will result in [4, 5, 1, 2, 3], which isn't sorted.
    fmt.Println(maxChunksToSorted([]int{5,4,3,2,1})) // 1
    // Example 2:
    // Input: arr = [2,1,3,4,4]
    // Output: 4
    // Explanation:
    // We can split into two chunks, such as [2, 1], [3, 4, 4].
    // However, splitting into [2, 1], [3], [4], [4] is the highest number of chunks possible.
    fmt.Println(maxChunksToSorted([]int{2,1,3,4,4})) // 4
}