package main

// 937. Reorder Data in Log Files
// You are given an array of logs. Each log is a space-delimited string of words, where the first word is the identifier.

// There are two types of logs:
//     Letter-logs: All words (except the identifier) consist of lowercase English letters.
//     Digit-logs: All words (except the identifier) consist of digits.

// Reorder these logs so that:
//     The letter-logs come before all digit-logs.
//     The letter-logs are sorted lexicographically by their contents. If their contents are the same, then sort them lexicographically by their identifiers.
//     The digit-logs maintain their relative ordering.

// Return the final order of the logs.

// Example 1:
// Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
// Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
// Explanation:
// The letter-log contents are all different, so their ordering is "art can", "art zero", "own kit dig".
// The digit-logs have a relative order of "dig1 8 1 5 1", "dig2 3 6".

// Example 2:
// Input: logs = ["a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"]
// Output: ["g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"]

// Constraints:
//     1 <= logs.length <= 100
//     3 <= logs[i].length <= 100
//     All the tokens of logs[i] are separated by a single space.
//     logs[i] is guaranteed to have an identifier and at least one word after the identifier.

import "fmt"
import "strings"
import "sort"

func reorderLogFiles(logs []string) []string {
    letterLogs, digitLogs := []string{}, []string{}
    for _, v := range logs {
        flag := strings.Fields(v)[1][0]
        if flag >= '0' && flag <= '9' {
            digitLogs = append(digitLogs, v)
        } else {
            letterLogs = append(letterLogs, v)
        }
    }
    sort.Slice(letterLogs, func(i, j int) bool {
        fieldsI, fieldsJ := strings.Fields(letterLogs[i])[1:], strings.Fields(letterLogs[j])[1:]
        for i := 0; i < len(fieldsI) && i < len(fieldsJ); i++ {
            if fieldsI[i] != fieldsJ[i] {
                return fieldsI[i] < fieldsJ[i]
            }
        }
        if len(fieldsI) == len(fieldsJ) {
            return strings.Fields(letterLogs[i])[0] < strings.Fields(letterLogs[j])[0]
        }
        return len(fieldsI) < len(fieldsJ)
    })
    return append(letterLogs, digitLogs...)
}

func reorderLogFiles1(logs []string) []string {
    isDigitLog := func (log string) bool {
        index := strings.Index(log, " ")
        firstChar := log[index + 1]
        return firstChar >= '0' && firstChar <= '9'
    }
    getCharLogDetail := func (log string) (string, string) {
        index := strings.Index(log, " ")
        return log[:index], log[index+1:]
    }
    sort.SliceStable(logs, func(i, j int) bool {
        logI, logJ := logs[i], logs[j]
        if isDigitLog(logI) && isDigitLog(logJ) { // 都是数字日志，应该保留相对顺序
            return false
        } 
        if !isDigitLog(logI) && !isDigitLog(logJ) {
            logoI, contentI := getCharLogDetail(logI)
            logoJ, contentJ := getCharLogDetail(logJ)
            return contentI < contentJ || contentI == contentJ && logoI < logoJ
        }
        return !isDigitLog(logI)
    })
    return logs
}

func main() {
    // Example 1:
    // Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
    // Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
    // Explanation:
    // The letter-log contents are all different, so their ordering is "art can", "art zero", "own kit dig".
    // The digit-logs have a relative order of "dig1 8 1 5 1", "dig2 3 6".
    logs1 := []string{
        "dig1 8 1 5 1",
        "let1 art can",
        "dig2 3 6",
        "let2 own kit dig",
        "let3 art zero",
    }
    fmt.Println(reorderLogFiles(logs1)) // ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
    // Example 2:
    // Input: logs = ["a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"]
    // Output: ["g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"]
    logs2 := []string{
        "a1 9 2 3 1",
        "g1 act car",
        "zo4 4 7",
        "ab1 off key dog",
        "a8 act zoo",
    }
    fmt.Println(reorderLogFiles(logs2)) // ["g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"]

    fmt.Println(reorderLogFiles1(logs1)) // ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
    fmt.Println(reorderLogFiles1(logs2)) // ["g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"]
}