package main

// 990. Satisfiability of Equality Equations
// You are given an array of strings equations that represent relationships between variables 
// where each string equations[i] is of length 4 and takes one of two different forms: 
// "xi==yi" or "xi!=yi".Here, xi and yi are lowercase letters (not necessarily different) that represent one-letter variable names.

// Return true if it is possible to assign integers to variable names so as to satisfy all the given equations, or false otherwise.

// Example 1:
// Input: equations = ["a==b","b!=a"]
// Output: false
// Explanation: If we assign say, a = 1 and b = 1, then the first equation is satisfied, but not the second.
// There is no way to assign the variables to satisfy both equations.

// Example 2:
// Input: equations = ["b==a","a==b"]
// Output: true
// Explanation: We could assign a = 1 and b = 1 to satisfy both equations.

// Constraints:
//     1 <= equations.length <= 500
//     equations[i].length == 4
//     equations[i][0] is a lowercase letter.
//     equations[i][1] is either '=' or '!'.
//     equations[i][2] is '='.
//     equations[i][3] is a lowercase letter.

import "fmt"

func equationsPossible(equations []string) bool {
    nodesEqual := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
    root := func (nodes *[]int, p int) int {
        for p != (*nodes)[p] {
            p = (*nodes)[p]
        }
        return p
    }
    // Union-find
    union := func(nodes *[]int, p, q int) {
        pRoot := root(nodes, p)
        qRoot := root(nodes, q)
        (*nodes)[qRoot] = pRoot
    }
    connected := func(nodes *[]int, p, q int) bool {
        return root(nodes, p) == root(nodes, q)
    }
    for _, eq := range equations {
        if eq[1] == '=' {
            union(&nodesEqual, int(eq[0]-'a'), int(eq[3]-'a'))
        }
    }
    for _, eq := range equations {
        if eq[1] == '!' {
            if connected(&nodesEqual, int(eq[0]-'a'), int(eq[3]-'a')) {
                return false
            }
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: equations = ["a==b","b!=a"]
    // Output: false
    // Explanation: If we assign say, a = 1 and b = 1, then the first equation is satisfied, but not the second.
    // There is no way to assign the variables to satisfy both equations.
    fmt.Println(equationsPossible([]string{"a==b","b!=a"})) // false
    // Example 2:
    // Input: equations = ["b==a","a==b"]
    // Output: true
    // Explanation: We could assign a = 1 and b = 1 to satisfy both equations.
    fmt.Println(equationsPossible([]string{"b==a","a==b"})) // true
}