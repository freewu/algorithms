package main

// 4036. Lexicographically Largest String After Pair Transformations
// You are given an integer array nums.

// For each integer x in nums, start with a string consisting of exactly x lowercase 'a' characters.

// You may perform the following operation any number of times (including zero):
//     Choose two adjacent equal letters and replace them with the next letter in the alphabet.

// For example, "aa" can be replaced with "b", and "bb" can be replaced with "c". The pair "zz" cannot be replaced.

// For each x, determine the lexicographically largest string that can be obtained.

// Return an array of strings where the ith string is the answer for nums[i].

// A string a is lexicographically larger than a string b if, at the first position where they differ, a contains a letter that appears later in the alphabet than the corresponding letter in b. 
// If the first min(a.length, b.length) characters are equal, the longer string is lexicographically larger.

// Example 1:
// Input: nums = [2,5,7]
// Output: ["b","ca","cba"]
// Explanation:
// nums[0] = 2: "aa" → "b".
// nums[1] = 5: "aaaaa" → "baaa" → "bba" → "ca".
// nums[2] = 7: "aaaaaaa" → "baaaaa" → "bbaaa" → "bbba" → "cba".
// Therefore, ans = ["b", "ca", "cba"].

// Example 2:
// Input: nums = [3,9,1]
// Output: ["ba","da","a"]
// Explanation:
// nums[0] = 3: "aaa" → "ba".
// nums[1] = 9: "aaaaaaaaa" → "baaaaaaa" → "bbaaaaa" → "bbbaaa" → "bbbba" → "cbba" → "cca" → "da".
// nums[2] = 1: No transformation can be applied, so the result is "a".
// Therefore, ans = ["ba", "da", "a"].

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^8

import "fmt"
import "bytes"
import "math/bits"

func largestString(nums []int) []string {
    res := make([]string, len(nums))
    for j, x := range nums {
        // 单独处理 'z'
        s := bytes.Repeat([]byte{'z'}, x>>25)
        // 然后从 'y' 到 'a'
        for i := min(24, bits.Len(uint(x))-1); i >= 0; i-- {
            if x >> i & 0x1 > 0 { // x 的 i 位是 1，所以有字母 i
                s = append(s, 'a' + byte(i))
            }
        }
        res[j] = string(s)
    }
    return res
}

func largestString1(nums []int) []string {
    helper := func(x int) string {
        ch := byte('a')
        res := make([]byte, 0)
        for x > 0 {
            if ch == 'z' {
                prefix := make([]byte, x)
                for i := 0; i < x; i++ {
                    prefix[i] = 'z'
                }
                return string(append(prefix, res...))
            }
            if x & 1 == 1 {
                res = append([]byte{ch}, res...)
            }
            x /= 2
            ch++
        }
        return string(res)
    }
    res := make([]string, len(nums))
    for i, n := range nums {
        res[i] = helper(n)
    }
    return res
}

func largestString2(nums []int) []string {
    res := make([]string, len(nums))
    weight := 1 << 25
    for i, x := range nums {
        b := make([]byte, 0, 32)
        for x >= weight {
            b = append(b, 'z')
            x -= weight
        }
        for bit := 24; bit >= 0; bit-- {
            if x&(1<<bit) != 0 {
                b = append(b, byte('a'+bit))
            }
        }
        res[i] = string(b)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,5,7]
    // Output: ["b","ca","cba"]
    // Explanation:
    // nums[0] = 2: "aa" → "b".
    // nums[1] = 5: "aaaaa" → "baaa" → "bba" → "ca".
    // nums[2] = 7: "aaaaaaa" → "baaaaa" → "bbaaa" → "bbba" → "cba".
    // Therefore, ans = ["b", "ca", "cba"].
    fmt.Println(largestString([]int{2,5,7})) // ["b","ca","cba"] 
    // Example 2:
    // Input: nums = [3,9,1]
    // Output: ["ba","da","a"]
    // Explanation:
    // nums[0] = 3: "aaa" → "ba".
    // nums[1] = 9: "aaaaaaaaa" → "baaaaaaa" → "bbaaaaa" → "bbbaaa" → "bbbba" → "cbba" → "cca" → "da".
    // nums[2] = 1: No transformation can be applied, so the result is "a".
    // Therefore, ans = ["ba", "da", "a"].
    fmt.Println(largestString([]int{3,9,1})) // ["ba","da","a"] 

    fmt.Println(largestString([]int{1,2,3,4,5,6,7,8,9})) // [a b ba c ca cb cba d da]
    fmt.Println(largestString([]int{9,8,7,6,5,4,3,2,1})) // [da d cba cb ca c ba b a]

    fmt.Println(largestString1([]int{2,5,7})) // ["b","ca","cba"] 
    fmt.Println(largestString1([]int{3,9,1})) // ["ba","da","a"] 
    fmt.Println(largestString1([]int{1,2,3,4,5,6,7,8,9})) // [a b ba c ca cb cba d da]
    fmt.Println(largestString1([]int{9,8,7,6,5,4,3,2,1})) // [da d cba cb ca c ba b a]

    fmt.Println(largestString2([]int{2,5,7})) // ["b","ca","cba"] 
    fmt.Println(largestString2([]int{3,9,1})) // ["ba","da","a"] 
    fmt.Println(largestString2([]int{1,2,3,4,5,6,7,8,9})) // [a b ba c ca cb cba d da]
    fmt.Println(largestString2([]int{9,8,7,6,5,4,3,2,1})) // [da d cba cb ca c ba b a]
}