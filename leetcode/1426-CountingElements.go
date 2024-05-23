package main

// 1426. Counting Elements
// Given an integer array arr, count how many elements x there are, such that x + 1 is also in arr. 
// If there are duplicates in arr, count them separately.

// Example 1:
// Input: arr = [1,2,3]
// Output: 2
// Explanation: 1 and 2 are counted cause 2 and 3 are in arr.

// Example 2:
// Input: arr = [1,1,3,3,5,5,7,7]
// Output: 0
// Explanation: No numbers are counted, cause there is no 2, 4, 6, or 8 in arr.
 
// Constraints:
//     1 <= arr.length <= 1000
//     0 <= arr[i] <= 1000

import "fmt"

// map
func countElements(arr []int) int {
    res, m := 0, make(map[int]int)
    for _, v := range arr {
        m[v]++
    }
    for _, x := range arr {
        if _, ok := m[x + 1]; ok { // 对于元素 x ，只有当 x + 1 也在数组 arr 里时，才能记为 1 个数
            res++
        }
    }
    return res
}

// array
func countElements1(arr []int) int {
    res, m := 0, make([]int,1002)
    for _, v := range arr {
        m[v]++
    }
    for _, x := range arr {
        if  m[x + 1] != 0 { // 对于元素 x ，只有当 x + 1 也在数组 arr 里时，才能记为 1 个数
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [1,2,3]
    // Output: 2
    // Explanation: 1 and 2 are counted cause 2 and 3 are in arr.
    fmt.Println(countElements([]int{1,2,3})) // 2
    // Example 2:
    // Input: arr = [1,1,3,3,5,5,7,7]
    // Output: 0
    // Explanation: No numbers are counted, cause there is no 2, 4, 6, or 8 in arr.
    fmt.Println(countElements([]int{1,1,3,3,5,5,7,7})) // 0

    fmt.Println(countElements1([]int{1,2,3})) // 2
    fmt.Println(countElements1([]int{1,1,3,3,5,5,7,7})) // 0
}