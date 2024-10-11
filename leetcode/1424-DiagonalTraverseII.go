package main

// 1424. Diagonal Traverse II
// Given a 2D integer array nums, return all elements of nums in diagonal order as shown in the below images.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/04/08/sample_1_1784.png" />
// Input: nums = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,4,2,7,5,3,8,6,9]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/04/08/sample_2_1784.png" />
// Input: nums = [[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]
// Output: [1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i].length <= 10^5
//     1 <= sum(nums[i].length) <= 10^5
//     1 <= nums[i][j] <= 10^5

import "fmt"

func findDiagonalOrder(nums [][]int) []int {
    triplets := map[int][]int{}
    n := len(nums)
    for i := n-1; i >= 0; i-- {
        for j := range nums[i] {
            triplets[i+j] = append(triplets[i+j], nums[i][j])
        }
    }
    res := []int{}
    for i := 0; i < len(triplets); i++ {
        res = append(res, triplets[i]...)
    }
    return res
}

func findDiagonalOrder1(nums [][]int) []int {
    n, path, res := len(nums), [][]int{}, []int{}
    for i := 0; i < n; i++ {
        for j := 0; j <len(nums[i]); j++{
            if len(path) <= i+j {
                curpath := []int{nums[i][j]}
                path = append(path, curpath)
            } else {
                path[i+j] = append(path[i+j], nums[i][j])
            }
        }
    }
    for _, row := range path {
        for i := len(row)-1; i >= 0; i-- {
            res = append(res, row[i])
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/04/08/sample_1_1784.png" />
    // Input: nums = [[1,2,3],[4,5,6],[7,8,9]]
    // Output: [1,4,2,7,5,3,8,6,9]
    fmt.Println(findDiagonalOrder([][]int{{1,2,3},{4,5,6},{7,8,9}})) // [1,4,2,7,5,3,8,6,9]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/04/08/sample_2_1784.png" />
    // Input: nums = [[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]
    // Output: [1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]
    fmt.Println(findDiagonalOrder([][]int{{1,2,3,4,5},{6,7},{8},{9,10,11},{12,13,14,15,16}})) // [1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]

    fmt.Println(findDiagonalOrder1([][]int{{1,2,3},{4,5,6},{7,8,9}})) // [1,4,2,7,5,3,8,6,9]
    fmt.Println(findDiagonalOrder1([][]int{{1,2,3,4,5},{6,7},{8},{9,10,11},{12,13,14,15,16}})) // [1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]
}