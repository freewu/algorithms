package main

// 753. Cracking the Safe
// There is a safe protected by a password. The password is a sequence of n digits where each digit can be in the range [0, k - 1].
// The safe has a peculiar way of checking the password. When you enter in a sequence, 
// it checks the most recent n digits that were entered each time you type a digit.
// For example, the correct password is "345" and you enter in "012345":
//     After typing 0, the most recent 3 digits is "0", which is incorrect.
//     After typing 1, the most recent 3 digits is "01", which is incorrect.
//     After typing 2, the most recent 3 digits is "012", which is incorrect.
//     After typing 3, the most recent 3 digits is "123", which is incorrect.
//     After typing 4, the most recent 3 digits is "234", which is incorrect.
//     After typing 5, the most recent 3 digits is "345", which is correct and the safe unlocks.

// Return any string of minimum length that will unlock the safe at some point of entering it.

// Example 1:
// Input: n = 1, k = 2
// Output: "10"
// Explanation: The password is a single digit, so enter each digit. "01" would also unlock the safe.

// Example 2:
// Input: n = 2, k = 2
// Output: "01100"
// Explanation: For each possible password:
// - "00" is typed in starting from the 4th digit.
// - "01" is typed in starting from the 1st digit.
// - "10" is typed in starting from the 3rd digit.
// - "11" is typed in starting from the 2nd digit.
// Thus "01100" will unlock the safe. "10011", and "11001" would also unlock the safe.
 
// Constraints:
//     1 <= n <= 4
//     1 <= k <= 10
//     1 <= kn <= 4096

import "fmt"
import "math"
import "strings"
import "strconv"

func crackSafe(n int, k int) string {
    const N1_STR = "0123456789"  // N * K ^ N
    if n == 1 {
        return N1_STR[:k]
    }
    seen, total := make(map[string]bool), int(math.Pow(float64(k), float64(n)))
    res := make([]byte, 0, total + n - 1)
    for i := 1; i != n; i++ {
        res = append(res, '0')
    }
    var dfs func(depth int) bool
    dfs = func(depth int) bool {
        if depth == 0 {
            return true // we did it
        }
        for i := 0; i != k; i++ {
            res = append(res, byte('0'+i))
            cur := string(res[len(res)-n:])
            if _, haveseen := seen[cur]; haveseen != true {
                seen[cur] = true
                if dfs(depth - 1) {
                    return true // only in this case do we not delete
                } else {
                    delete(seen, cur)
                }
            }
            res = res[0 : len(res)-1] // del
        }
        return false
    }
    dfs(total)
    return string(res)
}

func crackSafe1(n int, k int) string {
    res, total, visited := strings.Builder{}, int(math.Pow10(n - 1)), make(map[int]bool)
    var dfs func(node int)
    dfs = func(node int) {
        for i := 0; i < k; i++ {
            x := node * 10 + i
            if _, exist := visited[x]; !exist {
                visited[x] = true
                dfs(x % total)
                res.WriteString(strconv.Itoa(i))
            }
        }
    }
    dfs(0)
    for i := 1; i < n; i++ {
        res.WriteString("0")
    }
    return res.String()
}

func main() {
    // Example 1:
    // Input: n = 1, k = 2
    // Output: "10"
    // Explanation: The password is a single digit, so enter each digit. "01" would also unlock the safe.
    fmt.Println(crackSafe(1,2)) // "10"
    // Example 2:
    // Input: n = 2, k = 2
    // Output: "01100"
    // Explanation: For each possible password:
    // - "00" is typed in starting from the 4th digit.
    // - "01" is typed in starting from the 1st digit.
    // - "10" is typed in starting from the 3rd digit.
    // - "11" is typed in starting from the 2nd digit.
    // Thus "01100" will unlock the safe. "10011", and "11001" would also unlock the safe.
    fmt.Println(crackSafe(2,2)) // "01100"

    fmt.Println(crackSafe1(1,2)) // "10"
    fmt.Println(crackSafe1(2,2)) // "01100"
}