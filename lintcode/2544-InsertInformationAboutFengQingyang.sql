-- 2544 · Insert information about Feng Qingyang
-- # Description
-- We want to insert a message to the teachers table for Feng Qingyang with email feng.qingyang@163.com, age 37 and nationality CN, 
-- but the teachers table is locked with a read lock, write an SQL statement to insert the information.

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

-- | id    | name             | email                     | age    | country |
-- | :---: | :--------------: | :----------------------:  | :----: | :-----: |
-- | 1     | Eastern Heretic  | eastern.heretic@gmail.com | 20     | UK      |
-- | 2     | Northern Beggar  | northern.beggar@qq.com    | 21     | CN      |
-- | 3     | Western Venom    | western.venom@163.com     | 28     | USA     |
-- | 4     | Southern Emperor | southern.emperor@qq.com   | 21     | JP      |
-- | 5     | Linghu Chong     |                           | 18     | CN      |

-- After running your SQL statement, the table should return.

-- | id    | name             | email                     | age    | country |
-- | :---: | :--------------: | :----------------------:  | :----: | :-----: |
-- | 1     | Eastern Heretic  | eastern.heretic@gmail.com | 20     | UK      |
-- | 2     | Northern Beggar  | northern.beggar@qq.com    | 21     | CN      |
-- | 3     | Western Venom    | western.venom@163.com     | 28     | USA     |
-- | 4     | Southern Emperor | southern.emperor@qq.com   | 21     | JP      |
-- | 5     | Linghu Chong     |                           | 18     | CN      |
-- | 6     | Feng Qingyang    | feng.qingyang@163.com     | 37     | CN      |

-- Sample 2:

-- | id    | name             | email                     | age    | country |
-- | :---: | :--------------: | :----------------------:  | :----: | :-----: |
-- | 1     | Eastern Heretic  | eastern.heretic@gmail.com | 20     | UK      |
-- | 2     | Northern Beggar  | northern.beggar@qq.com    | 21     | CN      |
-- | 3     | Western Venom    | western.venom@163.com     | 28     | USA     |
-- | 4     | Southern Emperor | southern.emperor@qq.com   | 21     | JP      |

-- After running your SQL statement, the table should return.

-- | id    | name             | email                     | age    | country |
-- | :---: | :--------------: | :----------------------:  | :----: | :-----: |
-- | 1     | Eastern Heretic  | eastern.heretic@gmail.com | 20     | UK      |
-- | 2     | Northern Beggar  | northern.beggar@qq.com    | 21     | CN      |
-- | 3     | Western Venom    | western.venom@163.com     | 28     | USA     |
-- | 4     | Southern Emperor | southern.emperor@qq.com   | 21     | JP      |
-- | 5     | Feng Qingyang    | feng.qingyang@163.com     | 37     | CN      |

-- 对 teachers 表上读锁，不要删除该代码 --
LOCK TABLES teachers READ;

-- Write your SQL Query here --
UNLOCK TABLES;

INSERT INTO 
	teachers (
		`name`,
		`email`,
		`age`,
		`country`
	)
VALUES (
	'Feng Qingyang',
	'feng.qingyang@163.com',
	37,
	'CN'
)