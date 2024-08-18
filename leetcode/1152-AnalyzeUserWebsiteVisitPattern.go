package main

// 1152. Analyze User Website Visit Pattern
// You are given two string arrays username and website and an integer array timestamp. 
// All the given arrays are of the same length and the tuple [username[i], 
// website[i], timestamp[i]] indicates that the user username[i] visited the website website[i] at time timestamp[i].

// A pattern is a list of three websites (not necessarily distinct).
//     For example, ["home", "away", "love"], ["leetcode", "love", "leetcode"], 
//     and ["luffy", "luffy", "luffy"] are all patterns.

// The score of a pattern is the number of users that visited all the websites in the pattern in the same order they appeared in the pattern.

//     1. For example, if the pattern is ["home", "away", "love"], 
//        the score is the number of users x such that x visited "home" then visited "away" and visited "love" after that.
//     2. Similarly, if the pattern is ["leetcode", "love", "leetcode"], 
//        the score is the number of users x such that x visited "leetcode" then visited "love" and visited "leetcode" one more time after that.
//     3. Also, if the pattern is ["luffy", "luffy", "luffy"], 
//        the score is the number of users x such that x visited "luffy" three different times at different timestamps.

// Return the pattern with the largest score. 
// If there is more than one pattern with the same largest score, return the lexicographically smallest such pattern.

// Example 1:
// Input: username = ["joe","joe","joe","james","james","james","james","mary","mary","mary"], timestamp = [1,2,3,4,5,6,7,8,9,10], website = ["home","about","career","home","cart","maps","home","home","about","career"]
// Output: ["home","about","career"]
// Explanation: The tuples in this example are:
// ["joe","home",1],["joe","about",2],["joe","career",3],["james","home",4],["james","cart",5],["james","maps",6],["james","home",7],["mary","home",8],["mary","about",9], and ["mary","career",10].
// The pattern ("home", "about", "career") has score 2 (joe and mary).
// The pattern ("home", "cart", "maps") has score 1 (james).
// The pattern ("home", "cart", "home") has score 1 (james).
// The pattern ("home", "maps", "home") has score 1 (james).
// The pattern ("cart", "maps", "home") has score 1 (james).
// The pattern ("home", "home", "home") has score 0 (no user visited home 3 times).

// Example 2:
// Input: username = ["ua","ua","ua","ub","ub","ub"], timestamp = [1,2,3,4,5,6], website = ["a","b","a","a","b","c"]
// Output: ["a","b","a"]

// Constraints:
//     3 <= username.length <= 50
//     1 <= username[i].length <= 10
//     timestamp.length == username.length
//     1 <= timestamp[i] <= 109
//     website.length == username.length
//     1 <= website[i].length <= 10
//     username[i] and website[i] consist of lowercase English letters.
//     It is guaranteed that there is at least one user who visited at least three websites.
//     All the tuples [username[i], timestamp[i], website[i]] are unique.

import "fmt"
import "sort"

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
    type Log struct {
        name string
        timestamp int
        website string
    }
    n := make([]Log, len(username))
    for i := 0; i < len(username); i++ {
        n[i] = Log{username[i], timestamp[i], website[i]}
    }
    sort.Slice(n, func(i, j int) bool {
        return n[i].timestamp < n[j].timestamp // 保证用户访问记录有序
    })
    mp := make(map[string][]Log)
    for i := 0; i < len(n); i++ {
        mp[n[i].name] = append(mp[n[i].name], n[i]) // 统计每个用户的所有访问记录
    }
    route := make(map[[3]string]int)
    for _, v := range mp {
        tmp := make(map[[3]string]int)
        for i := 0; i < len(v); i++ {
            for j := i + 1; j < len(v); j++ {
                for k := j + 1; k < len(v); k++ {
                    tmp[[3]string{v[i].website, v[j].website, v[k].website}] = 1 // 获取每个访问路径
                }
            }
        }
        for k1, v1 := range tmp {
            route[k1] += v1
        }
    }
    res, mx := [3]string{}, -1
    for k, v := range route {
        if v > mx {
            res = k
            mx = v
        } else if v == mx {
            if k[0] < res[0] || 
               (k[0] == res[0] && k[1] < res[1] || 
               (k[0] == res[0] && k[1] == res[1] && k[2] < res[2])) {
                res = k
            }
        }
    }
    return []string{res[0], res[1], res[2]}
}

func main() {
    // Example 1:
    // Input: username = ["joe","joe","joe","james","james","james","james","mary","mary","mary"], timestamp = [1,2,3,4,5,6,7,8,9,10], website = ["home","about","career","home","cart","maps","home","home","about","career"]
    // Output: ["home","about","career"]
    // Explanation: The tuples in this example are:
    // ["joe","home",1],["joe","about",2],["joe","career",3],["james","home",4],["james","cart",5],["james","maps",6],["james","home",7],["mary","home",8],["mary","about",9], and ["mary","career",10].
    // The pattern ("home", "about", "career") has score 2 (joe and mary).
    // The pattern ("home", "cart", "maps") has score 1 (james).
    // The pattern ("home", "cart", "home") has score 1 (james).
    // The pattern ("home", "maps", "home") has score 1 (james).
    // The pattern ("cart", "maps", "home") has score 1 (james).
    // The pattern ("home", "home", "home") has score 0 (no user visited home 3 times).
    fmt.Println(mostVisitedPattern(
        []string{"joe","joe","joe","james","james","james","james","mary","mary","mary"},
        []int{1,2,3,4,5,6,7,8,9,10},
        []string{"home","about","career","home","cart","maps","home","home","about","career"},
    )) // ["home","about","career"]
    // Example 2:
    // Input: username = ["ua","ua","ua","ub","ub","ub"], timestamp = [1,2,3,4,5,6], website = ["a","b","a","a","b","c"]
    // Output: ["a","b","a"]
    fmt.Println(mostVisitedPattern(
        []string{"ua","ua","ua","ub","ub","ub"},
        []int{1,2,3,4,5,6},
        []string{"a","b","a","a","b","c"},
    )) // ["a","b","a"]
}