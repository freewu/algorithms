package main

// 3387. Maximize Amount After Two Days of Conversions
// You are given a string initialCurrency, and you start with 1.0 of initialCurrency.

// You are also given four arrays with currency pairs (strings) and rates (real numbers):
//     1. pairs1[i] = [startCurrencyi, targetCurrencyi] denotes that you can convert from startCurrencyi to targetCurrencyi at a rate of rates1[i] on day 1.
//     2. pairs2[i] = [startCurrencyi, targetCurrencyi] denotes that you can convert from startCurrencyi to targetCurrencyi at a rate of rates2[i] on day 2.
//     3. Also, each targetCurrency can be converted back to its corresponding startCurrency at a rate of 1 / rate.

// You can perform any number of conversions, including zero, using rates1 on day 1, followed by any number of additional conversions, including zero, using rates2 on day 2.

// Return the maximum amount of initialCurrency you can have after performing any number of conversions on both days in order.

// Note: Conversion rates are valid, and there will be no contradictions in the rates for either day. 
// The rates for the days are independent of each other.

// Example 1:
// Input: initialCurrency = "EUR", pairs1 = [["EUR","USD"],["USD","JPY"]], rates1 = [2.0,3.0], pairs2 = [["JPY","USD"],["USD","CHF"],["CHF","EUR"]], rates2 = [4.0,5.0,6.0]
// Output: 720.00000
// Explanation:
// To get the maximum amount of EUR, starting with 1.0 EUR:
// On Day 1:
// Convert EUR to USD to get 2.0 USD.
// Convert USD to JPY to get 6.0 JPY.
// On Day 2:
// Convert JPY to USD to get 24.0 USD.
// Convert USD to CHF to get 120.0 CHF.
// Finally, convert CHF to EUR to get 720.0 EUR.

// Example 2:
// Input: initialCurrency = "NGN", pairs1 = [["NGN","EUR"]], rates1 = [9.0], pairs2 = [["NGN","EUR"]], rates2 = [6.0]
// Output: 1.50000
// Explanation:
// Converting NGN to EUR on day 1 and EUR to NGN using the inverse rate on day 2 gives the maximum amount.

// Example 3:
// Input: initialCurrency = "USD", pairs1 = [["USD","EUR"]], rates1 = [1.0], pairs2 = [["EUR","JPY"]], rates2 = [10.0]
// Output: 1.00000
// Explanation:
// In this example, there is no need to make any conversions on either day.

// Constraints:
//     1 <= initialCurrency.length <= 3
//     initialCurrency consists only of uppercase English letters.
//     1 <= n == pairs1.length <= 10
//     1 <= m == pairs2.length <= 10
//     pairs1[i] == [startCurrencyi, targetCurrencyi]
//     pairs2[i] == [startCurrencyi, targetCurrencyi]
//     1 <= startCurrencyi.length, targetCurrencyi.length <= 3
//     startCurrencyi and targetCurrencyi consist only of uppercase English letters.
//     rates1.length == n
//     rates2.length == m
//     1.0 <= rates1[i], rates2[i] <= 10.0
//     The input is generated such that there are no contradictions or cycles in the conversion graphs for either day.
//     The input is generated such that the output is at most 5 * 10^10.

import "fmt"

func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) float64 {
    type Conversion struct { 
        curr string
        rate float64 
    }
    genMap := func(pairs [][]string, rates []float64) map[string][]Conversion {
        res := make(map[string][]Conversion, len(pairs) * 2)
        for i := range pairs {
            res[pairs[i][0]] = append(res[pairs[i][0]], Conversion{pairs[i][1], rates[i]})
            res[pairs[i][1]] = append(res[pairs[i][1]], Conversion{pairs[i][0], 1 / rates[i]})
        }
        return res
    }
    mp1, mp2 := genMap(pairs1, rates1), genMap(pairs2, rates2)
    res, queue, visited := 1.0, []Conversion{}, make(map[string]bool)
    max := func (x, y float64) float64 { if x > y { return x; }; return y; }
    var dfs func(curr Conversion, mp map[string][]Conversion, push bool)
    dfs = func(curr Conversion, mp map[string][]Conversion, push bool) {
        if curr.curr == initialCurrency {
            res = max(res, curr.rate)
        }
        if visited[curr.curr] { return }
        visited[curr.curr] = true
        if push {
            queue = append(queue, curr)
        }
        for _, next := range mp[curr.curr] {
            next.rate *= curr.rate
            dfs(next, mp, push)
        }
        visited[curr.curr] = false
    }
    dfs(Conversion{ initialCurrency, 1.0 }, mp1, true)
    for _, v := range queue {
        dfs(v, mp2, false)
    }
    return res
}

func main() {
    // Example 1:
    // Input: initialCurrency = "EUR", pairs1 = [["EUR","USD"],["USD","JPY"]], rates1 = [2.0,3.0], pairs2 = [["JPY","USD"],["USD","CHF"],["CHF","EUR"]], rates2 = [4.0,5.0,6.0]
    // Output: 720.00000
    // Explanation:
    // To get the maximum amount of EUR, starting with 1.0 EUR:
    // On Day 1:
    // Convert EUR to USD to get 2.0 USD.
    // Convert USD to JPY to get 6.0 JPY.
    // On Day 2:
    // Convert JPY to USD to get 24.0 USD.
    // Convert USD to CHF to get 120.0 CHF.
    // Finally, convert CHF to EUR to get 720.0 EUR.
    fmt.Println(maxAmount("EUR", [][]string{{"EUR","USD"}, {"USD","JPY"}}, []float64{2.0,3.0}, [][]string{{"JPY","USD"}, {"USD","CHF"}, {"CHF","EUR"}}, []float64{4.0,5.0,6.0})) // 720.00000
    // Example 2:
    // Input: initialCurrency = "NGN", pairs1 = [["NGN","EUR"]], rates1 = [9.0], pairs2 = [["NGN","EUR"]], rates2 = [6.0]
    // Output: 1.50000
    // Explanation:
    // Converting NGN to EUR on day 1 and EUR to NGN using the inverse rate on day 2 gives the maximum amount.
    fmt.Println(maxAmount("NGN", [][]string{{"NGN","EUR"}}, []float64{9.0}, [][]string{{"NGN","EUR"}}, []float64{6.0})) // 1.50000
    // Example 3:
    // Input: initialCurrency = "USD", pairs1 = [["USD","EUR"]], rates1 = [1.0], pairs2 = [["EUR","JPY"]], rates2 = [10.0]
    // Output: 1.00000
    // Explanation:
    // In this example, there is no need to make any conversions on either day.
    fmt.Println(maxAmount("USD", [][]string{{"USD","EUR"}}, []float64{1.0}, [][]string{{"EUR","JPY"}}, []float64{10.0})) // 1.00000
}