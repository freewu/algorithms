package main

// 927. Three Equal Parts
// You are given an array arr which consists of only zeros and ones, 
// divide the array into three non-empty parts such that all of these parts represent the same binary value.

// If it is possible, return any [i, j] with i + 1 < j, such that:
//     arr[0], arr[1], ..., arr[i] is the first part,
//     arr[i + 1], arr[i + 2], ..., arr[j - 1] is the second part, and
//     arr[j], arr[j + 1], ..., arr[arr.length - 1] is the third part.
//     All three parts have equal binary values.

// If it is not possible, return [-1, -1].

// Note that the entire part is used when considering what binary value it represents. 
// For example, [1,1,0] represents 6 in decimal, not 3. Also, leading zeros are allowed, so [0,1,1] and [1,1] represent the same value.

// Example 1:
// Input: arr = [1,0,1,0,1]
// Output: [0,3]

// Example 2:
// Input: arr = [1,1,0,1,1]
// Output: [-1,-1]

// Example 3:
// Input: arr = [1,1,0,0,1]
// Output: [0,2]

// Constraints:
//     3 <= arr.length <= 3 * 10^4
//     arr[i] is 0 or 1

import "fmt"

func threeEqualParts(arr []int) []int {
    //if there are three equal parts, then the zero/one pattern of the number must be the same for the three number
    //first count the total one in the array
    //the number of ones for each number is cnt/=3
    //to get the exact number, traverse from the rightmost part
    //count one to the target number
    if len(arr) < 3 { return []int{-1, -1} }
    numOnes := 0
    for i := range arr {
        if arr[i] == 1 {
            numOnes++
        }
    }
    if numOnes == 0 { return []int{0, 2} } // 没有 1 出现
    if numOnes % 3 != 0 { return []int{-1, -1} } // 1 不够3等分
    count, r, curr := numOnes / 3, len(arr) - 1, 0
    for curr < count {
        if arr[r] == 1 { curr++ }
        r--
    }
    lp, l := len(arr) - r - 1, 0 // pattern length, then remove all leading zero for the leftmost part
    for arr[l] == 0 { l++ }
    for i := 0; i < lp; i++ { // check if the pattern matches (left and right)
        if arr[l+i] != arr[r+1+i] {
            return []int{-1, -1}
        }
    }
    m := l + lp
    for arr[m] == 0 { m++ } // remove all the leading zero for the middle part 
    for i := 0; i < lp; i++ { // check if the pattern matches (middle and right)
        if arr[m+i] != arr[r+1+i] { return []int{-1, -1} }
    }
    return []int{l + lp - 1, m + lp}
}

func threeEqualParts1(arr []int) []int {
    find := func(x int) int {
        s := 0
        for i, v := range arr {
            s += v
            if s == x { return i }
        }
        return 0
    }
    n, cnt := len(arr),  0
    for _, v := range arr {
        cnt += v
    }
    if cnt % 3 != 0 { return []int{ -1, -1} }
    if cnt == 0 { return []int{0, n - 1} }
    cnt /= 3
    i, j, k := find(1), find(cnt+1), find(cnt*2+1)
    for ; k < n && arr[i] == arr[j] && arr[j] == arr[k]; i, j, k = i+1, j+1, k+1 { }
    if k == n { return []int{i - 1, j} }
    return []int{-1, -1}
}

func main() {
    // Example 1:
    // Input: arr = [1,0,1,0,1]
    // Output: [0,3]
    fmt.Println(threeEqualParts([]int{1,0,1,0,1})) // [0,3]
    // Example 2:
    // Input: arr = [1,1,0,1,1]
    // Output: [-1,-1]
    fmt.Println(threeEqualParts([]int{1,1,0,1,1})) // [-1,-1]
    // Example 3:
    // Input: arr = [1,1,0,0,1]
    // Output: [0,2]
    fmt.Println(threeEqualParts([]int{1,1,0,0,1})) // [0,2]

    fmt.Println(threeEqualParts1([]int{1,0,1,0,1})) // [0,3]
    fmt.Println(threeEqualParts1([]int{1,1,0,1,1})) // [-1,-1]
    fmt.Println(threeEqualParts1([]int{1,1,0,0,1})) // [0,2]
}