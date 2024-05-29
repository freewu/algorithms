package main

// 370. Range Addition
// You are given an integer length and an array updates where updates[i] = [startIdxi, endIdxi, inci].
// You have an array arr of length length with all zeros, and you have some operation to apply on arr. 
// In the ith operation, you should increment all the elements arr[startIdxi], arr[startIdxi + 1], ..., arr[endIdxi] by inci.

// Return arr after applying all the updates.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/27/rangeadd-grid.jpg" />
// Input: length = 5, updates = [[1,3,2],[2,4,3],[0,2,-2]]
// Output: [-2,0,3,5,3]

// Example 2:
// Input: length = 10, updates = [[2,4,6],[5,6,8],[1,9,-4]]
// Output: [0,-4,2,2,2,4,4,-4,-4,-4]
 
// Constraints:
//     1 <= length <= 10^5
//     0 <= updates.length <= 10^4
//     0 <= startIdxi <= endIdxi < length
//     -1000 <= inci <= 1000

import "fmt"

// 差分数组
func getModifiedArray(length int, updates [][]int) []int {
    res, diff := make([]int, length + 1), make([]int, length + 1)
    for i := 0; i < len(updates); i++ {
        proc := updates[i]
        diff[proc[0]] += proc[2]
        diff[proc[1]+1] -= proc[2]
    }
    for i := 1; i < length + 1; i++ {
        res[i] = res[i-1] + diff[i-1]
    }
    return res[1:]
}

func getModifiedArray1(length int, updates [][]int) []int {
    res, diff := make([]int, length), make([]int, length + 1)
    for _, update := range updates {
        diff[update[0]] += update[2]
        diff[update[1]+1] -= update[2]
    }
    res[0] = diff[0]
    for i := 1; i < len(diff) - 1; i++ {
        res[i] = res[i-1] + diff[i]
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/27/rangeadd-grid.jpg" />
    // Input: length = 5, updates = [[1,3,2],[2,4,3],[0,2,-2]]
    // Output: [-2,0,3,5,3]
    fmt.Println(getModifiedArray(5,[][]int{{1,3,2},{2,4,3},{0,2,-2}})) // [-2,0,3,5,3]
    // Example 2:
    // Input: length = 10, updates = [[2,4,6],[5,6,8],[1,9,-4]]
    // Output: [0,-4,2,2,2,4,4,-4,-4,-4]
    fmt.Println(getModifiedArray(10,[][]int{{2,4,6},{5,6,8},{1,9,-4}})) // [0,-4,2,2,2,4,4,-4,-4,-4]

    fmt.Println(getModifiedArray1(5,[][]int{{1,3,2},{2,4,3},{0,2,-2}})) // [-2,0,3,5,3]
    fmt.Println(getModifiedArray1(10,[][]int{{2,4,6},{5,6,8},{1,9,-4}})) // [0,-4,2,2,2,4,4,-4,-4,-4]
}