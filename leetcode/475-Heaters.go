package main

// 475. Heaters
// Winter is coming! During the contest, your first job is to design a standard heater 
// with a fixed warm radius to warm all the houses.

// Every house can be warmed, as long as the house is within the heater's warm radius range. 

// Given the positions of houses and heaters on a horizontal line, 
// return the minimum radius standard of heaters so that those heaters could cover all houses.

// Notice that all the heaters follow your radius standard, and the warm radius will the same.

// Example 1:
// Input: houses = [1,2,3], heaters = [2]
// Output: 1
// Explanation: The only heater was placed in the position 2, and if we use the radius 1 standard, then all the houses can be warmed.

// Example 2:
// Input: houses = [1,2,3,4], heaters = [1,4]
// Output: 1
// Explanation: The two heaters were placed at positions 1 and 4. We need to use a radius 1 standard, then all the houses can be warmed.

// Example 3:
// Input: houses = [1,5], heaters = [2]
// Output: 3

// Constraints:
//     1 <= houses.length, heaters.length <= 3 * 10^4
//     1 <= houses[i], heaters[i] <= 10^9

import "fmt"
import "sort"

func findRadius(houses []int, heaters []int) int {
    sort.Ints(houses)
    sort.Ints(heaters)

    heaters = append([]int{-1 << 31 }, heaters...)
    heaters = append(heaters, 1 << 31 )
    radius := 0

    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, house := range houses {
        left, right := 0, len(heaters) - 1
        for left <= right {
            mid := left + (right-left) / 2
            if heaters[mid] < house {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        currRadius := abs(heaters[left] - house)
        if left-1 >= 0 {
            currRadius = min(currRadius, abs(house-heaters[left-1]))
        }
        radius = max(radius, currRadius)
    }
    return radius
}

func findRadius1(houses []int, heaters []int) int {
    sort.Ints(houses)
    sort.Ints(heaters)
    res, r, m, n := 0, 0, len(heaters), len(houses)

    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        for ; r < m-1 && abs(houses[i]-heaters[r]) >= abs(houses[i] - heaters[r+1]); {
            r++
        }
        res = max(res,abs(houses[i]-heaters[r]))
    }
    return res
}


func main() {
    // Example 1:
    // Input: houses = [1,2,3], heaters = [2]
    // Output: 1
    // Explanation: The only heater was placed in the position 2, and if we use the radius 1 standard, then all the houses can be warmed.
    fmt.Println(findRadius([]int{1,2,3},[]int{2})) // 1
    // Example 2:
    // Input: houses = [1,2,3,4], heaters = [1,4]
    // Output: 1
    // Explanation: The two heaters were placed at positions 1 and 4. We need to use a radius 1 standard, then all the houses can be warmed.
    fmt.Println(findRadius([]int{1,2,3,4},[]int{1,4})) // 1
    // Example 3:
    // Input: houses = [1,5], heaters = [2]
    // Output: 3
    fmt.Println(findRadius([]int{1,5},[]int{2})) // 3

    fmt.Println(findRadius1([]int{1,2,3},[]int{2})) // 1
    fmt.Println(findRadius1([]int{1,2,3,4},[]int{1,4})) // 1
    fmt.Println(findRadius1([]int{1,5},[]int{2})) // 3
}