package main

// 2457. Minimum Addition to Make Integer Beautiful
// You are given two positive integers n and target.

// An integer is considered beautiful if the sum of its digits is less than or equal to target.

// Return the minimum non-negative integer x such that n + x is beautiful. 
// The input will be generated such that it is always possible to make n beautiful.

// Example 1:
// Input: n = 16, target = 6
// Output: 4
// Explanation: Initially n is 16 and its digit sum is 1 + 6 = 7. After adding 4, n becomes 20 and digit sum becomes 2 + 0 = 2. It can be shown that we can not make n beautiful with adding non-negative integer less than 4.

// Example 2:
// Input: n = 467, target = 6
// Output: 33
// Explanation: Initially n is 467 and its digit sum is 4 + 6 + 7 = 17. After adding 33, n becomes 500 and digit sum becomes 5 + 0 + 0 = 5. It can be shown that we can not make n beautiful with adding non-negative integer less than 33.

// Example 3:
// Input: n = 1, target = 1
// Output: 0
// Explanation: Initially n is 1 and its digit sum is 1, which is already smaller than or equal to target.

// Constraints:
//     1 <= n <= 10^12
//     1 <= target <= 150
//     The input will be generated such that it is always possible to make n beautiful.

import "fmt"
import "math"
import "strconv"

func makeIntegerBeautiful(n int64, target int) int64 {
    sum := 0
    s := strconv.FormatInt(n, 10)
    // Calculate sum of all digits
    for n > 0 {
        sum += int(n % 10)
        n /= 10
    }
    if sum <= target { return 0 }
    res, i, carry := int64(0), len(s) - 1, 0
    for sum > target {
        cur := int(s[i] - '0')
        // If both cur and carry are 0 (cur+carry==0), that means the current digit is already minimized,
        // there is nothing we can do to it, so skip. This if condition is only for understanding, it can also be removed
        if cur+carry == 0 {
            i--
            continue
        } else if cur+carry == 10 { // if cur==9 && carry == 1 (cur+carry==10), we reduce by sum by 9
            sum -= 9
        } else { // This cur has potential to be reduced to 0
            // The digit to be added for it to hit 10
            digit := 10 - cur - carry
            res += int64(math.Pow(10, float64(len(s)-1-i))) * int64(digit)
            // This line is the key! We need to minus 1 because remember: we reached ten "1"0,
            // so basically this 1 is added to the next (left) digit, and sum should add this 1.
            // But at the same time, we have carry that is inherited from last (right) digit,
            // so we should also reduce carry because carry (altogether with cur) is cleared
            sum -= (cur + carry - 1)
            carry = 1
        }
        i--
    }
    return res
}

func makeIntegerBeautiful1(n int64, target int) int64 {
    check := func() bool {
        cnt := 0
        for t := n; t > 0; t /= 10  {
            cnt += int(t % 10)
        }
        return cnt <= target
    }
    res := int64(0)
    for i := int64(10); !check(); i *= 10 {
        res += (10 - n % 10) * (i / 10)
        n = n / 10 + 1
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 16, target = 6
    // Output: 4
    // Explanation: Initially n is 16 and its digit sum is 1 + 6 = 7. After adding 4, n becomes 20 and digit sum becomes 2 + 0 = 2. It can be shown that we can not make n beautiful with adding non-negative integer less than 4.
    fmt.Println(makeIntegerBeautiful(16, 6)) // 4
    // Example 2:
    // Input: n = 467, target = 6
    // Output: 33
    // Explanation: Initially n is 467 and its digit sum is 4 + 6 + 7 = 17. After adding 33, n becomes 500 and digit sum becomes 5 + 0 + 0 = 5. It can be shown that we can not make n beautiful with adding non-negative integer less than 33.
    fmt.Println(makeIntegerBeautiful(467, 6)) // 33
    // Example 3:
    // Input: n = 1, target = 1
    // Output: 0
    // Explanation: Initially n is 1 and its digit sum is 1, which is already smaller than or equal to target.
    fmt.Println(makeIntegerBeautiful(1, 1)) // 0

    fmt.Println(makeIntegerBeautiful1(16, 6)) // 4
    fmt.Println(makeIntegerBeautiful1(467, 6)) // 33
    fmt.Println(makeIntegerBeautiful1(1, 1)) // 0
}