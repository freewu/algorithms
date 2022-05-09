-- 2556 · Update the age of Linghu Chong
-- # Description
-- We want to update the age of Linghu Chong in the teachers table to 26. 
-- Please add the SQL statement to update the age of Linghu Chong.

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
-- 5	Linghu Chong		26	CN
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
-- Because there is no data for Linghu Chong in Example 2, the original table is returned

-- 不要删除预置代码 --
-- 开启一个事务 -- 
BEGIN;

-- 更新 Linghu Chong 的年龄 --
-- Write your SQL Query here --
UPDATE
	teachers
SET
	age = 26 
WHERE
	name = 'Linghu Chong';

COMMIT;