-- 2016 · Update the email address of the designated teacher
-- Description
-- Please write an SQL statement to update the email address of the teacher named Linghu Chong to "linghu.chong@lintcode.com" in the teacher table teachers.

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- The update record field type should match the table definition field type

-- Example
-- Eample 1 ：

-- Table content: teachers

-- | id   | name             | email                     | age  | country |
-- | ---- | ---------------- | ------------------------- | ---- | ------- |
-- | 1    | Eastern Heretic  | eastern.heretic@gmail.com | 20   | UK      |
-- | 2    | Northern Beggar  | northern.beggar@qq.com    | 21   | CN      |
-- | 3    | Western Venom    | western.venom@163.com     | 28   | USA     |
-- | 4    | Southern Emperor | southern.emperor@qq.com   | 21   | JP      |
-- | 5    | Linghu Chong     |                           | 18   | CN      |

-- After running your SQL statement, we will execute the following statement:

-- SELECT * 
-- FROM teachers 
-- WHERE name='Linghu Chong';
-- We will use the result of the execution of this statement to determine the final outcome of the run.

-- return to:

-- id	name	email	age	country
-- 1	Linghu Chong	linghu.chong@lintcode.com	18	CN
-- Eample 2 ：
-- Table content: teachers

-- | id   | name             | email                     | age  | country |
-- | ---- | ---------------- | ------------------------- | ---- | ------- |
-- | 1    | Eastern Heretic  | eastern.heretic@gmail.com | 20   | UK      |
-- | 2    | Northern Beggar  | northern.beggar@qq.com    | 21   | CN      |
-- | 3    | Western Venom    | western.venom@163.com     | 28   | USA     |
-- | 4    | Southern Emperor | southern.emperor@qq.com   | 21   | JP      |

-- After running your SQL statement, we will execute the following statement:

-- SELECT * 
-- FROM teachers 
-- WHERE name='Linghu Chong';
-- We will use the result of the execution of this statement to determine the final outcome of the run.

-- return to:

-- id	name	email	age	country
-- Sample 2 does not have a teacher named Linghu Chong, so the output is empty

UPDATE 
	teachers
SET
	email = 'linghu.chong@lintcode.com'
WHERE
	name = 'Linghu Chong'
