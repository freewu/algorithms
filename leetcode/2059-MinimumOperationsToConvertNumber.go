package main

// 2059. Minimum Operations to Convert Number
// You are given a 0-indexed integer array nums containing distinct numbers, an integer start, and an integer goal. 
// There is an integer x that is initially set to start, and you want to perform operations on x such that it is converted to goal. 
// You can perform the following operation repeatedly on the number x:

// If 0 <= x <= 1000, then for any index i in the array (0 <= i < nums.length), you can set x to any of the following:
//     x + nums[i]
//     x - nums[i]
//     x ^ nums[i] (bitwise-XOR)

// Note that you can use each nums[i] any number of times in any order. 
// Operations that set x to be out of the range 0 <= x <= 1000 are valid, but no more operations can be done afterward.

// Return the minimum number of operations needed to convert x = start into goal, and -1 if it is not possible.

// Example 1:
// Input: nums = [2,4,12], start = 2, goal = 12
// Output: 2
// Explanation: We can go from 2 → 14 → 12 with the following 2 operations.
// - 2 + 12 = 14
// - 14 - 2 = 12

// Example 2:
// Input: nums = [3,5,7], start = 0, goal = -4
// Output: 2
// Explanation: We can go from 0 → 3 → -4 with the following 2 operations. 
// - 0 + 3 = 3
// - 3 - 7 = -4
// Note that the last operation sets x out of the range 0 <= x <= 1000, which is valid.

// Example 3:
// Input: nums = [2,8,16], start = 0, goal = 1
// Output: -1
// Explanation: There is no way to convert 0 into 1.

// Constraints:
//     1 <= nums.length <= 1000
//     -10^9 <= nums[i], goal <= 10^9
//     0 <= start <= 1000
//     start != goal
//     All the integers in nums are distinct.

import "fmt"

func minimumOperations(nums []int, start int, goal int) int {
    res, queue, visited := 1, []int{ start }, make(map[int]bool)
    visited[start] = true
    for len(queue) > 0 {
        n := len(queue)
        for i := 0; i < n;i++{
            cur := queue[i]
            for _, v := range nums {
                val := cur + v
                if val == goal { return res }
                if val >= 0 && val <= 1000 && !visited[val] {
                    visited[val] = true
                    queue = append(queue, val)
                }
                val = cur - v
                if val == goal{ return res }
                if val >= 0 && val <= 1000 && !visited[val] {
                    visited[val] = true
                    queue = append(queue, val)
                }
                val = cur ^ v
                if val == goal { return res }
                if val >= 0 && val <= 1000 && !visited[val] {
                    visited[val] = true
                    queue = append(queue, val)
                }
            }
        }
        queue = queue[n:]
        res++
    }
    return -1
}

// 三种运算  + -  或者异或 1 0 1 0 1 1 分情况讨论 有的话 直接输出2  
func minimumOperations1(nums []int, start int, goal int) int {
    visited, queue := [1001]bool{}, []int{ start }
    visited[start] = true
    for step := 1; queue != nil; step++ {
        tmp := queue 
        queue = nil
        for _, v := range tmp {
            for _, num := range nums {
                for _, x := range []int{ v + num, v - num,v ^ num } {
                    if x == goal { return step  }
                    // 如果不是目标数 判断是不是超过1000. 或者小于 0
                    if 0 <= x && x <= 1000 && !visited[x] {
                        visited[x] = true 
                        queue = append(queue,x)
                    }
                }
            }
        }
    }
    return -1
}

func minimumOperations2(nums []int, start int, goal int) int {
    type Node struct {
        value, count int
        next     *Node
    }
    visited, ops := make([]bool, 1001), make([]int, 3)
    queue := &Node{ value: start, count: 0 }
    tail := queue
    visited[start] = true
    for ; queue != nil; queue = queue.next {
        for _, n := range nums {
            ops[0], ops[1], ops[2] = queue.value + n,  queue.value - n, queue.value ^ n
            for _, v := range ops {
                if v == goal { return queue.count + 1 }
                if v >= 0 && v <= 1000 && !visited[v] {
                    tail.next = &Node{ value: v, count: queue.count + 1 }
                    tail = tail.next
                    visited[v] = true
                }
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [2,4,12], start = 2, goal = 12
    // Output: 2
    // Explanation: We can go from 2 → 14 → 12 with the following 2 operations.
    // - 2 + 12 = 14
    // - 14 - 2 = 12
    fmt.Println(minimumOperations([]int{2,4,12}, 2, 12)) // 2
    // Example 2:
    // Input: nums = [3,5,7], start = 0, goal = -4
    // Output: 2
    // Explanation: We can go from 0 → 3 → -4 with the following 2 operations. 
    // - 0 + 3 = 3
    // - 3 - 7 = -4
    // Note that the last operation sets x out of the range 0 <= x <= 1000, which is valid.
    fmt.Println(minimumOperations([]int{3,5,7}, 0, -4)) // 2
    // Example 3:
    // Input: nums = [2,8,16], start = 0, goal = 1
    // Output: -1
    // Explanation: There is no way to convert 0 into 1.
    fmt.Println(minimumOperations([]int{2,8,16}, 0, 1)) // -1

    fmt.Println(minimumOperations1([]int{2,4,12}, 2, 12)) // 2
    fmt.Println(minimumOperations1([]int{3,5,7}, 0, -4)) // 2
    fmt.Println(minimumOperations1([]int{2,8,16}, 0, 1)) // -1

    fmt.Println(minimumOperations2([]int{2,4,12}, 2, 12)) // 2
    fmt.Println(minimumOperations2([]int{3,5,7}, 0, -4)) // 2
    fmt.Println(minimumOperations2([]int{2,8,16}, 0, 1)) // -1
}