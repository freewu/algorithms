package main

// 4053. Minimum Operations to Make Every Element Palindromic
// You are given an integer array nums.

// In one operation, you may choose an index i and either increment or decrement nums[i] by 2.

// Return the minimum number of operations required to make every element in nums a positive palindrome. 
// Different elements may be changed into different palindromic integers.

// Example 1:
// Input: nums = [10,12,14,16]
// Output: 9
// Explanation:
// One optimal sequence of operations is:
// Decrement nums[0] by 2 once to change it from 10 to 8.
// Decrement nums[1] by 2 twice to change it from 12 to 8.
// Decrement nums[2] by 2 three times to change it from 14 to 8.
// Increment nums[3] by 2 three times to change it from 16 to 22.
// After 1 + 2 + 3 + 3 = 9 operations, nums = [8, 8, 8, 22], and every element is a positive palindromic integer.
// It can be shown that fewer than 9 operations cannot achieve this.

// Example 2:
// Input: nums = [9,10,11,10]
// Output: 2
// Explanation:
// Decrement nums[1] and nums[3] by 2 once each.
// After 2 operations, nums = [9, 8, 11, 8], and every element is a positive palindromic integer.
// At least one operation is needed for each of these two elements, so the minimum number of operations is 2.

// Example 3:
// Input: nums = [125]
// Output: 2
// Explanation:
// Decrement nums[0] by 2 twice to change it from 125 to 121, which is a positive palindromic integer.
// A single operation would change it to 123 or 127, neither of which is palindromic. 
// Thus, the minimum number of operations is 2.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"
import "sort"

const MX = 2_000_000_002
var palindromes = [2][]int{{0}, {0}} // 哨兵

// 预处理 [1, MX] 中的回文数
func init() {
    for base := 1; ; base *= 10 {
        // 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
        for i := base; i < base * 10; i++ {
            x := i
            for t := i / 10; t > 0; t /= 10 {
                x = x * 10 + t % 10
            }
            if x > MX {
                return
            }
            // 按照 x 的奇偶性分组
            palindromes[x % 2] = append(palindromes[x % 2], x)
        }
        // 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
        for i := base; i < base*10; i++ {
            x := i
            for t := i; t > 0; t /= 10 {
                x = x * 10 + t % 10
            }
            if x > MX {
                return
            }
            palindromes[x % 2] = append(palindromes[x % 2], x)
        }
    }
}

func minOperations(nums []int) int64 {
    res := int64(0)
    for _, v := range nums {
        p := palindromes[v % 2]
        i := sort.SearchInts(p, v)
        res += int64(min(p[i] - v, v - p[i-1]))
    }
    return res / 2
}

func genPalindromes() (odd, even []int) {
    const MX = 1_000_000_001
    for l := 1; l <= 10; l++ {
        halfLen := (l + 1) / 2
        low := 1
        for i := 1; i < halfLen; i++ {
            low *= 10
        }
        high := low * 10
        for prefix := low; prefix < high; prefix++ {
            p := makePalindrome(prefix, l&1 == 1)
            if p > MX {     
                break
            }
            if p & 1 == 0 {
                even = append(even, p)
            } else {
                odd = append(odd, p)
            }
        }
    }
    return
}

func makePalindrome(prefix int, odd bool) int {
    res := prefix
    if odd {
        prefix /= 10
    }
    for prefix > 0 {
        res = res * 10 + prefix % 10
        prefix /= 10
    }
    return res
}

func lowerBound(a []int, x int) int {
    low, high := 0, len(a)
    for low < high {
        mid := low + (high-low)/2
        if a[mid] < x {
            low = mid + 1
        } else {
            high = mid
        }
    }
    return low
}

var odd, even = genPalindromes()

func minOperations1(nums []int) int64 {
    res := int64(0)
    for _, n := range nums {
        list := even
        if n & 1 != 0 {
            list = odd
        }
        i := lowerBound(list, n)
        if i == 0 {
            res += int64(list[0] - n) / 2
            continue
        }
        if i == len(list) {
            res += int64(n - list[i - 1] - n) / 2
            continue
        }
        left, right := int64(n - list[i-1]), int64(list[i] - n)
        if left < right {
            res += left / 2
        } else {
            res += right / 2
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [10,12,14,16]
    // Output: 9
    // Explanation:
    // One optimal sequence of operations is:
    // Decrement nums[0] by 2 once to change it from 10 to 8.
    // Decrement nums[1] by 2 twice to change it from 12 to 8.
    // Decrement nums[2] by 2 three times to change it from 14 to 8.
    // Increment nums[3] by 2 three times to change it from 16 to 22.
    // After 1 + 2 + 3 + 3 = 9 operations, nums = [8, 8, 8, 22], and every element is a positive palindromic integer.
    // It can be shown that fewer than 9 operations cannot achieve this.
    fmt.Println(minOperations([]int{10,12,14,16})) // 9
    // Example 2:
    // Input: nums = [9,10,11,10]
    // Output: 2
    // Explanation:
    // Decrement nums[1] and nums[3] by 2 once each.
    // After 2 operations, nums = [9, 8, 11, 8], and every element is a positive palindromic integer.
    // At least one operation is needed for each of these two elements, so the minimum number of operations is 2.
    fmt.Println(minOperations([]int{9,10,11,10})) // 2
    // Example 3:
    // Input: nums = [125]
    // Output: 2
    // Explanation:
    // Decrement nums[0] by 2 twice to change it from 125 to 121, which is a positive palindromic integer.
    // A single operation would change it to 123 or 127, neither of which is palindromic. 
    // Thus, the minimum number of operations is 2.
    fmt.Println(minOperations([]int{125})) // 2

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 0
    
    fmt.Println(minOperations1([]int{10,12,14,16})) // 9
    fmt.Println(minOperations1([]int{9,10,11,10})) // 2
    fmt.Println(minOperations1([]int{125})) // 2
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1})) // 0
}