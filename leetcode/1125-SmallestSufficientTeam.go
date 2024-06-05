package main

// 1125. Smallest Sufficient Team
// In a project, you have a list of required skills req_skills, and a list of people. 
// The ith person people[i] contains a list of skills that the person has.

// Consider a sufficient team: a set of people such that for every required skill in req_skills, 
// there is at least one person in the team who has that skill. 
// We can represent these teams by the index of each person.
//     For example, team = [0, 1, 3] represents the people with skills people[0], people[1], and people[3].

// Return any sufficient team of the smallest possible size, represented by the index of each person. 
// You may return the answer in any order.

// It is guaranteed an answer exists.

// Example 1:
// Input: req_skills = ["java","nodejs","reactjs"], people = [["java"],["nodejs"],["nodejs","reactjs"]]
// Output: [0,2]

// Example 2:
// Input: req_skills = ["algorithms","math","java","reactjs","csharp","aws"], people = [["algorithms","math","java"],["algorithms","math","reactjs"],["java","csharp","aws"],["reactjs","csharp"],["csharp","math"],["aws","java"]]
// Output: [1,2]
 
// Constraints:
//     1 <= req_skills.length <= 16
//     1 <= req_skills[i].length <= 16
//     req_skills[i] consists of lowercase English letters.
//     All the strings of req_skills are unique.
//     1 <= people.length <= 60
//     0 <= people[i].length <= 16
//     1 <= people[i][j].length <= 16
//     people[i][j] consists of lowercase English letters.
//     All the strings of people[i] are unique.
//     Every skill in people[i] is a skill in req_skills.
//     It is guaranteed a sufficient team exists.

import "fmt"

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
    skillNameToId := make(map[string]uint8, len(req_skills))
    for i, skill := range req_skills {
        skillNameToId[skill] = uint8(i)
    }
    type Pair struct {
        prev         uint16
        last         uint8
        count        uint8
    }
    dp := make([]Pair, 1<<len(req_skills))
    for i, skills := range people {
        if len(skills) == 0 {
            continue
        }
        var skillsBits uint16
        for _, skill := range skills {
            skillsBits |= 1 << skillNameToId[skill]
        }
        if dp[skillsBits].count == 1 {
            continue
        }
        dp[skillsBits] = Pair{
            last: uint8(i),
            count:  1,
        }
        for j := range dp {
            if dp[j].count != 0 {
                or := uint16(j) | skillsBits
                if dp[or].count == 0 || dp[j].count+1 < dp[or].count {
                    dp[or] = Pair{
                        prev: uint16(j),
                        last: uint8(i),
                        count: dp[j].count + 1,
                    }
                }
            }
        }
    }
    res := []int{}
    for i := uint16(len(dp) - 1); i != 0; i = dp[i].prev {
        res = append(res, int(dp[i].last))
    }
    return res
}

func main() {
    // Example 1:
    // Input: req_skills = ["java","nodejs","reactjs"], people = [["java"],["nodejs"],["nodejs","reactjs"]]
    // Output: [0,2]
    fmt.Println(smallestSufficientTeam([]string{"java","nodejs","reactjs"},[][]string{{"java"},{"nodejs"},{"nodejs","reactjs"}})) // [0,2]
    // Example 2:
    // Input: req_skills = ["algorithms","math","java","reactjs","csharp","aws"], people = [["algorithms","math","java"],["algorithms","math","reactjs"],["java","csharp","aws"],["reactjs","csharp"],["csharp","math"],["aws","java"]]
    // Output: [1,2]
    fmt.Println(smallestSufficientTeam([]string{"algorithms","math","java","reactjs","csharp","aws"},[][]string{{"algorithms","math","java"},{"algorithms","math","reactjs"},{"java","csharp","aws"},{"reactjs","csharp"},{"csharp","math"},{"aws","java"}})) // [1,2]
}