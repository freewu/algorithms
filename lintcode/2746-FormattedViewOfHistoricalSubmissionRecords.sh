# 2746 · Formatted view of historical submission records
# Description

# Go to the /nobody/my-repo directory and use the git log command to view the history of commits in the specified format.
cd /nobody/my-repo
git log

# Step 1 Export the commit history in the following format.
# the complete hash of the tree - the author's name : the author's email address
git log --pretty=format:"%T —— %an : %ae"

# Step 2 Please output the history of commits in the following format.
# Complete hash of the tree -- Submitter's name : Submitter's email address, commit description
git log --pretty=format:"%T —— %cn : %ce, %s"

# Step 3 Please output the history of commits in the following format.
# abbreviated hash of tree - name of committer, marks: commit description
git log --pretty=format:"%t —— %cn, marks: %s"

# https://git-scm.com/docs/pretty-formats
# git log --pretty=format:" "
# 控制显示的记录格式，常用的格式占位符写法及其代表的意义:
# '%H': commit hash
# '%h': abbreviated commit hash
# '%t': abbreviated tree hash
# '%P': parent hashes
# '%p': abbreviated parent hashes
# '%an': author name
# '%aN': author name (respecting .mailmap, see git-shortlog or git-blame)
# '%ae': author email
# '%aE': author email (respecting .mailmap, see git-shortlog or git-blame)
# '%ad': author date (format respects --date= option)
# '%aD': author date, RFC2822 style
# '%ar': author date, relative
# '%at': author date, UNIX timestamp
# '%ai': author date, ISO 8601-like format
# '%aI': author date, strict ISO 8601 format
# '%cn': committer name
# '%cN': committer name (respecting .mailmap, see git-shortlog or git-blame)
# '%ce': committer email
# '%cE': committer email (respecting .mailmap, see git-shortlog or git-blame)
# '%cd': committer date (format respects --date= option)
# '%cD': committer date, RFC2822 style
# '%cr': committer date, relative
# '%ct': committer date, UNIX timestamp
# '%ci': committer date, ISO 8601-like format
# '%cI': committer date, strict ISO 8601 format
# '%d': ref names, like the --decorate option of git-log
# '%D': ref names without the " (", ")" wrapping.
# '%e': encoding
# '%s': subject
# '%f': sanitized subject line, suitable for a filename
# '%b': body
# '%B': raw body (unwrapped subject and body)
# '%N': commit notes
# '%T': tree hash
# '%t': abbreviated tree hash
# '%P': parent hashes
# '%p': abbreviated parent hashes
# '%an': author name
# '%aN': author name (respecting .mailmap, see git-shortlog or git-blame)
# '%ae': author email
# '%aE': author email (respecting .mailmap, see git-shortlog or git-blame)
# '%ad': author date (format respects --date= option)
# '%aD': author date, RFC2822 style
# '%ar': author date, relative
# '%at': author date, UNIX timestamp
# '%ai': author date, ISO 8601-like format
# '%aI': author date, strict ISO 8601 format
# '%cn': committer name
# '%cN': committer name (respecting .mailmap, see git-shortlog or git-blame)
# '%ce': committer email
# '%cE': committer email (respecting .mailmap, see git-shortlog or git-blame)
# '%cd': committer date (format respects --date= option)
# '%cD': committer date, RFC2822 style
# '%cr': committer date, relative
# '%ct': committer date, UNIX timestamp
# '%ci': committer date, ISO 8601-like format
# '%cI': committer date, strict ISO 8601 format
# '%d': ref names, like the --decorate option of git-log
# '%D': ref names without the " (", ")" wrapping.
# '%e': encoding
# '%s': subject
# '%f': sanitized subject line, suitable for a filename
# '%b': body
# '%B': raw body (unwrapped subject and body)
# '%N': commit notes