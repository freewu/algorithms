package main

// 2659. Make Array Empty
// You are given an integer array nums containing distinct numbers, 
// and you can perform the following operations until the array is empty:
//     1. If the first element has the smallest value, remove it
//     2. Otherwise, put the first element at the end of the array.

// Return an integer denoting the number of operations it takes to make nums empty.

// Example 1:
// Input: nums = [3,4,-1]
// Output: 5
// Operation	Array
//     1       [4, -1, 3]
//     2       [-1, 3, 4]
//     3       [3, 4]
//     4       [4]
//     5       []

// Example 2:
// Input: nums = [1,2,4,3]
// Output: 5
// Operation	Array
//     1       [2, 4, 3]
//     2       [4, 3]
//     3       [3, 4]
//     4       [4]
//     5       []

// Example 3:
// Input: nums = [1,2,3]
// Output: 3
// Operation	Array
//     1       [2, 3]
//     2       [3]
//     3       []
 
// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9
//     All values in nums are distinct.

import "fmt"
import "sort"

func countOperationsToEmptyArray(nums []int) int64 {
    res, n := len(nums), len(nums)
    index := make([][2]int, n)
    for i, v := range nums {
        index[i] = [2]int{ i + 1, v}
    }
    sort.Slice(index, func(i, j int) bool { 
        return index[i][1] < index[j][1] 
    })
    for i := 1; i < n; i++ {
        if index[i][0] < index[i-1][0] {
            res += (n - i)
        }
    }
    return int64(res)
}

func countOperationsToEmptyArray1(nums []int) int64 {
    n := len(nums)
    index := make([]int, n)
    for i := range index {
        index[i] = i
    }
    sort.Slice(index, func(i, j int) bool { 
        return nums[index[i]] < nums[index[j]] 
    })
    res := n
    for i := 0; i < n - 1; i++ {
        if index[i + 1] < index[i] {
            res += (n - i - 1)
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [3,4,-1]
    // Output: 5
    // Operation	Array
    //     1       [4, -1, 3]
    //     2       [-1, 3, 4]
    //     3       [3, 4]
    //     4       [4]
    //     5       []
    fmt.Println(countOperationsToEmptyArray([]int{3,4,-1})) // 5
    // Example 2:
    // Input: nums = [1,2,4,3]
    // Output: 5
    // Operation	Array
    //     1       [2, 4, 3]
    //     2       [4, 3]
    //     3       [3, 4]
    //     4       [4]
    //     5       []
    fmt.Println(countOperationsToEmptyArray([]int{1,2,4,3})) // 5
    // Example 3:
    // Input: nums = [1,2,3]
    // Output: 3
    // Operation	Array
    //     1       [2, 3]
    //     2       [3]
    //     3       []
    fmt.Println(countOperationsToEmptyArray([]int{1,2,3})) // 3

    fmt.Println(countOperationsToEmptyArray([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(countOperationsToEmptyArray([]int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(countOperationsToEmptyArray1([]int{3,4,-1})) // 5
    fmt.Println(countOperationsToEmptyArray1([]int{1,2,4,3})) // 5
    fmt.Println(countOperationsToEmptyArray1([]int{1,2,3})) // 3
    fmt.Println(countOperationsToEmptyArray1([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(countOperationsToEmptyArray1([]int{9,8,7,6,5,4,3,2,1})) // 45
}