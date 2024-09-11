package main

// 1096. Brace Expansion II
// Under the grammar given below, strings can represent a set of lowercase words. 
// Let R(expr) denote the set of words the expression represents.

// The grammar can best be understood through simple examples:
//     Single letters represent a singleton set containing that word.
//         R("a") = {"a"}
//         R("w") = {"w"}
//     When we take a comma-delimited list of two or more expressions, we take the union of possibilities.
//         R("{a,b,c}") = {"a","b","c"}
//         R("{{a,b},{b,c}}") = {"a","b","c"} (notice the final set only contains each word at most once)
//     When we concatenate two expressions, we take the set of possible concatenations between two words where the first word comes from the first expression and the second word comes from the second expression.
//         R("{a,b}{c,d}") = {"ac","ad","bc","bd"}
//         R("a{b,c}{d,e}f{g,h}") = {"abdfg", "abdfh", "abefg", "abefh", "acdfg", "acdfh", "acefg", "acefh"}

// Formally, the three rules for our grammar:
//     For every lowercase letter x, we have R(x) = {x}.
//     For expressions e1, e2, ... , ek with k >= 2, we have R({e1, e2, ...}) = R(e1) ∪ R(e2) ∪ ...
//     For expressions e1 and e2, we have R(e1 + e2) = {a + b for (a, b) in R(e1) × R(e2)}, where + denotes concatenation, and × denotes the cartesian product.

// Given an expression representing a set of words under the given grammar, 
// return the sorted list of words that the expression represents.

// Example 1:
// Input: expression = "{a,b}{c,{d,e}}"
// Output: ["ac","ad","ae","bc","bd","be"]

// Example 2:
// Input: expression = "{{a,z},a{b,c},{ab,z}}"
// Output: ["a","ab","ac","z"]
// Explanation: Each distinct word is written only once in the final answer.

// Constraints:
//     1 <= expression.length <= 60
//     expression[i] consists of '{', '}', ','or lowercase English letters.
//     The given expression represents a set of words based on the grammar given in the description.

import "fmt"
import "strings"
import "sort"

// bfs 
func braceExpansionII(expression string) []string {
    OPENING_BRACE, CLOSING_BRACE := "{", "}"
    res, mp, queue := []string{}, make(map[string]bool), []string{ expression }
    for len(queue) > 0 {
        cur := queue[0] // pop
        queue = queue[1:]
        left := strings.Index(cur, OPENING_BRACE)
        if left == -1 {
            mp[cur] = true
            continue
        }
        i := left
        for i < len(cur) && string(cur[i]) != CLOSING_BRACE {
            if string(cur[i]) == OPENING_BRACE { left = i }
            i++
        }
        right := i
        processed := cur[:left]
        processing := strings.Split(cur[left + 1:right], ",")
        unprocessed := cur[right+1:]
        for _, part := range processing {
            sb := strings.Builder{}
            sb.WriteString(processed)
            sb.WriteString(part)
            sb.WriteString(unprocessed)
            queue = append(queue, sb.String())
        }
    }
    for v := range mp {
        res = append(res, v)
    }
    sort.Strings(res)
    return res
}

func braceExpansionII1(expression string) []string {
    getToken := func(expression string) []string {
        res, n := []string{}, len(expression)
        for i := 0; i < n; {
            if expression[i] >= 'a' && expression[i] <= 'z' {
                j := i
                for j < n && expression[j] >= 'a' && expression[j] <= 'z' {
                    j++
                }
                res = append(res, expression[i:j])
                if j < n && expression[j] == '{' {
                    res = append(res, "*")
                }
                i = j-1
            } else if expression[i] == '{' {
                res = append(res, string([]byte{expression[i]}))
            } else if expression[i] == '}' {
                res = append(res, string([]byte{expression[i]}))
                if i + 1 < n && (expression[i+1] >= 'a' && expression[i+1] <= 'z' || expression[i+1] == '{') {
                    res = append(res, "*")
                }
            } else if expression[i] == ',' {
                res = append(res, string([]byte{expression[i]}))
            }
            i++
        }
        return res
    }
    getPost := func (token []string) []string {
        res, os := []string{}, []string{}
        for _, x := range token {
            if x == "{" {
                os = append(os, x)
            } else if x == "}" {
                for os[len(os)-1] != "{" {
                    res = append(res, os[len(os)-1])
                    os = os[:len(os)-1]
                }
                os = os[:len(os)-1]
            } else if x == "," {
                for len(os) > 0 && os[len(os)-1] == "," || os[len(os)-1] == "*" {
                    res = append(res, os[len(os)-1])
                    os = os[:len(os)-1]
                }
                os = append(os, x)
            } else if x == "*" {
                for len(os) > 0 && os[len(os)-1] == "*" {
                    res = append(res, os[len(os)-1])
                    os = os[:len(os)-1]
                }
                os = append(os, x)
            } else {
                res = append(res, x)
            }
        }
        for len(os) > 0 {
            res = append(res, os[len(os)-1])
            os = os[:len(os)-1]
        }
        return res
    }
    calc := func(post []string) []string {
        st := []map[string]int{}
        for _, x := range post {
            n := len(st)
            if x == "*" {
                h := map[string]int{}
                for k1, _ := range st[n-2] {
                    for k2, _ :=range st[n-1] {
                        h[k1+k2] = 1
                    }
                }
                st[n-2] = h
                st = st[:n-1]
            } else if x == "," {
                for k, _ := range st[n-1] {
                    st[n-2][k]++
                }
                st = st[:n-1]
            } else {
                h := map[string]int{x:1}
                st = append(st, h)
            }
        }
        res := make([]string, 0, len(st[0]))
        for k, _ := range st[0] {
            res = append(res, k)
        }
        sort.Strings(res)
        return res
    }
    token := getToken(expression)
    post := getPost(token)
    return calc(post)
}

func main() {
    // Example 1:
    // Input: expression = "{a,b}{c,{d,e}}"
    // Output: ["ac","ad","ae","bc","bd","be"]
    fmt.Println(braceExpansionII("{a,b}{c,{d,e}}")) // ["ac","ad","ae","bc","bd","be"]
    // Example 2:
    // Input: expression = "{{a,z},a{b,c},{ab,z}}"
    // Output: ["a","ab","ac","z"]
    // Explanation: Each distinct word is written only once in the final answer.
    fmt.Println(braceExpansionII("{{a,z},a{b,c},{ab,z}}")) // ["a","ab","ac","z"]

    fmt.Println(braceExpansionII1("{a,b}{c,{d,e}}")) // ["ac","ad","ae","bc","bd","be"]
    fmt.Println(braceExpansionII1("{{a,z},a{b,c},{ab,z}}")) // ["a","ab","ac","z"]
}