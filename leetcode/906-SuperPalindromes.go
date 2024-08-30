package main

// 906. Super Palindromes
// Let's say a positive integer is a super-palindrome if it is a palindrome, 
// and it is also the square of a palindrome.

// Given two positive integers left and right represented as strings, 
// return the number of super-palindromes integers in the inclusive range [left, right].

// Example 1:
// Input: left = "4", right = "1000"
// Output: 4
// Explanation: 4, 9, 121, and 484 are superpalindromes.
// Note that 676 is not a superpalindrome: 26 * 26 = 676, but 26 is not a palindrome.

// Example 2:
// Input: left = "1", right = "2"
// Output: 1

// Constraints:
//     1 <= left.length, right.length <= 18
//     left and right consist of only digits.
//     left and right cannot have leading zeros.
//     left and right represent integers in the range [1, 10^18 - 1].
//     left is less than or equal to right.

import "fmt"
import "math"
import "strconv"

func superpalindromesInRange(left string, right string) int {
    res := 0
    isPalindrome := func(x int) bool {
        y := 0
        for ; x > y; x /= 10 {
            y *= 10
            y += x % 10
        }

        return x == y || x == y/10
    }
    sqrt := func (s string, roundFunc func(float64) float64) int {
        f, _ := strconv.ParseFloat(s, 64)
        return int(roundFunc(math.Sqrt(f)))
    }
    l, r := sqrt(left, math.Ceil), sqrt(right, math.Floor)
    countSuperpalindromes := func(isEven bool) {
        for i := 1; ; i++ {
            p, j := i, i
            if !isEven {
                j /= 10
            }
            for ; j > 0; j /= 10 {
                p *= 10
                p += j % 10
            }
            if p > r {
                break
            }
            if p >= l && isPalindrome(p * p) {
                res++
            }
        }
    }
    countSuperpalindromes(true)
    countSuperpalindromes(false)
    return res
}

func superpalindromesInRange1(left string, right string) int {
    l, _ := strconv.ParseInt(left, 10, 64)
    r, _ := strconv.ParseInt(right, 10, 64)
    res := 0
    arr := []int64{1,4,9,121,484,10201,12321,14641,40804,44944,1002001,1234321,4008004,100020001,102030201,104060401,121242121,123454321,125686521,400080004,404090404,10000200001,10221412201,12102420121,12345654321,40000800004,1000002000001,1002003002001,1004006004001,1020304030201,1022325232201,1024348434201,1210024200121,1212225222121,1214428244121,1232346432321,1234567654321,4000008000004,4004009004004,100000020000001,100220141022001,102012040210201,102234363432201,121000242000121,121242363242121,123212464212321,123456787654321,400000080000004,10000000200000001,10002000300020001,10004000600040001,10020210401202001,10022212521222001,10024214841242001,10201020402010201,10203040504030201,10205060806050201,10221432623412201,10223454745432201,12100002420000121,12102202520220121,12104402820440121,12122232623222121,12124434743442121,12321024642012321,12323244744232321,12343456865434321,12345678987654321,40000000800000004,40004000900040004}
    for i := 0; i < len(arr); i++ {
        if r < arr[i] { break }
        if l <= arr[i] {  res++ }
    }
    return res
}

func main() {
    // Example 1:
    // Input: left = "4", right = "1000"
    // Output: 4
    // Explanation: 4, 9, 121, and 484 are superpalindromes.
    // Note that 676 is not a superpalindrome: 26 * 26 = 676, but 26 is not a palindrome.
    fmt.Println(superpalindromesInRange("4", "1000")) // 4
    // Example 2:
    // Input: left = "1", right = "2"
    // Output: 1
    fmt.Println(superpalindromesInRange("1", "2")) // 1

    fmt.Println(superpalindromesInRange1("4", "1000")) // 4
    fmt.Println(superpalindromesInRange1("1", "2")) // 1
}