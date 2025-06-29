package main

// 3597. Partition String 
// Given a string s, partition it into unique segments according to the following procedure:
//     1. Start building a segment beginning at index 0.
//     2. Continue extending the current segment character by character until the current segment has not been seen before.
//     3. Once the segment is unique, add it to your list of segments, mark it as seen, and begin a new segment from the next index.
//     4. Repeat until you reach the end of s.

// Return an array of strings segments, where segments[i] is the ith segment created.

// Example 1:
// Input: s = "abbccccd"
// Output: ["a","b","bc","c","cc","d"]
// Explanation:
// Index	Segment After Adding	Seen Segments	             Current Segment Seen Before?	New Segment	        Updated Seen Segments
// 0           "a"                     []                              No                          ""                  ["a"]
// 1           "b"                     ["a"]                           No                          ""                  ["a", "b"]
// 2           "b"                     ["a", "b"]                      Yes                         "b"                 ["a", "b"]
// 3           "bc"                    ["a", "b"]                      No                          ""                  ["a", "b", "bc"]
// 4           "c"	                    ["a", "b", "bc"]                No                          ""                  ["a", "b", "bc", "c"]
// 5           "c"	                    ["a", "b", "bc", "c"]           Yes                         "c"                 ["a", "b", "bc", "c"]
// 6           "cc"                    ["a", "b", "bc", "c"]           No                          ""                  ["a", "b", "bc", "c", "cc"]
// 7           "d"                     ["a", "b", "bc", "c", "cc"]     No                          ""                  ["a", "b", "bc", "c", "cc", "d"]
// Hence, the final output is ["a", "b", "bc", "c", "cc", "d"].

// Example 2:
// Input: s = "aaaa"
// Output: ["a","aa"]
// Explanation:
// Index	Segment After Adding	Seen Segments	Current Segment Seen Before?	New Segment	Updated Seen Segments
// 0	"a"	[]	No	""	["a"]
// 1	"a"	["a"]	Yes	"a"	["a"]
// 2	"aa"	["a"]	No	""	["a", "aa"]
// 3	"a"	["a", "aa"]	Yes	"a"	["a", "aa"]
// Hence, the final output is ["a", "aa"].

// Constraints:
//     1 <= s.length <= 10^5
//     s contains only lowercase English letters.

import "fmt"

func partitionString(s string) []string {
    seen := make(map[string]bool)
    res, cur := []string{}, ""
    for _, v := range s {
        cur += string(v)
        if !seen[cur] {
            res = append(res, cur)
            seen[cur] = true
            cur = ""
        }
    }
    return res
}

func partitionString1(s string) []string {
    res, n, cur := []string{}, len(s), ""
    visited := make(map[string]bool, n)
    for i := 0; i < n; {
        j := i + 1
        for j < n && visited[s[i: j]] {
            j++
        }
        cur = s[i: j]
        if !visited[cur] {
            visited[cur] = true
            res = append(res, cur)
        }
        i = j
    }
    return res
}

func partitionString2(s string) []string {
    res, seen, start := []string{}, make(map[string]bool), 0
    for i := 0; i < len(s); i++ {
        cur := s[start:i+1]
        if !seen[cur] {
            res = append(res, cur)
            seen[cur] = true
            start = i+1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abbccccd"
    // Output: ["a","b","bc","c","cc","d"]
    // Explanation:
    // Index	Segment After Adding	Seen Segments	             Current Segment Seen Before?	New Segment	        Updated Seen Segments
    // 0           "a"                     []                              No                          ""                  ["a"]
    // 1           "b"                     ["a"]                           No                          ""                  ["a", "b"]
    // 2           "b"                     ["a", "b"]                      Yes                         "b"                 ["a", "b"]
    // 3           "bc"                    ["a", "b"]                      No                          ""                  ["a", "b", "bc"]
    // 4           "c"	                    ["a", "b", "bc"]                No                          ""                  ["a", "b", "bc", "c"]
    // 5           "c"	                    ["a", "b", "bc", "c"]           Yes                         "c"                 ["a", "b", "bc", "c"]
    // 6           "cc"                    ["a", "b", "bc", "c"]           No                          ""                  ["a", "b", "bc", "c", "cc"]
    // 7           "d"                     ["a", "b", "bc", "c", "cc"]     No                          ""                  ["a", "b", "bc", "c", "cc", "d"]
    // Hence, the final output is ["a", "b", "bc", "c", "cc", "d"].
    fmt.Println(partitionString("abbccccd")) // ["a","b","bc","c","cc","d"]
    // Example 2:
    // Input: s = "aaaa"
    // Output: ["a","aa"]
    // Explanation:
    // Index	Segment After Adding	Seen Segments	Current Segment Seen Before?	New Segment	Updated Seen Segments
    // 0	"a"	[]	No	""	["a"]
    // 1	"a"	["a"]	Yes	"a"	["a"]
    // 2	"aa"	["a"]	No	""	["a", "aa"]
    // 3	"a"	["a", "aa"]	Yes	"a"	["a", "aa"]
    // Hence, the final output is ["a", "aa"].
    fmt.Println(partitionString("aaaa")) // ["a","aa"]

    fmt.Println(partitionString("leetcode")) // [l e et c o d]
    fmt.Println(partitionString("bluefrog")) // [b l u e f r o g]

    fmt.Println(partitionString1("abbccccd")) // ["a","b","bc","c","cc","d"]
    fmt.Println(partitionString1("aaaa")) // ["a","aa"]
    fmt.Println(partitionString1("leetcode")) // [l e et c o d]
    fmt.Println(partitionString1("bluefrog")) // [b l u e f r o g]

    fmt.Println(partitionString2("abbccccd")) // ["a","b","bc","c","cc","d"]
    fmt.Println(partitionString2("aaaa")) // ["a","aa"]
    fmt.Println(partitionString2("leetcode")) // [l e et c o d]
    fmt.Println(partitionString2("bluefrog")) // [b l u e f r o g]
}