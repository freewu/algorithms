package main

// 面试题 17.05. Find Longest Subarray LCCI
// Given an array filled with letters and numbers, find the longest subarray with an equal number of letters and numbers.

// Return the subarray. 
// If there are more than one answer, return the one which has the smallest index of its left endpoint. 
// If there is no answer, return an empty arrary.

// Example 1:
// Input: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7","H","I","J","K","L","M"]
// Output: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7"]

// Example 2:
// Input: ["A","A"]
// Output: []

// Note:
//     array.length <= 100000

import "fmt"

func findLongestSubarray(array []string) []string {
    sum, mx := 0, 0
    res, mp := []string{}, map[int]int{0:-1} // 需要初始化一个哨兵，处理边界情况，例如[1,-1]
    isLetter := func (char byte) bool { return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' }
    for i, v := range array {
        if isLetter(v[0]) {
            sum++
        } else {
            sum--
        }
        if j, ok := mp[sum]; ok {
            n := i - j
            if n <= mx {
                // 不用max函数，区分开是否存在多个最长子数组，根据题目要求，像如果是相等的情况，就不能更新res，因为相同长度的数组，前面的左端点下标值更小
            } else { 
                res, mx = array[j + 1:i + 1], n
            }
        } else {
            mp[sum] = i
        }
    }
    return res
}

func findLongestSubarray1(array []string) []string {
    left, mx, score := 0, 0, 0
    mp := map[int]int{0: -1}
    isDigit := func(ch byte) bool { return '0' <= ch && ch <= '9' }
    for right := 0; right < len(array); right++ {
        if isDigit(array[right][0]) {
            score++
        } else {
            score--
        }
        if i, ok := mp[score]; ok {
            if mx < right - i {
                left, mx = i, right - i
            }
            continue
        }
        mp[score] = right
    }
    if mx == 0 { return []string{} }
    return array[left + 1 : left + mx + 1]
}

func main() {
    // Example 1:
    // Input: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7","H","I","J","K","L","M"]
    // Output: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7"]
    fmt.Println(findLongestSubarray([]string{"A","1","B","C","D","2","3","4","E","5","F","G","6","7","H","I","J","K","L","M"})) // ["A","1","B","C","D","2","3","4","E","5","F","G","6","7"]
    // Example 2:
    // Input: ["A","A"]
    // Output: []
    fmt.Println(findLongestSubarray([]string{"A","A"})) // []

    fmt.Println(findLongestSubarray1([]string{"A","1","B","C","D","2","3","4","E","5","F","G","6","7","H","I","J","K","L","M"})) // ["A","1","B","C","D","2","3","4","E","5","F","G","6","7"]
    fmt.Println(findLongestSubarray1([]string{"A","A"})) // []
}