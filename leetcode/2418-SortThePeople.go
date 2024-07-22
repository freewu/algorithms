package main

// 2418. Sort the People
// You are given an array of strings names, and an array heights that consists of distinct positive integers. 
// Both arrays are of length n.

// For each index i, names[i] and heights[i] denote the name and height of the ith person.
// Return names sorted in descending order by the people's heights.

// Example 1:
// Input: names = ["Mary","John","Emma"], heights = [180,165,170]
// Output: ["Mary","Emma","John"]
// Explanation: Mary is the tallest, followed by Emma and John.

// Example 2:
// Input: names = ["Alice","Bob","Bob"], heights = [155,185,150]
// Output: ["Bob","Alice","Bob"]
// Explanation: The first Bob is the tallest, followed by Alice and the second Bob.

// Constraints:
//     n == names.length == heights.length
//     1 <= n <= 10^3
//     1 <= names[i].length <= 20
//     1 <= heights[i] <= 10^5
//     names[i] consists of lower and upper case English letters.
//     All the values of heights are distinct.

import "fmt"
import "slices"

func sortPeople(names []string, heights []int) []string {
    type pair struct {
        name string
        height int
    }
    res, arr := []string{}, []pair{}
    for i := range names {
        arr = append(arr, pair{names[i], heights[i]})
    }
    slices.SortFunc(arr, func(a, b pair) int { // 使用 slices 排序
        return b.height - a.height
    })
    for _, v := range arr {
        res = append(res, v.name)
    }
    return res
}

func main() {
    // Example 1:
    // Input: names = ["Mary","John","Emma"], heights = [180,165,170]
    // Output: ["Mary","Emma","John"]
    // Explanation: Mary is the tallest, followed by Emma and John.
    fmt.Println(sortPeople([]string{"Mary","John","Emma"},[]int{180,165,170})) // ["Mary","Emma","John"]
    // Example 2:
    // Input: names = ["Alice","Bob","Bob"], heights = [155,185,150]
    // Output: ["Bob","Alice","Bob"]
    // Explanation: The first Bob is the tallest, followed by Alice and the second Bob.
    fmt.Println(sortPeople([]string{"Alice","Bob","Bob"},[]int{155,185,150})) // ["Bob","Alice","Bob"]
}