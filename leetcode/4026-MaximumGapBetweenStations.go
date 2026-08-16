package main

// 4026. Maximum Gap Between Stations
// You are given two strings skill and station of lengths n and m, respectively.

// skill[i] represents the skill of worker i, and station[j] represents the skill supported by station j.

// You must assign every worker to a distinct station. Let ji be the index of the station assigned to worker i. 
// A valid assignment must satisfy:
//     1. station[ji] == skill[i] for every 0 <= i < n.
//     2. The assigned station indices must be strictly increasing in worker order, meaning j0 < j1 < ... < jn - 1.

// The gap of an assignment is the maximum difference between the station indices assigned to two consecutive workers. 
// In other words, it is max(ji - ji - 1) over all 1 <= i < n.

// If there is only one worker, the gap is 0.

// Return the maximum possible gap among all valid assignments. It is guaranteed that at least one valid assignment exists.

// Example 1:
// Input: skill = "aa", station = "aaaa"
// Output: 3
// Explanation:
// The two workers must be assigned to two different 'a' stations.
// Assigning them to stations [0, 3] gives a gap of 3.

// Example 2:
// Input: skill = "xyz", station = "xyzz"
// Output: 2
// Explanation:
// Assign worker 0 to station j = 0, and worker 1 to station j = 1.
// To maximize the gap, assign worker 2 to station j = 3.
// This gives the assignment [0, 1, 3] with gaps [1, 2], so the gap is 2.

// Example 3:
// Input: skill = "cbc", station = "cbcdbc"
// Output: 4
// Explanation:
// Assign worker 0 to station j = 0, and worker 1 to station j = 1.
// To maximize the gap, assign worker 2 to station j = 5.
// This gives the assignment [0, 1, 5] with gaps [1, 4], so the gap is 4.

// Constraints:
//     skill.length == n
//     station.length == m
//     1 <= n <= m <= 10^5
//     skill and station consist of lowercase English letters.
//     It is guaranteed that a valid assignment exists for every worker.

import "fmt"
import "sort"

func maximumGap(skill string, station string) int {
    res, n, m := 0, len(skill), len(station)
    if n <= 1 {
        return 0
    }
    pos := make([][]int, 26)
    for i := 0; i < 26; i++ {
        pos[i] = []int{}
    }
    for i, ch := range station {
        pos[ch-'a'] = append(pos[ch-'a'], i)
    }
    prev, left := -1, make([]int, n)
    for i := 0; i < n; i++ {
        c := skill[i] - 'a'
        arr := pos[c]
        index := sort.Search(len(arr), func(j int) bool { 
            return arr[j] > prev 
        })
        left[i] = arr[index]
        prev = left[i]
    }
    next, right := m, make([]int, n)
    for i := n - 1; i >= 0; i-- {
        c := skill[i] - 'a'
        arr := pos[c]
        index := sort.Search(len(arr), func(j int) bool { 
            return arr[j] >= next 
        })
        index--
        right[i] = arr[index]
        next = right[i] 
    }
    for i := 1; i < n; i++ {
        gap := right[i] - left[i-1]
        if gap > res {
            res = gap
        }
    }
    return res
}

func maximumGap1(skill string, station string) int {
    n, m := len(station),len(skill)
    last := make([]int, n)
    for i := range n {
        last[i] = -1
    }
    j := n - 1
    for i := m - 1; i >= 0; i-- {
        for skill[i] != station[j] {
            j--
        }
        last[j] = i
        j--
    }
    res, low, high := 0, 0, 0
    for i := range m-1 {
        for skill[i] != station[low] {
            low++
        }
        for i+1 != last[high] {
            high++
        }
        res = max(res, high - low)
        low++
    }
    return res
}

func main() {
    // Example 1:
    // Input: skill = "aa", station = "aaaa"
    // Output: 3
    // Explanation:
    // The two workers must be assigned to two different 'a' stations.
    // Assigning them to stations [0, 3] gives a gap of 3.
    fmt.Println(maximumGap("aa", "aaaa")) // 3 
    // Example 2:
    // Input: skill = "xyz", station = "xyzz"
    // Output: 2
    // Explanation:
    // Assign worker 0 to station j = 0, and worker 1 to station j = 1.
    // To maximize the gap, assign worker 2 to station j = 3.
    // This gives the assignment [0, 1, 3] with gaps [1, 2], so the gap is 2.
    fmt.Println(maximumGap("xyz", "xyzz")) // 2
    // Example 3:
    // Input: skill = "cbc", station = "cbcdbc"
    // Output: 4
    // Explanation:
    // Assign worker 0 to station j = 0, and worker 1 to station j = 1.
    // To maximize the gap, assign worker 2 to station j = 5.
    // This gives the assignment [0, 1, 5] with gaps [1, 4], so the gap is 4.  
    fmt.Println(maximumGap("cbc", "cbcdbc")) // 4

    fmt.Println(maximumGap("let", "leetcode")) // 2
    fmt.Println(maximumGap("bfg", "bluefrog")) // 2
    fmt.Println(maximumGap("fw", "freewu")) // 4

    fmt.Println(maximumGap1("aa", "aaaa")) // 3 
    fmt.Println(maximumGap1("xyz", "xyzz")) // 2
    fmt.Println(maximumGap1("cbc", "cbcdbc")) // 4
    fmt.Println(maximumGap1("let", "leetcode")) // 2
    fmt.Println(maximumGap1("bfg", "bluefrog")) // 2
    fmt.Println(maximumGap1("fw", "freewu")) // 4
}