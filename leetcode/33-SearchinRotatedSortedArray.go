package main

// 33. Search in Rotated Sorted Array
// There is an integer array nums sorted in ascending order (with distinct values).
// Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such 
// that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). 
// For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

// Given the array nums after the possible rotation and an integer target,
// return the index of target if it is in nums, or -1 if it is not in nums.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:
// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4

// Example 2:
// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1

// Example 3:
// Input: nums = [1], target = 0
// Output: -1
 
// Constraints:
//     1 <= nums.length <= 5000
//     -10^4 <= nums[i] <= 10^4
//     All values of nums are unique.
//     nums is an ascending array that is possibly rotated.
//     -10^4 <= target <= 10^4

import "fmt"

func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    low, high := 0, len(nums)-1
    for low <= high {
        mid := low + (high-low)>>1
        if nums[mid] == target {
            return mid
        } else if nums[mid] > nums[low] { // 在数值大的一部分区间里
            if nums[low] <= target && target < nums[mid] {
                high = mid - 1
            } else {
                low = mid + 1
            }
        } else if nums[mid] < nums[high] { // 在数值小的一部分区间里
            if nums[mid] < target && target <= nums[high] {
                low = mid + 1
            } else {
                high = mid - 1
            }
        } else {
            if nums[low] == nums[mid] {
                low++
            }
            if nums[high] == nums[mid] {
                high--
            }
        }
    }
    return -1
}

// best solution
func search1(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    l, r := 0, len(nums) - 1
    for l < r {
        mid := l + (r - l) / 2
        if nums[mid] > nums[r] { // 先找到中的点
            l = mid + 1
        } else {
            r = mid
        }
    }
    //fmt.Printf("l = %v\n",l) // 找到了 两段数组的中间
    rot := l
    l, r = 0, len(nums) - 1
    if nums[rot] <= target && target <= nums[r] { // 确定值在哪段数组上
        l = rot
    } else {
        r = rot - 1
    }
    for l <= r {
        mid := l + (r - l) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return -1
}

func search2(nums []int, target int) int {
    l, r := 0,len(nums)-1
    for l <= r {
        mid := (l + r)/2
        if nums[mid] == target {
            return mid
        }
        if nums[0] <= nums[mid] {
            if nums[0] <= target && target < nums[mid] {
                r = mid -1
            } else {
                l = mid + 1
            }
        } else {
            if nums[mid] < target && target <= nums[len(nums)-1] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    return -1
}

func main() {
    fmt.Printf("search([]int{4,5,6,7,0,1,2},0)  = %v\n",search([]int{4,5,6,7,0,1,2},0)) // 4
    fmt.Printf("search([]int{4,5,6,7,0,1,2},3)  = %v\n",search([]int{4,5,6,7,0,1,2},3)) // -1
    fmt.Printf("search([]int{1},0)  = %v\n",search([]int{1},0)) // -1
    
    fmt.Printf("search1([]int{4,5,6,7,0,1,2},0)  = %v\n",search1([]int{4,5,6,7,0,1,2},0)) //4
    fmt.Printf("search1([]int{4,5,6,7,0,1,2},3)  = %v\n",search1([]int{4,5,6,7,0,1,2},3)) // -1
    fmt.Printf("search1([]int{1},0)  = %v\n",search1([]int{1},0)) // -1

    fmt.Printf("search2([]int{4,5,6,7,0,1,2},0)  = %v\n",search2([]int{4,5,6,7,0,1,2},0)) //4
    fmt.Printf("search2([]int{4,5,6,7,0,1,2},3)  = %v\n",search2([]int{4,5,6,7,0,1,2},3)) // -1
    fmt.Printf("search2([]int{1},0)  = %v\n",search2([]int{1},0)) // -1
}