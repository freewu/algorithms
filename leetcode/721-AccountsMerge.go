package main

// 721. Accounts Merge
// Given a list of accounts where each element accounts[i] is a list of strings, 
// where the first element accounts[i][0] is a name, and the rest of the elements are emails representing emails of the account.

// Now, we would like to merge these accounts. 
// Two accounts definitely belong to the same person if there is some common email to both accounts. 
// Note that even if two accounts have the same name, they may belong to different people as people could have the same name. 
// A person can have any number of accounts initially, but all of their accounts definitely have the same name.

// After merging the accounts, return the accounts in the following format: 
//     the first element of each account is the name, and the rest of the elements are emails in sorted order. 
// The accounts themselves can be returned in any order.

// Example 1:
// Input: accounts = [["John","johnsmith@mail.com","john_newyork@mail.com"],["John","johnsmith@mail.com","john00@mail.com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
// Output: [["John","john00@mail.com","john_newyork@mail.com","johnsmith@mail.com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
// Explanation:
// The first and second John's are the same person as they have the common email "johnsmith@mail.com".
// The third John and Mary are different people as none of their email addresses are used by other accounts.
// We could return these lists in any order, for example the answer [['Mary', 'mary@mail.com'], ['John', 'johnnybravo@mail.com'], 
// ['John', 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com']] would still be accepted.

// Example 2:
// Input: accounts = [["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"],["Kevin","Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"],["Ethan","Ethan5@m.co","Ethan4@m.co","Ethan0@m.co"],["Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"],["Fern","Fern5@m.co","Fern1@m.co","Fern0@m.co"]]
// Output: [["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],["Gabe","Gabe0@m.co","Gabe1@m.co","Gabe3@m.co"],["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.co"],["Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],["Fern","Fern0@m.co","Fern1@m.co","Fern5@m.co"]]

// Constraints:
//     1 <= accounts.length <= 1000
//     2 <= accounts[i].length <= 10
//     1 <= accounts[i][j].length <= 30
//     accounts[i][0] consists of English letters.
//     accounts[i][j] (for j > 0) is a valid email.

import "fmt"
import "sort"

// dfs
func accountsMerge(accounts [][]string) [][]string {
    etoa := make(map[string][]int)
    for i, account := range accounts {
        for j := 1; j < len(account); j++ {
            email := account[j]
            etoa[email] = append(etoa[email], i)
        }
    }
    res, visited := [][]string{}, make([]bool, len(accounts))
    var dfs func(visited []bool, i int, accounts [][]string, etoa map[string][]int, emails map[string]bool)
    dfs = func (visited []bool, i int, accounts [][]string, etoa map[string][]int, emails map[string]bool) {
        if visited[i] { return }
        visited[i] = true
        account := accounts[i]
        for j := 1; j < len(account); j++ {
            email := account[j]
            emails[email] = true
            for _, neighbour := range etoa[email] {
                dfs(visited, neighbour, accounts, etoa, emails)
            }
        }
    }
    for i := 0; i < len(accounts); i++ {
        if visited[i] { continue }
        name, emails := accounts[i][0], make(map[string]bool)
        dfs(visited, i, accounts, etoa, emails)
        vals, index := make([]string, len(emails)), 0
        for email, _ := range emails {
            vals[index] = email
            index++
        }
        sort.Strings(vals)
        merged := make([]string, len(vals) + 1)
        merged[0] = name
        copy(merged[1:], vals)
        res = append(res, merged)
    }
    return res
}

func accountsMerge1(accounts [][]string) [][]string {
    emailToID := make(map[string]int) // Initialize a map to store the email to account ID mapping
    emailToName := make(map[string]string) // Initialize a map to store the email to name mapping
    unionFind := make(map[int][]string) // Initialize a map to store the account ID to emails mapping
    id := 0 // Assign unique IDs to each email and store the name
    for _, account := range accounts {
        name := account[0] // First element is the name
        for _, email := range account[1:] {
            // Assign a unique ID to each email
            if _, ok := emailToID[email]; !ok {
                emailToID[email] = id
                id++
            }
            // Store the email to name mapping
            emailToName[email] = name
        }
    }
    // Initialize the parent array for Union-Find algorithm
    parent := make([]int, id)
    for i := range parent {
        parent[i] = i
    }
    // Function to find the root of a given element in Union-Find
    find := func(x int) int {
        for x != parent[x] {
            parent[x] = parent[parent[x]] // Path compression
            x = parent[x]
        }
        return x
    }
    // Function to union two elements in Union-Find
    union := func(x, y int) {
        rootX, rootY := find(x), find(y)
        if rootX != rootY {
            parent[rootX] = rootY
        }
    }
    // Step 2: Union emails in the same account
    for _, account := range accounts {
        firstEmailID := emailToID[account[1]]
        for _, email := range account[2:] {
            union(firstEmailID, emailToID[email]) // Union emails in the same account
        }
    }
    // Step 3: Group emails by account
    for email, id := range emailToID {
        rootID := find(id)
        unionFind[rootID] = append(unionFind[rootID], email)
    }
    // Step 4: Format the result
    res := [][]string{}
    for _, emails := range unionFind {
        sort.Strings(emails)                         // Sort emails
        name := emailToName[emails[0]]               // Get the name associated with the first email
        account := append([]string{name}, emails...) // Construct the account
        res = append(res, account)             // Add account to result
    }
    return res // Return the merged accounts
}

func main() {
    // Example 1:
    // Input: accounts = [["John","johnsmith@mail.com","john_newyork@mail.com"],["John","johnsmith@mail.com","john00@mail.com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
    // Output: [["John","john00@mail.com","john_newyork@mail.com","johnsmith@mail.com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
    // Explanation:
    // The first and second John's are the same person as they have the common email "johnsmith@mail.com".
    // The third John and Mary are different people as none of their email addresses are used by other accounts.
    // We could return these lists in any order, for example the answer [['Mary', 'mary@mail.com'], ['John', 'johnnybravo@mail.com'], 
    // ['John', 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com']] would still be accepted.
    accounts1 := [][]string{
        {"John","johnsmith@mail.com","john_newyork@mail.com"},
        {"John","johnsmith@mail.com","john00@mail.com"},
        {"Mary","mary@mail.com"},
        {"John","johnnybravo@mail.com"},
    }
    fmt.Println(accountsMerge(accounts1)) // [["John","john00@mail.com","john_newyork@mail.com","johnsmith@mail.com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
    // Explanation:
    // Example 2:
    // Input: accounts = [["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"],["Kevin","Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"],["Ethan","Ethan5@m.co","Ethan4@m.co","Ethan0@m.co"],["Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"],["Fern","Fern5@m.co","Fern1@m.co","Fern0@m.co"]]
    // Output: [["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],["Gabe","Gabe0@m.co","Gabe1@m.co","Gabe3@m.co"],["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.co"],["Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],["Fern","Fern0@m.co","Fern1@m.co","Fern5@m.co"]]
    accounts2 := [][]string{
        {"Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"},
        {"Kevin","Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"},
        {"Ethan","Ethan5@m.co","Ethan4@m.co","Ethan0@m.co"},
        {"Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"},
        {"Fern","Fern5@m.co","Fern1@m.co","Fern0@m.co"},
    }
    fmt.Println(accountsMerge(accounts2)) // [["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],["Gabe","Gabe0@m.co","Gabe1@m.co","Gabe3@m.co"],["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.co"],["Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],["Fern","Fern0@m.co","Fern1@m.co","Fern5@m.co"]]

    fmt.Println(accountsMerge1(accounts1)) // [["John","john00@mail.com","john_newyork@mail.com","johnsmith@mail.com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
    fmt.Println(accountsMerge1(accounts2)) // [["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],["Gabe","Gabe0@m.co","Gabe1@m.co","Gabe3@m.co"],["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.co"],["Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],["Fern","Fern0@m.co","Fern1@m.co","Fern5@m.co"]]

}