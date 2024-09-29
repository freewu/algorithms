package main

// 1505. Minimum Possible Integer After at Most K Adjacent Swaps On Digits
// You are given a string num representing the digits of a very large integer and an integer k. 
// You are allowed to swap any two adjacent digits of the integer at most k times.

// Return the minimum integer you can obtain also as a string.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/06/17/q4_1.jpg" />
// Input: num = "4321", k = 4
// Output: "1342"
// Explanation: The steps to obtain the minimum integer from 4321 with 4 adjacent swaps are shown.

// Example 2:
// Input: num = "100", k = 1
// Output: "010"
// Explanation: It's ok for the output to have leading zeros, but the input is guaranteed not to have any leading zeros.

// Example 3:
// Input: num = "36789", k = 1000
// Output: "36789"
// Explanation: We can keep the number without any swaps.

// Constraints:
//     1 <= num.length <= 3 * 10^4
//     num consists of only digits and does not contain leading zeros.
//     1 <= k <= 10^9

import "fmt"
import "strings"

func minInteger(num string, k int) string {
    targetIndex, swaps, n := 0,  k, len(num)
    res, bytes := make([]byte, n), []byte(num)
    // will return index -1 if none found
    findMinDigit := func(arr []byte, threshold byte) (byte, int) {
        b, index := threshold, -1
        for i, ch := range arr {
            if ch < b {
                index, b = i, ch
            }
        }
        return b, index
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for swaps > 0 && targetIndex <= n - 1 {
        rangeEnd, target := min(targetIndex + swaps, n-1), bytes[targetIndex]
        b, index := findMinDigit(bytes[targetIndex:rangeEnd + 1], target)
        if index > -1 {
            index += targetIndex
            res[targetIndex] = b
            copy(bytes[targetIndex+1:index+1], bytes[targetIndex:index])
            // perform all swaps 
            swaps -= index-targetIndex 
        } else { // no swap
            res[targetIndex] = bytes[targetIndex]
        }
        targetIndex++
    }
    copy(res[targetIndex:], bytes[targetIndex:])
    return string(res)
}

func minInteger1(num string, k int) string {
    stat := [10][]int{}
    for i := 0; i < len(num); i++ {
        c := num[i]
        stat[c-'0'] = append(stat[c-'0'], i)
    }
    bb := make([]int, len(num) + 1)
    update := func(i int) {
        i++
        for i < len(bb) {
            bb[i]++
            i += i & (-i)
        }
    }
    sums := func(i int) int {
        tt := 0
        i++
        for i > 0 {
            tt += bb[i]
            i -= i & (-i)
        }
        return tt
    }
    indice := [10]int{}
    res := make([]byte, len(num))
    for i := 0; i < len(num); i++ {
        for j := 0; j < 10; j++ {
            index := indice[j]
            if index >= len(stat[j]) { continue }
            cd := stat[j][index]
            ss := sums(cd)
            tmp := cd - ss
            if tmp > k { continue }
            k -= tmp
            update(cd)
            res[i] = num[cd]
            indice[j]++
            break
        }
    }
    return string(res)
}

func minInteger2(num string, k int) string {
    n := len(num)
    visit, tree, digit := make([]bool, n), make([]int, n + 1), [10][]int{}
    for i := 0; i < n; i++ {
        digit[num[i]-'0'] = append(digit[num[i]-'0'], i)
    }
    add := func(tree []int, x, v int) {
        for x < len(tree) {
            tree[x] += v
            x += x&-x
        }
    }
    getPre := func (tree []int, x int) int {
        var sum int
        for x>0 {
            sum += tree[x]
            x -= x&-x
        }
        return sum
    }
    getSum := func (tree []int, a, b int) int {
        return getPre(tree, b) - getPre(tree, a-1)
    }
    var str strings.Builder
    for i := 0; i < n;  {
        if visit[i] {
            i++
            continue
        }
        if k <= 0 {
            str.WriteByte(num[i])
            i++
            continue
        }
        for j := 0; j <= int(num[i] - '0'); j++ {
            if len(digit[j]) == 0 { continue }
            t := digit[j][0]
            tmpk := t - i - getSum(tree, i+1, t+1)
            if tmpk > k { continue }
            digit[j] = digit[j][1:]
            k -= tmpk
            visit[t] = true
            str.WriteByte(num[t])
            add(tree, t+1, 1)
            break
        }
    }
    return str.String()
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/06/17/q4_1.jpg" />
    // Input: num = "4321", k = 4
    // Output: "1342"
    // Explanation: The steps to obtain the minimum integer from 4321 with 4 adjacent swaps are shown.
    fmt.Println(minInteger("4321", 4)) // "1342"
    // Example 2:
    // Input: num = "100", k = 1
    // Output: "010"
    // Explanation: It's ok for the output to have leading zeros, but the input is guaranteed not to have any leading zeros.
    fmt.Println(minInteger("100", 1)) // "010"
    // Example 3:
    // Input: num = "36789", k = 1000
    // Output: "36789"
    // Explanation: We can keep the number without any swaps.
    fmt.Println(minInteger("36789", 1000)) // "36789"

    fmt.Println(minInteger1("4321", 4)) // "1342"
    fmt.Println(minInteger1("100", 1)) // "010"
    fmt.Println(minInteger1("36789", 1000)) // "36789"

    fmt.Println(minInteger2("4321", 4)) // "1342"
    fmt.Println(minInteger2("100", 1)) // "010"
    fmt.Println(minInteger2("36789", 1000)) // "36789"
}