package main

// 374. Guess Number Higher or Lower
// We are playing the Guess Game. The game is as follows:
// I pick a number from 1 to n. You have to guess which number I picked.
// Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.
// You call a pre-defined API int guess(int num), which returns three possible results:
//     -1: Your guess is higher than the number I picked (i.e. num > pick).
//     1: Your guess is lower than the number I picked (i.e. num < pick).
//     0: your guess is equal to the number I picked (i.e. num == pick).

// Return the number that I picked.

// Example 1:
// Input: n = 10, pick = 6
// Output: 6

// Example 2:
// Input: n = 1, pick = 1
// Output: 1

// Example 3:
// Input: n = 2, pick = 1
// Output: 1

// Constraints:
//     1 <= n <= 2^31 - 1
//     1 <= pick <= n

import "fmt"

var pick int
func guess(num int) int {
    if pick == num {
        return 0
    }
    if pick > num {
        return 1
    }
    return -1
}
 /** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guessNumber(n int) int {
    left, right := 1, n
    for left < right {
        mid := left + (right-left)/2
        res := guess(mid)
        if res == 0 {
            return mid
        } else if res == -1 {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return left
}

func main() {
    pick = 6
    fmt.Println(guessNumber(10)) // 6
    pick = 1
    fmt.Println(guessNumber(1)) // 1
    pick = 1
    fmt.Println(guessNumber(2)) // 1
    pick = 2
    fmt.Println(guessNumber(2)) // 2
}