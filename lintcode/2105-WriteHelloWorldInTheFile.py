# 2105 · Write Hello World! in the file
# Description
# Please write the relevant Python code in the file write_hello_world.py, 
# use the open() function to read the file content from the given file path filepath, and write'Hello World!' to the file.

# Example
# The evaluation opportunity executes your code by executing the command python write_hello_world.py {path} 
# and passing the path as a command line parameter. You can learn how the code runs in write_hello_world.py.

# Example 1
# When the input file path is:
#   /output/1.txt
# The output data is:
#   Hello World!

# Example 2
# When the input file path is:
#   /output/2.txt
# The output data is:
#   Hello World!

import sys

def write_to_file(filepath):
    # write your code here
    # you code should write "Hello World!" to the give filepath
    # do not change the function name
    with open(filepath,"w") as f:
        f.write("Hello World!")

write_to_file(sys.argv[1])
with open(sys.argv[1], 'r') as f:
    print(f.read())