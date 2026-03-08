package main

// 3863. Minimum Operations to Sort a String
// You are given a string s consisting of lowercase English letters.

// In one operation, you can select any substring of s that is not the entire string and sort it in non-descending alphabetical order.

// Return the minimum number of operations required to make s sorted in non-descending order. 
// If it is not possible, return -1.

// Example 1:
// Input: s = "dog"
// Output: 1
// Explanation:​​​​​​​
// Sort substring "og" to "go".
// Now, s = "dgo", which is sorted in ascending order. Thus, the answer is 1.

// Example 2:
// Input: s = "card"
// Output: 2
// Explanation:
// Sort substring "car" to "acr", so s = "acrd".
// Sort substring "rd" to "dr", making s = "acdr", which is sorted in ascending order. Thus, the answer is 2.

// Example 3:
// Input: s = "gf"
// Output: -1
// Explanation:
// It is impossible to sort s under the given constraints. Thus, the answer is -1.
 
// Constraints:
//     1 <= s.length <= 10^5
//     s consists of only lowercase English letters.

import "fmt"
import "slices"

func minOperations(s string) int {
    arr := []byte(s)
    if slices.IsSorted(arr) { return 0 } // s 已经是升序
    n := len(arr)
    if n == 2 { return -1 } // 长为 2，无法排序
    mn, mx := slices.Min(arr), slices.Max(arr)
    // 如果 s[0] 是最小值，排序 [1,n-1] 即可
    // 如果 s[n-1] 是最大值，排序 [0,n-2] 即可
    if arr[0] == mn || arr[n-1] == mx { return 1 }
    // 如果 [1,n-2] 中有最小值，那么先排序 [0,n-2]，把最小值排在最前面，然后排序 [1,n-1] 即可
    // 如果 [1,n-2] 中有最大值，那么先排序 [1,n-1]，把最大值排在最后面，然后排序 [0,n-2] 即可
    for _, v := range arr[1 : n-1] {   
        if v == mn || v == mx {
            return 2
        }
    }
    // 现在只剩下一种情况：s[0] 是最大值，s[n-1] 是最小值，且 [1,n-2] 不含最小值和最大值
    // 先排序 [0,n-2]，把最大值排到 n-2
    // 然后排序 [1,n-1]，把最大值排在最后面，且最小值排在 1
    // 最后排序 [0,n-2]，把最小值排在最前面
    return 3
}


func main() {
    // Example 1:
    // Input: s = "dog"
    // Output: 1
    // Explanation:​​​​​​​
    // Sort substring "og" to "go".
    // Now, s = "dgo", which is sorted in ascending order. Thus, the answer is 1.
    fmt.Println(minOperations("dog")) // 1
    // Example 2:
    // Input: s = "card"
    // Output: 2
    // Explanation:
    // Sort substring "car" to "acr", so s = "acrd".
    // Sort substring "rd" to "dr", making s = "acdr", which is sorted in ascending order. Thus, the answer is 2.
    fmt.Println(minOperations("card")) // 2
    // Example 3:
    // Input: s = "gf"
    // Output: -1
    // Explanation:
    // It is impossible to sort s under the given constraints. Thus, the answer is -1.
    fmt.Println(minOperations("gf")) // -1

    fmt.Println(minOperations("bluefrog")) // 1
    fmt.Println(minOperations("leetcode")) // 2
    fmt.Println(minOperations("abc")) // 0
    fmt.Println(minOperations("freewu")) // 2
}