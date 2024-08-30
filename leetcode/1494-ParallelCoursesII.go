package main

// 1494. Parallel Courses II
// You are given an integer n, which indicates that there are n courses labeled from 1 to n. 
// You are also given an array relations where relations[i] = [prevCoursei, nextCoursei], 
// representing a prerequisite relationship between course prevCoursei and course nextCoursei: 
// course prevCoursei has to be taken before course nextCoursei.
// Also, you are given the integer k.

// In one semester, you can take at most k courses as long as you have taken all the prerequisites 
// in the previous semesters for the courses you are taking.

// Return the minimum number of semesters needed to take all courses. 
// The testcases will be generated such that it is possible to take every course.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/05/22/leetcode_parallel_courses_1.png" />
// Input: n = 4, relations = [[2,1],[3,1],[1,4]], k = 2
// Output: 3
// Explanation: The figure above represents the given graph.
// In the first semester, you can take courses 2 and 3.
// In the second semester, you can take course 1.
// In the third semester, you can take course 4.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/05/22/leetcode_parallel_courses_2.png" />
// Input: n = 5, relations = [[2,1],[3,1],[4,1],[1,5]], k = 2
// Output: 4
// Explanation: The figure above represents the given graph.
// In the first semester, you can only take courses 2 and 3 since you cannot take more than two per semester.
// In the second semester, you can take course 4.
// In the third semester, you can take course 1.
// In the fourth semester, you can take course 5.

// Constraints:
//     1 <= n <= 15
//     1 <= k <= n
//     0 <= relations.length <= n * (n-1) / 2
//     relations[i].length == 2
//     1 <= prevCoursei, nextCoursei <= n
//     prevCoursei != nextCoursei
//     All the pairs [prevCoursei, nextCoursei] are unique.
//     The given graph is a directed acyclic graph.

import "fmt"
import "math/bits"

func minNumberOfSemesters(n int, relations [][]int, k int) int {
    // prerequisites
    pre := make([]int, n)
    for _, relation := range relations {
        pre[relation[1]-1] |= 1 << (relation[0] - 1)
    }
    // deque of states, costs and set of added states
    states, costs, added := []int{0}, []int{0}, map[int]struct{}{ 0: struct{}{} }
    for cur := 0; cur < (1 << n); cur++ {
        // take current state from the deque
        state, cost := states[cur], costs[cur]
        if state == (1 << n) - 1 { // took all courses
            return cost
        }
        // find courses that are already unblocked and not yet taken 
        courses, next := []int{}, state
        for course := 0; course < n; course++ {
            bit := 1 << course
            if state & bit != 0 || pre[course] & state != pre[course] {
                continue
            }
            courses = append(courses, course)
            next |= bit
        }
        // if <= k, take them all
        if len(courses) <= k {
            if _, ok := added[next]; ok {
                continue
            }
            added[next] = struct{}{}
            states = append(states, next)
            costs = append(costs, cost+1)
            continue
        }
        // else take all subsets of size k
        var addStates func(next, k, from int)
        addStates = func(next, k, from int) {
            if k == 0 {
                if _, ok := added[next]; ok {
                    return
                }
                added[next] = struct{}{}
                states = append(states, next)
                costs = append(costs, cost+1)
                return
            }
            if k > len(courses)-from {
                return
            }
            for index := from; index < len(courses)-k+1; index++ {
                addStates(next|(1<<courses[index]), k-1, index+1)
            }
        }
        addStates(state, k, 0)
    }
    return 0
}

func minNumberOfSemesters1(n int, relations [][]int, k int) int {
    g := make([]int, n)
    for _, x := range relations {
        g[x[0]-1] |= 1<<(x[1]-1)
    }
    f := make([]int, 1 << n)
    for i := range f {
        f[i] = -1
    }
    u := (1 << n)-1
    var dfs func(int) int
    dfs = func(m int) int {
        s := u ^ m
        if s == 0 {
            return 0
        }
        if f[m] != -1 {
            return f[m]
        }
        for j, y := range g {
            if (y^m)&y != 0 {
                s ^= (1<<j)  
            }
        }
        res := n
        if bits.OnesCount(uint(s)) <= k {
            res = dfs(m|s) + 1
        } else {
            for sub := s; sub > 0; sub = s&(sub-1) {
                if bits.OnesCount(uint(sub)) == k {
                    res = min(res, dfs(m|sub) + 1)
                }
            }
        }
        f[m] = res
        return res
    }
    return dfs(0)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/05/22/leetcode_parallel_courses_1.png" />
    // Input: n = 4, relations = [[2,1],[3,1],[1,4]], k = 2
    // Output: 3
    // Explanation: The figure above represents the given graph.
    // In the first semester, you can take courses 2 and 3.
    // In the second semester, you can take course 1.
    // In the third semester, you can take course 4.
    fmt.Println(minNumberOfSemesters(4, [][]int{{2,1},{3,1},{1,4}}, 2)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/05/22/leetcode_parallel_courses_2.png" />
    // Input: n = 5, relations = [[2,1],[3,1],[4,1],[1,5]], k = 2
    // Output: 4
    // Explanation: The figure above represents the given graph.
    // In the first semester, you can only take courses 2 and 3 since you cannot take more than two per semester.
    // In the second semester, you can take course 4.
    // In the third semester, you can take course 1.
    // In the fourth semester, you can take course 5.
    fmt.Println(minNumberOfSemesters(5, [][]int{{2,1},{3,1},{4,1},{1,5}}, 2)) // 4

    fmt.Println(minNumberOfSemesters1(4, [][]int{{2,1},{3,1},{1,4}}, 2)) // 3
    fmt.Println(minNumberOfSemesters1(5, [][]int{{2,1},{3,1},{4,1},{1,5}}, 2)) // 4
}