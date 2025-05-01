package main

// 838. Push Dominoes
// There are n dominoes in a line, and we place each domino vertically upright. 
// In the beginning, we simultaneously push some of the dominoes either to the left or to the right.

// After each second, each domino that is falling to the left pushes the adjacent domino on the left. 
// Similarly, the dominoes falling to the right push their adjacent dominoes standing on the right.

// When a vertical domino has dominoes falling on it from both sides, it stays still due to the balance of the forces.

// For the purposes of this question, we will consider that a falling domino expends no additional force to a falling or already fallen domino.

// You are given a string dominoes representing the initial state where:
//     dominoes[i] = 'L', if the ith domino has been pushed to the left,
//     dominoes[i] = 'R', if the ith domino has been pushed to the right, and
//     dominoes[i] = '.', if the ith domino has not been pushed.

// Return a string representing the final state.

// Example 1:
// Input: dominoes = "RR.L"
// Output: "RR.L"
// Explanation: The first domino expends no additional force on the second domino.

// Example 2:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/05/18/domino.png" />
// Input: dominoes = ".L.R...LR..L.."
// Output: "LL.RR.LLRRLL.."

// Constraints:
//     n == dominoes.length
//     1 <= n <= 10^5
//     dominoes[i] is either 'L', 'R', or '.'.

import "fmt"

func pushDominoes(dominoes string) string {
    arr := []byte(dominoes)
    left, right := -1, -1
    for i := 0; i <= len(dominoes); i++ {
        if i == len(arr) || arr[i] == 'R' {
            if right > left {
                for right < i {
                    arr[right] = 'R'
                    right++
                }
            }
            right = i
        } else if arr[i] == 'L' {
            if left > right || right == -1 && left == -1 {
                left++
                for left < i {
                    arr[left] = 'L'
                    left++
                }
            } else {
                left = i
                low, high := right + 1, left - 1
                for low < high {
                    arr[low] = 'R'
                    arr[high] = 'L'
                    low++
                    high--
                }
            }
        }
    }
    return string(arr)
}

func pushDominoes1(deminoes string) string {
    s := []byte(deminoes)
    i, n, left := 0, len(s), byte('L')
    for i < n {
        j := i
        for j < n && s[j] == '.' { // 找到一段连续的没有被推动的骨牌
            j++
        }
        right := byte('R')
        if j < n {
            right = s[j]
        }
        if left == right { // 方向相同，那么这些竖立的骨牌也会倒向同一方向
            for i < j {
                s[i] = right
                i++
            }
        } else if left == 'R' && right == 'L' {
            k := j - 1
            for i < k {
                s[i] = 'R'
                s[k] = 'L'
                i++
                k--
            }
        }
        left = right
        i = j + 1
    }
    return string(s)
}

func main() {
    // Example 1:
    // Input: dominoes = "RR.L"
    // Output: "RR.L"
    // Explanation: The first domino expends no additional force on the second domino.
    fmt.Println(pushDominoes("RR.L")) // "RR.L"
    // Example 2:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/05/18/domino.png" />
    // Input: dominoes = ".L.R...LR..L.."
    // Output: "LL.RR.LLRRLL.."
    fmt.Println(pushDominoes(".L.R...LR..L..")) // "LL.RR.LLRRLL.."

    fmt.Println(pushDominoes("LLLLLLLLLLLLLLLLLL")) // LLLLLLLLLLLLLLLLLL
    fmt.Println(pushDominoes("RRRRRRRRRRRRRRRRRR")) // RRRRRRRRRRRRRRRRRR
    fmt.Println(pushDominoes("..................")) // ..................

    fmt.Println(pushDominoes1("RR.L")) // "RR.L"
    fmt.Println(pushDominoes1(".L.R...LR..L..")) // "LL.RR.LLRRLL.."
    fmt.Println(pushDominoes1("LLLLLLLLLLLLLLLLLL")) // LLLLLLLLLLLLLLLLLL
    fmt.Println(pushDominoes1("RRRRRRRRRRRRRRRRRR")) // RRRRRRRRRRRRRRRRRR
    fmt.Println(pushDominoes1("..................")) // ..................
}