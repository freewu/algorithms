-- 2542 · Update Linghu Chong's email
-- # Description
-- We want to update the mailbox of Linghu Chong in the teachers table to linghu.chong@ninechapter.com, 
-- but the teachers table is under a read lock, write an SQL statement to update the mailbox of Linghu Chong.

-- Table definition: teachers (teachers table)

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- Please note that the teachers table is read locked

-- Example
-- Sample 1:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	linghu.chong@ninechapter.com	18	CN
-- Sample 2:

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP

-- Because there is no data for Linghu Chong in Example 2, the original table is returned

-- 对 teachers 表上读锁，不要删除该代码 --
LOCK TABLES teachers READ;

-- Write your SQL Query here --
-- 先要释放锁
UNLOCK TABLES;

UPDATE
	teachers
SET
	email = 'linghu.chong@ninechapter.com'
WHERE
	name = 'Linghu Chong'

-- LOCK TABLES为当前线程锁定表。 UNLOCK TABLES释放被当前线程持有的任何锁。当线程发出另外一个LOCK TABLES时，或当服务器的连接被关闭时，当前线程锁定的所有表会自动被解锁。 
-- 如果一个线程获得在一个表上的一个READ锁，该线程和所有其他线程只能从表中读。
--      lock tables <table-name> READ;

-- 如果一个线程获得一个表上的一个WRITE锁，那么只有持锁的线程READ或WRITE表，其他线程被阻止。
--      lock tables <table-name> WRITE;

-- 释放被当前线程持有的任何锁
--      unlock tables;