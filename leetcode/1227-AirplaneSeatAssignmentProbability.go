package main

// 1227. Airplane Seat Assignment Probability
// n passengers board an airplane with exactly n seats. 
// The first passenger has lost the ticket and picks a seat randomly. 
// But after that, the rest of the passengers will:
//     Take their own seat if it is still available, and
//     Pick other seats randomly when they find their seat occupied

// Return the probability that the nth person gets his own seat.

// Example 1:
// Input: n = 1
// Output: 1.00000
// Explanation: The first person can only get the first seat.

// Example 2:
// Input: n = 2
// Output: 0.50000
// Explanation: The second person has a probability of 0.5 to get the second seat (when first person gets the first seat).
 
// Constraints:
//     1 <= n <= 10^5

import "fmt"

func nthPersonGetsNthSeat(n int) float64 {
    // 1/n p1 sits correctly
    // 1/n p1 takes pn's seat
    // (n-2)/n * 1/2 pn gets his seats when p1 is in some else's seat
    // 1/2 becoz the poor passenge pi either sits at p1 or pn
    // if his seat is taken by p1
    if n == 1 {
        return 1.0
    }
    return 0.5
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: 1.00000
    // Explanation: The first person can only get the first seat.
    fmt.Println(nthPersonGetsNthSeat(1)) // 1.00000
    // Example 2:
    // Input: n = 2
    // Output: 0.50000
    // Explanation: The second person has a probability of 0.5 to get the second seat (when first person gets the first seat).
    fmt.Println(nthPersonGetsNthSeat(2)) // 0.50000
}