# 2696 · Cache and restore across branches
# Description
# You find that you have edited the temp.md file on the master branch by mistake. 
# Please cache your current changes, create a branch1 branch and restore the changes on this branch. 
# Please follow the five-step exercise below:

# Step 1 View the current status.
cd /nobody/my-repo
git status

# Step 2 Cache the current changes.
git stash

# Step 3 Create a new branch and switch to the branch1 branch.
git checkout -b branch1

# Step 4 Restore the previous modification.
git stash pop

# Step 5 View the current status.
git status