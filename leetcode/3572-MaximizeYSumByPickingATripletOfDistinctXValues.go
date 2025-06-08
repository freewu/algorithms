package main

// 3572. Maximize Y‑Sum by Picking a Triplet of Distinct X‑Values
// You are given two integer arrays x and y, each of length n. 
// You must choose three distinct indices i, j, and k such that:
//     x[i] != x[j]
//     x[j] != x[k]
//     x[k] != x[i]

// Your goal is to maximize the value of y[i] + y[j] + y[k] under these conditions. 
// Return the maximum possible sum that can be obtained by choosing such a triplet of indices.

// If no such triplet exists, return -1.

// Example 1:
// Input: x = [1,2,1,3,2], y = [5,3,4,6,2]
// Output: 14
// Explanation:
// Choose i = 0 (x[i] = 1, y[i] = 5), j = 1 (x[j] = 2, y[j] = 3), k = 3 (x[k] = 3, y[k] = 6).
// All three values chosen from x are distinct. 5 + 3 + 6 = 14 is the maximum we can obtain. 
// Hence, the output is 14.

// Example 2:
// Input: x = [1,2,1,2], y = [4,5,6,7]
// Output: -1
// Explanation:
// There are only two distinct values in x. Hence, the output is -1.

// Constraints:
//     n == x.length == y.length
//     3 <= n <= 10^5
//     1 <= x[i], y[i] <= 10^6

import "fmt"
import "sort"

func maxSumDistinctTriplet(x []int, y []int) int {
    type Pair struct { index, value int }
    arr := []Pair{}
    for i, v := range y {
        arr = append(arr, Pair{ i, v })
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].value > arr[j].value
    })
    i, j := 0, 1
    for j < len(x) && x[arr[i].index] == x[arr[j].index] {
        j++
    }
    if j == len(x) {
        return -1
    }
    k := j + 1
    for k < len(x) && (x[arr[k].index] == x[arr[j].index] || x[arr[k].index] == x[arr[i].index]) {
        k++
    }
    if k == len(x) {
        return -1
    }
    return arr[i].value + arr[j].value + arr[k].value
}

func maxSumDistinctTriplet1(x []int, y []int) int {
    mp := make(map[int]int)
    for i, v1 := range x {
        if v2, ok := mp[v1]; !ok || y[i] > v2 {
            mp[v1] = y[i]
        }
    }
    arr := make([]int, 0, len(mp))
    for _, v := range mp {
        arr = append(arr, v)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(arr)))
    if len(arr) < 3 {
        return -1
    }
    return arr[0] + arr[1] + arr[2]
}

func main() {
    // Example 1:
    // Input: x = [1,2,1,3,2], y = [5,3,4,6,2]
    // Output: 14
    // Explanation:
    // Choose i = 0 (x[i] = 1, y[i] = 5), j = 1 (x[j] = 2, y[j] = 3), k = 3 (x[k] = 3, y[k] = 6).
    // All three values chosen from x are distinct. 5 + 3 + 6 = 14 is the maximum we can obtain. 
    // Hence, the output is 14.
    fmt.Println(maxSumDistinctTriplet([]int{1,2,1,3,2}, []int{5,3,4,6,2})) // 14
    // Example 2:
    // Input: x = [1,2,1,2], y = [4,5,6,7]
    // Output: -1
    // Explanation:
    // There are only two distinct values in x. Hence, the output is -1.
    fmt.Println(maxSumDistinctTriplet([]int{1,2,1,2}, []int{4,5,6,7})) // -1

    fmt.Println(maxSumDistinctTriplet([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 24
    fmt.Println(maxSumDistinctTriplet([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 24
    fmt.Println(maxSumDistinctTriplet([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 24
    fmt.Println(maxSumDistinctTriplet([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 24

    fmt.Println(maxSumDistinctTriplet1([]int{1,2,1,3,2}, []int{5,3,4,6,2})) // 14
    fmt.Println(maxSumDistinctTriplet1([]int{1,2,1,2}, []int{4,5,6,7})) // -1
    fmt.Println(maxSumDistinctTriplet1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 24
    fmt.Println(maxSumDistinctTriplet1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 24
    fmt.Println(maxSumDistinctTriplet1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 24
    fmt.Println(maxSumDistinctTriplet1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 24
}