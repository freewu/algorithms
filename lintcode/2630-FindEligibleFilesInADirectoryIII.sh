# 2630 · Find eligible files in a directory (iii)
# Description
# Using the find command as required, complete the following three steps to find files in the current folder that match the requirements.

# Step 1 ignore the case of the file name and search for all files in the current directory 
# and its subdirectories that start with apple and list them
find -iname "apple*"

# Step 2 search for all files in the current directory and its subdirectories that are not empty and list them
find -type f ! -size 0
find -size +0 -type f


# Step 3 in the current directory, look up all files in the current directory's subdirectory sub_dir 
# that are of length 0 and write them to the file file (one file name per line)
find sub_dir -size 0 -type f > file
