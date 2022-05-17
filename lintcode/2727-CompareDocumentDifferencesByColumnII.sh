# 2727 · Compare document differences by column (ii)
# Description
# The comm command compares two sorted files. 
# This command compares the differences between the two sorted files, column by column, 
# and displays the results in three columns if no parameters are specified: 
# column 1 is only the columns that appear in the first file, 
# column 2 is only the columns that appear in the second file, and
# column 3 is the columns that appear in both the first and second files. 
# Column 3 is the column that appears in both file 1 and file 2. 
# You are now asked to compare the differences between files file1 and file2 using the parameters of this command, 
# following the 3-step procedure as follows.

# Step 1 show the columns that appear only in the first file. 2.
comm -23 file1 file2

# Step 2 do not display the columns that only appear in the second file. 3.
comm -2 file1 file2

# Step 3 show the columns that appear in both file 1 and file 2.
comm -12 file1 file2
