package main

// 2097. Valid Arrangement of Pairs
// You are given a 0-indexed 2D integer array pairs where pairs[i] = [starti, endi]. 
// An arrangement of pairs is valid if for every index i where 1 <= i < pairs.length, we have endi-1 == starti.

// Return any valid arrangement of pairs.

// Note: The inputs will be generated such that there exists a valid arrangement of pairs.

// Example 1:
// Input: pairs = [[5,1],[4,5],[11,9],[9,4]]
// Output: [[11,9],[9,4],[4,5],[5,1]]
// Explanation:
// This is a valid arrangement since endi-1 always equals starti.
// end0 = 9 == 9 = start1 
// end1 = 4 == 4 = start2
// end2 = 5 == 5 = start3

// Example 2:
// Input: pairs = [[1,3],[3,2],[2,1]]
// Output: [[1,3],[3,2],[2,1]]
// Explanation:
// This is a valid arrangement since endi-1 always equals starti.
// end0 = 3 == 3 = start1
// end1 = 2 == 2 = start2
// The arrangements [[2,1],[1,3],[3,2]] and [[3,2],[2,1],[1,3]] are also valid.

// Example 3:
// Input: pairs = [[1,2],[1,3],[2,1]]
// Output: [[1,2],[2,1],[1,3]]
// Explanation:
// This is a valid arrangement since endi-1 always equals starti.
// end0 = 2 == 2 = start1
// end1 = 1 == 1 = start2
 
// Constraints:
//     1 <= pairs.length <= 10^5
//     pairs[i].length == 2
//     0 <= starti, endi <= 10^9
//     starti != endi
//     No two pairs are exactly the same.
//     There exists a valid arrangement of pairs.

import "fmt"

func validArrangement(pairs [][]int) [][]int {
    cnt, graph, start := make(map[int]int), make(map[int][]int), pairs[0][0]
    for _, p := range pairs {
        graph[p[0]] = append(graph[p[0]], p[1])
        cnt[p[0]]++
        cnt[p[1]]--
    }
    for vertex, inOutDegree := range cnt {
        if inOutDegree > 0 {
            start = vertex
            break
        }
    }
    stack, path := []int{start}, make([]int, 0, len(graph))
    for len(stack) > 0 {
        cur := stack[len(stack)-1]
        if len(graph[cur]) > 0 {
            last := len(graph[cur]) - 1
            stack = append(stack, graph[cur][last])
            graph[cur] = graph[cur][:last]
        } else {
            last := len(stack) - 1
            path = append(path, stack[last])
            stack = stack[:last]
        }
    }
    res := make([][]int, 0, len(graph))
    for i := len(path) - 1; i > 0; i-- {
        res = append(res, []int{path[i], path[i-1]})
    }
    return res
}

func main() {
    // Example 1:
    // Input: pairs = [[5,1],[4,5],[11,9],[9,4]]
    // Output: [[11,9],[9,4],[4,5],[5,1]]
    // Explanation:
    // This is a valid arrangement since endi-1 always equals starti.
    // end0 = 9 == 9 = start1 
    // end1 = 4 == 4 = start2
    // end2 = 5 == 5 = start3
    fmt.Println(validArrangement([][]int{{5,1},{4,5},{11,9},{9,4}})) // [[11,9],[9,4],[4,5],[5,1]]
    // Example 2:
    // Input: pairs = [[1,3],[3,2],[2,1]]
    // Output: [[1,3],[3,2],[2,1]]
    // Explanation:
    // This is a valid arrangement since endi-1 always equals starti.
    // end0 = 3 == 3 = start1
    // end1 = 2 == 2 = start2
    // The arrangements [[2,1],[1,3],[3,2]] and [[3,2],[2,1],[1,3]] are also valid.
    fmt.Println(validArrangement([][]int{{1,3},{3,2},{2,1}})) // [[1,3],[3,2],[2,1]]
    // Example 3:
    // Input: pairs = [[1,2],[1,3],[2,1]]
    // Output: [[1,2],[2,1],[1,3]]
    // Explanation:
    // This is a valid arrangement since endi-1 always equals starti.
    // end0 = 2 == 2 = start1
    // end1 = 1 == 1 = start2
    fmt.Println(validArrangement([][]int{{1,2},{1,3},{2,1}})) //  [[1,2],[2,1],[1,3]]
}