import sys
from solution import remove_empty_lines

input_path, output_path = sys.argv[1], sys.argv[2]

lines = remove_empty_lines(input_path,output_path)
