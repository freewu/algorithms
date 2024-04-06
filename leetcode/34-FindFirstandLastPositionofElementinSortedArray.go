package main

// 34. Find First and Last Position of Element in Sorted Array
// Given an array of integers nums sorted in non-decreasing order, 
// find the starting and ending position of a given target value.
// If target is not found in the array, return [-1, -1].
// You must write an algorithm with O(log n) runtime complexity.

// Example 1:
// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]

// Example 2:
// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]

// Example 3:
// Input: nums = [], target = 0
// Output: [-1,-1]
 
// Constraints:
//     0 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9
//     nums is a non-decreasing array.
//     -10^9 <= target <= 10^9

// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。
// 找出给定目标值在数组中的开始位置和结束位置。你的算法时间复杂度必须是 O(log n) 级别。
// 如果数组中不存在目标值，返回 [-1, -1]

import "fmt"

func searchRange(nums []int, target int) []int {
    // 二分查找最后一个与 target 相等的元素，时间复杂度 O(logn)
    searchLastEqualElement := func (nums []int, target int) int {
        low, high := 0, len(nums)-1
        for low <= high {
            mid := low + ((high - low) >> 1)
            if nums[mid] > target {
                high = mid - 1
            } else if nums[mid] < target {
                low = mid + 1
            } else {
                if (mid == len(nums)-1) || (nums[mid+1] != target) { // 找到最后一个与 target 相等的元素
                    return mid
                }
                low = mid + 1
            }
        }
        return -1
    }
    // 二分查找第一个与 target 相等的元素，时间复杂度 O(logn)
    searchFirstEqualElement := func (nums []int, target int) int {
        low, high := 0, len(nums)-1
        for low <= high {
            mid := low + ((high - low) >> 1)
            if nums[mid] > target {
                high = mid - 1
            } else if nums[mid] < target {
                low = mid + 1
            } else {
                if (mid == 0) || (nums[mid-1] != target) { // 找到第一个与 target 相等的元素
                    return mid
                }
                high = mid - 1
            }
        }
        return -1
    }
    return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}
}

// best solution
func searchRange1(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    findLeft := func (nums []int, target int) int {
        left, right := 0, len(nums)-1
        for left <= right {
            mid := (left+right) / 2
            if nums[mid] == target {
                right = mid - 1
            } else if nums[mid] < target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        if left >= 0 && left < len(nums) && nums[left] == target {
            return left
        }
        return -1
    }
    findRight := func (nums []int, target int) int {
        left, right := 0, len(nums) - 1
        for left <= right {
            mid := (left+right) / 2
            if nums[mid] == target {
                left = mid + 1
            } else if nums[mid] < target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        if right >= 0 && right < len(nums) && nums[right] == target {
            return right
        }
        return -1
    }
    return []int{ findLeft(nums, target), findRight(nums, target) }
}

func searchRange2(nums []int, target int) []int {
    if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
        return []int{-1,-1}
    }
    var start, end = -1, -1
    for i, j := 0, len(nums) - 1; (start == -1 || end == -1) && i <= j; {
        if nums[i] == target {
            start = i
        } else {
            i++
        }
        if nums[j] == target {
            end = j
        } else {
            j--
        }
    }
    return []int{ start, end }
}

func searchRange3(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    binarySearch := func(nums []int, target int, lower bool) int {
        i, j, res := 0,len(nums) - 1, -1
        for i <= j {
            mid := i + (j-i)/2
            if nums[mid] > target || (lower && nums[mid] >= target) {
                j = mid - 1
                if lower {
                    res = mid
                }
            } else {
                i = mid + 1
                if !lower {
                    res = mid
                }
            }
        }
        return res
    }
    left := binarySearch(nums, target, true)
    right := binarySearch(nums, target, false)
    if left >= 0 && right >= 0 && nums[left] == nums[right] && nums[left] == target {
        return []int{left, right}
    }
    return []int{-1, -1}
}

func main() {
    fmt.Printf("searchRange([]int{5,7,7,8,8,10},8) = %v\n",searchRange([]int{5,7,7,8,8,10}, 8)) // [3,4]
    fmt.Printf("searchRange([]int{5,7,7,8,8,10},6) = %v\n",searchRange([]int{5,7,7,8,8,10}, 6)) // [-1,-1]
    fmt.Printf("searchRange([]int{},0) = %v\n",searchRange([]int{}, 0)) // [-1,-1]

    fmt.Printf("searchRange1([]int{5,7,7,8,8,10},8) = %v\n",searchRange1([]int{5,7,7,8,8,10}, 8)) // [3,4]
    fmt.Printf("searchRange1([]int{5,7,7,8,8,10},6) = %v\n",searchRange1([]int{5,7,7,8,8,10}, 6)) // [-1,-1]
    fmt.Printf("searchRange1([]int{},0) = %v\n",searchRange1([]int{}, 0)) // [-1,-1]

    fmt.Printf("searchRange2([]int{5,7,7,8,8,10},8) = %v\n",searchRange2([]int{5,7,7,8,8,10}, 8)) // [3,4]
    fmt.Printf("searchRange2([]int{5,7,7,8,8,10},6) = %v\n",searchRange2([]int{5,7,7,8,8,10}, 6)) // [-1,-1]
    fmt.Printf("searchRange2([]int{},0) = %v\n",searchRange2([]int{}, 0)) // [-1,-1]

    fmt.Printf("searchRange3([]int{5,7,7,8,8,10},8) = %v\n",searchRange3([]int{5,7,7,8,8,10}, 8)) // [3,4]
    fmt.Printf("searchRange3([]int{5,7,7,8,8,10},6) = %v\n",searchRange3([]int{5,7,7,8,8,10}, 6)) // [-1,-1]
    fmt.Printf("searchRange3([]int{},0) = %v\n",searchRange3([]int{}, 0)) // [-1,-1]
}
