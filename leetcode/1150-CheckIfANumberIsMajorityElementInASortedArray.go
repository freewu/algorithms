package main

// 1150. Check If a Number Is Majority Element in a Sorted Array
// Given an integer array nums sorted in non-decreasing order and an integer target, 
// return true if target is a majority element, or false otherwise.

// A majority element in an array nums is an element that appears more than nums.length / 2 times in the array.

// Example 1:
// Input: nums = [2,4,5,5,5,5,5,6,6], target = 5
// Output: true
// Explanation: The value 5 appears 5 times and the length of the array is 9.
// Thus, 5 is a majority element because 5 > 9/2 is true.

// Example 2:
// Input: nums = [10,100,101,101], target = 101
// Output: false
// Explanation: The value 101 appears 2 times and the length of the array is 4.
// Thus, 101 is not a majority element because 2 > 4/2 is false.
 
// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i], target <= 10^9
//     nums is sorted in non-decreasing order.

import "fmt"
import "sort"

// O(logn)
func isMajorityElement(nums []int, target int) bool {
    n, index := len(nums), sort.SearchInts(nums, target) // 得到 target 出现的第一个位置
    //fmt.Println("index: ", index)
    // 判断 从 target 出现的起始位置 + 上一半数组长度的 的值 是否还是 target
    return index < n && nums[index] == target && index + n / 2 < n && nums[index + n/2] == target
}

// O(n)
func isMajorityElement1(nums []int, target int) bool {
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == target { // 找到目标累加
            count++
        }
    }
    return count > len(nums) / 2 // 判断是否超过一半数量
}

func isMajorityElement2(nums []int, target int) bool {
    searchLeft := func (nums []int, target int) int {
        l, r := 0, len(nums) - 1
        for l <= r {
            mid := (l + r)/2
            if nums[mid] >= target {
                r = mid -1
            }else {
                l = mid+1
            }
        }
        if l < len(nums) && nums[l] == target {
            return l
        }
        return -1
    }
    searchRight := func (nums []int, target int) int {
        l, r := 0, len(nums) - 1
        for l <= r {
            mid := (l+r)/2
            if nums[mid] <= target {
                l = mid + 1
            }else {
                r = mid - 1
            }
        }
        if r >= 0 && nums[r] == target {
            return r
        }
        return r - 1
    }
    left, right := searchLeft(nums, target), searchRight(nums, target)
    return left != -1 && right != -1 && right - left + 1 > len(nums) / 2
}

func main() {
    // Example 1:
    // Input: nums = [2,4,5,5,5,5,5,6,6], target = 5
    // Output: true
    // Explanation: The value 5 appears 5 times and the length of the array is 9.
    // Thus, 5 is a majority element because 5 > 9/2 is true.
    fmt.Println(isMajorityElement([]int{2,4,5,5,5,5,5,6,6}, 5)) // true
    // Example 2:
    // Input: nums = [10,100,101,101], target = 101
    // Output: false
    // Explanation: The value 101 appears 2 times and the length of the array is 4.
    // Thus, 101 is not a majority element because 2 > 4/2 is false.
    fmt.Println(isMajorityElement([]int{10,100,101,101}, 101)) // false

    fmt.Println(isMajorityElement([]int{1,2,3,4,5,6,7,8,9}, 2)) // false
    fmt.Println(isMajorityElement([]int{9,8,7,6,5,4,3,2,1}, 2)) // false

    fmt.Println(isMajorityElement1([]int{2,4,5,5,5,5,5,6,6}, 5)) // true
    fmt.Println(isMajorityElement1([]int{10,100,101,101}, 101)) // false
    fmt.Println(isMajorityElement1([]int{1,2,3,4,5,6,7,8,9}, 2)) // false
    fmt.Println(isMajorityElement1([]int{9,8,7,6,5,4,3,2,1}, 2)) // false

    fmt.Println(isMajorityElement2([]int{2,4,5,5,5,5,5,6,6}, 5)) // true
    fmt.Println(isMajorityElement2([]int{10,100,101,101}, 101)) // false
    fmt.Println(isMajorityElement2([]int{1,2,3,4,5,6,7,8,9}, 2)) // false
    fmt.Println(isMajorityElement2([]int{9,8,7,6,5,4,3,2,1}, 2)) // false
}