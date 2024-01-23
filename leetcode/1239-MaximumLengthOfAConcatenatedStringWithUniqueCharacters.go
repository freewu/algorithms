package main

import "fmt"

// 1239. Maximum Length of a Concatenated String with Unique Characters
// You are given an array of strings arr. 
// A string s is formed by the concatenation of a subsequence of arr that has unique characters.
// Return the maximum possible length of s.
// A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.

// Example 1:
// Input: arr = ["un","iq","ue"]
// Output: 4
// Explanation: All the valid concatenations are:
// - ""
// - "un"
// - "iq"
// - "ue"
// - "uniq" ("un" + "iq")
// - "ique" ("iq" + "ue")
// Maximum length is 4.

// Example 2:
// Input: arr = ["cha","r","act","ers"]
// Output: 6
// Explanation: Possible longest valid concatenations are "chaers" ("cha" + "ers") and "acters" ("act" + "ers").

// Example 3:
// Input: arr = ["abcdefghijklmnopqrstuvwxyz"]
// Output: 26
// Explanation: The only string in arr has all 26 characters.
 
// Constraints:
// 		1 <= arr.length <= 16
// 		1 <= arr[i].length <= 26
// 		arr[i] contains only lowercase English letters.

import (
	"math/bits"
)

func maxLength(arr []string) int {
	c, res := []uint32{}, 0
	for _, s := range arr {
		var mask uint32
		for _, c := range s {
			mask = mask | 1<<(c-'a')
		}
		if len(s) != bits.OnesCount32(mask) { // 如果字符串本身带有重复的字符，需要排除
			continue
		}
		c = append(c, mask)
	}
	dfs(c, 0, 0, &res)
	return res
}

func dfs(c []uint32, index int, mask uint32, res *int) {
	*res = max(*res, bits.OnesCount32(mask))
	for i := index; i < len(c); i++ {
		if mask&c[i] == 0 {
			dfs(c, i+1, mask|c[i], res)
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// best solution 
func maxLength1(arr []string) int {
    mask := make([][]int, 0, len(arr))
    helper := func(s string) (int, bool) {
        m := 0
        for _, c := range s {
            t := 1 << (c - 'a')
            if m & t != 0 {
                return 0, false
            }
            m |= t
        }
        return m, true
    }
    for i, s := range arr {
        if m, ok := helper(s); ok {
            mask = append(mask, []int{m, i})
        }
    }
    ans, n := 0, len(mask)
    var f func(int, int, int)
    f = func(state, sum, idx int) {
        if idx == n {
            return
        }
        if state & mask[idx][0] == 0 {
            t := sum + len(arr[mask[idx][1]])
            ans = max(ans, t)
            f(state | mask[idx][0], t, idx + 1)
        }
        f(state, sum, idx + 1)
    }
    f(0, 0, 0)
    return ans
}

func main() {
	fmt.Println(maxLength([]string{"un","iq","ue"})) // 4
	fmt.Println(maxLength([]string{"cha","r","act","ers"})) // 6
	fmt.Println(maxLength([]string{"abcdefghijklmnopqrstuvwxyz"})) // 26

	fmt.Println(maxLength1([]string{"un","iq","ue"})) // 4
	fmt.Println(maxLength1([]string{"cha","r","act","ers"})) // 6
	fmt.Println(maxLength1([]string{"abcdefghijklmnopqrstuvwxyz"})) // 26
}