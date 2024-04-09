package main

// 448. Find All Numbers Disappeared in an Array
// Given an array nums of n integers where nums[i] is in the range [1, n], 
// return an array of all the integers in the range [1, n] that do not appear in nums.

// Example 1:
// Input: nums = [4,3,2,7,8,2,3,1]
// Output: [5,6]

// Example 2:
// Input: nums = [1,1]
// Output: [2]

// Constraints:
//     n == nums.length
//     1 <= n <= 10^5
//     1 <= nums[i] <= n
    
// Follow up: Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

import "fmt"

// Space: O(n) Time: O(n)
func findDisappearedNumbers(nums []int) []int {
    m := make(map[int]int,len(nums)) 
    for _, v := range nums {
        m[v]++
    } 
    res := []int{}
    for i := 1; i < len(nums) + 1; i++ {
        if _, ok := m[i]; !ok {
            res = append(res, i)
        }
    }
    return res
}

func findDisappearedNumbers1(nums []int) []int {
    n := len(nums) + 1
    for _, v := range nums {
        // fmt.Printf("n= %v v = %v v %% n-1 = %v \n",n, v, v % n-1)
        // 这步太骚了
        // 通过取(n-1)的 余数 v 一定是在 [1-n] 之间 加 + n
        // 如果是丢失了数字的下标 + 不 n,下面遍历的时候 < n 就是丢失的数字了
        nums[v % n-1] += n
    }
    fmt.Println(nums)
    res := []int{}
    for i, v := range nums {
        if v < n {
            res = append(res, i+1)
        }
    }
    return res
}

func main() {
    fmt.Println(findDisappearedNumbers([]int{4,3,2,7,8,2,3,1})) // [5,6]
    fmt.Println(findDisappearedNumbers([]int{1,1})) // [2]

    fmt.Println(findDisappearedNumbers1([]int{4,3,2,7,8,2,3,1})) // [5,6]
    fmt.Println(findDisappearedNumbers1([]int{1,1})) // [2]
}