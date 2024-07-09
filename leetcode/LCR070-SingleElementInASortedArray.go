package main

// LCR 070. 有序数组中的单一元素
// 给定一个只包含整数的有序数组 nums ，每个元素都会出现两次，唯有一个数只会出现一次，请找出这个唯一的数字。

// 示例 1:
// 输入: nums = [1,1,2,3,3,4,4,8,8]
// 输出: 2

// 示例 2:
// 输入: nums =  [3,3,7,7,10,11,11]
// 输出: 10

// 提示:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^5

// 进阶: 采用的方案可以在 O(log n) 时间复杂度和 O(1) 空间复杂度中运行吗？

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