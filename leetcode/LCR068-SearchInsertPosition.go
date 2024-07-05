package main

// LCR 068. 搜索插入位置
// 给定一个排序的整数数组 nums 和一个整数目标值 target ，请在数组中找到 target ，并返回其下标。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

// 请必须使用时间复杂度为 O(log n) 的算法。

// 示例 1:
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2

// 示例 2:
// 输入: nums = [1,3,5,6], target = 2
// 输出: 1

// 示例 3:
// 输入: nums = [1,3,5,6], target = 7
// 输出: 4

// 示例 4:
// 输入: nums = [1,3,5,6], target = 0
// 输出: 0

// 示例 5:
// 输入: nums = [1], target = 0
// 输出: 0

// 提示:
//     1 <= nums.length <= 10^4
//     -10^4 <= nums[i] <= 10^4
//     nums 为无重复元素的升序排列数组
//     -10^4 <= target <= 10^4


// O(n)
func searchInsert1(nums []int, target int) int {
    for i := 0; i < len(nums); i++ {
        // 找到了位置
        if nums[i] == target {
            return i
        }
        // 找到了大于目标的位置
        if nums[i] > target {
            return i
        }
    }
    // 插在队尾
    return len(nums)
}

// O(log n)
func searchInsert(nums []int, target int) int {
    s, l := 0, len(nums) - 1
    if nums[s] >= target {
        return s
    }
    if nums[l] == target {
        return l
    }
    if nums[l] < target {
        return l + 1
    }
    for {
        if s > l {
            break
        }
        m := (s + l) / 2
        if nums[m] == target {
            return m
        }
        if nums[m] < target {
            s = m + 1
        } else {
            l = m - 1
        }
    }
    return s
}

func searchInsert2(nums []int, target int) int {
    low, high := 0, len(nums)-1
    for low <= high {
        mid := low + (high-low) >> 1 // 找到中间位置
        if nums[mid] >= target { // 如果中间位置 >= 目标值,  结束位置 设置为中间值  low -- mid -- high 说明目标值 存在 low -- mid 段中
            high = mid - 1
        } else {
            if (mid == len(nums)-1) || (nums[mid+1] >= target) {
                return mid + 1
            }
            low = mid + 1
        }
    }
    return 0
}

// best solution
func searchInsertBest(nums []int, target int) int {
    right, left := len(nums) - 1, 0
    for left <= right {
        mid := left + (right - left) / 2
        if nums[mid] == target {
            return mid
        }
        if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return left
}

func searchInsert3(nums []int, target int) int {
    // 二分查找，使得该位置左边的数比目标值小，右边的值比目标值大，如果找到则返回索引
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        if nums[0] < target {
            return 1
        } else {
            return 0
        }
    }
    mid := len(nums) / 2
    if nums[mid] == target {
        return mid
    }
    if nums[mid] > target {
        return searchInsert3(nums[:mid], target)
    }
    return searchInsert3(nums[mid+1:], target) + mid + 1
}

func main() {
    fmt.Printf("searchInsert([]int{1,3,5,6},5) = %v\n",searchInsert([]int{1,3,5,6},5)) // 2
    fmt.Printf("searchInsert([]int{1,3,5,6},2) = %v\n",searchInsert([]int{1,3,5,6},2)) // 1
    fmt.Printf("searchInsert([]int{1,3,5,6},7) = %v\n",searchInsert([]int{1,3,5,6},7)) // 4
    fmt.Printf("searchInsert([]int{1,3,5,6},0) = %v\n",searchInsert([]int{1,3,5,6},0)) // 0
    fmt.Printf("searchInsert1([]int{1,3,5,6},5) = %v\n",searchInsert1([]int{1,3,5,6},5)) // 2
    fmt.Printf("searchInsert1([]int{1,3,5,6},2) = %v\n",searchInsert1([]int{1,3,5,6},2)) // 1
    fmt.Printf("searchInsert1([]int{1,3,5,6},7) = %v\n",searchInsert1([]int{1,3,5,6},7)) // 4
    fmt.Printf("searchInsert1([]int{1,3,5,6},0) = %v\n",searchInsert1([]int{1,3,5,6},0)) // 0
    fmt.Printf("searchInsert2([]int{1,3,5,6},5) = %v\n",searchInsert2([]int{1,3,5,6},5)) // 2
    fmt.Printf("searchInsert2([]int{1,3,5,6},2) = %v\n",searchInsert2([]int{1,3,5,6},2)) // 1
    fmt.Printf("searchInsert2([]int{1,3,5,6},7) = %v\n",searchInsert2([]int{1,3,5,6},7)) // 4
    fmt.Printf("searchInsert2([]int{1,3,5,6},0) = %v\n",searchInsert2([]int{1,3,5,6},0)) // 0

    fmt.Printf("searchInsertBest([]int{1,3,5,6},5) = %v\n",searchInsertBest([]int{1,3,5,6},5)) // 2
    fmt.Printf("searchInsertBest([]int{1,3,5,6},2) = %v\n",searchInsertBest([]int{1,3,5,6},2)) // 1
    fmt.Printf("searchInsertBest([]int{1,3,5,6},7) = %v\n",searchInsertBest([]int{1,3,5,6},7)) // 4
    fmt.Printf("searchInsertBest([]int{1,3,5,6},0) = %v\n",searchInsertBest([]int{1,3,5,6},0)) // 0

    fmt.Printf("searchInsert3([]int{1,3,5,6},5) = %v\n",searchInsert3([]int{1,3,5,6},5)) // 2
    fmt.Printf("searchInsert3([]int{1,3,5,6},2) = %v\n",searchInsert3([]int{1,3,5,6},2)) // 1
    fmt.Printf("searchInsert3([]int{1,3,5,6},7) = %v\n",searchInsert3([]int{1,3,5,6},7)) // 4
    fmt.Printf("searchInsert3([]int{1,3,5,6},0) = %v\n",searchInsert3([]int{1,3,5,6},0)) // 0
}