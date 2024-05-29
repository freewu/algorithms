package main

// 401. Binary Watch
// A binary watch has 4 LEDs on the top to represent the hours (0-11), and 6 LEDs on the bottom to represent the minutes (0-59). 
// Each LED represents a zero or one, with the least significant bit on the right.
//     For example, the below binary watch reads "4:51".
//     <img src="https://assets.leetcode.com/uploads/2021/04/08/binarywatch.jpg" />

// Given an integer turnedOn which represents the number of LEDs that are currently on (ignoring the PM), 
// return all possible times the watch could represent. You may return the answer in any order.

// The hour must not contain a leading zero.
//     For example, "01:00" is not valid. It should be "1:00".

// The minute must consist of two digits and may contain a leading zero.
//     For example, "10:2" is not valid. It should be "10:02".
 
// Example 1:
// Input: turnedOn = 1
// Output: ["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]

// Example 2:
// Input: turnedOn = 9
// Output: []

// Constraints:
//     0 <= turnedOn <= 10

import "fmt"
import "strconv"
import "math/bits"

func readBinaryWatch(num int) []string {
    res, memo := []string{}, make([]int, 60)
    count := func(n int) int { // count the number of 1 in a binary number
        if memo[n] != 0 {
            return memo[n]
        }
        originN, res := n, 0
        for n != 0 {
            n = n & (n - 1)
            res++
        }
        memo[originN] = res
        return res
    }
    fmtMinute := func(m int) string { // fmtMinute format minute 0:1 -> 0:01
        if m < 10 {
            return "0" + strconv.Itoa(m)
        }
        return strconv.Itoa(m)
    }
    for i := 0; i < 12; i++ { // traverse 0:00 -> 12:00
        for j := 0; j < 60; j++ {
            if count(i) + count(j) == num {
                res = append(res, strconv.Itoa(i) + ":" + fmtMinute(j))
            }
        }
    }
    return res
}

func readBinaryWatch1(turnedOn int) []string {
    res := []string{}
    for h := uint8(0); h < 12; h++ {
        for m := uint8(0); m < 60; m++ {
            if bits.OnesCount8(h) + bits.OnesCount8(m) == turnedOn {
                res = append(res, fmt.Sprintf("%d:%02d", h, m))
            }
        }
    }
    return res
}


var (
    hour    = []string{"1", "2", "4", "8"}
    minute  = []string{"01", "02", "04", "08", "16", "32"}
    hourMap = map[int][]string{
        0: {"0"},
        1: {"1", "2", "4", "8"},
        2: {"3", "5", "9", "6", "10"},
        3: {"7", "11"},
    }
    minuteMap = map[int][]string{
        0: {"00"},
        1: {"01", "02", "04", "08", "16", "32"},
        2: {"03", "05", "09", "17", "33", "06", "10", "18", "34", "12", "20", "36", "24", "40", "48"},
        3: {"07", "11", "19", "35", "13", "21", "37", "25", "41", "49", "14", "22", "38", "26", "42", "50", "28", "44", "52", "56"},
        4: {"15", "23", "39", "27", "43", "51", "29", "45", "53", "57", "30", "46", "54", "58"},
        5: {"31", "47", "55", "59"},
    }
)

// 打表大法
func readBinaryWatch2(num int) []string {
    res := []string{}
    if num > 8 {
        return res
    }
    for i := 0; i <= num; i++ {
        for j := 0; j < len(hourMap[i]); j++ {
            for k := 0; k < len(minuteMap[num-i]); k++ {
                res = append(res, hourMap[i][j]+":"+minuteMap[num-i][k])
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: turnedOn = 1
    // Output: ["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
    fmt.Println(readBinaryWatch(1)) // ["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
    // Example 2:
    // Input: turnedOn = 9
    // Output: []
    fmt.Println(readBinaryWatch(9)) // []

    fmt.Println(readBinaryWatch1(1)) // ["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
    fmt.Println(readBinaryWatch1(9)) // []

    fmt.Println(readBinaryWatch2(1)) // ["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
    fmt.Println(readBinaryWatch2(9)) // []
}