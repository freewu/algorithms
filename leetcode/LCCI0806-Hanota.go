package main

// 面试题 08.06. Hanota LCCI
// In the classic problem of the Towers of Hanoi, you have 3 towers and N disks of different sizes which can slide onto any tower. 
// The puzzle starts with disks sorted in ascending order of size from top to bottom (i.e., each disk sits on top of an even larger one). 
// You have the following constraints:
//     (1) Only one disk can be moved at a time.
//     (2) A disk is slid off the top of one tower onto another tower.
//     (3) A disk cannot be placed on top of a smaller disk.

// Write a program to move the disks from the first tower to the last using stacks.

// Example1:
// Input: A = [2, 1, 0], B = [], C = []
// Output: C = [2, 1, 0]

// Example2:
// Input: A = [1, 0], B = [], C = []
// Output: C = [1, 0]

// Note:
//     A.length <= 14

import "fmt"

// 移动的过程中, 柱子的属性发生了变化 (例如 辅助柱子变为了目标柱子)
func hanota(A []int, B []int, C []int) []int {
    var dfs func(int, *[]int, *[]int, *[]int)
    dfs = func(n int, from, mid, to *[]int) {
        if n == 1 {
            *to = append(*to, (*from)[len(*from)-1])
            *from = (*from)[:len(*from)-1]
            return
        }
        dfs(n-1, from, to, mid) // 先将 n-1 个盘子从 from 通过 to 移动到 mid
        dfs(1, from, mid, to)   // 将 最后一个盘子移动到 to
        dfs(n-1, mid, from, to) // 将在 mid 的 n-1 个盘子通过 from 移动到 to
    }
    dfs(len(A), &A, &B, &C)
    return C
}

// class Solution {
//     public void hanota(List<Integer> A, List<Integer> B, List<Integer> C) {
//         C.addAll(A);
//     }
// }

func main() {
    // Example1:
    // Input: A = [2, 1, 0], B = [], C = []
    // Output: C = [2, 1, 0]
    a1 := []int{2, 1, 0}
    c1 := make([]int, 0)
    fmt.Println("before a1:", a1, " c1: ", c1)
    hanota(a1, []int{}, c1)
    fmt.Println("after a1:", a1, " c1: ", c1)
    // Example2:
    // Input: A = [1, 0], B = [], C = []
    // Output: C = [1, 0]
    a2 := []int{1, 0}
    c2 := make([]int, 0)
    fmt.Println("before a2:", a2, " c2: ", c2)
    hanota(a2, []int{}, c2)
    fmt.Println("after a2:", a2, " c2: ", c2)
}