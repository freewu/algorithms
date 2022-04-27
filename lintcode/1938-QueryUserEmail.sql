-- 1938 · Query User Email
-- # Description
-- The users table records the user name (username), password (password) and email (email) of the user. Now the user "zhangsan" logs in, and logs in according to the user name and password
-- Please write a SQL statement to query the email of user "zhangsan" based on the user name (zhangsan) and password (zs789852).

-- Table definition: users

-- columns_name	type	explaination
-- username	varchar	user name (primary key)
-- password	varchar	user's password
-- email	varchar	user's email
-- Tip:
-- What to return is username and email

-- Example
-- Example 1:

-- Table content: users

-- username	password	email
-- zhangsan	zs789852	zhangsan@gmail.com
-- lisi	ls654852	lisi@126.com
-- wanger	we951753	wanger@163.com
-- mazouri	mzr753951	mazouri@outlook.com
-- xiangfeitian	xft159357	xiangfeitian@qq.com
-- After running your SQL statement, the table should return:

-- username	email
-- zhangsan	zhangsan@gmail.com
-- Example 2:

-- Table content: users

-- username	password	email
-- zhangsan	zs789852	zhangsan@qq.com
-- lisi	ls654852	lisi@126.com
-- wanger	we951753	wanger@163.com
-- After running your SQL statement, the table should return:

-- username	email
-- zhangsan	zhangsan@qq.com
SELECT 
	username,
	email
FROM
	users
WHERE
	username = 'zhangsan' AND
	password = 'zs789852'