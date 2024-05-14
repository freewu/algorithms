package main

// 332. Reconstruct Itinerary
// You are given a list of airline tickets where tickets[i] = [fromi, toi] represent the departure and the arrival airports of one flight. 
// Reconstruct the itinerary in order and return it.

// All of the tickets belong to a man who departs from "JFK", thus, the itinerary must begin with "JFK". 
// If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order when read as a single string.
//     For example, the itinerary ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"].

// You may assume all tickets form at least one valid itinerary. 
// You must use all the tickets once and only once.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/14/itinerary1-graph.jpg" />
// Input: tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
// Output: ["JFK","MUC","LHR","SFO","SJC"]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/03/14/itinerary2-graph.jpg" />
// Input: tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
// Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
// Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"] but it is larger in lexical order.
 
// Constraints:
//     1 <= tickets.length <= 300
//     tickets[i].length == 2
//     fromi.length == 3
//     toi.length == 3
//     fromi and toi consist of uppercase English letters.
//     fromi != toi

import "fmt"
import "sort"

func findItinerary(tickets [][]string) []string {
    graph := make(map[string][]string)
    for _, ticket := range tickets {
        graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
    }
    for key := range graph {
        sort.Sort(sort.Reverse(sort.StringSlice(graph[key])))
    }
    itinerary := []string{}
    var dfs func(airport string)
    dfs = func(airport string) {
        for len(graph[airport]) > 0 {
            next := graph[airport][len(graph[airport])-1]
            graph[airport] = graph[airport][:len(graph[airport])-1]
            dfs(next)
        }
        itinerary = append(itinerary, airport)
    }
    dfs("JFK") // the itinerary must begin with "JFK". 
    for i := 0; i < len(itinerary) / 2; i++ {
        itinerary[i], itinerary[len(itinerary)-1-i] = itinerary[len(itinerary)-1-i], itinerary[i]
    }
    return itinerary
}

func findItinerary1(tickets [][]string) []string {
    m, res := make(map[string][]string), []string{}
    for _, ticket := range tickets {
        src, dst := ticket[0], ticket[1]
        m[src] = append(m[src], dst)
    }
    for key := range m {
        sort.Strings(m[key])
    }
    var dfs func(curr string)
    dfs = func(curr string) {
        for {
            if v, ok := m[curr]; !ok || len(v) == 0 {
                break
            }
            tmp := m[curr][0]
            m[curr] = m[curr][1:]
            dfs(tmp)
        }
        res = append(res, curr)
    }
    dfs("JFK")
    for i := 0; i < len(res)/2; i++ {
        res[i], res[len(res) - 1 - i] = res[len(res) - 1 - i], res[i]
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/14/itinerary1-graph.jpg" />
    // Input: tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
    // Output: ["JFK","MUC","LHR","SFO","SJC"]
    fmt.Println(findItinerary([][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}})) // ["JFK","MUC","LHR","SFO","SJC"]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/03/14/itinerary2-graph.jpg" />
    // Input: tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
    // Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
    // Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"] but it is larger in lexical order.
    fmt.Println(findItinerary([][]string{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}})) // ["JFK","MUC","LHR","SFO","SJC"]
    
    fmt.Println(findItinerary1([][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}})) // ["JFK","MUC","LHR","SFO","SJC"]
    fmt.Println(findItinerary1([][]string{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}})) // ["JFK","MUC","LHR","SFO","SJC"]
}