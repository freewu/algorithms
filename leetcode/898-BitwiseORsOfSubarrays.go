package main

// 898. Bitwise ORs of Subarrays
// Given an integer array arr, return the number of distinct bitwise ORs of all the non-empty subarrays of arr.

// The bitwise OR of a subarray is the bitwise OR of each integer in the subarray. 
// The bitwise OR of a subarray of one integer is that integer.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: arr = [0]
// Output: 1
// Explanation: There is only one possible result: 0.

// Example 2:
// Input: arr = [1,1,2]
// Output: 3
// Explanation: The possible subarrays are [1], [1], [2], [1, 1], [1, 2], [1, 1, 2].
// These yield the results 1, 1, 2, 1, 3, 3.
// There are 3 unique values, so the answer is 3.

// Example 3:
// Input: arr = [1,2,4]
// Output: 6
// Explanation: The possible results are 1, 2, 3, 4, 6, and 7.

// Constraints:
//     1 <= arr.length <= 5 * 10^4
//     0 <= arr[i] <= 10^9

import "fmt"

func subarrayBitwiseORs(arr []int) int {
    res, mp, cache := 0, make(map[int]bool), []int{}
    for _, v := range arr {
        if !mp[v] {
            res++
            mp[v] = true
        }
        t := []int{v}
        for i := 0; i < len(cache); i++ {
            cache[i] |= v
            if !mp[cache[i]] {
                res++
                mp[cache[i]] = true
            }
            if t[len(t)-1] != cache[i] {
                t = append(t, cache[i])
            }
        }
        cache = t
    }
    return res
}

func subarrayBitwiseORs1(arr []int) int {
    set := map[int]struct{}{}
    for i, v := range arr {
        set[v] = struct{}{}
        for j := i-1; j >= 0 && arr[j]|v != arr[j]; j-- {
            arr[j] |= v
            set[arr[j]] = struct{}{}
        }
    }
    return len(set)
}

func main() {
    // Example 1:
    // Input: arr = [0]
    // Output: 1
    // Explanation: There is only one possible result: 0.
    fmt.Println(subarrayBitwiseORs([]int{0})) // 1
    // Example 2:
    // Input: arr = [1,1,2]
    // Output: 3
    // Explanation: The possible subarrays are [1], [1], [2], [1, 1], [1, 2], [1, 1, 2].
    // These yield the results 1, 1, 2, 1, 3, 3.
    // There are 3 unique values, so the answer is 3.
    fmt.Println(subarrayBitwiseORs([]int{1,1,2})) // 3
    // Example 3:
    // Input: arr = [1,2,4]
    // Output: 6
    // Explanation: The possible results are 1, 2, 3, 4, 6, and 7.
    fmt.Println(subarrayBitwiseORs([]int{1,2,4})) // 6

    fmt.Println(subarrayBitwiseORs1([]int{0})) // 1
    fmt.Println(subarrayBitwiseORs1([]int{1,1,2})) // 3
    fmt.Println(subarrayBitwiseORs1([]int{1,2,4})) // 6
}