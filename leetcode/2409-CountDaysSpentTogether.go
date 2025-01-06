package main

// 2409. Count Days Spent Together
// Alice and Bob are traveling to Rome for separate business meetings.

// You are given 4 strings arriveAlice, leaveAlice, arriveBob, and leaveBob. 
// Alice will be in the city from the dates arriveAlice to leaveAlice (inclusive), 
// while Bob will be in the city from the dates arriveBob to leaveBob (inclusive). 
// Each will be a 5-character string in the format "MM-DD", corresponding to the month and day of the date.

// Return the total number of days that Alice and Bob are in Rome together.

// You can assume that all dates occur in the same calendar year, which is not a leap year. 
// Note that the number of days per month can be represented as: [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31].

// Example 1:
// Input: arriveAlice = "08-15", leaveAlice = "08-18", arriveBob = "08-16", leaveBob = "08-19"
// Output: 3
// Explanation: Alice will be in Rome from August 15 to August 18. Bob will be in Rome from August 16 to August 19. They are both in Rome together on August 16th, 17th, and 18th, so the answer is 3.

// Example 2:
// Input: arriveAlice = "10-01", leaveAlice = "10-31", arriveBob = "11-01", leaveBob = "12-31"
// Output: 0
// Explanation: There is no day when Alice and Bob are in Rome together, so we return 0.

// Constraints:
//     All dates are provided in the format "MM-DD".
//     Alice and Bob's arrival dates are earlier than or equal to their leaving dates.
//     The given dates are valid dates of a non-leap year.

import "fmt"
import "strconv"

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
    month := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    min := func (x, y string) string { if x < y { return x; }; return y; }
    max := func (x, y string) string { if x > y { return x; }; return y; }
    parseInt := func(s string) int { n, _ := strconv.Atoi(s); return n; }
    if parseInt(leaveAlice[:2]) >= parseInt(arriveBob[:2]) && parseInt(arriveAlice[:2]) <= parseInt(leaveBob[:2]) {
        count := (parseInt(min(leaveAlice, leaveBob)[3:]) - parseInt( max(arriveAlice, arriveBob)[3:])) + 1
        for i := parseInt(max(arriveAlice, arriveBob)[:2]); i < parseInt(min(leaveAlice, leaveBob)[:2]); i++ {
            count += month[i-1]
        }
        if count > 0 {
            return count
        }
        return 0
    }
    return 0
}

func countDaysTogether1(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
    if leaveAlice < arriveBob || leaveBob < arriveAlice   { return 0 } // 没有交集
    month := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    sub := func(start string, end string) int {
        sm, _ := strconv.Atoi(string(start[0:2]))
        sd, _ := strconv.Atoi(string(start[3:5]))
        em, _ := strconv.Atoi(string(end[0:2]))
        ed, _ := strconv.Atoi(string(end[3:5]))
        total := 0
        for i := sm; i < em; i++ {
            total = total + month[i-1]
        }
        total = total - sd + ed +1
        return total
    }
    if arriveAlice <= arriveBob && leaveAlice >= leaveBob { return sub(arriveBob, leaveBob) } // 取bob
    if arriveBob <= arriveAlice && leaveBob >= leaveAlice { return sub(arriveAlice, leaveAlice) } // 取alice
    if arriveAlice <= arriveBob { return sub(arriveBob, leaveAlice) } // 取 leaveAlice - arriveBob
    return sub(arriveAlice, leaveBob)
}

func main() {
    // Example 1:
    // Input: arriveAlice = "08-15", leaveAlice = "08-18", arriveBob = "08-16", leaveBob = "08-19"
    // Output: 3
    // Explanation: Alice will be in Rome from August 15 to August 18. Bob will be in Rome from August 16 to August 19. They are both in Rome together on August 16th, 17th, and 18th, so the answer is 3.
    fmt.Println(countDaysTogether("08-15", "08-18", "08-16", "08-19")) // 3
    // Example 2:
    // Input: arriveAlice = "10-01", leaveAlice = "10-31", arriveBob = "11-01", leaveBob = "12-31"
    // Output: 0
    // Explanation: There is no day when Alice and Bob are in Rome together, so we return 0.
    fmt.Println(countDaysTogether("10-01", "10-31", "11-01", "12-31")) // 0

    fmt.Println(countDaysTogether1("08-15", "08-18", "08-16", "08-19")) // 3
    fmt.Println(countDaysTogether1("10-01", "10-31", "11-01", "12-31")) // 0
}