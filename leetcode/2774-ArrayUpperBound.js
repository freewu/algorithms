// 2774. Array Upper Bound
// Write code that enhances all arrays such that you can call the upperBound() method on any array and it will return the last index of a given target number. 
// nums is a sorted ascending array of numbers that may contain duplicates. 
// If the target number is not found in the array, return -1.

// Example 1:
// Input: nums = [3,4,5], target = 5
// Output: 2
// Explanation: Last index of target value is 2

// Example 2:
// Input: nums = [1,4,5], target = 2
// Output: -1
// Explanation: Because there is no digit 2 in the array, return -1.

// Example 3:
// Input: nums = [3,4,6,6,6,6,7], target = 6
// Output: 5
// Explanation: Last index of target value is 5
 
// Constraints:
//         1 <= nums.length <= 10^4
//         -10^4 <= nums[i], target <= 10^4
//         nums is sorted in ascending order.
 
// Follow up: Can you write an algorithm with O(log n) runtime complexity?

/** 
 * @param {number} target
 * @return {number}
 */
Array.prototype.upperBound = function(target) {
    return this.lastIndexOf(target);
}

Array.prototype.upperBound1 = function(target) {
    for(let i = this.length - 1; i >= 0; i--) {
        if(target === this[i]) return i
    }
    return -1;
}

Array.prototype.upperBound2 = function(target) {
    let L = 0, R = this.length - 1;
    while(L <= R) {
      const mid = (R + L) >> 1;
      if(this[mid] <= target) {
        L = mid + 1;
      }else{
        R = mid - 1;
      }
    }
    return this[L - 1] === target ? L - 1 : -1;
}


console.log([3,4,5].upperBound(5)); // 2
console.log([1,4,5].upperBound(2)); // -1
console.log([3,4,6,6,6,6,7].upperBound(6)); // 5

console.log([3,4,5].upperBound1(5)); // 2
console.log([1,4,5].upperBound1(2)); // -1
console.log([3,4,6,6,6,6,7].upperBound1(6)); // 5

console.log([3,4,5].upperBound2(5)); // 2
console.log([1,4,5].upperBound2(2)); // -1
console.log([3,4,6,6,6,6,7].upperBound2(6)); // 5