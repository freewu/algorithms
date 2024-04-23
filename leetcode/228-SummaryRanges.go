package main

// 228. Summary Ranges
// You are given a sorted unique integer array nums.
// A range [a,b] is the set of all integers from a to b (inclusive).
// Return the smallest sorted list of ranges that cover all the numbers in the array exactly. 
// That is, each element of nums is covered by exactly one of the ranges, 
// and there is no integer x such that x is in one of the ranges but not in nums.
// Each range [a,b] in the list should be output as:
//     "a->b" if a != b
//     "a" if a == b
 
// Example 1:
// Input: nums = [0,1,2,4,5,7]
// Output: ["0->2","4->5","7"]
// Explanation: The ranges are:
// [0,2] --> "0->2"
// [4,5] --> "4->5"
// [7,7] --> "7"

// Example 2:
// Input: nums = [0,2,3,4,6,8,9]
// Output: ["0","2->4","6","8->9"]
// Explanation: The ranges are:
// [0,0] --> "0"
// [2,4] --> "2->4"
// [6,6] --> "6"
// [8,9] --> "8->9"
 
// Constraints:
//     0 <= nums.length <= 20
//     -2^31 <= nums[i] <= 2^31 - 1
//     All the values of nums are unique.
//     nums is sorted in ascending order.

import "fmt"
import "strconv"

func summaryRanges(nums []int) []string {
    res := []string{}
    for i, n := 0, len(nums); i < n; {
        left := i
        for i++; i < n && nums[i-1]+1 == nums[i]; i++ { // 判断是否连续，连续就空转掉  nums[i-1]+1 == nums[i]
        }
        s := strconv.Itoa(nums[left])
        if left != i-1 { // 如果是连续的区间 取 left -> i - 1
            s += "->" + strconv.Itoa(nums[i-1]) // 记得是  nums[i-1]
        }
        res = append(res, s)
    }
    return res
}

func main() {
    fmt.Printf("summaryRanges([]int{0,1,2,4,5,7}) = %v\n",summaryRanges([]int{0,1,2,4,5,7})) // ["0->2","4->5","7"]
    fmt.Printf("summaryRanges([]int{0,2,3,4,6,8,9}) = %v\n",summaryRanges([]int{0,2,3,4,6,8,9})) // ["0","2->4","6","8->9"]
    fmt.Printf("summaryRanges([]int{}) = %v\n",summaryRanges([]int{})) // []
    fmt.Printf("summaryRanges([]int{-1}) = %v\n",summaryRanges([]int{-1})) // ["-1"]
    fmt.Printf("summaryRanges([]int{0}) = %v\n",summaryRanges([]int{0})) // ["0"]
}