package main

// 540. Single Element in a Sorted Array
// You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.
// Return the single element that appears only once.

// Your solution must run in O(log n) time and O(1) space.

// Example 1:
// Input: nums = [1,1,2,3,3,4,4,8,8]
// Output: 2

// Example 2:
// Input: nums = [3,3,7,7,10,11,11]
// Output: 10

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^5

import "fmt"

func singleNonDuplicate(nums []int) int {
    low, high := 0, len(nums) - 1
    for low < high {
        mid := (low + high) / 2
        if mid % 2 == 0 { // nums[mid] == nums[mid+1]
            if nums[mid] != nums[mid + 1] {
                high = mid
            } else {
                low = mid + 2
            }
        } else {
            if nums[mid-1] != nums[mid] {
                high = mid - 1
            } else {
                low = mid + 1
            }
        }
    }
    return nums[low]
}

func singleNonDuplicate1(nums []int) int {
    l, r := 0 ,len(nums) - 1
    for l < r {
        mid := l + (r - l ) / 2
        mid -= mid & 1
        if nums[mid] == nums[mid+1] {
            l = mid + 2
        } else {
            r = mid
        }
    }
    return nums[l]
}

func main() {
    fmt.Println(singleNonDuplicate([]int{1,1,2,3,3,4,4,8,8})) // 2
    fmt.Println(singleNonDuplicate([]int{3,3,7,7,10,11,11})) // 10

    fmt.Println(singleNonDuplicate1([]int{1,1,2,3,3,4,4,8,8})) // 2
    fmt.Println(singleNonDuplicate1([]int{3,3,7,7,10,11,11})) // 10
}