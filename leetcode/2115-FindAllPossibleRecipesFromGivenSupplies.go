package main

// 2115. Find All Possible Recipes from Given Supplies
// You have information about n different recipes. 
// You are given a string array recipes and a 2D string array ingredients. 
// The ith recipe has the name recipes[i], and you can create it if you have all the needed ingredients from ingredients[i]. 
// Ingredients to a recipe may need to be created from other recipes, 
// i.e., ingredients[i] may contain a string that is in recipes.

// You are also given a string array supplies containing all the ingredients that you initially have, 
// and you have an infinite supply of all of them.

// Return a list of all the recipes that you can create. 
// You may return the answer in any order.

// Note that two recipes may contain each other in their ingredients.

// Example 1:
// Input: recipes = ["bread"], ingredients = [["yeast","flour"]], supplies = ["yeast","flour","corn"]
// Output: ["bread"]
// Explanation:
// We can create "bread" since we have the ingredients "yeast" and "flour".

// Example 2:
// Input: recipes = ["bread","sandwich"], ingredients = [["yeast","flour"],["bread","meat"]], supplies = ["yeast","flour","meat"]
// Output: ["bread","sandwich"]
// Explanation:
// We can create "bread" since we have the ingredients "yeast" and "flour".
// We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".

// Example 3:
// Input: recipes = ["bread","sandwich","burger"], ingredients = [["yeast","flour"],["bread","meat"],["sandwich","meat","bread"]], supplies = ["yeast","flour","meat"]
// Output: ["bread","sandwich","burger"]
// Explanation:
// We can create "bread" since we have the ingredients "yeast" and "flour".
// We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
// We can create "burger" since we have the ingredient "meat" and can create the ingredients "bread" and "sandwich".

// Constraints:
//     n == recipes.length == ingredients.length
//     1 <= n <= 100
//     1 <= ingredients[i].length, supplies.length <= 100
//     1 <= recipes[i].length, ingredients[i][j].length, supplies[k].length <= 10
//     recipes[i], ingredients[i][j], and supplies[k] consist only of lowercase English letters.
//     All the values of recipes and supplies combined are unique.
//     Each ingredients[i] does not contain any duplicate values.

import "fmt"

// BFS Topo Sort 
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
    adjList, ref := make(map[string][]string), make(map[string]int)
    for i, greets := range ingredients {
        for _, g := range greets {
            adjList[g] = append(adjList[g], recipes[i])
            ref[recipes[i]]++
        }        
    }
    res := []string{}
    for len(supplies) > 0 {
        cur := supplies[0]
        supplies = supplies[1:]
        for _, n := range adjList[cur] {
            ref[n]--
            if ref[n] == 0 {
                res = append(res, n)
                supplies = append(supplies, n)
            }
        }
    }
    return res
}

func findAllRecipes1(recipes []string, ingredients [][]string, supplies []string) []string {
    supplySet := make(map[string]bool)
    for _, supply := range supplies {
        supplySet[supply] = true
    }
    graph := make(map[string][]string)  // 菜指向依赖这道菜的菜
    indegree := make(map[string]int)    // 记录每道菜未准备好的原材料数量
    for i, recipe := range recipes {
        for _, ingredient := range ingredients[i] {
            if !supplySet[ingredient] {  // 如果原材料不在supplies里，它可能是其他菜
                graph[ingredient] = append(graph[ingredient], recipe)
                indegree[recipe]++
            }
        }
    }
    queue := make([]string, 0)
    for _, recipe := range recipes { // 初始可制作的菜
        if indegree[recipe] == 0 {
            queue = append(queue, recipe)
        }
    }
    res := make([]string, 0)
    for len(queue) > 0 { // 拓扑排序处理队列
        recipe := queue[0]
        queue = queue[1:]
        res = append(res, recipe)
        for _, next := range graph[recipe] { // 对于当前菜，将其作为原材料的菜的入度减一
            indegree[next]--
            if indegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }
    return res
}

func findAllRecipes2(recipes []string, ingredients [][]string, supplies []string) []string {
    // Step 1: Initialize the supplies set
    available := make(map[string]bool)
    for _, supply := range supplies {
        available[supply] = true
    }
    // Step 2: Build the graph and in-degree map
    graph, inDegree := make(map[string][]string), make(map[string]int)
    for i, recipe := range recipes {
        for _, ingredient := range ingredients[i] {
            if !available[ingredient] {
                graph[ingredient] = append(graph[ingredient], recipe)
                inDegree[recipe]++
            }
        }
    }
    // Step 3: Initialize the queue with recipes that have an in-degree of 0 (can be made with initial supplies)
    queue := []string{}
    for _, recipe := range recipes {
        if inDegree[recipe] == 0 {
            queue = append(queue, recipe)
        }
    }
    // Step 4: Process each recipe in the queue
    res := []string{}
    for len(queue) > 0 {
        recipe := queue[0]
        queue = queue[1:]
        res = append(res, recipe)
        available[recipe] = true // Mark this recipe as available
        for _, dependentRecipe := range graph[recipe] { // Update in-degrees for recipes that depend on this one
            inDegree[dependentRecipe]--
            if inDegree[dependentRecipe] == 0 {
                queue = append(queue, dependentRecipe)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: recipes = ["bread"], ingredients = [["yeast","flour"]], supplies = ["yeast","flour","corn"]
    // Output: ["bread"]
    // Explanation:
    // We can create "bread" since we have the ingredients "yeast" and "flour".
    fmt.Println(findAllRecipes([]string{"bread"}, [][]string{{"yeast","flour"}}, []string{"yeast","flour","corn"})) // ["bread"]
    // Example 2:
    // Input: recipes = ["bread","sandwich"], ingredients = [["yeast","flour"],["bread","meat"]], supplies = ["yeast","flour","meat"]
    // Output: ["bread","sandwich"]
    // Explanation:
    // We can create "bread" since we have the ingredients "yeast" and "flour".
    // We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
    fmt.Println(findAllRecipes([]string{"bread","sandwich"}, [][]string{{"yeast","flour"},{"bread","meat"}}, []string{"yeast","flour","meat"})) // ["bread","sandwich"]
    // Example 3:
    // Input: recipes = ["bread","sandwich","burger"], ingredients = [["yeast","flour"],["bread","meat"],["sandwich","meat","bread"]], supplies = ["yeast","flour","meat"]
    // Output: ["bread","sandwich","burger"]
    // Explanation:
    // We can create "bread" since we have the ingredients "yeast" and "flour".
    // We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
    // We can create "burger" since we have the ingredient "meat" and can create the ingredients "bread" and "sandwich".
    fmt.Println(findAllRecipes([]string{"bread","sandwich","burger"}, [][]string{{"yeast","flour"},{"bread","meat"},{"sandwich","meat","bread"}}, []string{"yeast","flour","meat"})) // ["bread","sandwich","burger"]
    
    fmt.Println(findAllRecipes1([]string{"bread"}, [][]string{{"yeast","flour"}}, []string{"yeast","flour","corn"})) // ["bread"]
    fmt.Println(findAllRecipes1([]string{"bread","sandwich"}, [][]string{{"yeast","flour"},{"bread","meat"}}, []string{"yeast","flour","meat"})) // ["bread","sandwich"]
    fmt.Println(findAllRecipes1([]string{"bread","sandwich","burger"}, [][]string{{"yeast","flour"},{"bread","meat"},{"sandwich","meat","bread"}}, []string{"yeast","flour","meat"})) // ["bread","sandwich","burger"]
    
    fmt.Println(findAllRecipes2([]string{"bread"}, [][]string{{"yeast","flour"}}, []string{"yeast","flour","corn"})) // ["bread"]
    fmt.Println(findAllRecipes2([]string{"bread","sandwich"}, [][]string{{"yeast","flour"},{"bread","meat"}}, []string{"yeast","flour","meat"})) // ["bread","sandwich"]
    fmt.Println(findAllRecipes2([]string{"bread","sandwich","burger"}, [][]string{{"yeast","flour"},{"bread","meat"},{"sandwich","meat","bread"}}, []string{"yeast","flour","meat"})) // ["bread","sandwich","burger"]
}