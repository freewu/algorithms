package main

// 599. Minimum Index Sum of Two Lists
// Given two arrays of strings list1 and list2, find the common strings with the least index sum.
// A common string is a string that appeared in both list1 and list2.

// A common string with the least index sum is a common string such that if it appeared at list1[i] and list2[j] then i + j should be the minimum value among all the other common strings.
// Return all the common strings with the least index sum. Return the answer in any order.

// Example 1:
// Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]
// Output: ["Shogun"]
// Explanation: The only common string is "Shogun".

// Example 2:
// Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["KFC","Shogun","Burger King"]
// Output: ["Shogun"]
// Explanation: The common string with the least index sum is "Shogun" with index sum = (0 + 1) = 1.

// Example 3:
// Input: list1 = ["happy","sad","good"], list2 = ["sad","happy","good"]
// Output: ["sad","happy"]
// Explanation: There are three common strings:
// "happy" with index sum = (0 + 1) = 1.
// "sad" with index sum = (1 + 0) = 1.
// "good" with index sum = (2 + 2) = 4.
// The strings with the least index sum are "sad" and "happy".

// Constraints:
//     1 <= list1.length, list2.length <= 1000
//     1 <= list1[i].length, list2[i].length <= 30
//     list1[i] and list2[i] consist of spaces ' ' and English letters.
//     All the strings of list1 are unique.
//     All the strings of list2 are unique.
//     There is at least a common string between list1 and list2.

import "fmt"

// O(n^2)
func findRestaurant(list1 []string, list2 []string) []string {
    m, n := make(map[int][]string), 1 << 32 -1
    for i1, v1 := range list1 {
        for i2, v2 := range list2 {
            if v1 == v2 { // 相同的
                index := i1+ i2 // 两个索引相加
                m[index] = append(m[index], v1)
                if index < n {
                    n = index
                }
            }
        }
    }
    return m[n]
}

// O(n)
func findRestaurant1(list1 []string, list2 []string) []string {
    m := make(map[string]int)
    for i, v := range list1 { // list 转成 map
        m[v] = i
    }
    minSum, res := 1 << 32 - 1, make([]string, 0)
    for i, v := range list2 {
        if index, ok := m[v]; ok {
            if index + i < minSum {
                res = []string{ v }
                minSum = index + i
            } else if index + i == minSum {
                res = append(res, v)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]
    // Output: ["Shogun"]
    // Explanation: The only common string is "Shogun".
    fmt.Println(findRestaurant([]string{"Shogun","Tapioca Express","Burger King","KFC"},[]string{"Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"})) // ["Shogun"]
    // Example 2:
    // Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["KFC","Shogun","Burger King"]
    // Output: ["Shogun"]
    // Explanation: The common string with the least index sum is "Shogun" with index sum = (0 + 1) = 1.
    fmt.Println(findRestaurant([]string{"Shogun","Tapioca Express","Burger King","KFC"},[]string{"KFC","Shogun","Burger King"})) // ["Shogun"]
    // Example 3:
    // Input: list1 = ["happy","sad","good"], list2 = ["sad","happy","good"]
    // Output: ["sad","happy"]
    // Explanation: There are three common strings:
    // "happy" with index sum = (0 + 1) = 1.
    // "sad" with index sum = (1 + 0) = 1.
    // "good" with index sum = (2 + 2) = 4.
    // The strings with the least index sum are "sad" and "happy".
    fmt.Println(findRestaurant([]string{"happy","sad","good"},[]string{"sad","happy","good"})) //  ["sad","happy"]

    fmt.Println(findRestaurant1([]string{"Shogun","Tapioca Express","Burger King","KFC"},[]string{"Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"})) // ["Shogun"]
    fmt.Println(findRestaurant1([]string{"Shogun","Tapioca Express","Burger King","KFC"},[]string{"KFC","Shogun","Burger King"})) // ["Shogun"]
    fmt.Println(findRestaurant1([]string{"happy","sad","good"},[]string{"sad","happy","good"})) //  ["sad","happy"]
}