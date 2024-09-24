package main

// 2306. Naming a Company
// You are given an array of strings ideas that represents a list of names to be used in the process of naming a company. 
// The process of naming a company is as follows:
//     1. Choose 2 distinct names from ideas, call them ideaA and ideaB.
//     2. Swap the first letters of ideaA and ideaB with each other.
//     3. If both of the new names are not found in the original ideas, then the name ideaA ideaB (the concatenation of ideaA and ideaB, separated by a space) is a valid company name.
//     4. Otherwise, it is not a valid name.

// Return the number of distinct valid names for the company.

// Example 1:
// Input: ideas = ["coffee","donuts","time","toffee"]
// Output: 6
// Explanation: The following selections are valid:
// - ("coffee", "donuts"): The company name created is "doffee conuts".
// - ("donuts", "coffee"): The company name created is "conuts doffee".
// - ("donuts", "time"): The company name created is "tonuts dime".
// - ("donuts", "toffee"): The company name created is "tonuts doffee".
// - ("time", "donuts"): The company name created is "dime tonuts".
// - ("toffee", "donuts"): The company name created is "doffee tonuts".
// Therefore, there are a total of 6 distinct company names.
// The following are some examples of invalid selections:
// - ("coffee", "time"): The name "toffee" formed after swapping already exists in the original array.
// - ("time", "toffee"): Both names are still the same after swapping and exist in the original array.
// - ("coffee", "toffee"): Both names formed after swapping already exist in the original array.

// Example 2:
// Input: ideas = ["lack","back"]
// Output: 0
// Explanation: There are no valid selections. Therefore, 0 is returned.

// Constraints:
//     2 <= ideas.length <= 5 * 10^4
//     1 <= ideas[i].length <= 10
//     ideas[i] consists of lowercase English letters.
//     All the strings in ideas are unique.

import "fmt"

func distinctNames(ideas []string) int64 {
    res, words := 0, [26]map[string]bool{}
    for i := range words {
        words[i] = make(map[string]bool)
    }
    for _, idea := range ideas {
        bi := []byte(idea)
        words[int(bi[0]-'a')][string(bi[1:])] = true
    }
    for i := 0; i < 25; i++ {
        for j := i + 1; j < 26; j++ {
            mut := 0
            for suf := range words[i] {
                if words[j][suf] {
                    mut++
                }
            }
            res += 2 * (len(words[i]) - mut) * (len(words[j]) - mut)
        }
    }
    return int64(res)
}

func distinctNames1(ideas []string) int64 {
    size := [26]int{} // 集合大小
    intersection := [26][26]int{} // 交集大小
    groups := map[string]int{} // 后缀 -> 首字母
    for _, s := range ideas {
        b := s[0] - 'a'
        size[b]++ // 增加集合大小
        suffix := s[1:]
        mask := groups[suffix]
        groups[suffix] = mask | 1<<b // 把 b 加到 mask 中
        for a := 0; a < 26; a++ { // a 是和 s 有着相同后缀的字符串的首字母
            if mask >> a & 1 > 0 { // a 在 mask 中
                intersection[b][a]++ // 增加交集大小
                intersection[a][b]++
            }
        }
    }
    res := 0
    for a := 1; a < 26; a++ { // 枚举所有组对
        for b := 0; b < a; b++ {
            m := intersection[a][b]
            res += (size[a] - m) * (size[b] - m)
        }
    }
    return int64(res * 2) // 乘 2 放到最后
}

func main() {
    // Example 1:
    // Input: ideas = ["coffee","donuts","time","toffee"]
    // Output: 6
    // Explanation: The following selections are valid:
    // - ("coffee", "donuts"): The company name created is "doffee conuts".
    // - ("donuts", "coffee"): The company name created is "conuts doffee".
    // - ("donuts", "time"): The company name created is "tonuts dime".
    // - ("donuts", "toffee"): The company name created is "tonuts doffee".
    // - ("time", "donuts"): The company name created is "dime tonuts".
    // - ("toffee", "donuts"): The company name created is "doffee tonuts".
    // Therefore, there are a total of 6 distinct company names.
    // The following are some examples of invalid selections:
    // - ("coffee", "time"): The name "toffee" formed after swapping already exists in the original array.
    // - ("time", "toffee"): Both names are still the same after swapping and exist in the original array.
    // - ("coffee", "toffee"): Both names formed after swapping already exist in the original array.
    fmt.Println(distinctNames([]string{"coffee","donuts","time","toffee"})) // 6
    // Example 2: 
    // Input: ideas = ["lack","back"]
    // Output: 0
    // Explanation: There are no valid selections. Therefore, 0 is returned.
    fmt.Println(distinctNames([]string{"lack","back"})) // 0

    fmt.Println(distinctNames1([]string{"coffee","donuts","time","toffee"})) // 6
    fmt.Println(distinctNames1([]string{"lack","back"})) // 0
}