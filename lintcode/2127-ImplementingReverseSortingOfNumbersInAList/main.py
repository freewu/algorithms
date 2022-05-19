import sys
from solution import list_sort

input_path = sys.argv[1]
with open(input_path, 'r', encoding = 'utf-8') as f:
    list_in = eval(f.readline())
print(list_sort(list_in))