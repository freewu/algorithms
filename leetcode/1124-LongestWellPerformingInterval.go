package main

// 1124. Longest Well-Performing Interval
// We are given hours, a list of the number of hours worked per day for a given employee.

// A day is considered to be a tiring day if and only if the number of hours worked is (strictly) greater than 8.

// A well-performing interval is an interval of days for which the number of tiring days is strictly larger than the number of non-tiring days.

// Return the length of the longest well-performing interval.

// Example 1:
// Input: hours = [9,9,6,0,6,6,9]
// Output: 3
// Explanation: The longest well-performing interval is [9,9,6].

// Example 2:
// Input: hours = [6,6,6]
// Output: 0
 
// Constraints:
//     1 <= hours.length <= 10^4
//     0 <= hours[i] <= 16

import "fmt"

func longestWPI(hours []int) int {
    res, count, n := 0, 0, len(hours)
    mp := make(map[int]int)
    for i := 0; i < n; i++ {
        if hours[i] > 8 { // 当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」
            count++
        } else {
            count--
        }
        if count > 0 {
            res = i + 1
        } else {
            if v, ok := mp[count - 1]; ok {
                if i - v > res {
                    res = i - v
                }
            }
            if _, ok := mp[count]; !ok {
                mp[count] = i
            }
        }
    }
    return res
}

func longestWPI1(hours []int) int {
    res, n := 0, len(hours)
    prefixSum := make([]int, n+1)
    stack := []int{0} // 栈初始值为0，表示前缀和初始点

    for i := 1; i <= n; i++ {
        if hours[i-1] > 8 { // 当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」
            prefixSum[i] = prefixSum[i-1] + 1
        } else {
            prefixSum[i] = prefixSum[i-1] - 1
        }
        // 记录 prefixSum 出现的位置，只保留最早的，栈为单调递减
        if prefixSum[stack[len(stack)-1]] > prefixSum[i] {
            stack = append(stack, i)
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n; i > 0; i-- { // 从右往左遍历，寻找最大长度的符合条件的子序列
        for len(stack) > 0 && prefixSum[i] > prefixSum[stack[len(stack)-1]] {
            res = max(res, i - stack[len(stack)-1])
            stack = stack[:len(stack)-1]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: hours = [9,9,6,0,6,6,9]
    // Output: 3
    // Explanation: The longest well-performing interval is [9,9,6].
    fmt.Println(longestWPI([]int{9,9,6,0,6,6,9})) // 3
    // Example 2:
    // Input: hours = [6,6,6]
    // Output: 0
    fmt.Println(longestWPI([]int{6,6,6})) // 0

    fmt.Println(longestWPI1([]int{9,9,6,0,6,6,9})) // 3
    fmt.Println(longestWPI1([]int{6,6,6})) // 0
}