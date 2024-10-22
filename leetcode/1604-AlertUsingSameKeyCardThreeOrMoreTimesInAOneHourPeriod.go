package main

// 1604. Alert Using Same Key-Card Three or More Times in a One Hour Period
// LeetCode company workers use key-cards to unlock office doors. 
// Each time a worker uses their key-card, the security system saves the worker's name and the time when it was used. 
// The system emits an alert if any worker uses the key-card three or more times in a one-hour period.

// You are given a list of strings keyName and keyTime where [keyName[i], 
// keyTime[i]] corresponds to a person's name and the time when their key-card was used in a single day.

// Access times are given in the 24-hour time format "HH:MM", such as "23:51" and "09:49".

// Return a list of unique worker names who received an alert for frequent keycard use. 
// Sort the names in ascending order alphabetically.

// Notice that "10:00" - "11:00" is considered to be within a one-hour period, 
// while "22:51" - "23:52" is not considered to be within a one-hour period.

// Example 1:
// Input: keyName = ["daniel","daniel","daniel","luis","luis","luis","luis"], keyTime = ["10:00","10:40","11:00","09:00","11:00","13:00","15:00"]
// Output: ["daniel"]
// Explanation: "daniel" used the keycard 3 times in a one-hour period ("10:00","10:40", "11:00").

// Example 2:
// Input: keyName = ["alice","alice","alice","bob","bob","bob","bob"], keyTime = ["12:01","12:00","18:00","21:00","21:20","21:30","23:00"]
// Output: ["bob"]
// Explanation: "bob" used the keycard 3 times in a one-hour period ("21:00","21:20", "21:30").

// Constraints:
//     1 <= keyName.length, keyTime.length <= 10^5
//     keyName.length == keyTime.length
//     keyTime[i] is in the format "HH:MM".
//     [keyName[i], keyTime[i]] is unique.
//     1 <= keyName[i].length <= 10
//     keyName[i] contains only lowercase English letters.

import "fmt"
import "sort"
import "strconv"
import "strings"

func alertNames(keyName []string, keyTime []string) []string {
    tracker, nameSet := map[string][]int{}, map[string]struct{}{}
    for i, name := range keyName {
        timeParts := strings.Split(keyTime[i], ":")
        hour, _ := strconv.Atoi(timeParts[0])
        minute, _ := strconv.Atoi(timeParts[1])
        tracker[name] = append(tracker[name], 60 * hour + minute)
    }
    for name, times := range tracker {
        sort.Ints(times)
        for i := 2; i < len(times); i++ {
            diff := times[i] - times[i-2]
            if diff <= 60 {
                nameSet[name] = struct{}{}
            }
        }
    }
    names := []string{}
    for name := range nameSet {
        names = append(names, name)
    }
    sort.Strings(names) // Sort the names in ascending order alphabetically.
    return names
}

func alertNames1(keyName []string, keyTime []string) []string {
    res, mp := []string{}, map[string][]int{}
    for i, name := range keyName {
        t := keyTime[i]
        hour, minute := (int(t[0] - '0') * 10 + int(t[1] - '0')), (int(t[3] - '0') * 10 + int(t[4] - '0'))
        mp[name] = append(mp[name], hour * 60 + minute)
    }
    for name, ts := range mp {
        n := len(ts)
        if n <= 2 { continue }
        sort.Ints(ts)
        for i := 0; i < n - 2; i ++ {
            if ts[i + 2] - ts[i] <= 60 {
                res = append(res, name)
                break
            }
        }
    }
    sort.Strings(res) // Sort the names in ascending order alphabetically.
    return res
}

func main() {
    // Example 1:
    // Input: keyName = ["daniel","daniel","daniel","luis","luis","luis","luis"], keyTime = ["10:00","10:40","11:00","09:00","11:00","13:00","15:00"]
    // Output: ["daniel"]
    // Explanation: "daniel" used the keycard 3 times in a one-hour period ("10:00","10:40", "11:00").
    fmt.Println(alertNames([]string{"daniel","daniel","daniel","luis","luis","luis","luis"}, []string{"10:00","10:40","11:00","09:00","11:00","13:00","15:00"}))  // ["daniel"]
    // Example 2:
    // Input: keyName = ["alice","alice","alice","bob","bob","bob","bob"], keyTime = ["12:01","12:00","18:00","21:00","21:20","21:30","23:00"]
    // Output: ["bob"]
    // Explanation: "bob" used the keycard 3 times in a one-hour period ("21:00","21:20", "21:30").
    fmt.Println(alertNames([]string{"alice","alice","alice","bob","bob","bob","bob"}, []string{"12:01","12:00","18:00","21:00","21:20","21:30","23:00"}))  // ["bob"]

    fmt.Println(alertNames1([]string{"daniel","daniel","daniel","luis","luis","luis","luis"}, []string{"10:00","10:40","11:00","09:00","11:00","13:00","15:00"}))  // ["daniel"]
    fmt.Println(alertNames1([]string{"alice","alice","alice","bob","bob","bob","bob"}, []string{"12:01","12:00","18:00","21:00","21:20","21:30","23:00"}))  // ["bob"]
}