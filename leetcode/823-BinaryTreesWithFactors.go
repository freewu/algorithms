package main

// 823. Binary Trees With Factors
// Given an array of unique integers, arr, where each integer arr[i] is strictly greater than 1.

// We make a binary tree using these integers, and each number may be used for any number of times. 
// Each non-leaf node's value should be equal to the product of the values of its children.

// Return the number of binary trees we can make. The answer may be too large so return the answer modulo 10^9 + 7.

// Example 1:
// Input: arr = [2,4]
// Output: 3
// Explanation: We can make these trees: [2], [4], [4, 2, 2]

// Example 2:
// Input: arr = [2,4,5,10]
// Output: 7
// Explanation: We can make these trees: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].

// Constraints:
//     1 <= arr.length <= 1000
//     2 <= arr[i] <= 10^9
//     All the values of arr are unique.

import "fmt"
import "sort"

func numFactoredBinaryTrees(arr []int) int {
    res, mod, mp := 0, 1_000_000_007, make(map[int]int)
    sort.Ints(arr)
    for i, v := range arr {
        mp[v] += 1
        for j := 0; j <= i; j++ {
            if arr[j] * arr[j] > v {
                break
            }
            k := arr[j]
            if v % k == 0 {
                a := v / k
                v1, ok1 := mp[k]
                v2, ok2 := mp[a]
                if ok1 && ok2 {
                    if k == a {
                        mp[v] += v1 * v2 % mod
                    } else {
                        mp[v] += 2 * v1 * v2 % mod
                    }
                }
            }
        }
    }
    for _, v := range mp {
        res = (res + v) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [2,4]
    // Output: 3
    // Explanation: We can make these trees: [2], [4], [4, 2, 2]
    fmt.Println(numFactoredBinaryTrees([]int{2,4})) // 3   [2], [4], [4, 2, 2]
    // Example 2:
    // Input: arr = [2,4,5,10]
    // Output: 7
    // Explanation: We can make these trees: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].
    fmt.Println(numFactoredBinaryTrees([]int{2,4,5,10})) // 7 [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].
}