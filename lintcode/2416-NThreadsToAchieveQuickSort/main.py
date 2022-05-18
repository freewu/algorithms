from quick_sort import QuickSort
from solution import Solution
import json
import sys


input_path = sys.argv[1]
output_path = sys.argv[2]
input_file = open(input_path, 'r')
output_file = open(output_path, 'w')

n = json.loads(input_file.readline())
nums = json.loads(input_file.readline())
input_file.close()

solution = Solution()
sorted_nums = solution.quick_sort_in_threadings(n, nums)

if not QuickSort.has_been_called_in_sub_thread:
    raise Exception('You should call sort_range in a sub thread.')

output_file.write(json.dumps(sorted_nums))
output_file.close()