package main

// 2126. Destroying Asteroids
// You are given an integer mass, which represents the original mass of a planet. 
// You are further given an integer array asteroids, where asteroids[i] is the mass of the ith asteroid.

// You can arrange for the planet to collide with the asteroids in any arbitrary order. 
// If the mass of the planet is greater than or equal to the mass of the asteroid, the asteroid is destroyed and the planet gains the mass of the asteroid. 
// Otherwise, the planet is destroyed.

// Return true if all asteroids can be destroyed. Otherwise, return false.

// Example 1:
// Input: mass = 10, asteroids = [3,9,19,5,21]
// Output: true
// Explanation: One way to order the asteroids is [9,19,5,3,21]:
// - The planet collides with the asteroid with a mass of 9. New planet mass: 10 + 9 = 19
// - The planet collides with the asteroid with a mass of 19. New planet mass: 19 + 19 = 38
// - The planet collides with the asteroid with a mass of 5. New planet mass: 38 + 5 = 43
// - The planet collides with the asteroid with a mass of 3. New planet mass: 43 + 3 = 46
// - The planet collides with the asteroid with a mass of 21. New planet mass: 46 + 21 = 67
// All asteroids are destroyed.

// Example 2:
// Input: mass = 5, asteroids = [4,9,23,4]
// Output: false
// Explanation: 
// The planet cannot ever gain enough mass to destroy the asteroid with a mass of 23.
// After the planet destroys the other asteroids, it will have a mass of 5 + 4 + 9 + 4 = 22.
// This is less than 23, so a collision would not destroy the last asteroid.

// Constraints:
//     1 <= mass <= 10^5
//     1 <= asteroids.length <= 10^5
//     1 <= asteroids[i] <= 10^5

import "fmt"
import "sort"
import "math/bits"
import "slices"

func asteroidsDestroyed(mass int, asteroids []int) bool {
    sort.Ints(asteroids)
    for _, asteroid := range asteroids {
        if mass < asteroid {
            return false
        }
        mass += asteroid
    }
    return true
}

func asteroidsDestroyed1(mass int, asteroids []int) bool {
    maxWidth := bits.Len(uint(slices.Max(asteroids)))
    sum, mn := make([]int, maxWidth), make([]int, maxWidth)
    for i := range mn {
        mn[i] = 1 << 61
    }
    for _, x := range asteroids {
        i := bits.Len(uint(x)) - 1
        sum[i] += x
        mn[i] = min(mn[i], x)
    }
    for i, v := range mn {
        if v == 1 << 61 { continue }
        if mass < v { // 无法摧毁这组的任意小行星
            return false
        }
        mass += sum[i] // 获得这组小行星的质量
    }
    return true
}

func main() {
    // Example 1:
    // Input: mass = 10, asteroids = [3,9,19,5,21]
    // Output: true
    // Explanation: One way to order the asteroids is [9,19,5,3,21]:
    // - The planet collides with the asteroid with a mass of 9. New planet mass: 10 + 9 = 19
    // - The planet collides with the asteroid with a mass of 19. New planet mass: 19 + 19 = 38
    // - The planet collides with the asteroid with a mass of 5. New planet mass: 38 + 5 = 43
    // - The planet collides with the asteroid with a mass of 3. New planet mass: 43 + 3 = 46
    // - The planet collides with the asteroid with a mass of 21. New planet mass: 46 + 21 = 67
    // All asteroids are destroyed.
    fmt.Println(asteroidsDestroyed(10, []int{3,9,19,5,21})) // true
    // Example 2:
    // Input: mass = 5, asteroids = [4,9,23,4]
    // Output: false
    // Explanation: 
    // The planet cannot ever gain enough mass to destroy the asteroid with a mass of 23.
    // After the planet destroys the other asteroids, it will have a mass of 5 + 4 + 9 + 4 = 22.
    // This is less than 23, so a collision would not destroy the last asteroid.
    fmt.Println(asteroidsDestroyed(5, []int{4,9,23,4})) // false

    fmt.Println(asteroidsDestroyed(5, []int{1,2,3,4,5,6,7,8,9})) // true
    fmt.Println(asteroidsDestroyed(5, []int{9,8,7,6,5,4,3,2,1})) // true

    fmt.Println(asteroidsDestroyed1(10, []int{3,9,19,5,21})) // true
    fmt.Println(asteroidsDestroyed1(5, []int{4,9,23,4})) // false
    fmt.Println(asteroidsDestroyed1(5, []int{1,2,3,4,5,6,7,8,9})) // true
    fmt.Println(asteroidsDestroyed1(5, []int{9,8,7,6,5,4,3,2,1})) // true
}