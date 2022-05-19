import sys
from solution import get_len, get_max, get_min, pop_list

input_path = sys.argv[1]
with open(input_path, 'r', encoding = 'utf-8') as f:
    list_in = eval(f.readline())
print(get_len(list_in))
print(get_max(list_in))
print(get_min(list_in))
print(pop_list(list_in))