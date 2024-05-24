package main

// 2248. Intersection of Multiple Arrays
// Given a 2D integer array nums where nums[i] is a non-empty array of distinct positive integers, 
// return the list of integers that are present in each array of nums sorted in ascending order.
 
// Example 1:
// Input: nums = [[3,1,2,4,5],[1,2,3,4],[3,4,5,6]]
// Output: [3,4]
// Explanation: 
// The only integers present in each of nums[0] = [3,1,2,4,5], nums[1] = [1,2,3,4], and nums[2] = [3,4,5,6] are 3 and 4, so we return [3,4].

// Example 2:
// Input: nums = [[1,2,3],[4,5,6]]
// Output: []
// Explanation: 
// There does not exist any integer present both in nums[0] and nums[1], so we return an empty list [].
 
// Constraints:
//     1 <= nums.length <= 1000
//     1 <= sum(nums[i].length) <= 1000
//     1 <= nums[i][j] <= 1000
//     All the values of nums[i] are unique.

import "fmt"
import "sort"

func intersection(nums [][]int) []int {
    n, res, m := len(nums), []int{}, make(map[int]int) 
    for i := 0; i < n; i++ { // 遍历所有数组写到 map 中
        for j := 0; j < len(nums[i]); j++ {
            m[nums[i][j]]++
        }
    }
    for k, v := range m {
        if v == n { // 所有数组存在
            res = append(res, k)
        }
    }  
    sort.Ints(res)
    return res  
}

// 使用数组
func intersection1(nums [][]int) []int {
    n, res, arr := len(nums), []int{}, make([]int, 1001) // 1 <= nums[i][j] <= 1000
    for _, rows := range nums {
        for _, v := range rows {
            arr[v]++
        }
    }
    for k, v := range arr {
        if v == n { // 所有数组存在
            res = append(res, k)
        }
    }
    return res  
}

func main() {
    // Example 1:
    // Input: nums = [[3,1,2,4,5],[1,2,3,4],[3,4,5,6]]
    // Output: [3,4]
    // Explanation: 
    // The only integers present in each of nums[0] = [3,1,2,4,5], nums[1] = [1,2,3,4], and nums[2] = [3,4,5,6] are 3 and 4, so we return [3,4].
    fmt.Println(intersection([][]int{{3,1,2,4,5},{1,2,3,4},{3,4,5,6}})) // [3,4]
    // Example 2:
    // Input: nums = [[1,2,3],[4,5,6]]
    // Output: []
    // Explanation: 
    // There does not exist any integer present both in nums[0] and nums[1], so we return an empty list [].
    fmt.Println(intersection([][]int{{1,2,3},{4,5,6}})) // []

    fmt.Println(intersection1([][]int{{3,1,2,4,5},{1,2,3,4},{3,4,5,6}})) // [3,4]
    fmt.Println(intersection1([][]int{{1,2,3},{4,5,6}})) // []
}