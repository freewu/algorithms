-- 2559 · Withdrawal of age update for Linghu Chong
-- Description
-- We need to undo the update of Linghu Chong's age in the teachers table. 
-- Please add the SQL statement to undo the update of Linghu Chong's age.

-- Table Definition: teachers (Teachers Table)

-- Column Name	Type	Comments
-- id	int	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Tutor's email
-- age	int	Tutor's age
-- country	varchar	Tutor's nationality

-- Example
-- Sample 1:

-- Table content : teachers

-- | id    | name             | email                     | age   | country |
-- | ----: | :--------------: | :-----------------------: | :---: | :-----: |
-- | 1     | Eastern Heretic  | eastern.heretic@gmail.com | 20    | UK      |
-- | 2     | Northern Beggar  | northern.beggar@qq.com    | 21    | CN      |
-- | 3     | Western Venom    | western.venom@163.com     | 28    | USA     |
-- | 4     | Southern Emperor | southern.emperor@qq.com   | 21    | JP      |
-- | 5     | Linghu Chong     |                           | 18    | CN      |

-- After running your SQL statement, the table should return.

-- | id    | name             | email                     | age   | country |
-- | ----: | :--------------: | :-----------------------: | :---: | :-----: |
-- | 1     | Eastern Heretic  | eastern.heretic@gmail.com | 20    | UK      |
-- | 2     | Northern Beggar  | northern.beggar@qq.com    | 21    | CN      |
-- | 3     | Western Venom    | western.venom@163.com     | 28    | USA     |
-- | 4     | Southern Emperor | southern.emperor@qq.com   | 21    | JP      |
-- | 5     | Linghu Chong     |                           | 18    | CN      |

-- Example 2:

-- | id    | name             | email                     | age   | country |
-- | ----: | :--------------: | :-----------------------: | :---: | :-----: |
-- | 1     | Eastern Heretic  | eastern.heretic@gmail.com | 20    | UK      |
-- | 2     | Northern Beggar  | northern.beggar@qq.com    | 21    | CN      |
-- | 3     | Western Venom    | western.venom@163.com     | 28    | USA     |
-- | 4     | Southern Emperor | southern.emperor@qq.com   | 21    | JP      |

-- After running your SQL statement, the table should return.

-- | id    | name             | email                     | age   | country |
-- | ----: | :--------------: | :-----------------------: | :---: | :-----: |
-- | 1     | Eastern Heretic  | eastern.heretic@gmail.com | 20    | UK      |
-- | 2     | Northern Beggar  | northern.beggar@qq.com    | 21    | CN      |
-- | 3     | Western Venom    | western.venom@163.com     | 28    | USA     |
-- | 4     | Southern Emperor | southern.emperor@qq.com   | 21    | JP      |

-- Because there is no data for Linghu Chong in Example 2, the original table is returned

-- 不要删除预置代码 --
-- 开启一个事务 -- 
BEGIN;

-- 更新 Linghu Chong 的年龄 --
UPDATE teachers
SET age = 26
WHERE name = 'Linghu Chong';

-- Write your SQL Query here --
ROLLBACK;