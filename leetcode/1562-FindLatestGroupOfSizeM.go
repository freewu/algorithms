package main

// 1562. Find Latest Group of Size M
// Given an array arr that represents a permutation of numbers from 1 to n.

// You have a binary string of size n that initially has all its bits set to zero. 
// At each step i (assuming both the binary string and arr are 1-indexed) from 1 to n, the bit at position arr[i] is set to 1.

// You are also given an integer m. 
// Find the latest step at which there exists a group of ones of length m. 
// A group of ones is a contiguous substring of 1's such that it cannot be extended in either direction.

// Return the latest step at which there exists a group of ones of length exactly m. 
// If no such group exists, return -1.

// Example 1:
// Input: arr = [3,5,1,2,4], m = 1
// Output: 4
// Explanation: 
// Step 1: "00100", groups: ["1"]
// Step 2: "00101", groups: ["1", "1"]
// Step 3: "10101", groups: ["1", "1", "1"]
// Step 4: "11101", groups: ["111", "1"]
// Step 5: "11111", groups: ["11111"]
// The latest step at which there exists a group of size 1 is step 4.

// Example 2:
// Input: arr = [3,1,5,4,2], m = 2
// Output: -1
// Explanation: 
// Step 1: "00100", groups: ["1"]
// Step 2: "10100", groups: ["1", "1"]
// Step 3: "10101", groups: ["1", "1", "1"]
// Step 4: "10111", groups: ["1", "111"]
// Step 5: "11111", groups: ["11111"]
// No group of size 2 exists during any step.

// Constraints:
//     n == arr.length
//     1 <= m <= n <= 10^5
//     1 <= arr[i] <= n
//     All integers in arr are distinct.

import "fmt"

func findLatestStep(arr []int, m int) int {
    mp, list := make(map[int]int), make([]int, len(arr) + 1)
    res, count := -1, 0
    find := func(arr []int, k int) int {
        for arr[k] != k {
            k = arr[k]
        }
        return k
    }
    for i, v := range arr {
        list[v] = v
        if v == 1 || list[v - 1] == 0 {
            mp[v] = 1
            if m == 1 {
                count++
            }
        }
        if v - 1 >= 0 && list[v - 1] > 0 {
            list[v] = find(list, v - 1)
            if mp[list[v]] == m {
                count--
            }
            mp[list[v]]++
            if mp[list[v]] == m {
                count++
            }
        }
        if v <= len(arr) - 1 && list[v + 1] > 0 {
            t := mp[list[v + 1]]
            if t == m {
                count--
            }
            delete(mp, list[v + 1])
            list[v + 1] = list[v]
            if mp[list[v]] == m {
                count--
            }
            mp[list[v]] += t
            if mp[list[v]] == m {
                count++
            }
        }
        if count > 0 {
            res = i + 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [3,5,1,2,4], m = 1
    // Output: 4
    // Explanation: 
    // Step 1: "00100", groups: ["1"]
    // Step 2: "00101", groups: ["1", "1"]
    // Step 3: "10101", groups: ["1", "1", "1"]
    // Step 4: "11101", groups: ["111", "1"]
    // Step 5: "11111", groups: ["11111"]
    // The latest step at which there exists a group of size 1 is step 4.
    fmt.Println(findLatestStep([]int{3,5,1,2,4}, 1)) // 4
    // Example 2:
    // Input: arr = [3,1,5,4,2], m = 2
    // Output: -1
    // Explanation: 
    // Step 1: "00100", groups: ["1"]
    // Step 2: "10100", groups: ["1", "1"]
    // Step 3: "10101", groups: ["1", "1", "1"]
    // Step 4: "10111", groups: ["1", "111"]
    // Step 5: "11111", groups: ["11111"]
    // No group of size 2 exists during any step.
    fmt.Println(findLatestStep([]int{3,1,5,4,2}, 2)) // -1
}