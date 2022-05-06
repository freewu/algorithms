# 2540 · Your First Git Commit
# Please use git command to implement the following 3 steps:

# Step 1 Go to the repository directory of my-repo, create a new branch named new-branch, and switch to it.
cd my-repo
git branch new-branch
git checkout new-branch

# Step 2 Create a new file names new-file.txt under the repository my-repo.
touch new-file.txt

# Step 3 Use command git add and git commit to commit your changes to the repository.
git add .
git commit -m'new-file.txt'