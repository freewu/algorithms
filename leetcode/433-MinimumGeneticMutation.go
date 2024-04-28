package main

// 433. Minimum Genetic Mutation
// A gene string can be represented by an 8-character long string, with choices from 'A', 'C', 'G', and 'T'.
// Suppose we need to investigate a mutation from a gene string startGene to a gene string endGene where one mutation is defined as one single character changed in the gene string.
//     For example, "AACCGGTT" --> "AACCGGTA" is one mutation.

// There is also a gene bank bank that records all the valid gene mutations. 
// A gene must be in bank to make it a valid gene string.

// Given the two gene strings startGene and endGene and the gene bank bank, 
// return the minimum number of mutations needed to mutate from startGene to endGene. 
// If there is no such a mutation, return -1.

// Note that the starting point is assumed to be valid, so it might not be included in the bank.

// Example 1:
// Input: startGene = "AACCGGTT", endGene = "AACCGGTA", bank = ["AACCGGTA"]
// Output: 1

// Example 2:
// Input: startGene = "AACCGGTT", endGene = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
// Output: 2
 
// Constraints:
//     0 <= bank.length <= 10
//     startGene.length == endGene.length == bank[i].length == 8
//     startGene, endGene, and bank[i] consist of only the characters ['A', 'C', 'G', 'T']. 

import "fmt"

// bfs
func minMutation(startGene string, endGene string, bank []string) int {
    queue, notVisitBank, counter := make([]string, 0), make(map[string]struct{}), 0
    queue = append(queue, startGene)
    distance := func (a, b string) int { // distance between two strings
        count := 0
        for i := 0; i < len(a); i++ {
            if a[i] != b[i] {
                count++
            }
        }
        return count
    }
    for _, s := range bank {
        notVisitBank[s] = struct{}{}
    }
    for len(queue) > 0 {
        queueLen := len(queue)
        for i := 0; i < queueLen; i++ {
            if queue[i] == endGene {
                return counter
            }
            for s, _ := range notVisitBank {
                if distance(s, queue[i]) == 1 {
                    queue = append(queue, s)
                    delete(notVisitBank, s)
                }
            }
        }
        counter++
        queue = queue[queueLen:]
    }
    return -1
}

func main() {
    // Example 1:
    // Input: startGene = "AACCGGTT", endGene = "AACCGGTA", bank = ["AACCGGTA"]
    // Output: 1
    fmt.Println(minMutation("AACCGGTT","AACCGGTA",[]string{"AACCGGTA"})) // 1
    // Example 2:
    // Input: startGene = "AACCGGTT", endGene = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
    // Output: 2
    fmt.Println(minMutation("AACCGGTT","AAACGGTA",[]string{"AACCGGTA","AACCGCTA","AAACGGTA"})) // 2
    fmt.Println(minMutation("AAAAACCC","AACCCCCC",[]string{"AAAACCCC","AAACCCCC","AACCCCCC"})) // 3
}