# 2730 · Delete file contents
# In this topic, we will practice using the sed command to delete the contents of a file 
# by completing the following 4 steps in order.

# Step 1 using the nl command, view the contents of the file file and display the line numbers
nl file

# Step 2 delete line 88 of the file file and output the result to standard output
sed 88d file

# Step 3 delete lines 88 - 99 of the file file, displaying the line numbers and outputting the result to standard output
nl file | sed 88,99d