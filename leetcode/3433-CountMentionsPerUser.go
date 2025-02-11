package main

// 3433. Count Mentions Per User
// You are given an integer numberOfUsers representing the total number of users and an array events of size n x 3.

// Each events[i] can be either of the following two types:

// 1. Message Event: ["MESSAGE", "timestampi", "mentions_stringi"]
//     1.1 This event indicates that a set of users was mentioned in a message at timestampi.
//     1.2 The mentions_stringi string can contain one of the following tokens:
//         1.2.1 id<number>: where <number> is an integer in range [0,numberOfUsers - 1]. 
//               There can be multiple ids separated by a single whitespace and may contain duplicates. 
//               This can mention even the offline users.
//         1.2.2 ALL: mentions all users.
//         1.2.3 HERE: mentions all online users.
// 2. Offline Event: ["OFFLINE", "timestampi", "idi"]
//     2.1 This event indicates that the user idi had become offline at timestampi for 60 time units. 
//         The user will automatically be online again at time timestampi + 60.

// Return an array mentions where mentions[i] represents the number of mentions the user with id i has across all MESSAGE events.

// All users are initially online, and if a user goes offline or comes back online, 
// their status change is processed before handling any message event that occurs at the same timestamp.

// Note that a user can be mentioned multiple times in a single message event, 
// and each mention should be counted separately.

// Example 1:
// Input: numberOfUsers = 2, events = [["MESSAGE","10","id1 id0"],["OFFLINE","11","0"],["MESSAGE","71","HERE"]]
// Output: [2,2]
// Explanation:
// Initially, all users are online.
// At timestamp 10, id1 and id0 are mentioned. mentions = [1,1]
// At timestamp 11, id0 goes offline.
// At timestamp 71, id0 comes back online and "HERE" is mentioned. mentions = [2,2]

// Example 2:
// Input: numberOfUsers = 2, events = [["MESSAGE","10","id1 id0"],["OFFLINE","11","0"],["MESSAGE","12","ALL"]]
// Output: [2,2]
// Explanation:
// Initially, all users are online.
// At timestamp 10, id1 and id0 are mentioned. mentions = [1,1]
// At timestamp 11, id0 goes offline.
// At timestamp 12, "ALL" is mentioned. This includes offline users, so both id0 and id1 are mentioned. mentions = [2,2]

// Example 3:
// Input: numberOfUsers = 2, events = [["OFFLINE","10","0"],["MESSAGE","12","HERE"]]
// Output: [0,1]
// Explanation:
// Initially, all users are online.
// At timestamp 10, id0 goes offline.
// At timestamp 12, "HERE" is mentioned. Because id0 is still offline, they will not be mentioned. mentions = [0,1]

// Constraints:
//     1 <= numberOfUsers <= 100
//     1 <= events.length <= 100
//     events[i].length == 3
//     events[i][0] will be one of MESSAGE or OFFLINE.
//     1 <= int(events[i][1]) <= 10^5
//     The number of id<number> mentions in any "MESSAGE" event is between 1 and 100.
//     0 <= <number> <= numberOfUsers - 1
//     It is guaranteed that the user id referenced in the OFFLINE event is online at the time the event occurs.

import "fmt"
import "sort"
import "strconv"
import "strings"

func countMentions(numberOfUsers int, events [][]string) []int {
    mentions, online, offline := make([]int, numberOfUsers), make([]bool, numberOfUsers), make([]int, numberOfUsers)
    for i := range online {
        online[i] = true
    }
    // 排序事件，确保同一时间戳下OFFLINE事件先处理
    sort.Slice(events, func(i, j int) bool {
        time1, _ := strconv.Atoi(events[i][1])
        time2, _ := strconv.Atoi(events[j][1])
        if time1 != time2 { return time1 < time2 }
        // 时间相同，OFFLINE事件排在前面
        if events[i][0] == "OFFLINE" && events[j][0] != "OFFLINE" { return true  }
        if events[i][0] != "OFFLINE" && events[j][0] == "OFFLINE" { return false }
        // 若两个事件类型相同，保持原顺序（不影响结果）
        return i < j
    })
    for _, event := range events {
        timestamp, _ := strconv.Atoi(event[1])
        // 检查并更新在线状态
        for i := range online {
            if !online[i] && timestamp >= offline[i] {
                online[i] = true
            }
        }
        if event[0] == "OFFLINE" {
            userID, _ := strconv.Atoi(event[2])
            online[userID] = false
            offline[userID] = timestamp + 60
        } else {
            mentionsStr := event[2]
            switch mentionsStr {
            case "ALL":
                for i := 0; i < numberOfUsers; i++ {
                    mentions[i]++
                }
            case "HERE":
                for i := 0; i < numberOfUsers; i++ {
                    if online[i] {
                        mentions[i]++
                    }
                }
            default:
                ids := strings.Fields(mentionsStr)
                for _, idStr := range ids {
                    if strings.HasPrefix(idStr, "id") {
                        userID, _ := strconv.Atoi(idStr[2:])
                        mentions[userID]++
                    }
                }
            }
        }
    }
    return mentions
}

func main() {
    // Example 1:
    // Input: numberOfUsers = 2, events = [["MESSAGE","10","id1 id0"],["OFFLINE","11","0"],["MESSAGE","71","HERE"]]
    // Output: [2,2]
    // Explanation:
    // Initially, all users are online.
    // At timestamp 10, id1 and id0 are mentioned. mentions = [1,1]
    // At timestamp 11, id0 goes offline.
    // At timestamp 71, id0 comes back online and "HERE" is mentioned. mentions = [2,2]
    fmt.Println(countMentions(2, [][]string{{"MESSAGE","10","id1 id0"},{"OFFLINE","11","0"},{"MESSAGE","71","HERE"}})) // [2,2]
    // Example 2:
    // Input: numberOfUsers = 2, events = [["MESSAGE","10","id1 id0"],["OFFLINE","11","0"],["MESSAGE","12","ALL"]]
    // Output: [2,2]
    // Explanation:
    // Initially, all users are online.
    // At timestamp 10, id1 and id0 are mentioned. mentions = [1,1]
    // At timestamp 11, id0 goes offline.
    // At timestamp 12, "ALL" is mentioned. This includes offline users, so both id0 and id1 are mentioned. mentions = [2,2]
    fmt.Println(countMentions(2, [][]string{{"MESSAGE","10","id1 id0"},{"OFFLINE","11","0"},{"MESSAGE","12","ALL"}})) // [2,2]
    // Example 3:
    // Input: numberOfUsers = 2, events = [["OFFLINE","10","0"],["MESSAGE","12","HERE"]]
    // Output: [0,1]
    // Explanation:
    // Initially, all users are online.
    // At timestamp 10, id0 goes offline.
    // At timestamp 12, "HERE" is mentioned. Because id0 is still offline, they will not be mentioned. mentions = [0,1]
    fmt.Println(countMentions(2, [][]string{{"OFFLINE","10","0"},{"MESSAGE","12","HERE"}})) // [0,1]
}