from quick_sort import QuickSort
import heapq
# you can import any package that you need here
# write your code here
from threading import Thread

class Solution:
    def quick_sort_in_threadings(self, n, nums):
        size = (len(nums) - 1) // n + 1

        threads = []
        subarrays = []

        for i in range(n):
            start = i * size
            end = min(start + size - 1, len(nums) - 1)

            subarray = nums[start: end + 1]
            thread = Thread(target=lambda: self.quick_sort(subarray))
            subarrays.append(subarray)

            threads.append(thread)
            thread.start()

        for thread in threads:
            thread.join()

        return self.merge_n_sorted_arrays(subarrays)

    def quick_sort(self, nums):
        instance = QuickSort(nums)
        instance.sort()

    def merge_n_sorted_arrays(self, arrays):
        result = []
        heap = []
        for index, array in enumerate(arrays):
            if len(array) == 0:
                continue
            heapq.heappush(heap, (array[0], index, 0))

        while len(heap):
            val, x, y = heap[0]
            heapq.heappop(heap)
            result.append(val)
            if y + 1 < len(arrays[x]):
                heapq.heappush(heap, (arrays[x][y + 1], x, y + 1))

        return result