package main

// 2672. Number of Adjacent Elements With the Same Color
// You are given an integer n representing an array colors of length n where all elements are set to 0's meaning uncolored. You are also given a 2D integer array queries where queries[i] = [indexi, colori]. 
// For the ith query:
//     1. Set colors[indexi] to colori.
//     2. Count adjacent pairs in colors set to the same color (regardless of colori).

// Return an array answer of the same length as queries where answer[i] is the answer to the ith query.

// Example 1:
// Input: n = 4, queries = [[0,2],[1,2],[3,1],[1,1],[2,1]]
// Output: [0,1,1,0,2]
// Explanation:
// Initially array colors = [0,0,0,0], where 0 denotes uncolored elements of the array.
// After the 1st query colors = [2,0,0,0]. The count of adjacent pairs with the same color is 0.
// After the 2nd query colors = [2,2,0,0]. The count of adjacent pairs with the same color is 1.
// After the 3rd query colors = [2,2,0,1]. The count of adjacent pairs with the same color is 1.
// After the 4th query colors = [2,1,0,1]. The count of adjacent pairs with the same color is 0.
// After the 5th query colors = [2,1,1,1]. The count of adjacent pairs with the same color is 2.

// Example 2:
// Input: n = 1, queries = [[0,100000]]
// Output: [0]
// Explanation:
// After the 1st query colors = [100000]. The count of adjacent pairs with the same color is 0.

// Constraints:
//     1 <= n <= 10^5
//     1 <= queries.length <= 10^5
//     queries[i].length == 2
//     0 <= indexi <= n - 1
//     1 <=  colori <= 10^5

import "fmt"

func colorTheArray(n int, queries [][]int) []int {
    res, nums := make([]int, len(queries)), make([]int, n)
    count := 0
    for i, q := range queries {
        j, color := q[0], q[1]
        if nums[j] == color {
            res[i] = count
            continue 
        }
        if j > 0 && nums[j] > 0 && nums[j - 1] == nums[j]     { count-- }
        if j < n - 1 && nums[j] > 0 && nums[j + 1] == nums[j] { count-- }
        if j > 0 && nums[j - 1] == color   { count++ }
        if j < n - 1 && nums[j + 1] == color { count++ }
        res[i], nums[j] = count, color
    }
    return res
}

func colorTheArray1(n int, queries [][]int) []int {
    res, arr := make([]int, len(queries)), make([]int, n + 2) // 避免讨论下标出界的情况
    count := 0
    for j, q := range queries {
        i, color := q[0] + 1, q[1] // 下标改成从 1 开始
        if arr[i] > 0 {
            if arr[i] == arr[i - 1] { count-- }
            if arr[i] == arr[i + 1] { count-- }
        }
        arr[i] = color
        if arr[i] == arr[i - 1] { count++ }
        if arr[i] == arr[i + 1] { count++ }
        res[j] = count
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, queries = [[0,2],[1,2],[3,1],[1,1],[2,1]]
    // Output: [0,1,1,0,2]
    // Explanation:
    // Initially array colors = [0,0,0,0], where 0 denotes uncolored elements of the array.
    // After the 1st query colors = [2,0,0,0]. The count of adjacent pairs with the same color is 0.
    // After the 2nd query colors = [2,2,0,0]. The count of adjacent pairs with the same color is 1.
    // After the 3rd query colors = [2,2,0,1]. The count of adjacent pairs with the same color is 1.
    // After the 4th query colors = [2,1,0,1]. The count of adjacent pairs with the same color is 0.
    // After the 5th query colors = [2,1,1,1]. The count of adjacent pairs with the same color is 2.
    fmt.Println(colorTheArray(4, [][]int{{0,2},{1,2},{3,1},{1,1},{2,1}})) // [0,1,1,0,2]
    // Example 2:
    // Input: n = 1, queries = [[0,100000]]
    // Output: [0]
    // Explanation:
    // After the 1st query colors = [100000]. The count of adjacent pairs with the same color is 0.
    fmt.Println(colorTheArray(1, [][]int{{0,100000}})) // [0]

    fmt.Println(colorTheArray1(4, [][]int{{0,2},{1,2},{3,1},{1,1},{2,1}})) // [0,1,1,0,2]
    fmt.Println(colorTheArray1(1, [][]int{{0,100000}})) // [0]
}