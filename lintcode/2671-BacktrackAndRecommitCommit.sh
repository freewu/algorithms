# 2671 · Backtrack and recommit commit
# Description
# Now that you have an erroneous commit in git, you are asked to fall back to the previous version, 
# track down the new file solution.h, and recommit by following this 8-step exercise:

# Step 1 go to the /nobody/my-repo directory and use git log to view the git commit log
cd /nobody/my-repo
git log

# Step 2 view and print all file names in the current directory
ls

# Step 3 use the git reset command to fall back to the first commit and reset the contents of the staging area and working directory
# 我们用Git的时候有可能 commit 提交代码后，发现这一次commit的内容是有错误的，那么有两种处理方法：
# 1、修改错误内容，再次 commit一次 
# 2、使用 git reset 命令撤销这一次错误的commit 错误的commit没必要保留下来
# git reset 三种模式
#   1 git reset --hard：重置stage区和工作目录:
#       reset --hard 会在重置 HEAD 和branch的同时，重置stage区和工作目录里的内容。
#       当你在 reset 后面加了 --hard 参数时，你的stage区和工作目录里的内容会被完全重置为和HEAD的新位置相同的内容。换句话说，就是你的没有commit的修改会被全部擦掉
#   2 git reset --soft：保留工作目录，并把重置 HEAD 所带来的新的差异放进暂存区
#       reset --soft 会在重置 HEAD 和 branch 时，保留工作目录和暂存区中的内容，并把重置 HEAD 所带来的新的差异放进暂存区。
#   3 git reset 不加参数(mixed)：保留工作目录，并清空暂存区
#       reset 如果不加参数，那么默认使用 --mixed 参数。
#       它的行为是：保留工作目录，并且清空暂存区。
#       也就是说，工作目录的修改、暂存区的内容以及由 reset 所导致的新的文件差异，都会被放进工作目录。
#       就是「把所有差异都混合（mixed）放在工作目录中」
git reset --hard HEAD^

# Step 4 view and print all the filenames in the directory
ls

# Step 5 create a new file solution.h
touch solution.h

# Step 6 use git add to track the new file solution.h
git add solution.h

# Step 7 commit the update using git commit, you should remark 'This is a correct commit.'
git commit -m'This is a correct commit.'

# Step 8 use git log to view the git commit log
git log