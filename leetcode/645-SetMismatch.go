package main

// 645. Set Mismatch
// You have a set of integers s, which originally contains all the numbers from 1 to n. 
// Unfortunately, due to some error, one of the numbers in s got duplicated to another number in the set, which results in repetition of one number and loss of another number.
// You are given an integer array nums representing the data status of this set after the error.
// Find the number that occurs twice and the number that is missing and return them in the form of an array.

// Example 1:
// Input: nums = [1,2,2,4]
// Output: [2,3]

// Example 2:
// Input: nums = [1,1]
// Output: [1,2]

// Constraints:
//     2 <= nums.length <= 10^4
//     1 <= nums[i] <= 10^4

import "fmt"

func findErrorNums(nums []int) []int {
    m, res := make([]int, len(nums)), make([]int, 2)
    // 把列表放入到一个 map 中
    for _, n := range nums {
        if m[n-1] == 0 { // 初始化都为 0 
            m[n-1] = 1
        } else {
            res[0] = n // 如果不为0说明已出现过，就是我们要找的
        }
    }
    // 找出的数字
    for i := range m {
        if m[i] == 0 {
            res[1] = i + 1
            break
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,4]
    // Output: [2,3]
    fmt.Println(findErrorNums([]int{1,2,2,4})) // [2 3]
    // Example 2:
    // Input: nums = [1,1]
    // Output: [1,2]
    fmt.Println(findErrorNums([]int{1,1})) // [1 2]

    fmt.Println(findErrorNums([]int{1,2,3,4,5,6,7,8,9})) // [0 0]
    fmt.Println(findErrorNums([]int{9,8,7,6,5,4,3,2,1})) // [0 0]
}