// 2677. Chunk Array
// Given an array arr and a chunk size size, return a chunked array. 
// A chunked array contains the original elements in arr, but consists of subarrays each of length size. 
// The length of the last subarray may be less than size if arr.length is not evenly divisible by size.
// You may assume the array is the output of JSON.parse. In other words, it is valid JSON.
// Please solve it without using lodash's _.chunk function.

// Example 1:
// Input: arr = [1,2,3,4,5], size = 1
// Output: [[1],[2],[3],[4],[5]]
// Explanation: The arr has been split into subarrays each with 1 element.

// Example 2:
// Input: arr = [1,9,6,3,2], size = 3
// Output: [[1,9,6],[3,2]]
// Explanation: The arr has been split into subarrays with 3 elements. However, only two elements are left for the 2nd subarray.

// Example 3:
// Input: arr = [8,5,3,2,6], size = 6
// Output: [[8,5,3,2,6]]
// Explanation: Size is greater than arr.length thus all elements are in the first subarray.

// Example 4:
// Input: arr = [], size = 1
// Output: []
// Explanation: There are no elements to be chunked so an empty array is returned.

// Constraints:
//         arr is a valid JSON array
//         2 <= JSON.stringify(arr).length <= 10^5
//         1 <= size <= arr.length + 1

/**
 * @param {Array} arr
 * @param {number} size
 * @return {Array}
 */
var chunk = function(arr, size) {
    if(arr.length == 0) return [];
    if (size >= arr.length) return [arr];
    let res = [];
    for ( let i = 0; i < arr.length; i = i + size) {
        let temp = [];
        let j = 0;
        while(j < size && i + j < arr.length) {
            temp.push(arr[i + j])
            j++;
        }
        res.push(temp);
    }
    return res;
};

// use arr.slice
var chunk = function (arr, size) {
    let num = Math.ceil(arr.length/size)
    let newArr = []
    for(let i=0;i<num;i++){
        newArr.push(arr.slice(i*size,(i+1)*size))
    }
    return newArr
};