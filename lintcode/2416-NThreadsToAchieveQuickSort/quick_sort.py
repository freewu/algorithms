import threading

class QuickSort:
    has_been_called_in_sub_thread = False

    def __init__(self, nums):
        self.n = len(nums)
        self.nums = nums

    def sort(self):
        self.sort_range(0, self.n - 1)

    def sort_range(self, start, end):
        if threading.current_thread() is threading.main_thread():
            raise Exception('You should call sort_range in a sub thread.')
        QuickSort.has_been_called_in_sub_thread = True

        if start >= end:
            return

        left, right = start, end
        pivot = self.nums[(start + end) >> 1]
        while left <= right:
            while left <= right and self.nums[left] < pivot:
                left += 1
            while left <= right and self.nums[right] > pivot:
                right -= 1
            if left <= right:
                self.nums[left], self.nums[right] = self.nums[right], self.nums[left]
                left += 1
                right -= 1

        self.sort_range(start, right)
        self.sort_range(left, end)
