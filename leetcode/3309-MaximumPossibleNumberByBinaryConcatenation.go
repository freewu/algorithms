package main

// 3309. Maximum Possible Number by Binary Concatenation
// You are given an array of integers nums of size 3.

// Return the maximum possible number whose binary representation can be formed by concatenating the binary representation of all elements in nums in some order.

// Note that the binary representation of any number does not contain leading zeros.

// Example 1:
// Input: nums = [1,2,3]
// Output: 30
// Explanation:
// Concatenate the numbers in the order [3, 1, 2] to get the result "11110", which is the binary representation of 30.

// Example 2:
// Input: nums = [2,8,16]
// Output: 1296
// Explanation:
// Concatenate the numbers in the order [2, 8, 16] to get the result "10100010000", which is the binary representation of 1296.

// Constraints:
//     nums.length == 3
//     1 <= nums[i] <= 127

import "fmt"
import "strconv"
import "sort"

func maxGoodNumber(nums []int) int {
    i2bs := func(n int) string { // Int -> Binary String
        if n == 0 { return "0" }
        res := ""
        for n > 0 {
            if n % 2 == 0 {
                res = "0" + res 
            } else {
                res = "1" + res 
            }
            n /= 2
        }
        return res
    }
    bs2i := func(s string) int { // Binary String => Int
        res := 0
        for i := 0; i < len(s); i++ {
            res = res * 2 + int(s[i] - '0')
        }
        return res
    }
    a, b, c := i2bs(nums[0]), i2bs(nums[1]), i2bs(nums[2])
    res, arr := bs2i(a+b+c), []int{ bs2i(a+c+b), bs2i(b+a+c), bs2i(b+c+a), bs2i(c+a+b), bs2i(c+b+a) }
    for _, v := range arr {
        if v > res { res = v }
    }
    return res
}

func maxGoodNumber1(nums []int) int {
    a, b, c := fmt.Sprintf("%b", nums[0]), fmt.Sprintf("%b", nums[1]), fmt.Sprintf("%b", nums[2])
    res, t := 0, int64(0)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    t, _ = strconv.ParseInt(fmt.Sprintf("%s%s%s", a, b, c), 2, 64) // 012
    res = max(res, int(t))
    t, _ = strconv.ParseInt(fmt.Sprintf("%s%s%s", a, c, b), 2, 64) // 021
    res = max(res, int(t))
    t, _ = strconv.ParseInt(fmt.Sprintf("%s%s%s", b, a, c), 2, 64) // 102
    res = max(res, int(t))
    t, _ = strconv.ParseInt(fmt.Sprintf("%s%s%s", b, c, a), 2, 64) // 120
    res = max(res, int(t))
    t, _ = strconv.ParseInt(fmt.Sprintf("%s%s%s", c, a, b), 2, 64) // 201
    res = max(res, int(t))
    t, _ = strconv.ParseInt(fmt.Sprintf("%s%s%s", c, b, a), 2, 64) // 210
    res = max(res, int(t))
    return res
}

func maxGoodNumber2(nums []int) int {
    join := func (x, y int) int {
        t, l := y, 0
        for ; t > 0; t /= 2 { l++ }
        return x << l + y
    }
    sort.Slice(nums, func(i, j int) bool {
        return join(nums[i], nums[j]) > join(nums[j], nums[i])
    })
    res := 0
    for _, v := range nums {
        res = join(res, v)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 30
    // Explanation:
    // Concatenate the numbers in the order [3, 1, 2] to get the result "11110", which is the binary representation of 30.
    fmt.Println(maxGoodNumber([]int{1,2,3})) // 30
    // Example 2:
    // Input: nums = [2,8,16]
    // Output: 1296
    // Explanation:
    // Concatenate the numbers in the order [2, 8, 16] to get the result "10100010000", which is the binary representation of 1296.
    fmt.Println(maxGoodNumber([]int{2,8,16})) // 1296

    fmt.Println(maxGoodNumber([]int{1,1,1})) // 7
    fmt.Println(maxGoodNumber([]int{127,127,127})) // 2097151

    fmt.Println(maxGoodNumber1([]int{1,2,3})) // 30
    fmt.Println(maxGoodNumber1([]int{2,8,16})) // 1296
    fmt.Println(maxGoodNumber1([]int{1,1,1})) // 7
    fmt.Println(maxGoodNumber1([]int{127,127,127})) // 2097151

    fmt.Println(maxGoodNumber2([]int{1,2,3})) // 30
    fmt.Println(maxGoodNumber2([]int{2,8,16})) // 1296
    fmt.Println(maxGoodNumber2([]int{1,1,1})) // 7
    fmt.Println(maxGoodNumber2([]int{127,127,127})) // 2097151
}