package main

// 3265. Count Almost Equal Pairs I
// You are given an array nums consisting of positive integers.

// We call two integers x and y in this problem almost equal if both integers can become equal after performing the following operation at most once:
//     Choose either x or y and swap any two digits within the chosen number.

// Return the number of indices i and j in nums where i < j such that nums[i] and nums[j] are almost equal.

// Note that it is allowed for an integer to have leading zeros after performing an operation.

// Example 1:
// Input: nums = [3,12,30,17,21]
// Output: 2
// Explanation:
// The almost equal pairs of elements are:
// 3 and 30. By swapping 3 and 0 in 30, you get 3.
// 12 and 21. By swapping 1 and 2 in 12, you get 21.

// Example 2:
// Input: nums = [1,1,1,1,1]
// Output: 10
// Explanation:
// Every two elements in the array are almost equal.

// Example 3:
// Input: nums = [123,231]
// Output: 0
// Explanation:
// We cannot swap any two digits of 123 or 231 to reach the other.

// Constraints:
//     2 <= nums.length <= 100
//     1 <= nums[i] <= 10^6

import "fmt"
import "strconv"

func countPairs(nums []int) int {
    res, n := 0, len(nums)
    same := func(a, b string) bool {
        bi, _ := strconv.Atoi(b)
        for i := 0; i < len(a); i++ {
            for j := i + 1; j < len(a); j++ {
                ar := []byte(a)
                ar[i], ar[j] = ar[j], ar[i]
                ci, _ := strconv.Atoi(string(ar))
                if ci == bi {
                    return true
                }
                // ar[i], ar[j] = ar[j], ar[i]
            }
        }
        return false
    }
    check := func(x, y int) bool {
        a, b := strconv.Itoa(x), strconv.Itoa(y)
        if a == b || same(a, b) || same(b, a) {
            return true
        }
        return false
    }
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if(check(nums[i],nums[j])) {
                res++
            }
        }
    }
    return res
}

func countPairs1(nums []int) int {
    res, mp := 0, make(map[int][]int)
    for _, v := range nums {
        a, flag := v, 0
        for i := 0; i <= 6; i++ {
            num := v % 10
            v /= 10
            flagNum := (flag >> (num * 3)) & 0x7
            flag ^= flagNum << (num * 3)
            flagNum++
            flag |= flagNum << (num * 3)
        }
        mp[flag] = append(mp[flag], a)
    }
    valid := func (num1, num2 int) bool{
        if num1 == num2 { return true }
        gap := 0
        for num1 != 0 || num2 != 0 {
            if num1 % 10 != num2 % 10 { gap++ }
            if gap == 3 { return false }
            num1 /= 10
            num2 /= 10
        }
        return true
    }
    for _, arr := range mp {
        n := len(arr)
        for i := 0; i < n - 1; i++ {
            for j := i + 1; j < n; j++ {
                if valid(arr[i], arr[j]) {
                    res++
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,12,30,17,21]
    // Output: 2
    // Explanation:
    // The almost equal pairs of elements are:
    // 3 and 30. By swapping 3 and 0 in 30, you get 3.
    // 12 and 21. By swapping 1 and 2 in 12, you get 21.
    fmt.Println(countPairs([]int{3,12,30,17,21})) // 2
    // Example 2:
    // Input: nums = [1,1,1,1,1]
    // Output: 10
    // Explanation:
    // Every two elements in the array are almost equal.
    fmt.Println(countPairs([]int{1,1,1,1,1})) // 10
    // Example 3:
    // Input: nums = [123,231]
    // Output: 0
    // Explanation:
    // We cannot swap any two digits of 123 or 231 to reach the other.
    fmt.Println(countPairs([]int{123,231})) // 0

    fmt.Println(countPairs1([]int{3,12,30,17,21})) // 2
    fmt.Println(countPairs1([]int{1,1,1,1,1})) // 10
    fmt.Println(countPairs1([]int{123,231})) // 0
}