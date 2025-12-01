package main

// 2141. Maximum Running Time of N Computers
// You have n computers. 
// You are given the integer n and a 0-indexed integer array batteries where the ith battery can run a computer for batteries[i] minutes. 
// You are interested in running all n computers simultaneously using the given batteries.

// Initially, you can insert at most one battery into each computer. 
// After that and at any integer time moment, you can remove a battery from a computer and insert another battery any number of times. 
// The inserted battery can be a totally new battery or a battery from another computer. 
// You may assume that the removing and inserting processes take no time.

// Note that the batteries cannot be recharged.

// Return the maximum number of minutes you can run all the n computers simultaneously.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/01/06/example1-fit.png" />
// Input: n = 2, batteries = [3,3,3]
// Output: 4
// Explanation: 
// Initially, insert battery 0 into the first computer and battery 1 into the second computer.
// After two minutes, remove battery 1 from the second computer and insert battery 2 instead. Note that battery 1 can still run for one minute.
// At the end of the third minute, battery 0 is drained, and you need to remove it from the first computer and insert battery 1 instead.
// By the end of the fourth minute, battery 1 is also drained, and the first computer is no longer running.
// We can run the two computers simultaneously for at most 4 minutes, so we return 4.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/01/06/example2.png" />
// Input: n = 2, batteries = [1,1,1,1]
// Output: 2
// Explanation: 
// Initially, insert battery 0 into the first computer and battery 2 into the second computer. 
// After one minute, battery 0 and battery 2 are drained so you need to remove them and insert battery 1 into the first computer and battery 3 into the second computer. 
// After another minute, battery 1 and battery 3 are also drained so the first and second computers are no longer running.
// We can run the two computers simultaneously for at most 2 minutes, so we return 2.

// Constraints:
//     1 <= n <= batteries.length <= 10^5
//     1 <= batteries[i] <= 10^9

import "fmt"
import "sort"

func maxRunTime(n int, batteries []int) int64 {
    sort.Ints(batteries)
    live := make([]int64, n)
    for i := 0; i < n; i++{
        live[i] = int64(batteries[len(batteries) - n + i])
    }
    extra := int64(0)
    for i := 0; i < len(batteries) - n; i++ {
        extra += int64(batteries[i])
    }
    for i := 0; i < n-1; i++ {
        if extra < int64(i + 1) * (live[i + 1] - live[i]) {
            return live[i] + extra / int64(i + 1)
        }
        extra -= int64(i + 1) * (live[i + 1] - live[i])
    }
    return live[n - 1] + extra / int64(n)
}

func maxRunTime1(n int, batteries []int) int64 {
    sum := 0
    for _, battery := range batteries {
        sum += battery
    }
    var check = func(checkNum int) bool {
        num, reduce := n, 0
        for _, battery := range batteries {
            if battery >= checkNum {
                num--
            } else {
                reduce += battery
            }
            if reduce >= num*checkNum {
                return true
            }
        }
        return false
    }
    left, right := 0, sum
    for left < right {
        mid := (left + right + 1) >> 1
        if check(mid) {
            left = mid
        } else {
            right = mid - 1
        }
    }
    return int64(left)
}

func maxRunTime2(n int, batteries []int) int64 {
    res, sum := 0, 0
    for _, v := range batteries {
        sum += v
    }
    left, right := 0, sum / n
    for left <= right {
        total, mid := 0, left + (right - left) / 2
        for _, v := range batteries {
            if v < mid {
                total += v
            } else {
                total += mid
            }
        }
        if total >= n * mid {
            res, left = mid, mid + 1
        } else {
            right = mid - 1
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/01/06/example1-fit.png" />
    // Input: n = 2, batteries = [3,3,3]
    // Output: 4
    // Explanation: 
    // Initially, insert battery 0 into the first computer and battery 1 into the second computer.
    // After two minutes, remove battery 1 from the second computer and insert battery 2 instead. Note that battery 1 can still run for one minute.
    // At the end of the third minute, battery 0 is drained, and you need to remove it from the first computer and insert battery 1 instead.
    // By the end of the fourth minute, battery 1 is also drained, and the first computer is no longer running.
    // We can run the two computers simultaneously for at most 4 minutes, so we return 4.
    fmt.Println(maxRunTime(2, []int{3,3,3})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/01/06/example2.png" />
    // Input: n = 2, batteries = [1,1,1,1]
    // Output: 2
    // Explanation: 
    // Initially, insert battery 0 into the first computer and battery 2 into the second computer. 
    // After one minute, battery 0 and battery 2 are drained so you need to remove them and insert battery 1 into the first computer and battery 3 into the second computer. 
    // After another minute, battery 1 and battery 3 are also drained so the first and second computers are no longer running.
    // We can run the two computers simultaneously for at most 2 minutes, so we return 2.
    fmt.Println(maxRunTime(2, []int{1,1,1,1})) // 2

    fmt.Println(maxRunTime(2, []int{1,2,3,4,5,6,7,8,9})) // 22
    fmt.Println(maxRunTime(2, []int{9,8,7,6,5,4,3,2,1})) // 22

    fmt.Println(maxRunTime1(2, []int{3,3,3})) // 4
    fmt.Println(maxRunTime1(2, []int{1,1,1,1})) // 2
    fmt.Println(maxRunTime1(2, []int{1,2,3,4,5,6,7,8,9})) // 22
    fmt.Println(maxRunTime1(2, []int{9,8,7,6,5,4,3,2,1})) // 22

    fmt.Println(maxRunTime2(2, []int{3,3,3})) // 4
    fmt.Println(maxRunTime2(2, []int{1,1,1,1})) // 2
    fmt.Println(maxRunTime2(2, []int{1,2,3,4,5,6,7,8,9})) // 22
    fmt.Println(maxRunTime2(2, []int{9,8,7,6,5,4,3,2,1})) // 22
}