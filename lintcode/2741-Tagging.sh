# 2741 · Tagging
# Description
# git can tag a commit in the repository history to indicate its importance. 
# In this topic, please use the git tag command to tag the release node by completing the following steps in order.

# Step 1 go to the /nobody/my-repo folder and look at the history of commits
cd /nobody/my-repo
git log

# Step 2 use the git tag command to see what tags are currently in git.
git tag

# Step 3 tag the last commit with a comment v1.0.
git tag v1.0

# Step 4 view the commits corresponding to tag v1.0.
git show v1.0

# Step 5 update the current commit with the comment This is version 1.1.
git commit --amend -m'This is version 1.1'

# Step 6 tag the changes you just made v1.1
git tag v1.1

# Step 7 check the commit for the tag v1.1.
git show v1.1

# 查看本地分支标签
git tag
git tag -l
git tag --list

# 查看远程所有标签
git ls-remote --tags
git ls-remote --tag

# 给当前分支打标签
git tag <tag-name>
git tag v1.1.0

# 给特定的某个 commit 版本打标签，比如现在某次提交的id为 039bf8b 使用 -m 添加注释
git tag v1.0.0 039bf8b
git tag v1.0.0 -m "add tags information" 039bf8b
git tag v1.0.0 039bf8b -m "add tags information"

# 删除本地某个标签
git tag --delete v1.0.0
git tag -d v1.0.0
git tag --d v1.0.0

# 删除远程的某个标签
git push -d origin v1.0.0
git push --delete origin v1.0.0
git push origin -d v1.0.0
git push origin --delete v1.0.0
git push origin :v1.0.0

# 将本地标签一次性推送到远程
git push origin --tags
git push origin --tag
git push --tags
git push --tag

# 将本地某个特定标签推送到远程
git push origin v1.0.0

# 查看某一个标签的提交信息
git show v1.0.0