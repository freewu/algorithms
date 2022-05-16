# 2628 · Find eligible files in a directory (ii)
# Description
# Use the find command as required to find the files that match the requirements in the current folder, 
# completing the following 3 steps.

# Step 1 search for all files in the current directory and its subdirectories with the suffix .c and list them
find -name "*.c"

# Step 2 search for all files in the current directory and its subdirectories that start with A and list them
find -name "A*"

# Step 3 search for all empty files in the current directory and its subdirectories and list them
find -size 0