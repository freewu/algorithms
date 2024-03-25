package main

// 442. Find All Duplicates in an Array
// Given an integer array nums of length n where all the integers of nums are in the range [1, n] and each integer appears once or twice, return an array of all the integers that appears twice.
// You must write an algorithm that runs in O(n) time and uses only constant extra space.

// Example 1:
// Input: nums = [4,3,2,7,8,2,3,1]
// Output: [2,3]

// Example 2:
// Input: nums = [1,1,2]
// Output: [1]

// Example 3:
// Input: nums = [1]
// Output: []
 
// Constraints:
//     n == nums.length
//     1 <= n <= 10^5
//     1 <= nums[i] <= n
//     Each element in nums appears once or twice.

import "fmt"

// map
func findDuplicates(nums []int) []int {
    m := make(map[int]int)
    res :=[]int{}
    for _,v := range nums {
        m[v]++
        if m[v] >= 2 {
            res = append(res,v)
        }
    }
    return res 
}

func findDuplicates1(nums []int) []int {
    var res []int
    for _, num := range nums {
        if num < 0 {
            num *= -1
        }
        //     n == nums.length
        //     1 <= nums[i] <= n
        if nums[num - 1] > 0 {
            // 将下标的设置为负, 输入的为只为 1 <= n <= 10^5 
            nums[num - 1] *= -1
        } else { // 如果小于 0 说明已在有一次存在了
            res = append(res, num)
        }
    }
    return res
}

func main() {
    fmt.Println(findDuplicates([]int{4,3,2,7,8,2,3,1})) // [2,3]
    fmt.Println(findDuplicates([]int{1,1,2})) // [1]
    fmt.Println(findDuplicates([]int{1})) // []

    fmt.Println(findDuplicates1([]int{4,3,2,7,8,2,3,1})) // [2,3]
    fmt.Println(findDuplicates1([]int{1,1,2})) // [1]
    fmt.Println(findDuplicates1([]int{1})) // []
}