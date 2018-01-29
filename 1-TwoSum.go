package main
/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

import "fmt"

func twoSum(nums []int, target int) []int {
    var l = len(nums)

    if l < 2 {
        return nil
    }
    if 2 == l && ((nums[0] + nums[1]) != target) {
        return nil
    }
    
    for i := 0; i < l; i++  {
        for j := i + 1; j < l; j++ {
            if (nums[i] + nums[j]) == target {
                return []int{i,j}
            }
        }
    }
    return nil
}

func main() {
    var list = []int{2, 7, 11, 15}
    fmt.Println(twoSum(list,9))
}