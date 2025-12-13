package main

// 3606. Coupon Code Validator
// You are given three arrays of length n that describe the properties of n coupons: code, businessLine, and isActive. 
// The ith coupon has:
//     1. code[i]: a string representing the coupon identifier.
//     2. businessLine[i]: a string denoting the business category of the coupon.
//     3. isActive[i]: a boolean indicating whether the coupon is currently active.

// A coupon is considered valid if all of the following conditions hold:
//     1. code[i] is non-empty and consists only of alphanumeric characters (a-z, A-Z, 0-9) and underscores (_).
//     2. businessLine[i] is one of the following four categories: "electronics", "grocery", "pharmacy", "restaurant".
//     3. isActive[i] is true.

// Return an array of the codes of all valid coupons, sorted first by their businessLine in the order: "electronics", "grocery", "pharmacy", "restaurant", and then by code in lexicographical (ascending) order within each category.

// Example 1:
// Input: code = ["SAVE20","","PHARMA5","SAVE@20"], businessLine = ["restaurant","grocery","pharmacy","restaurant"], isActive = [true,true,true,true]
// Output: ["PHARMA5","SAVE20"]
// Explanation:
// First coupon is valid.
// Second coupon has empty code (invalid).
// Third coupon is valid.
// Fourth coupon has special character @ (invalid).

// Example 2:
// Input: code = ["GROCERY15","ELECTRONICS_50","DISCOUNT10"], businessLine = ["grocery","electronics","invalid"], isActive = [false,true,true]
// Output: ["ELECTRONICS_50"]
// Explanation:
// First coupon is inactive (invalid).
// Second coupon is valid.
// Third coupon has invalid business line (invalid).

// Constraints:
//     n == code.length == businessLine.length == isActive.length
//     1 <= n <= 100
//     0 <= code[i].length, businessLine[i].length <= 100
//     code[i] and businessLine[i] consist of printable ASCII characters.
//     isActive[i] is either true or false.

import "fmt"
import "slices"
import "unicode"
import "sort"

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
    res, groups := []string{}, [4][]string{}
    categorys := map[string]int{
        "electronics": 0,
        "grocery":     1,
        "pharmacy":    2,
        "restaurant":  3,
    }
    isValid := func(s string) bool {
        for _, c := range s {
            if c != '_' && !unicode.IsLetter(c) && !unicode.IsDigit(c) {
                return false
            }
        }
        return true
    }
    for i, s := range code {
        cate, ok := categorys[businessLine[i]]
        if s != "" && ok && isActive[i] && isValid(s) {
            groups[cate] = append(groups[cate], s) // 相同类别的优惠码分到同一组
        }
    }
    for _, g := range groups {
        slices.Sort(g) // 每一组内部排序
        res = append(res, g...)
    }
    return res
}

func validateCoupons1(code []string, businessLine []string, isActive []bool) []string {
    res := make([]string, 0, len(code))
    business := []string{"electronics", "grocery", "pharmacy", "restaurant"}
    l, r := 0, 0
    isValidCode := func(code string) bool {
        if len(code) == 0 { return false }
        for _, r := range code {
            if !unicode.IsNumber(r) && !unicode.IsLetter(r) && r != '_' {
                return false
            }
        }
        return true
    }
    for _, s := range business {
        for i := 0; i < len(code); i++ {
            if isActive[i] && businessLine[i] == s && isValidCode(code[i]) {
                res = append(res, code[i])
                r++
            }
        }
        sort.Strings(res[l:r])
        l = r
    }
    return res
}

func main() {
    // Example 1:
    // Input: code = ["SAVE20","","PHARMA5","SAVE@20"], businessLine = ["restaurant","grocery","pharmacy","restaurant"], isActive = [true,true,true,true]
    // Output: ["PHARMA5","SAVE20"]
    // Explanation:
    // First coupon is valid.
    // Second coupon has empty code (invalid).
    // Third coupon is valid.
    // Fourth coupon has special character @ (invalid).
    fmt.Println(validateCoupons([]string{"SAVE20","","PHARMA5","SAVE@20"}, []string{"restaurant","grocery","pharmacy","restaurant"}, []bool{true,true,true,true})) // ["PHARMA5","SAVE20"]
    // Example 2:
    // Input: code = ["GROCERY15","ELECTRONICS_50","DISCOUNT10"], businessLine = ["grocery","electronics","invalid"], isActive = [false,true,true]
    // Output: ["ELECTRONICS_50"]
    // Explanation:
    // First coupon is inactive (invalid).
    // Second coupon is valid.
    // Third coupon has invalid business line (invalid).
    fmt.Println(validateCoupons([]string{"GROCERY15","ELECTRONICS_50","DISCOUNT10"}, []string{"grocery","electronics","invalid"}, []bool{false,true,true})) //  ["ELECTRONICS_50"]

    fmt.Println(validateCoupons1([]string{"SAVE20","","PHARMA5","SAVE@20"}, []string{"restaurant","grocery","pharmacy","restaurant"}, []bool{true,true,true,true})) // ["PHARMA5","SAVE20"]
    fmt.Println(validateCoupons1([]string{"GROCERY15","ELECTRONICS_50","DISCOUNT10"}, []string{"grocery","electronics","invalid"}, []bool{false,true,true})) //  ["ELECTRONICS_50"]
}