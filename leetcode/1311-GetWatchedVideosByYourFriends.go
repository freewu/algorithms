package main

// 1311. Get Watched Videos by Your Friends
// There are n people, each person has a unique id between 0 and n-1. 
// Given the arrays watchedVideos and friends, where watchedVideos[i] and friends[i] contain the list of watched videos 
// and the list of friends respectively for the person with id = i.

// Level 1 of videos are all watched videos by your friends, 
// level 2 of videos are all watched videos by the friends of your friends and so on. 
// In general, the level k of videos are all watched videos by people with the shortest path exactly equal to k with you.
// Given your id and the level of videos, return the list of videos ordered by their frequencies (increasing). 
// For videos with the same frequency order them alphabetically from least to greatest. 

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/01/02/leetcode_friends_1.png" />
// Input: watchedVideos = [["A","B"],["C"],["B","C"],["D"]], friends = [[1,2],[0,3],[0,3],[1,2]], id = 0, level = 1
// Output: ["B","C"] 
// Explanation: 
// You have id = 0 (green color in the figure) and your friends are (yellow color in the figure):
// Person with id = 1 -> watchedVideos = ["C"] 
// Person with id = 2 -> watchedVideos = ["B","C"] 
// The frequencies of watchedVideos by your friends are: 
// B -> 1 
// C -> 2

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/01/02/leetcode_friends_2.png" />
// Input: watchedVideos = [["A","B"],["C"],["B","C"],["D"]], friends = [[1,2],[0,3],[0,3],[1,2]], id = 0, level = 2
// Output: ["D"]
// Explanation: 
// You have id = 0 (green color in the figure) and the only friend of your friends is the person with id = 3 (yellow color in the figure).

// Constraints:
//     n == watchedVideos.length == friends.length
//     2 <= n <= 100
//     1 <= watchedVideos[i].length <= 100
//     1 <= watchedVideos[i][j].length <= 8
//     0 <= friends[i].length < n
//     0 <= friends[i][j] < n
//     0 <= id < n
//     1 <= level < n
//     if friends[i] contains j, then friends[j] contains i

import "fmt"
import "sort"

// bfs
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
    n := len(friends)
    visited, freq := make([]int, n), make(map[string]int)
    queue := make([]int, 0, 10)
    queue = append(queue, friends[id]...)
    visited[id] = 1
    for _, f := range friends[id] {
        visited[f] = 1
    }
    for len(queue) > 0 && level > 0 {
        level--
        for i := len(queue) - 1; i >= 0; i-- {
            f := queue[0]
            queue = queue[1:]
            if level == 0 {
                for _, v := range watchedVideos[f] {
                    freq[v]++
                }
            }
            for _, fof := range friends[f] { // 朋友的朋友
                if visited[fof] == 0 {
                    queue = append(queue, fof)
                    visited[fof] = 1
                }
            } 
        }  
    }
    res := make([]string, 0, len(freq))
    for k, _ := range freq {
        res = append(res, k)
    }
    sort.Slice(res, func(i, j int) bool {
        if freq[res[i]] == freq[res[j]] {
            return res[i] < res[j]
        }
        return freq[res[i]] < freq[res[j]]
    })
    return res
}

func watchedVideosByFriends1(watchedVideos [][]string, friends [][]int, id int, level int) []string {
    visited := make([]bool, len(friends))
    queue := []int{ id }
    visited[id] = true
    
    for ; level > 0; level-- {
        for n := len(queue); n > 0; n-- {
            for _, v := range friends[queue[0]] {
                if !visited[v] {
                    queue = append(queue, v)
                    visited[v] = true
                }
            }
            queue = queue[1:len(queue)] // pop
        }
    }
    count := map[string]int{}
    for _, r := range queue {
        for _, c := range watchedVideos[r] {
            count[c]++
        }
    }
    type pair struct {
        n int
        s string
    }
    ps := []pair{}
    for s, n := range count {
        ps = append(ps, pair{n, s})
    }
    sort.Slice(ps, func(i, j int) bool {
        return ps[i].n < ps[j].n || ps[i].n == ps[j].n && ps[i].s < ps[j].s
    })
    res := make([]string, 0, len(ps))
    for _, v := range ps {
        res = append(res, v.s)
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/01/02/leetcode_friends_1.png" />
    // Input: watchedVideos = [["A","B"],["C"],["B","C"],["D"]], friends = [[1,2],[0,3],[0,3],[1,2]], id = 0, level = 1
    // Output: ["B","C"] 
    // Explanation: 
    // You have id = 0 (green color in the figure) and your friends are (yellow color in the figure):
    // Person with id = 1 -> watchedVideos = ["C"] 
    // Person with id = 2 -> watchedVideos = ["B","C"] 
    // The frequencies of watchedVideos by your friends are: 
    // B -> 1 
    // C -> 2
    fmt.Println(watchedVideosByFriends([][]string{{"A","B"},{"C"},{"B","C"},{"D"}}, [][]int{{1,2},{0,3},{0,3},{1,2}}, 0, 1)) // ["B","C"] 
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/01/02/leetcode_friends_2.png" />
    // Input: watchedVideos = [["A","B"],["C"],["B","C"],["D"]], friends = [[1,2],[0,3],[0,3],[1,2]], id = 0, level = 2
    // Output: ["D"]
    // Explanation: 
    // You have id = 0 (green color in the figure) and the only friend of your friends is the person with id = 3 (yellow color in the figure).
    fmt.Println(watchedVideosByFriends([][]string{{"A","B"},{"C"},{"B","C"},{"D"}}, [][]int{{1,2},{0,3},{0,3},{1,2}}, 0, 2)) // ["D"]

    fmt.Println(watchedVideosByFriends1([][]string{{"A","B"},{"C"},{"B","C"},{"D"}}, [][]int{{1,2},{0,3},{0,3},{1,2}}, 0, 1)) // ["B","C"] 
    fmt.Println(watchedVideosByFriends1([][]string{{"A","B"},{"C"},{"B","C"},{"D"}}, [][]int{{1,2},{0,3},{0,3},{1,2}}, 0, 2)) // ["D"]
}