package main

// 2610. Convert an Array Into a 2D Array With Conditions
// You are given an integer array nums. 
// You need to create a 2D array from nums satisfying the following conditions:
//     1. The 2D array should contain only the elements of the array nums.
//     2. Each row in the 2D array contains distinct integers.
//     3. The number of rows in the 2D array should be minimal.

// Return the resulting array. 
// If there are multiple answers, return any of them.

// Note that the 2D array can have a different number of elements on each row.

// Example 1:
// Input: nums = [1,3,4,1,2,3,1]
// Output: [[1,3,4,2],[1,3],[1]]
// Explanation: We can create a 2D array that contains the following rows:
// - 1,3,4,2
// - 1,3
// - 1
// All elements of nums were used, and each row of the 2D array contains distinct integers, so it is a valid answer.
// It can be shown that we cannot have less than 3 rows in a valid array.

// Example 2:
// Input: nums = [1,2,3,4]
// Output: [[4,3,2,1]]
// Explanation: All elements of the array are distinct, so we can keep all of them in the first row of the 2D array.

// Constraints:
//     1 <= nums.length <= 200
//     1 <= nums[i] <= nums.length

import "fmt"

func findMatrix(nums []int) [][]int {
    res, freq := [][]int{}, make(map[int]int)
    for _, v := range nums {
        if freq[v] >= len(res) {
            res = append(res,[]int{})
        }
        res[freq[v]] = append(res[freq[v]], v)
        freq[v]++
    }
    return res
}

func findMatrix1(nums []int) [][]int {
    res, freq := [][]int{}, make(map[int]int)
    for _, v := range nums {
        freq[v]++
    }
    for len(freq) != 0 {
        arr := []int{}
        for v := range freq {
            arr = append(arr, v)
        }
        for _, v := range arr {
            freq[v]--
            if freq[v] == 0 {
                delete(freq, v)
            }
        }
        res = append(res, arr)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,4,1,2,3,1]
    // Output: [[1,3,4,2],[1,3],[1]]
    // Explanation: We can create a 2D array that contains the following rows:
    // - 1,3,4,2
    // - 1,3
    // - 1
    // All elements of nums were used, and each row of the 2D array contains distinct integers, so it is a valid answer.
    // It can be shown that we cannot have less than 3 rows in a valid array.
    fmt.Println(findMatrix([]int{1,3,4,1,2,3,1})) // [[1,3,4,2],[1,3],[1]]
    // Example 2:
    // Input: nums = [1,2,3,4]
    // Output: [[4,3,2,1]]
    // Explanation: All elements of the array are distinct, so we can keep all of them in the first row of the 2D array.
    fmt.Println(findMatrix([]int{1,2,3,4})) // [[4,3,2,1]]

    fmt.Println(findMatrix([]int{1,2,3,4,5,6,7,8,9})) // [[1 2 3 4 5 6 7 8 9]]
    fmt.Println(findMatrix([]int{9,8,7,6,5,4,3,2,1})) // [[9 8 7 6 5 4 3 2 1]]

    fmt.Println(findMatrix1([]int{1,3,4,1,2,3,1})) // [[1,3,4,2],[1,3],[1]]
    fmt.Println(findMatrix1([]int{1,2,3,4})) // [[4,3,2,1]]
    fmt.Println(findMatrix1([]int{1,2,3,4,5,6,7,8,9})) // [[1 2 3 4 5 6 7 8 9]]
    fmt.Println(findMatrix1([]int{9,8,7,6,5,4,3,2,1})) // [[9 8 7 6 5 4 3 2 1]]
}