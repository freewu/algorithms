package main

// 1603. Design Parking System
// Design a parking system for a parking lot. The parking lot has three kinds of parking spaces: big, medium, and small, with a fixed number of slots for each size.
// Implement the ParkingSystem class:
//     ParkingSystem(int big, int medium, int small) Initializes object of the ParkingSystem class. 
//          The number of slots for each parking space are given as part of the constructor.
//     bool addCar(int carType) Checks whether there is a parking space of carType for the car that wants to get into the parking lot. 
//          carType can be of three kinds: big, medium, or small, which are represented by 1, 2, and 3 respectively. 
//          A car can only park in a parking space of its carType. 
//          If there is no space available, return false, else park the car in that size space and return true.
    
// Example 1:
// Input
// ["ParkingSystem", "addCar", "addCar", "addCar", "addCar"]
// [[1, 1, 0], [1], [2], [3], [1]]
// Output
// [null, true, true, false, false]
// Explanation
// ParkingSystem parkingSystem = new ParkingSystem(1, 1, 0);
// parkingSystem.addCar(1); // return true because there is 1 available slot for a big car
// parkingSystem.addCar(2); // return true because there is 1 available slot for a medium car
// parkingSystem.addCar(3); // return false because there is no available slot for a small car
// parkingSystem.addCar(1); // return false because there is no available slot for a big car. It is already occupied.
 
// Constraints:
//     0 <= big, medium, small <= 1000
//     carType is 1, 2, or 3
//     At most 1000 calls will be made to addCar

import "fmt"

type ParkingSystem struct {
    bigLimit int
    mediumLimit int
    smallLimit int

    big int
    medium int
    small int
}

func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{ big, medium, small, 0, 0, 0 }
}

func (this *ParkingSystem) AddCar(carType int) bool {
    switch carType {
    case 1: // big
        if this.big < this.bigLimit {
            this.big++
            return true
        }
    case 2: // medium
        if this.medium < this.mediumLimit {
            this.medium++
            return true
        }
    case 3: // small
        if this.small < this.smallLimit {
            this.small++
            return true
        }
    }
    return false
}


type ParkingSystem1 struct {
    arr [4]int 
}

func Constructor1(big int, medium int, small int) ParkingSystem1 {
    return ParkingSystem1{[4]int {0,big,medium,small}}
}

func (this *ParkingSystem1) AddCar(carType int) bool {
    if this.arr[carType] > 0 {
        this.arr[carType]--
        return true
    }
    return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

func main() {
    // ParkingSystem parkingSystem = new ParkingSystem(1, 1, 0);
    // parkingSystem.addCar(1); // return true because there is 1 available slot for a big car
    // parkingSystem.addCar(2); // return true because there is 1 available slot for a medium car
    // parkingSystem.addCar(3); // return false because there is no available slot for a small car
    // parkingSystem.addCar(1); // return false because there is no available slot for a big car. It is already occupied.
    obj := Constructor(1,1,0)
    fmt.Println(obj.AddCar(1)) // true
    fmt.Println(obj.AddCar(2)) // true
    fmt.Println(obj.AddCar(3)) // false
    fmt.Println(obj.AddCar(1)) // false

    obj1 := Constructor1(1,1,0)
    fmt.Println(obj1.AddCar(1)) // true
    fmt.Println(obj1.AddCar(2)) // true
    fmt.Println(obj1.AddCar(3)) // false
    fmt.Println(obj1.AddCar(1)) // false
}