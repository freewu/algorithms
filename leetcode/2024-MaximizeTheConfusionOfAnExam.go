package main

// 2024. Maximize the Confusion of an Exam
// A teacher is writing a test with n true/false questions, with 'T' denoting true and 'F' denoting false. 
// He wants to confuse the students by maximizing the number of consecutive questions with the same answer (multiple trues or multiple falses in a row).

// You are given a string answerKey, where answerKey[i] is the original answer to the ith question. 
// In addition, you are given an integer k, the maximum number of times you may perform the following operation:

// Change the answer key for any question to 'T' or 'F' (i.e., set answerKey[i] to 'T' or 'F').
// Return the maximum number of consecutive 'T's or 'F's in the answer key after performing the operation at most k times.

// Example 1:
// Input: answerKey = "TTFF", k = 2
// Output: 4
// Explanation: We can replace both the 'F's with 'T's to make answerKey = "TTTT".
// There are four consecutive 'T's.

// Example 2:
// Input: answerKey = "TFFT", k = 1
// Output: 3
// Explanation: We can replace the first 'T' with an 'F' to make answerKey = "FFFT".
// Alternatively, we can replace the second 'T' with an 'F' to make answerKey = "TFFF".
// In both cases, there are three consecutive 'F's.

// Example 3:
// Input: answerKey = "TTFTTFTT", k = 1
// Output: 5
// Explanation: We can replace the first 'F' to make answerKey = "TTTTTFTT"
// Alternatively, we can replace the second 'F' to make answerKey = "TTFTTTTT". 
// In both cases, there are five consecutive 'T's.

// Constraints:
//     n == answerKey.length
//     1 <= n <= 5 * 10^4
//     answerKey[i] is either 'T' or 'F'
//     1 <= k <= n

import "fmt"

// Double Sliding Window
func maxConsecutiveAnswers(answerKey string, k int) int {
    if k >= len(answerKey) {
        return k
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    slidingWindow := func(answerKey string , k int , target rune) int {
        left, count, res := 0, 0, 0
        for right, item := range answerKey {
            if item != target {
                count++
            }
            for count > k {
                if answerKey[left] != byte(target) {
                    count--
                }
                left++
            }
            res = max(res, right - left + 1)
        }
        return res
    }
    return max(slidingWindow(answerKey, k ,'T'),slidingWindow(answerKey, k ,'F')) 
}

func maxConsecutiveAnswers1(answerKey string, k int) int {
    next, trues, falses, res := 0, 0, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := range answerKey {
        if answerKey[i] == 'T' {
            trues++
        } else {
            falses++
        }
        for min(trues, falses) > k {
            if answerKey[next] == 'T' {
                trues--
            } else {
                falses--
            }
            next++
        }
        res = max(res, trues + falses)
    }
    return res
}

func main() {
    // Example 1:
    // Input: answerKey = "TTFF", k = 2
    // Output: 4
    // Explanation: We can replace both the 'F's with 'T's to make answerKey = "TTTT".
    // There are four consecutive 'T's.
    fmt.Println(maxConsecutiveAnswers("TTFF", 2)) // 4
    // Example 2:
    // Input: answerKey = "TFFT", k = 1
    // Output: 3
    // Explanation: We can replace the first 'T' with an 'F' to make answerKey = "FFFT".
    // Alternatively, we can replace the second 'T' with an 'F' to make answerKey = "TFFF".
    // In both cases, there are three consecutive 'F's.
    fmt.Println(maxConsecutiveAnswers("TFFT", 1)) // 3
    // Example 3:
    // Input: answerKey = "TTFTTFTT", k = 1
    // Output: 5
    // Explanation: We can replace the first 'F' to make answerKey = "TTTTTFTT"
    // Alternatively, we can replace the second 'F' to make answerKey = "TTFTTTTT". 
    // In both cases, there are five consecutive 'T's.
    fmt.Println(maxConsecutiveAnswers("TTFTTFTT", 1)) // 5

    fmt.Println(maxConsecutiveAnswers1("TTFF", 2)) // 4
    fmt.Println(maxConsecutiveAnswers1("TFFT", 1)) // 3
    fmt.Println(maxConsecutiveAnswers1("TTFTTFTT", 1)) // 5
}