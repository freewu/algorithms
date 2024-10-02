package main

// 1912. Design Movie Rental System
// You have a movie renting company consisting of n shops. 
// You want to implement a renting system that supports searching for, booking, and returning movies. 
// The system should also support generating a report of the currently rented movies.

// Each movie is given as a 2D integer array entries where entries[i] = [shopi, moviei, pricei] indicates 
// that there is a copy of movie moviei at shop shopi with a rental price of pricei.
// Each shop carries at most one copy of a movie moviei.

// The system should support the following functions:
//     Search: 
//         Finds the cheapest 5 shops that have an unrented copy of a given movie. 
//         The shops should be sorted by price in ascending order, and in case of a tie, the one with the smaller shopi should appear first. 
//         If there are less than 5 matching shops, then all of them should be returned. 
//         If no shop has an unrented copy, then an empty list should be returned.
//     Rent: 
//         Rents an unrented copy of a given movie from a given shop.
//     Drop: 
//         Drops off a previously rented copy of a given movie at a given shop.
//     Report: 
//         Returns the cheapest 5 rented movies (possibly of the same movie ID) as a 2D list res where res[j] = [shopj, moviej] describes that the jth cheapest rented movie moviej was rented from the shop shopj. 
//         The movies in res should be sorted by price in ascending order, and in case of a tie, the one with the smaller shopj should appear first, 
//         and if there is still tie, the one with the smaller moviej should appear first. 
//         If there are fewer than 5 rented movies, then all of them should be returned. 
//         If no movies are currently being rented, then an empty list should be returned.

// Implement the MovieRentingSystem class:
//     MovieRentingSystem(int n, int[][] entries) 
//         Initializes the MovieRentingSystem object with n shops and the movies in entries.
//     List<Integer> search(int movie) 
//         Returns a list of shops that have an unrented copy of the given movie as described above.
//     void rent(int shop, int movie) 
//         Rents the given movie from the given shop.
//     void drop(int shop, int movie) 
//         Drops off a previously rented movie at the given shop.
//     List<List<Integer>> report() 
//         Returns a list of cheapest rented movies as described above.

// Note: The test cases will be generated such that rent will only be called if the shop has an unrented copy of the movie, and drop will only be called if the shop had previously rented out the movie.


// Example 1:
// Input
// ["MovieRentingSystem", "search", "rent", "rent", "report", "drop", "search"]
// [[3, [[0, 1, 5], [0, 2, 6], [0, 3, 7], [1, 1, 4], [1, 2, 7], [2, 1, 5]]], [1], [0, 1], [1, 2], [], [1, 2], [2]]
// Output
// [null, [1, 0, 2], null, null, [[0, 1], [1, 2]], null, [0, 1]]
// Explanation
// MovieRentingSystem movieRentingSystem = new MovieRentingSystem(3, [[0, 1, 5], [0, 2, 6], [0, 3, 7], [1, 1, 4], [1, 2, 7], [2, 1, 5]]);
// movieRentingSystem.search(1);  // return [1, 0, 2], Movies of ID 1 are unrented at shops 1, 0, and 2. Shop 1 is cheapest; shop 0 and 2 are the same price, so order by shop number.
// movieRentingSystem.rent(0, 1); // Rent movie 1 from shop 0. Unrented movies at shop 0 are now [2,3].
// movieRentingSystem.rent(1, 2); // Rent movie 2 from shop 1. Unrented movies at shop 1 are now [1].
// movieRentingSystem.report();   // return [[0, 1], [1, 2]]. Movie 1 from shop 0 is cheapest, followed by movie 2 from shop 1.
// movieRentingSystem.drop(1, 2); // Drop off movie 2 at shop 1. Unrented movies at shop 1 are now [1,2].
// movieRentingSystem.search(2);  // return [0, 1]. Movies of ID 2 are unrented at shops 0 and 1. Shop 0 is cheapest, followed by shop 1.

// Constraints:
//     1 <= n <= 3 * 10^5
//     1 <= entries.length <= 10^5
//     0 <= shopi < n
//     1 <= moviei, pricei <= 10^4
//     Each shop carries at most one copy of a movie moviei.
//     At most 10^5 calls in total will be made to search, rent, drop and report.

import "fmt"
import "container/heap"

type IndexKey struct {
    Shop  int
    Movie int
}

func (k IndexKey) String() string {
    return fmt.Sprintf("%d_%d", k.Shop, k.Movie)
}

type MovieRentingSystem struct {
    rentedMovieHeap   *Movies
    unrentedMovies    map[int]*Movies
    rentedMovies      map[string]*RentingMovie
    indexingKeyOffset map[string]*RentingMovie
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
    unrentedMovies := make(map[int]*Movies)
    indexingMovies := make(map[string]*RentingMovie)

    for _, entry := range entries {
        if _, ok := unrentedMovies[entry[1]]; !ok {
            unrentedMovies[entry[1]] = &Movies{}
        }
        m := &RentingMovie{
            Shop:  entry[0],
            Movie: entry[1],
            Price: entry[2],
        }
        heap.Push(unrentedMovies[entry[1]], m)
        indexingMovies[IndexKey{
            Shop:  entry[0],
            Movie: entry[1],
        }.String()] = m
    }
    return MovieRentingSystem{
        indexingKeyOffset: indexingMovies,
        unrentedMovies:    unrentedMovies,
        rentedMovieHeap:   &Movies{},
        rentedMovies:      make(map[string]*RentingMovie),
    }
}

func (this *MovieRentingSystem) Search(movie int) []int {
    movies, ok := this.unrentedMovies[movie]
    if !ok {
        return nil
    }
    result := make([]int, 0, 5)
    eligibleMovies := make([]*RentingMovie, 0, 5)
    for movies.Len() > 0 && len(eligibleMovies) < 5 {
        item := heap.Pop(movies)
        movie := item.(*RentingMovie)
        eligibleMovies = append(eligibleMovies, movie)
        result = append(result, movie.Shop)

    }
    if len(eligibleMovies) == 0 {
        return nil
    }
    for _, e := range eligibleMovies {
        heap.Push(movies, e)
    }
    return result
}

func (this *MovieRentingSystem) Rent(shop int, movie int) {
    key := IndexKey{
        Shop:  shop,
        Movie: movie,
    }
    rentingMovie := this.indexingKeyOffset[key.String()]
    if rentingMovie == nil {
        return
    }
    heap.Remove(this.unrentedMovies[movie], rentingMovie.Offset)
    //remove from heap min
    // add into rentedMovie List
    rentedMovie := &RentingMovie{
        Shop:     shop,
        Movie:    movie,
        Price:    rentingMovie.Price,
        IsRented: true,
    }
    this.rentedMovies[key.String()] = rentedMovie
    heap.Push(this.rentedMovieHeap, rentedMovie)
}

func (this *MovieRentingSystem) Drop(shop int, movie int) {
    key := IndexKey{
        Shop:  shop,
        Movie: movie,
    }
    m, ok := this.rentedMovies[key.String()]
    if !ok {
        return
    }
    rentingMovie := &RentingMovie{
        Shop:  shop,
        Movie: movie,
        Price: m.Price,
    }
    this.indexingKeyOffset[key.String()] = rentingMovie
    heap.Push(this.unrentedMovies[movie], rentingMovie)
    heap.Remove(this.rentedMovieHeap, m.Offset)
    delete(this.rentedMovies, key.String())
}

func (this *MovieRentingSystem) Report() [][]int {
    result := make([][]int, 0, 5)
    historicals := make([]*RentingMovie, 0, 5)
    for this.rentedMovieHeap.Len() > 0 && len(historicals) < 5 {
        item := heap.Pop(this.rentedMovieHeap)
        m := item.(*RentingMovie)
        historicals = append(historicals, m)
        result = append(result, []int{m.Shop, m.Movie})
    }
    if len(historicals) == 0 {
        return nil
    }
    for _, item := range historicals {
        heap.Push(this.rentedMovieHeap, item)
    }
    return result
}

type RentingMovie struct {
    Shop     int
    Movie    int
    Price    int
    Offset   int
    IsRented bool
}

type Movies []*RentingMovie

// Len implements heap.Interface.
func (m Movies) Len() int {
    return len(m)
}

// Less implements heap.Interface.
func (m Movies) Less(i int, j int) bool {
    if !m[i].IsRented {
        if m[i].Price == m[j].Price {
            return m[i].Shop < m[j].Shop
        }
    } else if m[i].Price == m[j].Price {
        if m[i].Shop == m[j].Shop {
            return m[i].Movie < m[j].Movie
        }
        return m[i].Shop < m[j].Shop
    }

    return m[i].Price < m[j].Price
}

// Pop implements heap.Interface.
func (m *Movies) Pop() any {
    item := (*m)[m.Len()-1]
    (*m) = (*m)[:m.Len()-1]
    return item
}

// Push implements heap.Interface.
func (m *Movies) Push(x any) {
    item := x.(*RentingMovie)
    item.Offset = m.Len()
    (*m) = append((*m), item)
}

// Swap implements heap.Interface.
func (m Movies) Swap(i int, j int) {
    m[i].Offset, m[j].Offset = j, i
    m[i], m[j] = m[j], m[i]
}


/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * obj := Constructor(n, entries);
 * param_1 := obj.Search(movie);
 * obj.Rent(shop,movie);
 * obj.Drop(shop,movie);
 * param_4 := obj.Report();
 */

func main() {
    // MovieRentingSystem movieRentingSystem = new MovieRentingSystem(3, [[0, 1, 5], [0, 2, 6], [0, 3, 7], [1, 1, 4], [1, 2, 7], [2, 1, 5]]);
    obj := Constructor(3,[][]int{{0, 1, 5}, {0, 2, 6}, {0, 3, 7}, {1, 1, 4}, {1, 2, 7}, {2, 1, 5}})
    fmt.Println(obj)
    // movieRentingSystem.search(1);  // return [1, 0, 2], Movies of ID 1 are unrented at shops 1, 0, and 2. Shop 1 is cheapest; shop 0 and 2 are the same price, so order by shop number.
    fmt.Println(obj.Search(1)) // [1, 0, 2]
    // movieRentingSystem.rent(0, 1); // Rent movie 1 from shop 0. Unrented movies at shop 0 are now [2,3].
    obj.Rent(0, 1)
    fmt.Println(obj)
    // movieRentingSystem.rent(1, 2); // Rent movie 2 from shop 1. Unrented movies at shop 1 are now [1].
    obj.Rent(1, 2)
    fmt.Println(obj)
    // movieRentingSystem.report();   // return [[0, 1], [1, 2]]. Movie 1 from shop 0 is cheapest, followed by movie 2 from shop 1.
    fmt.Println(obj.Report()) // [[0, 1], [1, 2]]
    // movieRentingSystem.drop(1, 2); // Drop off movie 2 at shop 1. Unrented movies at shop 1 are now [1,2].
    obj.Drop(1, 2)
    fmt.Println(obj)
    // movieRentingSystem.search(2);  // return [0, 1]. Movies of ID 2 are unrented at shops 0 and 1. Shop 0 is cheapest, followed by shop 1.
    fmt.Println(obj.Search(2)) // [0, 1]
}