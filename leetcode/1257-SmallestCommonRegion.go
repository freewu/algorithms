package main

// 1257. Smallest Common Region
// You are given some lists of regions where the first region of each list includes all other regions in that list.

// Naturally, if a region x contains another region y then x is bigger than y. 
// Also, by definition, a region x contains itself.

// Given two regions: region1 and region2, return the smallest region that contains both of them.

// If you are given regions r1, r2, and r3 such that r1 includes r3, it is guaranteed there is no r2 such that r2 includes r3.

// It is guaranteed the smallest region exists.

// Example 1:
// Input:
// regions = [["Earth","North America","South America"],
// ["North America","United States","Canada"],
// ["United States","New York","Boston"],
// ["Canada","Ontario","Quebec"],
// ["South America","Brazil"]],
// region1 = "Quebec",
// region2 = "New York"
// Output: "North America"

// Example 2:
// Input: regions = [["Earth", "North America", "South America"],["North America", "United States", "Canada"],["United States", "New York", "Boston"],["Canada", "Ontario", "Quebec"],["South America", "Brazil"]], region1 = "Canada", region2 = "South America"
// Output: "Earth"

// Constraints:
//     2 <= regions.length <= 10^4
//     2 <= regions[i].length <= 20
//     1 <= regions[i][j].length, region1.length, region2.length <= 20
//     region1 != region2
//     regions[i][j], region1, and region2 consist of English letters.

import "fmt"

func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
    region := map[string]string{}
    for _, strings := range regions {
        v := strings[0]
        for i := 1; i < len(strings); i++ {
            region[strings[i]] = v
        }
    }
    visited := map[string]bool{ region1: true, region2: true }
    for region1 != "" {
        region1 = region[region1]
        if visited[region1] {
            return region1
        }
        visited[region1]=true
    }
    for region2 != "" {
        region2 = region[region2]
        if visited[region2] {
            return region2
        }
        visited[region2] = true
    }
    return ""
}

func main() {
    // Example 1:
    // Input:
    // regions = [["Earth","North America","South America"],
    // ["North America","United States","Canada"],
    // ["United States","New York","Boston"],
    // ["Canada","Ontario","Quebec"],
    // ["South America","Brazil"]],
    // region1 = "Quebec",
    // region2 = "New York"
    // Output: "North America"
    regions1 := [][]string{
        {"Earth","North America","South America"},
        {"North America","United States","Canada"},
        {"United States","New York","Boston"},
        {"Canada","Ontario","Quebec"},
        {"South America","Brazil"},
    }
    fmt.Println(findSmallestRegion(regions1, "Quebec", "New York")) // "North America"
    // Example 2:
    // Input: regions = [["Earth", "North America", "South America"],["North America", "United States", "Canada"],["United States", "New York", "Boston"],["Canada", "Ontario", "Quebec"],["South America", "Brazil"]], region1 = "Canada", region2 = "South America"
    // Output: "Earth"
    regions2 := [][]string{
        {"Earth", "North America", "South America"},
        {"North America", "United States", "Canada"},
        {"United States", "New York", "Boston"},
        {"Canada", "Ontario", "Quebec"},
        {"South America", "Brazil"},
    }
    fmt.Println(findSmallestRegion(regions2, "Canada", "South America")) // "Earth"
}