/**
2416 · N threads to achieve quick sort
Please write the quick_sort_in_threadings method of the Solution class in the solution.py file. 
This method passes in an integer n and an integer array nums.
You need to divide nums into n sub-arrays, and create one for each sub-array The child thread calls the QuickSort class to sort,
and finally merges the n sorted sub-arrays into an ordered array and returns (you can call the merge_n_sorted_arrays method we wrote for you).
The functions of specific functions are shown in the table above. You can read the code in main.py to understand how your code is evaluated.


    1 ≤ nums ≤ 10 ^5
    0 ≤ nums[i] ≤ 10^6
    1≤ n≤ 10, nums 

Example
We will get the sorted array and print it by calling the quick_sort_in_threadings method under the solution.py file.

When n=5, nums=[10, 47, 36, 47, 22], your code should output:
[10, 22, 36, 47, 47]

When n=1, nums=[1], your code should output:
[1]
 */

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Solution {

    public int[] quickSortInThreadings(int n, int[] arr) throws Exception {
        // write your code here
        return null;
    }

    // Put the arrays to be merged into the list
    public int[] mergeNSortedArrays(List<Object> list) throws Exception {
        List<Integer> ls = new ArrayList<>();
        for (Object oj :list) {
            if(oj == null) continue;
            int arr[] = (int[]) oj;
            for (int num : arr){
                ls.add(num);
            }
        }
        int[] arr = ls.stream().mapToInt(Integer::valueOf).toArray();
        Arrays.sort(arr);
        return arr;
    }
}