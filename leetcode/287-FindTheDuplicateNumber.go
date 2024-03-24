package main

// 287. Find the Duplicate Number
// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
// There is only one repeated number in nums, return this repeated number.
// You must solve the problem without modifying the array nums and uses only constant extra space.

// Example 1:
// Input: nums = [1,3,4,2,2]
// Output: 2

// Example 2:
// Input: nums = [3,1,3,4,2]
// Output: 3

// Example 3:
// Input: nums = [3,3,3,3,3]
// Output: 3
 
// Constraints:
//     1 <= n <= 105
//     nums.length == n + 1
//     1 <= nums[i] <= n
//     All the integers in nums appear only once except for precisely one integer which appears two or more times.

import "fmt"
import "sort"

// 快慢指针
func findDuplicate(nums []int) int {
    slow := nums[0]
    fast := nums[nums[0]]
    for fast != slow {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }
    res := 0
    for res != slow {
        res = nums[res]
        slow = nums[slow]
    }
    return res
}

// 二分搜索
func findDuplicate1(nums []int) int {
    low, high := 0, len(nums)-1
    for low < high {
        mid, count := low + ( high - low) >> 1, 0
        for _, num := range nums {
            if num <= mid {
                count++
            }
        }
        if count > mid {
            high = mid
        } else {
            low = mid + 1
        }
    }
    return low
}

func findDuplicate2(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    sort.Ints(nums)
    diff := -1
    for i := 0; i < len(nums); i++ {
        if nums[i] - i - 1 >= diff {
            diff = nums[i] - i - 1
        } else {
            return nums[i]
        }
    }
    return 0
}

// best solution
func findDuplicate3(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    tortoise := nums[0]
    hare := nums[0]
    for {
        hare = nums[nums[hare]]
        tortoise = nums[tortoise]
        if hare == tortoise {
            break
        }
    }
    tortoise = nums[0]
    for hare != tortoise {
        hare = nums[hare]
        tortoise = nums[tortoise]
    }
    return tortoise
}

func findDuplicate4(nums []int) int {
    for {
        vi := nums[0]
        // 找到了重复的值
        if vi == nums[vi] {
            return vi
        }
        // 用 0 来存储需要查找的下一个值 
        nums[0] = nums[vi]
        // 把值移动到对应的下标
        nums[vi] = vi
    }
}

func main()  {
    fmt.Printf("findDuplicate([]int{ 1,3,4,2,2 }) = %v\n",findDuplicate([]int{ 1,3,4,2,2 })) // 2
    fmt.Printf("findDuplicate([]int{ 3,1,3,4,2 }) = %v\n",findDuplicate([]int{ 3,1,3,4,2 })) // 3
    fmt.Printf("findDuplicate([]int{ 3,3,3,3,3 }) = %v\n",findDuplicate([]int{ 3,3,3,3,3 })) // 3

    fmt.Printf("findDuplicate1([]int{ 1,3,4,2,2 }) = %v\n",findDuplicate1([]int{ 1,3,4,2,2 })) // 2
    fmt.Printf("findDuplicate1([]int{ 3,1,3,4,2 }) = %v\n",findDuplicate1([]int{ 3,1,3,4,2 })) // 3
    fmt.Printf("findDuplicate1([]int{ 3,3,3,3,3 }) = %v\n",findDuplicate1([]int{ 3,3,3,3,3 })) // 3

    fmt.Printf("findDuplicate2([]int{ 1,3,4,2,2 }) = %v\n",findDuplicate2([]int{ 1,3,4,2,2 })) // 2
    fmt.Printf("findDuplicate2([]int{ 3,1,3,4,2 }) = %v\n",findDuplicate2([]int{ 3,1,3,4,2 })) // 3
    fmt.Printf("findDuplicate2([]int{ 3,3,3,3,3 }) = %v\n",findDuplicate2([]int{ 3,3,3,3,3 })) // 3

    fmt.Printf("findDuplicate3([]int{ 1,3,4,2,2 }) = %v\n",findDuplicate3([]int{ 1,3,4,2,2 })) // 2
    fmt.Printf("findDuplicate3([]int{ 3,1,3,4,2 }) = %v\n",findDuplicate3([]int{ 3,1,3,4,2 })) // 3
    fmt.Printf("findDuplicate3([]int{ 3,3,3,3,3 }) = %v\n",findDuplicate3([]int{ 3,3,3,3,3 })) // 3

    fmt.Printf("findDuplicate4([]int{ 1,3,4,2,2 }) = %v\n",findDuplicate4([]int{ 1,3,4,2,2 })) // 2
    fmt.Printf("findDuplicate4([]int{ 3,1,3,4,2 }) = %v\n",findDuplicate4([]int{ 3,1,3,4,2 })) // 3
    fmt.Printf("findDuplicate4([]int{ 3,3,3,3,3 }) = %v\n",findDuplicate4([]int{ 3,3,3,3,3 })) // 3
}
