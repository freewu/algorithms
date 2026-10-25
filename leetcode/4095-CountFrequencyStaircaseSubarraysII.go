package main

// 4095. Count Frequency Staircase Subarrays II
// You are given an integer array nums.

// Consider the distinct values of a subarray in the order in which they first appear.

// If there are k distinct values, the subarray is called a frequency staircase if the first distinct value occurs exactly once, 
// the second distinct value occurs exactly twice, and so on, with the k^{th} distinct value occurring exactly k times.

// Equivalently, the frequencies of the distinct values, in the order of their first occurrence, must be exactly 1, 2, 3, ..., k.

// Return the number of frequency staircase subarrays in nums.

// A subarray is a contiguous non-empty sequence of elements within an array.

// The frequency of a value in a subarray is the number of times it occurs in that subarray.

// Example 1:
// Input: nums = [5,8,8,2,2,2]
// Output: 9
// Explanation:
// Every single-element subarray has frequencies [1], and there are 6 of them. 3 longer subarrays also qualify:
//  Subarray Values by first appearance Frequencies
//  nums[0...2] = [5, 8, 8] 5, then 8 [1, 2]
//  nums[2...4] = [8, 2, 2] 8, then 2 [1, 2]
//  nums[0...5] = [5, 8, 8, 2, 2, 2] 5, then 8, then 2 [1, 2, 3]
// So the answer is 6 + 3 = 9.

// Example 2:
// Input: nums = [6,6,9,9]
// Output: 5
// Explanation:
// Every single-element subarray has frequencies [1], and there are 4 of them. 1 longer subarray also qualifies:
//  Subarray Values by first appearance Frequencies
//  nums[1...3] = [6, 9, 9] 6, then 9 [1, 2]
// So the answer is 4 + 1 = 5.

// Example 3:
// Input: nums = [2,3,4,4,3,4]
// Output: 8
// Explanation:
// Every single-element subarray has frequencies [1], and there are 6 of them. 2 longer subarrays also qualify:
//  Subarray Values by first appearance Frequencies
//  nums[1...3] = [3, 4, 4] 3, then 4 [1, 2]
//  nums[0...5] = [2, 3, 4, 4, 3, 4] 2, then 3, then 4 [1, 2, 3]
// So the answer is 6 + 2 = 8.

// Constraints:
//     - 1 <= nums.length <= 5 * 10^{4}
//     - 1 <= nums[i] <= 10^{5}

import "fmt"
import "slices"

func countStaircase(nums []int) int {
    res, n, mx := 0, len(nums), slices.Max(nums)
    pos := make([][]int, mx + 1)
    indexPos := make([]int, n) // nums[i] 在 pos[nums[i]] 中的下标
    for i, v := range nums {
        indexPos[i] = len(pos[v])
        pos[v] = append(pos[v], i)
    }
    firsts := []int{}
    firstPos := make([]int, mx+1)
    for i := n - 1; i >= 0; i-- {
        v := nums[i]
        firstPos[v] = i
        // 把 firsts 中的 x 移到最前面
        j := slices.Index(firsts, v)
        if j != 0 { // 优化：j = 0 时无需移动
            if j >= 0 {
                firsts = append(firsts[:j], firsts[j+1:]...)
            }
            firsts = slices.Insert(firsts, 0, v)
        }
        m := len(firsts)
        if m*(m+1)/2 > n { // firsts 不需要这么多元素
            firsts = firsts[:len(firsts)-1]
        }
        mx1, mn1 := 0, n
        r := i // 子数组右开端点
        for count, v := range firsts {
            r += count + 1
            // [i,r) 必须包含恰好 cnt+1 个 x
            ps := pos[v]
            k := indexPos[firstPos[v]] + count
            if k >= len(ps) { // [i,n) 中的 x 的个数小于 cnt+1
                break
            }
            // 更新 mx 和 mn 后，必须满足 mx < r <= mn
            if k < len(ps)-1 {
                mn1 = min(mn1, ps[k+1])
            }
            if r > mn1 { // 继续遍历，r 更大，mn 更小，无法找到合法子数组
                break
            }
            mx1 = max(mx1, ps[k])
            if mx1 < r {
                res++
            }
        }
    }
    return res
}

const MX = 100_001

var last [MX]int32

func countStaircase1(nums []int) int {
    res, n := 0, len(nums)
    nxt := make([]int32, n)
    for i := range last {
        last[i] = int32(n)
    }
    for i := n - 1; i >= 0; i-- {
        nxt[i] = last[nums[i]]
        last[nums[i]] = int32(i)
    }
    stamp, ord, count := make([]int32, MX), make([]int32, MX), make([]int32, MX)
    for i := range n {
        st := int32(i + 1)
        rem := n - i
        var d, ok int32
        for r := i; r < n; r++ {
            v := nums[r]
            if stamp[v] != st {
                d++
                if int(d)*(int(d)+1)/2 > rem {
                    break
                }
                stamp[v], ord[v], count[v] = st, d, 0
            }
            j := ord[v]
            if count[v] == j {
                ok--
            }
            count[v]++
            if count[v] > j {
                break
            }
            if count[v] == j {
                ok++
            } else if int(nxt[r]) == n {
                break
            }
            if ok == d {
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,8,8,2,2,2]
    // Output: 9
    // Explanation:
    // Every single-element subarray has frequencies [1], and there are 6 of them. 3 longer subarrays also qualify:
    //  Subarray Values by first appearance Frequencies
    //  nums[0...2] = [5, 8, 8] 5, then 8 [1, 2]
    //  nums[2...4] = [8, 2, 2] 8, then 2 [1, 2]
    //  nums[0...5] = [5, 8, 8, 2, 2, 2] 5, then 8, then 2 [1, 2, 3]
    // So the answer is 6 + 3 = 9.
    fmt.Println(countStaircase([]int{5,8,8,2,2,2})) // 9
    // Example 2:
    // Input: nums = [6,6,9,9]
    // Output: 5
    // Explanation:
    // Every single-element subarray has frequencies [1], and there are 4 of them. 1 longer subarray also qualifies:
    //  Subarray Values by first appearance Frequencies
    //  nums[1...3] = [6, 9, 9] 6, then 9 [1, 2]
    // So the answer is 4 + 1 = 5.
    fmt.Println(countStaircase([]int{6,6,9,9})) // 5
    // Example 3:
    // Input: nums = [2,3,4,4,3,4]
    // Output: 8
    // Explanation:
    // Every single-element subarray has frequencies [1], and there are 6 of them. 2 longer subarrays also qualify:
    //  Subarray Values by first appearance Frequencies
    //  nums[1...3] = [3, 4, 4] 3, then 4 [1, 2]
    //  nums[0...5] = [2, 3, 4, 4, 3, 4] 2, then 3, then 4 [1, 2, 3]
    // So the answer is 6 + 2 = 8.  
    fmt.Println(countStaircase([]int{2,3,4,4,3,4})) // 8

    fmt.Println(countStaircase([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(countStaircase([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(countStaircase1([]int{5,8,8,2,2,2})) // 9
    fmt.Println(countStaircase1([]int{6,6,9,9})) // 5
    fmt.Println(countStaircase1([]int{2,3,4,4,3,4})) // 8
    fmt.Println(countStaircase1([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(countStaircase1([]int{9,8,7,6,5,4,3,2,1})) // 9
}