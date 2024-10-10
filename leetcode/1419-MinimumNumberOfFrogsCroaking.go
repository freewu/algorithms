package main

// 1419. Minimum Number of Frogs Croaking
// You are given the string croakOfFrogs, 
// which represents a combination of the string "croak" from different frogs, 
// that is, multiple frogs can croak at the same time, so multiple "croak" are mixed.

// Return the minimum number of different frogs to finish all the croaks in the given string.

// A valid "croak" means a frog is printing five letters 'c', 'r', 'o', 'a', and 'k' sequentially. 
// The frogs have to print all five letters to finish a croak. 
// If the given string is not a combination of a valid "croak" return -1.

// Example 1:
// Input: croakOfFrogs = "croakcroak"
// Output: 1 
// Explanation: One frog yelling "croak" twice.

// Example 2:
// Input: croakOfFrogs = "crcoakroak"
// Output: 2 
// Explanation: The minimum number of frogs is two. 
// The first frog could yell "crcoakroak".
// The second frog could yell later "crcoakroak".

// Example 3:
// Input: croakOfFrogs = "croakcrook"
// Output: -1
// Explanation: The given string is an invalid combination of "croak" from different frogs.

// Constraints:
//     1 <= croakOfFrogs.length <= 10^5
//     croakOfFrogs is either 'c', 'r', 'o', 'a', or 'k'.

import "fmt"

func minNumberOfFrogs(croakOfFrogs string) int {
    res, count := 0, make([]int, 5) // croak
    for _, v := range croakOfFrogs {
        if v == 'c' {
            count[0]++
            if count[0] > res { res = count[0] }
        } else if v == 'r' {
            count[1]++
            if count[1] > count[0] { return -1  }
        } else if v == 'o' {
            count[2]++
            if count[2] > count[0] || count[2] > count[1] { return -1 }
        } else if v == 'a' {
            count[3]++
            if count[3] > count[0] || count[3] > count[1] || count[3] > count[2] { return -1 }
        } else if v == 'k' {
            count[4]++
            if count[4] > count[3] || count[4] > count[0] || count[4] > count[1] || count[4] > count[2] { return -1 }
            for i := 0; i < 5; i++ { count[i]-- }
        } else {
            return -1
        }
    }
    for _, v := range count {
        if v > 0 { // 没有清零的
            return -1
        }
    }
    return res
}

func minNumberOfFrogs1(croakOfFrogs string) int {
    maxFrogs := 0
    c, r, o, a, k := 0, 0, 0, 0, 0
    for _, v := range croakOfFrogs {
        if v == 'c' {
            c++
            k++
            if k > maxFrogs {
                maxFrogs = k
            }
        } else if v == 'r' {
            if c == 0 {
                return -1
            }
            c--
            r++
        } else if v == 'o' {
            if r == 0 {
                return -1
            } else {
                r--
                o++
            }
        } else if v == 'a' {
            if o == 0 {
                return -1
            } else {
                o--
                a++
            }
        } else {
            if a == 0 {
                return -1
            } else {
                a--
                k--
            }
        }
    }
    if c != 0 || r != 0 || o != 0 || a != 0 || k != 0{
        return -1
    }
    return maxFrogs
}

func main() {
    // Example 1:
    // Input: croakOfFrogs = "croakcroak"
    // Output: 1 
    // Explanation: One frog yelling "croak" twice.
    fmt.Println(minNumberOfFrogs("croakcroak")) // 1
    // Example 2:
    // Input: croakOfFrogs = "crcoakroak"
    // Output: 2 
    // Explanation: The minimum number of frogs is two. 
    // The first frog could yell "crcoakroak".
    // The second frog could yell later "crcoakroak".
    fmt.Println(minNumberOfFrogs("crcoakroak")) // 2
    // Example 3:
    // Input: croakOfFrogs = "croakcrook"
    // Output: 1
    // Explanation: The given string is an invalid combination of "croak" from different frogs.
    fmt.Println(minNumberOfFrogs("croakcrook")) // 1

    fmt.Println(minNumberOfFrogs1("croakcroak")) // 1
    fmt.Println(minNumberOfFrogs1("crcoakroak")) // 2
    fmt.Println(minNumberOfFrogs1("croakcrook")) // 1
}