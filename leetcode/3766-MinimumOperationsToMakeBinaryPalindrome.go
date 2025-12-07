package main

// 3766. Minimum Operations to Make Binary Palindrome
// You are given an integer array nums.

// For each element nums[i], you may perform the following operations any number of times (including zero):
//     1. Increase nums[i] by 1, or
//     2. Decrease nums[i] by 1.

// A number is called a binary palindrome if its binary representation without leading zeros reads the same forward and backward.

// Your task is to return an integer array ans, where ans[i] represents the minimum number of operations required to convert nums[i] into a binary palindrome.

// Example 1:
// Input: nums = [1,2,4]
// Output: [0,1,1]
// Explanation:
// One optimal set of operations:
// nums[i] | Binary(nums[i]) | Nearest Palindrome | Binary (Palindrome) | Operations Required | ans[i]
//     1   |   1             |     1              |    1                |  Already palindrome | 0
//     2   |   10            |     3              |    11               |  Increase by 1      | 1
//     4   |   100           |     3              |    11               |  Decrease by 1      | 1  
// Thus, ans = [0, 1, 1].

// Example 2:
// Input: nums = [6,7,12]
// Output: [1,0,3]
// Explanation:
// One optimal set of operations:
// nums[i] | Binary(nums[i]) | Nearest Palindrome | Binary (Palindrome) | Operations Required | ans[i]
//     6   |   110           |     5              |    101              |  Decrease by 1      | 1
//     7   |   111           |     7              |    111              |  Already palindrome | 0
//     12  |   1100          |     15             |    1111             |  Increase by 3      | 3
// Thus, ans = [1, 0, 3].

// Constraints:
//     1 <= nums.length <= 5000
//     ​​​​​​​1 <= nums[i] <= 5000

import "fmt"
import "math/bits"

// Brute Force
func minOperations(nums []int) []int {
    res := make([]int, len(nums))
    isPalindrome := func(x int) bool {
        arr := []int{}
        for x > 0 {
            arr = append(arr, x%2)
            x /= 2
        }
        i, j := 0, len(arr) - 1 
        for i <= j {
            if arr[i] != arr[j] {
                return false
            }
            i++
            j--
        }
        return true
    }
    getResult :=func (x int) int {
        res, delta,front,back := 1 << 31, 0, false, false
        for !front && !back {
            f, b := x + delta, x - delta
            if isPalindrome(f) {
                res = min(res, (f - x))
                front = true
            }
            if isPalindrome(b) {
                res = min(res, (x - b))
                back = true
            }
            delta++
        }
        return res
    }
    for i := 0; i < len(nums); i++ {
        res[i] = getResult(nums[i])
    }
    return res
}

// 位运算
func minOperations1(nums []int) []int {
    abs := func(x int) int { if x < 0 { return -x }; return x }
    for i, x := range nums {
        res := 1 << 31
        n := bits.Len(uint(x))
        m := n / 2
        left := x >> m
        for l := left - 1; l <= left+1; l++ {
            // 左半反转到右半
            // 如果 n 是奇数，那么去掉回文中心再反转
            right := bits.Reverse(uint(l>>(n%2))) >> (bits.UintSize - m)
            pal := l << m | int(right)
            res = min(res, abs(x-pal))
        }
        nums[i] = res
    }
    return nums
}

func main() {
    // Example 1:
    // Input: nums = [1,2,4]
    // Output: [0,1,1]
    // Explanation:
    // One optimal set of operations:
    // nums[i] | Binary(nums[i]) | Nearest Palindrome | Binary (Palindrome) | Operations Required | ans[i]
    //     1   |   1             |     1              |    1                |  Already palindrome | 0
    //     2   |   10            |     3              |    11               |  Increase by 1      | 1
    //     4   |   100           |     3              |    11               |  Decrease by 1      | 1  
    // Thus, ans = [0, 1, 1].
    fmt.Println(minOperations([]int{1,2,4})) // [0, 1, 1]
    // Example 2:
    // Input: nums = [6,7,12]
    // Output: [1,0,3]
    // Explanation:
    // One optimal set of operations:
    // nums[i] | Binary(nums[i]) | Nearest Palindrome | Binary (Palindrome) | Operations Required | ans[i]
    //     6   |   110           |     5              |    101              |  Decrease by 1      | 1
    //     7   |   111           |     7              |    111              |  Already palindrome | 0
    //     12  |   1100          |     15             |    1111             |  Increase by 3      | 3
    // Thus, ans = [1, 0, 3].
    fmt.Println(minOperations([]int{6,7,12})) // [1, 0, 3]

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // [0 1 0 1 0 1 0 1 0]
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // [0 1 0 1 0 1 0 1 0]

    fmt.Println(minOperations1([]int{1,2,4})) // [0, 1, 1]
    fmt.Println(minOperations1([]int{6,7,12})) // [1, 0, 3]
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9})) // [0 1 0 1 0 1 0 1 0]
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1})) // [0 1 0 1 0 1 0 1 0]
}