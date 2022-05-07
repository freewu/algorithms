# 2701 · Remove documents
# Description
# If we have added a file to the tracked list and later find that it is no longer needed, 
# we can use the git rm command to remove it from the tracked list (from the staging area, to be precise) 
# and the specified file will be removed from the working directory 
# and then committed so that it will not appear in the untracked list in the future, or you can choose to leave the file in the working directory.
# Please complete the following 7 steps in order.

# Step 1 go to the directory /nobody/my-repo and use git status to check the current status of the repository"
cd /nobody/my-repo 
git status

# Step 2 use the ls command to view and print the names of all the files in the current directory
ls 

# Step 3 use a command that delete the files in the staging area that have been modified
# git rm index.php这个命令把工作区的index.php删除并暂存了
# git rm命令的本质就是rm 和 git add 
# git rm命令本质上就是先执行了rm 文件名，然后执行git add把rm命令提交到暂存了
# git rm ： 同时从工作区和索引中删除文件。即本地的文件也被删除了。
# git rm --cached ： 从索引中删除文件。但是本地文件还存在， 只是不希望这个文件被版本控制。
git rm -f solution.h

# Step 4 use git status to see the current status of the repository
git status

# Step 5 use a command that delete main.cpp from the list of tracked files, but keep it in the working directory
git rm --cached main.cpp

# Step 6 use git status to see the current status of the repository".
git status

# Step 7 use the ls command to see and print the names of all the files in the current directory
ls
