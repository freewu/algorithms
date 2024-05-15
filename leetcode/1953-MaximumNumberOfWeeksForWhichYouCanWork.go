package main

// 1953. Maximum Number of Weeks for Which You Can Work
// There are n projects numbered from 0 to n - 1. 
// You are given an integer array milestones where each milestones[i] denotes the number of milestones the ith project has.
// You can work on the projects following these two rules:
//     Every week, you will finish exactly one milestone of one project. You must work every week.
//     You cannot work on two milestones from the same project for two consecutive weeks.

// Once all the milestones of all the projects are finished, 
// or if the only milestones that you can work on will cause you to violate the above rules, you will stop working. 
// Note that you may not be able to finish every project's milestones due to these constraints.

// Return the maximum number of weeks you would be able to work on the projects without violating the rules mentioned above.

// Example 1:
// Input: milestones = [1,2,3]
// Output: 6
// Explanation: One possible scenario is:
// ​​​​- During the 1st week, you will work on a milestone of project 0.
// - During the 2nd week, you will work on a milestone of project 2.
// - During the 3rd week, you will work on a milestone of project 1.
// - During the 4th week, you will work on a milestone of project 2.
// - During the 5th week, you will work on a milestone of project 1.
// - During the 6th week, you will work on a milestone of project 2.
// The total number of weeks is 6.

// Example 2:
// Input: milestones = [5,2,1]
// Output: 7
// Explanation: One possible scenario is:
// - During the 1st week, you will work on a milestone of project 0.
// - During the 2nd week, you will work on a milestone of project 1.
// - During the 3rd week, you will work on a milestone of project 0.
// - During the 4th week, you will work on a milestone of project 1.
// - During the 5th week, you will work on a milestone of project 0.
// - During the 6th week, you will work on a milestone of project 2.
// - During the 7th week, you will work on a milestone of project 0.
// The total number of weeks is 7.
// Note that you cannot work on the last milestone of project 0 on 8th week because it would violate the rules.
// Thus, one milestone in project 0 will remain unfinished.
 
// Constraints:
//     n == milestones.length
//     1 <= n <= 10^5
//     1 <= milestones[i] <= 10^9

import "fmt"

func numberOfWeeks(milestones []int) int64 {
    res, mx, sum := 0, 0, 0
    if len(milestones) < 2 {
        return 1
    }
    for i := 0; i < len(milestones); i++ { // 统计出最大的任数 & 总和任务数据
        if milestones[i] > mx {
            mx = milestones[i]
        }
        sum += milestones[i]
    }
    if sum - mx >= mx - 1 {
        res = sum
    } else {
        res = (sum - mx) * 2 + 1 // 任务不能连续，必须休息一周
    }
    return int64(res)
}

func numberOfWeeks1(milestones []int) int64 {
    longest, rest := 0, 0 // 耗时最长工作所需周数
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range milestones {
        longest = max(longest, v)
        rest += v
    }
    rest -= longest // 其余工作共计所需周数
    if longest > rest + 1 { // 此时无法完成所耗时最长的工作
        return int64(rest * 2 + 1)
    }
    return int64(longest + rest)  // 此时可以完成所有工作
}

func main() {
    // Example 1:
    // Input: milestones = [1,2,3]
    // Output: 6
    // Explanation: One possible scenario is:
    // ​​​​- During the 1st week, you will work on a milestone of project 0.
    // - During the 2nd week, you will work on a milestone of project 2.
    // - During the 3rd week, you will work on a milestone of project 1.
    // - During the 4th week, you will work on a milestone of project 2.
    // - During the 5th week, you will work on a milestone of project 1.
    // - During the 6th week, you will work on a milestone of project 2.
    // The total number of weeks is 6.
    fmt.Println(numberOfWeeks([]int{1,2,3})) // 6
    // Example 2:
    // Input: milestones = [5,2,1]
    // Output: 7
    // Explanation: One possible scenario is:
    // - During the 1st week, you will work on a milestone of project 0.
    // - During the 2nd week, you will work on a milestone of project 1.
    // - During the 3rd week, you will work on a milestone of project 0.
    // - During the 4th week, you will work on a milestone of project 1.
    // - During the 5th week, you will work on a milestone of project 0.
    // - During the 6th week, you will work on a milestone of project 2.
    // - During the 7th week, you will work on a milestone of project 0.
    // The total number of weeks is 7.
    // Note that you cannot work on the last milestone of project 0 on 8th week because it would violate the rules.
    // Thus, one milestone in project 0 will remain unfinished.
    fmt.Println(numberOfWeeks([]int{5,2,1})) // 7

    fmt.Println(numberOfWeeks1([]int{1,2,3})) // 6
    fmt.Println(numberOfWeeks1([]int{5,2,1})) // 7
}