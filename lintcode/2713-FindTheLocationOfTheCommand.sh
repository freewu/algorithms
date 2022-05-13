# 2713 · Find the location of the command
# Description
# Use the which command to search for the location of a system command in the path specified by the PATH variable 
# and return the first search result, complete the following 3 steps in order.

# Step 1 use the echo command to check the value of the PATH variable
echo $PATH

# Step 2 use the which command to find the location of the command echo
which  echo

# Step 3 use the which command to find the location of the command du
which  du