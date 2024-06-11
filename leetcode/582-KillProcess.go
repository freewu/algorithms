package main

// 582. Kill Process
// You have n processes forming a rooted tree structure. 
// You are given two integer arrays pid and ppid, where pid[i] is the ID of the ith process and ppid[i] is the ID of the ith process's parent process.

// Each process has only one parent process but may have multiple children processes. 
// Only one process has ppid[i] = 0, which means this process has no parent process (the root of the tree).

// When a process is killed, all of its children processes will also be killed.

// Given an integer kill representing the ID of a process you want to kill, 
// return a list of the IDs of the processes that will be killed. You may return the answer in any order.

// Example 1:
//         3
//       /   \
//      1   (10)
//           /
//         (5)
// <img src="https://assets.leetcode.com/uploads/2021/02/24/ptree.jpg" />
// Input: pid = [1,3,10,5], ppid = [3,0,5,3], kill = 5
// Output: [5,10]
// Explanation: The processes colored in red are the processes that should be killed.

// Example 2:
// Input: pid = [1], ppid = [0], kill = 1
// Output: [1]

// Constraints:
//     n == pid.length
//     n == ppid.length
//     1 <= n <= 5 * 10^4
//     1 <= pid[i] <= 5 * 10^4
//     0 <= ppid[i] <= 5 * 10^4
//     Only one process has no parent.
//     All the values of pid are unique.
//     kill is guaranteed to be in pid.

import "fmt"

// 并查集
func killProcess(pid []int, ppid []int, kill int) []int {
    res, m := []int{}, make(map[int]int)
    for i, p := range pid {
        m[p] = i
    }
    var find func(p int)bool
    find = func(p int)bool{
        if p == kill {
            return true
        }
        if p == 0 {
            return false
        }
        ppi := ppid[m[p]]
        if find(ppi) {
            return true
        }
        return false
    }
    for _, p := range pid {
        if find(p){
            res = append(res, p)
        }
    }
    return res
}

// bfs
func killProcess1(pid []int, ppid []int, kill int) []int {
    res, m := []int{}, make(map[int][]int)
    for i,p := range ppid {
        m[p] = append(m[p], pid[i])
    }
    q := []int{ kill }
    for len(q) > 0{
        t := []int{}
        for i := range q {
            t = append(t, m[q[i]]...)
        }
        res = append(res, q...)
        q = t
    }
    return res
}

func main() {
    // Example 1:
    //         3
    //       /   \
    //      1   (10)
    //           /
    //         (5)
    // <img src="https://assets.leetcode.com/uploads/2021/02/24/ptree.jpg" />
    // Input: pid = [1,3,10,5], ppid = [3,0,5,3], kill = 5
    // Output: [5,10]
    // Explanation: The processes colored in red are the processes that should be killed.
    fmt.Println(killProcess([]int{1,3,10,5},[]int{3,0,5,3},5)) // [5,10]
    // Example 2:
    // Input: pid = [1], ppid = [0], kill = 1
    // Output: [1]
    fmt.Println(killProcess([]int{1},[]int{0},1)) // [1]

    fmt.Println(killProcess1([]int{1,3,10,5},[]int{3,0,5,3},5)) // [5,10]
    fmt.Println(killProcess1([]int{1},[]int{0},1)) // [1]
}