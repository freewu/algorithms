package main

// 2910. Minimum Number of Groups to Create a Valid Assignment
// You are given a collection of numbered balls and instructed to sort them into boxes for a nearly balanced distribution. 
// There are two rules you must follow:
//     1. Balls with the same box must have the same value. 
//        But, if you have more than one ball with the same number, you can put them in different boxes.
//     2. The biggest box can only have one more ball than the smallest box.

// ​Return the fewest number of boxes to sort these balls following these rules.

// Example 1:
// Input: balls = [3,2,3,2,3]
// Output: 2
// Explanation:
// We can sort balls into boxes as follows:
// [3,3,3]
// [2,2]
// The size difference between the two boxes doesn't exceed one.

// Example 2:
// Input: balls = [10,10,10,3,1,1]
// Output: 4
// Explanation:
// We can sort balls into boxes as follows:
// [10]
// [10,10]
// [3]
// [1,1]
// You can't use fewer than four boxes while still following the rules. 
// For example, putting all three balls numbered 10 in one box would break the rule about the maximum size difference between boxes.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func minGroupsForValidAssignment(balls []int) int {
    mn, n := len(balls), len(balls)
    freq := make(map[int]int)
    for _, v := range balls {
        freq[v]++
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, v := range freq {
        mn = min(mn, v)
    }
    check := func(m int) (bool, int) {
        sum := 0
        for _, v := range freq {
            remain, count := v % (m + 1), v / (m + 1)
            if remain != 0 && m - remain > count { return false, 0 }
            if remain != 0 { count++ }
            sum += count
        }
        return true, sum
    }
    for i := mn; i >= 1; i-- {
        if ok, res := check(i); ok {
            return res
        }
    }
    return n
}

func minGroupsForValidAssignment1(balls []int) int {
    mn, freq := len(balls), make(map[int]int)
    for _, v := range balls {
        freq[v]++
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, v := range freq {
        mn = min(mn, v)
    }
    right := mn // 最大盒子的大小范围[1, mn + 1]
    for ; ; right-- {
        res := 0 
        for _, v := range freq {
            if v / right < v % right {
                res = 0
                break
            }
            res += (v + right) / (right + 1)
        }
        if res > 0 {
            return res
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: balls = [3,2,3,2,3]
    // Output: 2
    // Explanation:
    // We can sort balls into boxes as follows:
    // [3,3,3]
    // [2,2]
    // The size difference between the two boxes doesn't exceed one.
    fmt.Println(minGroupsForValidAssignment([]int{3,2,3,2,3})) // 2
    // Example 2:
    // Input: balls = [10,10,10,3,1,1]
    // Output: 4
    // Explanation:
    // We can sort balls into boxes as follows:
    // [10]
    // [10,10]
    // [3]
    // [1,1]
    // You can't use fewer than four boxes while still following the rules. 
    // For example, putting all three balls numbered 10 in one box would break the rule about the maximum size difference between boxes.
    fmt.Println(minGroupsForValidAssignment([]int{10,10,10,3,1,1})) // 4

    fmt.Println(minGroupsForValidAssignment([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(minGroupsForValidAssignment([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(minGroupsForValidAssignment1([]int{3,2,3,2,3})) // 2
    fmt.Println(minGroupsForValidAssignment1([]int{10,10,10,3,1,1})) // 4
    fmt.Println(minGroupsForValidAssignment1([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(minGroupsForValidAssignment1([]int{9,8,7,6,5,4,3,2,1})) // 9
}