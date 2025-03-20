package main

// 3488. Closest Equal Element Queries
// You are given a circular array nums and an array queries.

// For each query i, you have to find the following:
//     1. The minimum distance between the element at index queries[i] and any other index j in the circular array, where nums[j] == nums[queries[i]]. 
//        If no such index exists, the answer for that query should be -1.

// Return an array answer of the same size as queries, where answer[i] represents the result for query i.

// Example 1:
// Input: nums = [1,3,1,4,1,3,2], queries = [0,3,5]
// Output: [2,-1,3]
// Explanation:
// Query 0: The element at queries[0] = 0 is nums[0] = 1. The nearest index with the same value is 2, and the distance between them is 2.
// Query 1: The element at queries[1] = 3 is nums[3] = 4. No other index contains 4, so the result is -1.
// Query 2: The element at queries[2] = 5 is nums[5] = 3. The nearest index with the same value is 1, and the distance between them is 3 (following the circular path: 5 -> 6 -> 0 -> 1).

// Example 2:
// Input: nums = [1,2,3,4], queries = [0,1,2,3]
// Output: [-1,-1,-1,-1]
// Explanation:
// Each value in nums is unique, so no index shares the same value as the queried element. This results in -1 for all queries.

// Constraints:
//     1 <= queries.length <= nums.length <= 10^5
//     1 <= nums[i] <= 10^6
//     0 <= queries[i] < nums.length

import "fmt"

func solveQueries(nums []int, queries []int) []int {
    n := len(nums)
    left, right, vleft, vright := make([]int, n), make([]int, n), make(map[int]int, n), make(map[int]int, n)
    for i, j := 0,(n * 2) - 1; i < n * 2; i, j = i+1, j-1 {
        if v, ok :=vleft[nums[i % n]]; ok {
            if v != i % n { // If its not same element
                left[i % n] = i - v
                vleft[nums[i % n]] = i
            }
        } else {
            left[i % n] = -1
            vleft[nums[i % n]] = i
        }
        if v, ok := vright[nums[j % n]]; ok {
            if v != j + n { // If its not same element
                right[j % n] = v - j
                vright[nums[j % n]] = j
            }
        } else {
            right[j % n] = -1
            vright[nums[j % n]] = j
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res := make([]int, len(queries))
    for i, v := range queries {
        res[i] = min(left[v], right[v])
    }
    return res
}

func solveQueries1(nums []int, queries []int) []int {
    type Item struct { pre, first, end int }
    n := len(nums)
    mp, arr := make(map[int]Item), make([]int, n)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, v := range nums {
        if i == 0 {
            mp[v] = Item{ i, i, i }
            arr[i] = 1 << 31
            continue
        }
        if item, ok := mp[v]; ok {
            arr[i] = i - item.pre
            arr[item.pre] = min(arr[item.pre], arr[i])
            mp[v] = Item{ i, item.first, i }
        } else {
            mp[v] = Item{ i, i, i}
            arr[i] = 1 << 31
        }
    }
    res := make([]int, len(queries))
    for i, v := range queries {
        if item := mp[nums[v]]; item.first == item.end {
            arr[v] = -1
        } else if item.first == v || item.end == v {
            arr[v] = min(arr[v], item.first - item.end + n)
        }
        res[i] = arr[v]
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,1,4,1,3,2], queries = [0,3,5]
    // Output: [2,-1,3]
    // Explanation:
    // Query 0: The element at queries[0] = 0 is nums[0] = 1. The nearest index with the same value is 2, and the distance between them is 2.
    // Query 1: The element at queries[1] = 3 is nums[3] = 4. No other index contains 4, so the result is -1.
    // Query 2: The element at queries[2] = 5 is nums[5] = 3. The nearest index with the same value is 1, and the distance between them is 3 (following the circular path: 5 -> 6 -> 0 -> 1).
    fmt.Println(solveQueries([]int{1,3,1,4,1,3,2}, []int{0,3,5})) // [2,-1,3]
    // Example 2:
    // Input: nums = [1,2,3,4], queries = [0,1,2,3]
    // Output: [-1,-1,-1,-1]
    // Explanation:
    // Each value in nums is unique, so no index shares the same value as the queried element. This results in -1 for all queries.
    fmt.Println(solveQueries([]int{1,2,3,4}, []int{0,1,2,3})) // [-1,-1,-1,-1]

    fmt.Println(solveQueries([]int{0,1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]
    fmt.Println(solveQueries([]int{9,8,7,6,5,4,3,2,1,0}, []int{1,2,3,4,5,6,7,8,9})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]
    fmt.Println(solveQueries([]int{0,1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]
    fmt.Println(solveQueries([]int{9,8,7,6,5,4,3,2,1,0}, []int{9,8,7,6,5,4,3,2,1})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]

    fmt.Println(solveQueries1([]int{1,3,1,4,1,3,2}, []int{0,3,5})) // [2,-1,3]
    fmt.Println(solveQueries1([]int{1,2,3,4}, []int{0,1,2,3})) // [-1,-1,-1,-1]
    fmt.Println(solveQueries1([]int{0,1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]
    fmt.Println(solveQueries1([]int{9,8,7,6,5,4,3,2,1,0}, []int{1,2,3,4,5,6,7,8,9})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]
    fmt.Println(solveQueries1([]int{0,1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]
    fmt.Println(solveQueries1([]int{9,8,7,6,5,4,3,2,1,0}, []int{9,8,7,6,5,4,3,2,1})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]
}