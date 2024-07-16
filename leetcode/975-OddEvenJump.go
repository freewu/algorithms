package main

// 975. Odd Even Jump
// You are given an integer array arr. 
// From some starting index, you can make a series of jumps. 
// The (1st, 3rd, 5th, ...) jumps in the series are called odd-numbered jumps, and the (2nd, 4th, 6th, ...) jumps in the series are called even-numbered jumps. 
// Note that the jumps are numbered, not the indices.

// You may jump forward from index i to index j (with i < j) in the following way:
//     During odd-numbered jumps (i.e., jumps 1, 3, 5, ...), you jump to the index j such that arr[i] <= arr[j] and arr[j] is the smallest possible value. If there are multiple such indices j, you can only jump to the smallest such index j.
//     During even-numbered jumps (i.e., jumps 2, 4, 6, ...), you jump to the index j such that arr[i] >= arr[j] and arr[j] is the largest possible value. If there are multiple such indices j, you can only jump to the smallest such index j.
//     It may be the case that for some index i, there are no legal jumps.

// A starting index is good if, starting from that index, you can reach the end of the array (index arr.length - 1) by jumping some number of times (possibly 0 or more than once).
// Return the number of good starting indices.

// Example 1:
// Input: arr = [10,13,12,14,15]
// Output: 2
// Explanation: 
// From starting index i = 0, we can make our 1st jump to i = 2 (since arr[2] is the smallest among arr[1], arr[2], arr[3], arr[4] that is greater or equal to arr[0]), then we cannot jump any more.
// From starting index i = 1 and i = 2, we can make our 1st jump to i = 3, then we cannot jump any more.
// From starting index i = 3, we can make our 1st jump to i = 4, so we have reached the end.
// From starting index i = 4, we have reached the end already.
// In total, there are 2 different starting indices i = 3 and i = 4, where we can reach the end with some number of
// jumps.

// Example 2:
// Input: arr = [2,3,1,1,4]
// Output: 3
// Explanation: 
// From starting index i = 0, we make jumps to i = 1, i = 2, i = 3:
// During our 1st jump (odd-numbered), we first jump to i = 1 because arr[1] is the smallest value in [arr[1], arr[2], arr[3], arr[4]] that is greater than or equal to arr[0].
// During our 2nd jump (even-numbered), we jump from i = 1 to i = 2 because arr[2] is the largest value in [arr[2], arr[3], arr[4]] that is less than or equal to arr[1]. arr[3] is also the largest value, but 2 is a smaller index, so we can only jump to i = 2 and not i = 3
// During our 3rd jump (odd-numbered), we jump from i = 2 to i = 3 because arr[3] is the smallest value in [arr[3], arr[4]] that is greater than or equal to arr[2].
// We can't jump from i = 3 to i = 4, so the starting index i = 0 is not good.
// In a similar manner, we can deduce that:
// From starting index i = 1, we jump to i = 4, so we reach the end.
// From starting index i = 2, we jump to i = 3, and then we can't jump anymore.
// From starting index i = 3, we jump to i = 4, so we reach the end.
// From starting index i = 4, we are already at the end.
// In total, there are 3 different starting indices i = 1, i = 3, and i = 4, where we can reach the end with some
// number of jumps.

// Example 3:
// Input: arr = [5,1,3,4,2]
// Output: 3
// Explanation: We can reach the end from starting indices 1, 2, and 4.

// Constraints:
//     1 <= arr.length <= 2 * 10^4
//     0 <= arr[i] < 10^5

import "fmt"
import "sort"

// 超出时间限制 63 / 65 
func oddEvenJumps1(arr []int) int {
    res, n := 0, len(arr)
    dp1, dp2 := make([]bool, n), make([]bool, n)
    dp1[n-1], dp2[n-1] = true, true
    for i:= n - 2; i >= 0; i-- {
        jodd, jeven := -1, -1
        for j := i+1; j < n; j++ {
            if arr[j]>= arr[i] {
                if jodd == -1 || arr[j]< arr[jodd] || arr[j] == arr[jodd] && j < jodd {
                    jodd = j
                    dp1[i] = dp2[j]
                }
            }
            if arr[j] <= arr[i] {
                if jeven == -1 || arr[j] > arr[jeven] || arr[j] == arr[jeven] && j < jeven {
                    jeven = j
                    dp2[i] = dp1[j]
                }
            }
        }
    }
    for _, v := range dp1 {
        if v { 
            res++ 
        }
    }
    return res
}

func oddEvenJumps(arr []int) int {
    type Node struct {
        val int
        idx int
    }
    var jump func(idx, flag int, n []Node, odd, even map[int]bool) bool
    jump = func(idx, flag int, n []Node, odd, even map[int]bool) bool {
        if idx >= len(n) { return false }
        if n[idx].idx == len(n)-1 { return true }
        if flag % 2 == 1 {
            if res, ok := odd[idx]; ok { return res }
            for i := idx + 1; i < len(n); i++ {
                if n[i].idx < n[idx].idx { continue }
                odd[idx] = jump(i, flag+1, n, odd, even)
                return odd[idx]
            }
            odd[idx] = false
            return odd[idx]
        } else {
            if res, ok := even[idx]; ok { return res }
            for i := idx + 1; i < len(n); i++ {
                if n[i].val == n[idx].val {
                    even[idx] = jump(i, flag+1, n, odd, even)
                    return even[idx]
                }
            }
            for i := idx - 1; i >= 0; i-- {
                if n[i].idx < n[idx].idx { continue }
                for j := i - 1; j >= 0; j-- {
                    if n[j].idx < n[idx].idx { continue }
                    if n[j].val != n[i].val { break }
                    i = j
                }
                even[idx] = jump(i, flag+1, n, odd, even)
                return even[idx]
            }
            even[idx] = false
            return even[idx]
        }
        return false
    }
    n := make([]Node, len(arr))
    for i := 0; i < len(arr); i++ {
        n[i] = Node{arr[i], i}
    }
    sort.Slice(n, func(i, j int) bool {
        if n[i].val == n[j].val {
            return n[i].idx < n[j].idx
        }
        return n[i].val < n[j].val
    })
    res, odd, even := 0, make(map[int]bool), make(map[int]bool)
    for i := 0; i < len(n); i++ {
        if success, ok := odd[i]; ok {
            if success {
                res++
            }
            continue
        }
        if jump(i, 1, n, odd, even) {
            res++
        }
    }
    return res
}

func oddEvenJumps2(arr []int) int {
    res, n := 1, len(arr)
    up, down := make([]int, n), make([]int, n)
    for i := 0; i<n;i++ {
        up[i], down[i] = i, i
    }
    sort.Slice(up, func(a, b int) bool {
        if arr[up[a]] > arr[up[b]] || (arr[up[a]] == arr[up[b]] && up[a] > up[b]) {
            return true
        }
        return false
    })
    sort.Slice(down, func(a, b int) bool{
        if arr[down[a]] < arr[down[b]] || (arr[down[a]] == arr[down[b]] && down[a] > down[b]) {
            return true
        }
        return false
    })

    big, small, s := make([]int, n), make([]int, n), make([]int, 0)
    for _, v := range up {
        for len(s) > 0 && s[len(s)-1] < v {
            s = s[:len(s)-1]
        }
        if len(s) > 0 {
            big[v] = s[len(s)-1]
        }
        s = append(s, v)
    }
    s = make([]int, 0)
    for _, v := range down {
        for len(s) > 0 && s[len(s)-1] < v  {
            s = s[:len(s)-1]
        }
        if len(s) > 0 {
            small[v] = s[len(s)-1]
        }
        s = append(s, v)
    }
    o,j := make([]bool, n), make([]bool, n)
    o[n-1], j[n-1] = true, true
    for i:= n-2; i>=0; i-- {
        if big[i] > 0 {
            j[i] = o[big[i]]
        }
        if small[i] > 0 {
            o[i] = j[small[i]]
        }
        if j[i] {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [10,13,12,14,15]
    // Output: 2
    // Explanation: 
    // From starting index i = 0, we can make our 1st jump to i = 2 (since arr[2] is the smallest among arr[1], arr[2], arr[3], arr[4] that is greater or equal to arr[0]), then we cannot jump any more.
    // From starting index i = 1 and i = 2, we can make our 1st jump to i = 3, then we cannot jump any more.
    // From starting index i = 3, we can make our 1st jump to i = 4, so we have reached the end.
    // From starting index i = 4, we have reached the end already.
    // In total, there are 2 different starting indices i = 3 and i = 4, where we can reach the end with some number of
    // jumps.
    fmt.Println(oddEvenJumps([]int{10,13,12,14,15})) // 2
    // Example 2:
    // Input: arr = [2,3,1,1,4]
    // Output: 3
    // Explanation: 
    // From starting index i = 0, we make jumps to i = 1, i = 2, i = 3:
    // During our 1st jump (odd-numbered), we first jump to i = 1 because arr[1] is the smallest value in [arr[1], arr[2], arr[3], arr[4]] that is greater than or equal to arr[0].
    // During our 2nd jump (even-numbered), we jump from i = 1 to i = 2 because arr[2] is the largest value in [arr[2], arr[3], arr[4]] that is less than or equal to arr[1]. arr[3] is also the largest value, but 2 is a smaller index, so we can only jump to i = 2 and not i = 3
    // During our 3rd jump (odd-numbered), we jump from i = 2 to i = 3 because arr[3] is the smallest value in [arr[3], arr[4]] that is greater than or equal to arr[2].
    // We can't jump from i = 3 to i = 4, so the starting index i = 0 is not good.
    // In a similar manner, we can deduce that:
    // From starting index i = 1, we jump to i = 4, so we reach the end.
    // From starting index i = 2, we jump to i = 3, and then we can't jump anymore.
    // From starting index i = 3, we jump to i = 4, so we reach the end.
    // From starting index i = 4, we are already at the end.
    // In total, there are 3 different starting indices i = 1, i = 3, and i = 4, where we can reach the end with some
    // number of jumps.
    fmt.Println(oddEvenJumps([]int{2,3,1,1,4})) // 3
    // Example 3:
    // Input: arr = [5,1,3,4,2]
    // Output: 3
    // Explanation: We can reach the end from starting indices 1, 2, and 4.
    fmt.Println(oddEvenJumps([]int{5,1,3,4,2})) // 3

    fmt.Println(oddEvenJumps1([]int{10,13,12,14,15})) // 2
    fmt.Println(oddEvenJumps1([]int{2,3,1,1,4})) // 3
    fmt.Println(oddEvenJumps1([]int{5,1,3,4,2})) // 3

    fmt.Println(oddEvenJumps2([]int{10,13,12,14,15})) // 2
    fmt.Println(oddEvenJumps2([]int{2,3,1,1,4})) // 3
    fmt.Println(oddEvenJumps2([]int{5,1,3,4,2})) // 3
}