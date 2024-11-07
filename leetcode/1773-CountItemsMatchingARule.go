package main

// 1773. Count Items Matching a Rule
// You are given an array items, 
// where each items[i] = [typei, colori, namei] describes the type, color, and name of the ith item. 
// You are also given a rule represented by two strings, ruleKey and ruleValue.

// The ith item is said to match the rule if one of the following is true:
//     ruleKey == "type" and ruleValue == typei.
//     ruleKey == "color" and ruleValue == colori.
//     ruleKey == "name" and ruleValue == namei.

// Return the number of items that match the given rule.

// Example 1:
// Input: items = [["phone","blue","pixel"],["computer","silver","lenovo"],["phone","gold","iphone"]], ruleKey = "color", ruleValue = "silver"
// Output: 1
// Explanation: There is only one item matching the given rule, which is ["computer","silver","lenovo"].

// Example 2:
// Input: items = [["phone","blue","pixel"],["computer","silver","phone"],["phone","gold","iphone"]], ruleKey = "type", ruleValue = "phone"
// Output: 2
// Explanation: There are only two items matching the given rule, which are ["phone","blue","pixel"] and ["phone","gold","iphone"]. Note that the item ["computer","silver","phone"] does not match.
 
// Constraints:
//     1 <= items.length <= 10^4
//     1 <= typei.length, colori.length, namei.length, ruleValue.length <= 10
//     ruleKey is equal to either "type", "color", or "name".
//     All strings consist only of lowercase letters.

import "fmt"

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    type Item struct {
        Type, Color, Name string
    }
    data := []Item{}
    for _, v := range items {
        data = append(data, Item{ Type: v[0], Color: v[1], Name: v[2] })
    }
    res := 0
    for _, v := range data {
        if ruleKey == "type" && ruleValue == v.Type { res++ } // ruleKey == "type" and ruleValue == typei.
        if ruleKey == "color" && ruleValue == v.Color { res++ } // ruleKey == "color" and ruleValue == colori.
        if ruleKey == "name" && ruleValue == v.Name { res++ } // ruleKey == "name" and ruleValue == namei.
    }
    return res
}

func countMatches1(items [][]string, ruleKey string, ruleValue string) int {
    mp := map[string]int{ "type": 0, "color": 1, "name": 2 }
    res := 0
    for _, v := range items {
        if v[mp[ruleKey]] == ruleValue {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: items = [["phone","blue","pixel"],["computer","silver","lenovo"],["phone","gold","iphone"]], ruleKey = "color", ruleValue = "silver"
    // Output: 1
    // Explanation: There is only one item matching the given rule, which is ["computer","silver","lenovo"].
    fmt.Println(countMatches([][]string{{"phone","blue","pixel"},{"computer","silver","lenovo"},{"phone","gold","iphone"}}, "color", "silver")) // 1
    // Example 2:
    // Input: items = [["phone","blue","pixel"],["computer","silver","phone"],["phone","gold","iphone"]], ruleKey = "type", ruleValue = "phone"
    // Output: 2
    // Explanation: There are only two items matching the given rule, which are ["phone","blue","pixel"] and ["phone","gold","iphone"]. Note that the item ["computer","silver","phone"] does not match.
    fmt.Println(countMatches([][]string{{"phone","blue","pixel"},{"computer","silver","phone"},{"phone","gold","iphone"}}, "type", "phone")) // 2

    fmt.Println(countMatches1([][]string{{"phone","blue","pixel"},{"computer","silver","lenovo"},{"phone","gold","iphone"}}, "color", "silver")) // 1
    fmt.Println(countMatches1([][]string{{"phone","blue","pixel"},{"computer","silver","phone"},{"phone","gold","iphone"}}, "type", "phone")) // 2
}