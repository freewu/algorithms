package main

// 3838. Weighted Word Mapping
// You are given an array of strings words, where each string represents a word containing lowercase English letters.

// You are also given an integer array weights of length 26, where weights[i] represents the weight of the ith lowercase English letter.

// The weight of a word is defined as the sum of the weights of its characters.

// For each word, take its weight modulo 26 and map the result to a lowercase English letter using reverse alphabetical order (0 -> 'z', 1 -> 'y', ..., 25 -> 'a').

// Return a string formed by concatenating the mapped characters for all words in order.

// Example 1:
// Input: words = ["abcd","def","xyz"], weights = [5,3,12,14,1,2,3,2,10,6,6,9,7,8,7,10,8,9,6,9,9,8,3,7,7,2]
// Output: "rij"
// Explanation:
// The weight of "abcd" is 5 + 3 + 12 + 14 = 34. The result modulo 26 is 34 % 26 = 8, which maps to 'r'.
// The weight of "def" is 14 + 1 + 2 = 17. The result modulo 26 is 17 % 26 = 17, which maps to 'i'.
// The weight of "xyz" is 7 + 7 + 2 = 16. The result modulo 26 is 16 % 26 = 16, which maps to 'j'.
// Thus, the string formed by concatenating the mapped characters is "rij".

// Example 2:
// Input: words = ["a","b","c"], weights = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
// Output: "yyy"
// Explanation:
// Each word has weight 1. The result modulo 26 is 1 % 26 = 1, which maps to 'y'.
// Thus, the string formed by concatenating the mapped characters is "yyy".

// Example 3:
// Input: words = ["abcd"], weights = [7,5,3,4,3,5,4,9,4,2,2,7,10,2,5,10,6,1,2,2,4,1,3,4,4,5]
// Output: "g"
// Explanation:​​​​​​​
// The weight of "abcd" is 7 + 5 + 3 + 4 = 19. The result modulo 26 is 19 % 26 = 19, which maps to 'g'.
// Thus, the string formed by concatenating the mapped characters is "g".

// Constraints:
//     1 <= words.length <= 100
//     1 <= words[i].length <= 10
//     weights.length == 26
//     1 <= weights[i] <= 100
//     words[i] consists of lowercase English letters.

import "fmt"

func mapWordWeights(words []string, weights []int) string {
    res := ""
    for _, word := range words {
        w := 0
        for _, c := range word {
            w += weights[c - 'a']
        }
        res += string('z' - w % 26)
    }
    return res
}

func mapWordWeights1(words []string, weights []int) string {
    res := make([]byte, len(words))
    for i, w := range words {
        sum := 0
        for _, ch := range w {
            sum += weights[ch-'a']
        }
        res[i] = 'z' - byte(sum % 26)
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: words = ["abcd","def","xyz"], weights = [5,3,12,14,1,2,3,2,10,6,6,9,7,8,7,10,8,9,6,9,9,8,3,7,7,2]
    // Output: "rij"
    // Explanation:
    // The weight of "abcd" is 5 + 3 + 12 + 14 = 34. The result modulo 26 is 34 % 26 = 8, which maps to 'r'.
    // The weight of "def" is 14 + 1 + 2 = 17. The result modulo 26 is 17 % 26 = 17, which maps to 'i'.
    // The weight of "xyz" is 7 + 7 + 2 = 16. The result modulo 26 is 16 % 26 = 16, which maps to 'j'.
    // Thus, the string formed by concatenating the mapped characters is "rij".
    fmt.Println(mapWordWeights([]string{"abcd","def","xyz"}, []int{5,3,12,14,1,2,3,2,10,6,6,9,7,8,7,10,8,9,6,9,9,8,3,7,7,2})) // "rij"
    // Example 2:
    // Input: words = ["a","b","c"], weights = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
    // Output: "yyy"
    // Explanation:
    // Each word has weight 1. The result modulo 26 is 1 % 26 = 1, which maps to 'y'.
    // Thus, the string formed by concatenating the mapped characters is "yyy".
    fmt.Println(mapWordWeights([]string{"a","b","c"}, []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})) // "yyy"
    // Example 3:
    // Input: words = ["abcd"], weights = [7,5,3,4,3,5,4,9,4,2,2,7,10,2,5,10,6,1,2,2,4,1,3,4,4,5]
    // Output: "g"
    // Explanation:​​​​​​​
    // The weight of "abcd" is 7 + 5 + 3 + 4 = 19. The result modulo 26 is 19 % 26 = 19, which maps to 'g'.
    // Thus, the string formed by concatenating the mapped characters is "g".
    fmt.Println(mapWordWeights([]string{"abcd"}, []int{7,5,3,4,3,5,4,9,4,2,2,7,10,2,5,10,6,1,2,2,4,1,3,4,4,5})) // "g"

    fmt.Println(mapWordWeights([]string{"leetcode"}, []int{7,5,3,4,3,5,4,9,4,2,2,7,10,2,5,10,6,1,2,2,4,1,3,4,4,5})) // "v"
    fmt.Println(mapWordWeights([]string{"bluefrog"}, []int{7,5,3,4,3,5,4,9,4,2,2,7,10,2,5,10,6,1,2,2,4,1,3,4,4,5})) // "r"
    fmt.Println(mapWordWeights([]string{"freewu"}, []int{7,5,3,4,3,5,4,9,4,2,2,7,10,2,5,10,6,1,2,2,4,1,3,4,4,5})) // "g"

    fmt.Println(mapWordWeights1([]string{"abcd","def","xyz"}, []int{5,3,12,14,1,2,3,2,10,6,6,9,7,8,7,10,8,9,6,9,9,8,3,7,7,2})) // "rij"
    fmt.Println(mapWordWeights1([]string{"a","b","c"}, []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})) // "yyy"
    fmt.Println(mapWordWeights1([]string{"abcd"}, []int{7,5,3,4,3,5,4,9,4,2,2,7,10,2,5,10,6,1,2,2,4,1,3,4,4,5})) // "g"
    fmt.Println(mapWordWeights1([]string{"leetcode"}, []int{7,5,3,4,3,5,4,9,4,2,2,7,10,2,5,10,6,1,2,2,4,1,3,4,4,5})) // "v"
    fmt.Println(mapWordWeights1([]string{"bluefrog"}, []int{7,5,3,4,3,5,4,9,4,2,2,7,10,2,5,10,6,1,2,2,4,1,3,4,4,5})) // "r"
    fmt.Println(mapWordWeights1([]string{"freewu"}, []int{7,5,3,4,3,5,4,9,4,2,2,7,10,2,5,10,6,1,2,2,4,1,3,4,4,5})) // "g"
}