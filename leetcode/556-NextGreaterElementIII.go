package main

// 556. Next Greater Element III
// Given a positive integer n, 
// find the smallest integer which has exactly the same digits existing in the integer n and is greater in value than n. 
// If no such positive integer exists, return -1.

// Note that the returned integer should fit in 32-bit integer, 
// if there is a valid answer but it does not fit in 32-bit integer, return -1.

// Example 1:
// Input: n = 12
// Output: 21

// Example 2:
// Input: n = 21
// Output: -1
 
// Constraints:
//     1 <= n <= 2^31 - 1

import "fmt"
import "math"
import "sort"
import "strconv"

func nextGreaterElement(n int) int {
    digits := make([]int, 0)
    for n != 0 {
        digits = append(digits, n % 10)
        n /= 10
    }
    for i := 1; i < len(digits); i++ {
        if digits[i] >= digits[i-1] {
            continue
        }
        index, diff := 0, 10
        for j := 0; j < i; j++ {
            curDiff := digits[j] - digits[i]
            if curDiff > 0 && curDiff < diff {
                diff = curDiff
                index = j
            }
        }
        digits[i], digits[index] = digits[index], digits[i]
        sort.Sort(sort.Reverse(sort.IntSlice(digits[:i])))
        dest := 0
        for index, value := range digits {
            dest += value * int(math.Pow10(index))
        }
        if dest <= math.MaxInt32 {
            return dest
        }
    }
    return -1
}

func nextGreaterElement1(n int) int {
    nums := []byte(strconv.Itoa(n))
    index := -1
    for i := len(nums) - 2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            index = i
            break
        }
    }
    if index >= 0 {
        for i := len(nums) - 1; i >= 0; i-- {
            if nums[i] > nums[index] {
                nums[i], nums[index] = nums[index], nums[i]
                break
            }
        }
    }
    low, high := index+1, len(nums)-1
    for low < high {
        nums[low], nums[high] = nums[high], nums[low]
        low++
        high--
    }
    result, _ := strconv.Atoi(string(nums))
    if result > n && result <= math.MaxInt32 {
        return result
    }
    return -1
}

func main() {
    // Example 1:
    // Input: n = 12
    // Output: 21
    fmt.Println(nextGreaterElement(12)) // 21
    // Example 2:
    // Input: n = 21
    // Output: -1
    fmt.Println(nextGreaterElement(21)) // -1

    fmt.Println(nextGreaterElement1(12)) // 21
    fmt.Println(nextGreaterElement1(21)) // -1
}