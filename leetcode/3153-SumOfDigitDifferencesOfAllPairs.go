package main

// 3153. Sum of Digit Differences of All Pairs
// You are given an array nums consisting of positive integers where all integers have the same number of digits.
// The digit difference between two integers is the count of different digits that are in the same position in the two integers.
// Return the sum of the digit differences between all pairs of integers in nums.

// Example 1:
// Input: nums = [13,23,12]
// Output: 4
// Explanation:
// We have the following:
// - The digit difference between 13 and 23 is 1.
// - The digit difference between 13 and 12 is 1.
// - The digit difference between 23 and 12 is 2.
// So the total sum of digit differences between all pairs of integers is 1 + 1 + 2 = 4.

// Example 2:
// Input: nums = [10,10,10,10]
// Output: 0
// Explanation:
// All the integers in the array are the same. So the total sum of digit differences between all pairs of integers will be 0.

// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i] < 10^9
//     All integers in nums have the same number of digits.

import "fmt"
import "strconv"

func sumDigitDifferences(nums []int) int64 {
    res, n := 0, len(strconv.Itoa(nums[0])) // 数位长度都 相同
    freq := make([][]int, n)
    for i := range freq {
        freq[i] = make([]int, 10)
    }
    for k, v := range nums {
        str := strconv.Itoa(v)
        for i := 0; i < n; i++ {
            d, _ := strconv.Atoi(string(str[i])) // 每一位转成数字
            res += k - freq[i][d]
            freq[i][d]++
        }
    }
    return int64(res)
}

func sumDigitDifferences1(nums []int) int64 {
    res, cnt := 0, make([][10]int, len(strconv.Itoa(nums[0])))
    for k, v := range nums {
        for i := 0; v > 0; v /= 10 { // 循环 / 10 取余
            d := v % 10
            res += k - cnt[i][d]
            cnt[i][d]++
            i++
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [13,23,12]
    // Output: 4
    // Explanation:
    // We have the following:
    // - The digit difference between 13 and 23 is 1.
    // - The digit difference between 13 and 12 is 1.
    // - The digit difference between 23 and 12 is 2.
    // So the total sum of digit differences between all pairs of integers is 1 + 1 + 2 = 4.
    fmt.Println(sumDigitDifferences([]int{13,23,12})) // 4
    // Example 2:
    // Input: nums = [10,10,10,10]
    // Output: 0
    // Explanation:
    // All the integers in the array are the same. So the total sum of digit differences between all pairs of integers will be 0.
    fmt.Println(sumDigitDifferences([]int{10,10,10,10})) // 0

    fmt.Println(sumDigitDifferences1([]int{13,23,12})) // 4
    fmt.Println(sumDigitDifferences1([]int{10,10,10,10})) // 0
}