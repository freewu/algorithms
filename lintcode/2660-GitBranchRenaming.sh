# 2660 · git Branch Renaming
# For this topic, let's do a simple exercise on the git branch command. 
# Use the git command to complete the renaming of a branch, by following the 3 steps:

# Step 1 go to the directory /nobody/my-repo and see all the current branches
cd /nobody/my-repo
git branch -a

# Step 2 rename the existing branch old_branch to new_branch
git branch -m old_branch new_branch

# Step 3 View all current branches
git branch -a