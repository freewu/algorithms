-- 2560 · Undo the insertion of Xie Xun's message
-- # Description
-- We need to undo the insertion of Xie Xun into the teachers table. 
-- Please add the SQL statement to undo the insertion of Xie Xun.

-- Table definition: teachers (Teachers table)

-- Column Name	Type	Comments
-- id	int	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Tutor's email
-- age	int	Tutor's age
-- country	varchar	Tutor's nationality

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
-- 5	Linghu Chong		18	CN
-- Example 2:

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

-- 不要删除预置代码 --
-- 开启一个事务 -- 
BEGIN;

-- 插入 Xie Xun 的信息 --
INSERT INTO teachers (name, age, country)
VALUES ('Xie Xun', 49, 'CN');

-- Write your SQL Query here --
ROLLBACK;