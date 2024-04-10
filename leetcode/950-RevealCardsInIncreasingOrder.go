package main

// 950. Reveal Cards In Increasing Order
// You are given an integer array deck.
// There is a deck of cards where every card has a unique integer. 
// The integer on the ith card is deck[i].

// You can order the deck in any order you want. 
// Initially, all the cards start face down (unrevealed) in one deck.
// You will do the following steps repeatedly until all cards are revealed:
//     Take the top card of the deck, reveal it, and take it out of the deck.
//     If there are still cards in the deck then put the next top card of the deck at the bottom of the deck.
//     If there are still unrevealed cards, go back to step 1. Otherwise, stop.

// Return an ordering of the deck that would reveal the cards in increasing order.
// Note that the first entry in the answer is considered to be the top of the deck.

// Example 1:
// Input: deck = [17,13,11,2,3,5,7]
// Output: [2,13,3,11,5,17,7]
// Explanation: 
// We get the deck in the order [17,13,11,2,3,5,7] (this order does not matter), and reorder it.
// [2 3 5 7 11 13 17]
// After reordering, the deck starts as [2,13,3,11,5,17,7], where 2 is the top of the deck.
// We reveal 2, and move 13 to the bottom.  The deck is now [3,11,5,17,7,13].
// We reveal 3, and move 11 to the bottom.  The deck is now [5,17,7,13,11].
// We reveal 5, and move 17 to the bottom.  The deck is now [7,13,11,17].
// We reveal 7, and move 13 to the bottom.  The deck is now [11,17,13].
// We reveal 11, and move 17 to the bottom.  The deck is now [13,17].
// We reveal 13, and move 17 to the bottom.  The deck is now [17].
// We reveal 17.
// Since all the cards revealed are in increasing order, the answer is correct.

// Example 2:
// Input: deck = [1,1000]
// Output: [1,1000]

// Constraints:
//     1 <= deck.length <= 1000
//     1 <= deck[i] <= 10^6
//     All the values of deck are unique.
import "fmt"
import "sort"

func deckRevealedIncreasing(deck []int) []int {
    if len(deck) < 2 { // early return in case of an edge case (only 1 card in deck)
        return deck
    }
    sort.Ints(deck) // // sort incoming deck in ascending order
    // create helper slice of 2x lenght
    sorted := make([]int, len(deck)*2)
    // fill helper slice with sorted deck items leaving every second space empty
    y := 0
    for i := 0; i < len(sorted); i++ {
        if i % 2 == 0 {
            sorted[i] = deck[y]
            y++
        } 
    }
    sorted = sorted[:len(sorted)-1] // trim the last empty element as it is not needed
    // backtrack from the end, filling the empty items 
    lastEmpty,lastItem := len(sorted) - 2, len(sorted) - 1
    for lastEmpty > 0 {
        sorted[lastEmpty] = sorted[lastItem]
        lastItem--
        lastEmpty -= 2
    }
    return sorted[:len(deck)] // return only as many items as in the deck
}

func deckRevealedIncreasing1(deck []int) []int {
    n := len(deck)
    indices := make([]int, n)
    for i := 0; i < n; i++ {
        indices[i] = i
    }
    res := make([]int, n)
    sort.Ints(deck)
    fmt.Println(deck)
    for _, c := range deck {
        // 从牌组顶部抽一张牌，显示它，然后将其从牌组中移出。
        res[indices[0]] = c
        indices = indices[1:]
        // 如果牌组中仍有牌，则将下一张处于牌组顶部的牌放在牌组的底部。
        if len(indices) > 0 {
           indices = append(indices, indices[0])
           indices = indices[1:]
        }
    }
    return res
}

func main() {
    // We get the deck in the order [17,13,11,2,3,5,7] (this order does not matter), and reorder it.
    // After reordering, the deck starts as [2,13,3,11,5,17,7], where 2 is the top of the deck.
    // We reveal 2, and move 13 to the bottom.  The deck is now [3,11,5,17,7,13].
    // We reveal 3, and move 11 to the bottom.  The deck is now [5,17,7,13,11].
    // We reveal 5, and move 17 to the bottom.  The deck is now [7,13,11,17].
    // We reveal 7, and move 13 to the bottom.  The deck is now [11,17,13].
    // We reveal 11, and move 17 to the bottom.  The deck is now [13,17].
    // We reveal 13, and move 17 to the bottom.  The deck is now [17].
    // We reveal 17.
    // Since all the cards revealed are in increasing order, the answer is correct.
    fmt.Println(deckRevealedIncreasing([]int{17,13,11,2,3,5,7})) // [2,13,3,11,5,17,7]
    fmt.Println(deckRevealedIncreasing([]int{1,1000})) // [1,1000]

    fmt.Println(deckRevealedIncreasing1([]int{17,13,11,2,3,5,7})) // [2,13,3,11,5,17,7]
    fmt.Println(deckRevealedIncreasing1([]int{1,1000})) // [1,1000]
}