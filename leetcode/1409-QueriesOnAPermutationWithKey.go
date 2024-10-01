package main

// 1409. Queries on a Permutation With Key
// Given the array queries of positive integers between 1 and m, 
// you have to process all queries[i] (from i=0 to i=queries.length-1) according to the following rules:
//     1. In the beginning, you have the permutation P=[1,2,3,...,m].
//     2. For the current i, find the position of queries[i] in the permutation P (indexing from 0) and then move this at the beginning of the permutation P. 
//        Notice that the position of queries[i] in P is the result for queries[i].

// Return an array containing the result for the given queries.

// Example 1:
// Input: queries = [3,1,2,1], m = 5
// Output: [2,1,2,1] 
// Explanation: The queries are processed as follow: 
// For i=0: queries[i]=3, P=[1,2,3,4,5], position of 3 in P is 2, then we move 3 to the beginning of P resulting in P=[3,1,2,4,5]. 
// For i=1: queries[i]=1, P=[3,1,2,4,5], position of 1 in P is 1, then we move 1 to the beginning of P resulting in P=[1,3,2,4,5]. 
// For i=2: queries[i]=2, P=[1,3,2,4,5], position of 2 in P is 2, then we move 2 to the beginning of P resulting in P=[2,1,3,4,5]. 
// For i=3: queries[i]=1, P=[2,1,3,4,5], position of 1 in P is 1, then we move 1 to the beginning of P resulting in P=[1,2,3,4,5]. 
// Therefore, the array containing the result is [2,1,2,1].  

// Example 2:
// Input: queries = [4,1,2,2], m = 4
// Output: [3,1,2,0]

// Example 3:
// Input: queries = [7,5,5,8,3], m = 8
// Output: [6,5,0,7,5]

// Constraints:
//     1 <= m <= 10^3
//     1 <= queries.length <= m
//     1 <= queries[i] <= m

import "fmt"

func processQueries(queries []int, m int) []int {
    res, permutation := make([]int, len(queries)),  make([]int, m)
    for i := 0; i < m; i++ {
        permutation[i] = 1 + i
    }
    for i, query := range queries {
        for j := 0; j < len(permutation); j++ {
            if permutation[j] == query {
                res[i] = j
                copy(permutation[1:j+1], permutation[:j])
                permutation[0] = query
                break
            }
        }
    }
    return res
}

type Tree []int

func(t Tree) pre(idx int) int {
    var sum int
    for ; idx > 0; idx -= idx & -idx {
        sum += t[idx]
    }
    return sum
}

func(t Tree) update(idx, val int) {
    for ; idx < len(t); idx += idx & -idx {
        t[idx] += val
    }
}

func processQueries1(queries []int, m int) []int {
    n := len(queries)
    tree := make(Tree, m + n + 1, m + n + 1)
    res, mp := []int{}, map[int]int{}
    for i := n; i < m + n; i += 1 {
        tree.update(i + 1, 1)
        mp[i - n + 1] = i + 1
    }        
    for i := 0; i < n; i += 1 {
        num := queries[i]
        index := mp[num]
        tree.update(index, -1)
        res = append(res, tree.pre(index))
        tree.update(n - i, 1)
        mp[num] = n - i
        //fmt.Println(res, mp, tree)
    }
    return res
}

func main() {
    // Example 1:
    // Input: queries = [3,1,2,1], m = 5
    // Output: [2,1,2,1] 
    // Explanation: The queries are processed as follow: 
    // For i=0: queries[i]=3, P=[1,2,3,4,5], position of 3 in P is 2, then we move 3 to the beginning of P resulting in P=[3,1,2,4,5]. 
    // For i=1: queries[i]=1, P=[3,1,2,4,5], position of 1 in P is 1, then we move 1 to the beginning of P resulting in P=[1,3,2,4,5]. 
    // For i=2: queries[i]=2, P=[1,3,2,4,5], position of 2 in P is 2, then we move 2 to the beginning of P resulting in P=[2,1,3,4,5]. 
    // For i=3: queries[i]=1, P=[2,1,3,4,5], position of 1 in P is 1, then we move 1 to the beginning of P resulting in P=[1,2,3,4,5]. 
    // Therefore, the array containing the result is [2,1,2,1].  
    fmt.Println(processQueries([]int{3,1,2,1}, 5)) // [2,1,2,1] 
    // Example 2:
    // Input: queries = [4,1,2,2], m = 4
    // Output: [3,1,2,0]
    fmt.Println(processQueries([]int{4,1,2,2}, 4)) // [3,1,2,0]
    // Example 3:
    // Input: queries = [7,5,5,8,3], m = 8
    // Output: [6,5,0,7,5]
    fmt.Println(processQueries([]int{7,5,5,8,3}, 8)) // [6,5,0,7,5]

    fmt.Println(processQueries1([]int{3,1,2,1}, 5)) // [2,1,2,1] 
    fmt.Println(processQueries1([]int{4,1,2,2}, 4)) // [3,1,2,0]
    fmt.Println(processQueries1([]int{7,5,5,8,3}, 8)) // [6,5,0,7,5]
}