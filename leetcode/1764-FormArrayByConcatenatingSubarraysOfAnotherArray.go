package main

// 1764. Form Array by Concatenating Subarrays of Another Array
// You are given a 2D integer array groups of length n. 
// You are also given an integer array nums.

// You are asked if you can choose n disjoint subarrays from the array nums such that the ith subarray is equal to groups[i] (0-indexed), 
// and if i > 0, the (i-1)th subarray appears before the ith subarray in nums 
// (i.e. the subarrays must be in the same order as groups).

// Return true if you can do this task, and false otherwise.

// Note that the subarrays are disjoint if and only if there is no index k such that nums[k] belongs to more than one subarray.
// A subarray is a contiguous sequence of elements within an array.

// Example 1:
// Input: groups = [[1,-1,-1],[3,-2,0]], nums = [1,-1,0,1,-1,-1,3,-2,0]
// Output: true
// Explanation: You can choose the 0th subarray as [1,-1,0,1,-1,-1,3,-2,0] and the 1st one as [1,-1,0,1,-1,-1,3,-2,0].
// These subarrays are disjoint as they share no common nums[k] element.

// Example 2:
// Input: groups = [[10,-2],[1,2,3,4]], nums = [1,2,3,4,10,-2]
// Output: false
// Explanation: Note that choosing the subarrays [1,2,3,4,10,-2] and [1,2,3,4,10,-2] is incorrect because they are not in the same order as in groups.
// [10,-2] must come before [1,2,3,4].

// Example 3:
// Input: groups = [[1,2,3],[3,4]], nums = [7,7,1,2,3,4,7,7]
// Output: false
// Explanation: Note that choosing the subarrays [7,7,1,2,3,4,7,7] and [7,7,1,2,3,4,7,7] is invalid because they are not disjoint.
// They share a common elements nums[4] (0-indexed).

// Constraints:
//     groups.length == n
//     1 <= n <= 10^3
//     1 <= groups[i].length, sum(groups[i].length) <= 10^3
//     1 <= nums.length <= 10^3
//     -10^7 <= groups[i][j], nums[k] <= 10^7

import "fmt"
import "reflect"

func canChoose(groups [][]int, nums []int) bool {
    i, n := 0, len(nums)
    for _, group := range groups {
        j, m := i, len(group)
        for ; j <= n - m; j++ {
            if reflect.DeepEqual(nums[j:j + m], group) {
                i = j + m
                break
            }
        }
        if i != j + m { return false }
    }
    return true
}

func canChoose1(groups [][]int, nums []int) bool {
    calcPi := func(pattern []int) []int {
        pi, match := make([]int, len(pattern)), 0
        for i := 1; i < len(pattern); i++ {
            v := pattern[i]
            for match > 0 && pattern[match] != v {
                match = pi[match-1]
            }
            if pattern[match] == v {
                match++
            }
            pi[i] = match
        }
        return pi
    }
    kmpSearch := func(text, pattern []int) int {
        if len(pattern) > len(text) { return -1 }
        pi, match := calcPi(pattern), 0
        for i := range text {
            v := text[i]
            for match > 0 && pattern[match] != v {
                match = pi[match-1]
            }
            if pattern[match] == v { match++ }
            if match == len(pattern) {
                return i - len(pattern) + 1
            }
        }
        return -1
    }
    k := 0
    for _, g := range groups {
        k = kmpSearch(nums[k:], g)
        if k == -1 { return false }
        k += len(g)
    }
    return true
}

func main() {
    // Example 1:
    // Input: groups = [[1,-1,-1],[3,-2,0]], nums = [1,-1,0,1,-1,-1,3,-2,0]
    // Output: true
    // Explanation: You can choose the 0th subarray as [1,-1,0,1,-1,-1,3,-2,0] and the 1st one as [1,-1,0,1,-1,-1,3,-2,0].
    // These subarrays are disjoint as they share no common nums[k] element.
    fmt.Println(canChoose([][]int{{1,-1,-1},{3,-2,0}}, []int{1,-1,0,1,-1,-1,3,-2,0})) // true
    // Example 2:
    // Input: groups = [[10,-2],[1,2,3,4]], nums = [1,2,3,4,10,-2]
    // Output: false
    // Explanation: Note that choosing the subarrays [1,2,3,4,10,-2] and [1,2,3,4,10,-2] is incorrect because they are not in the same order as in groups.
    // [10,-2] must come before [1,2,3,4].
    fmt.Println(canChoose([][]int{{10,-2},{1,2,3,4}}, []int{1,2,3,4,10,-2})) // false
    // Example 3:
    // Input: groups = [[1,2,3],[3,4]], nums = [7,7,1,2,3,4,7,7]
    // Output: false
    // Explanation: Note that choosing the subarrays [7,7,1,2,3,4,7,7] and [7,7,1,2,3,4,7,7] is invalid because they are not disjoint.
    // They share a common elements nums[4] (0-indexed).
    fmt.Println(canChoose([][]int{{1,2,3},{3,4}}, []int{7,7,1,2,3,4,7,7})) // false

    fmt.Println(canChoose1([][]int{{1,-1,-1},{3,-2,0}}, []int{1,-1,0,1,-1,-1,3,-2,0})) // true
    fmt.Println(canChoose1([][]int{{10,-2},{1,2,3,4}}, []int{1,2,3,4,10,-2})) // false
    fmt.Println(canChoose1([][]int{{1,2,3},{3,4}}, []int{7,7,1,2,3,4,7,7})) // false
}