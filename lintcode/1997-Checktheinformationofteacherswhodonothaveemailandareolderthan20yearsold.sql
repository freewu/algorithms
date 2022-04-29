-- 1997 · Check the information of teachers who do not have email and are older than 20 years old
-- # Description
-- Write an SQL statement to find teachers who have no mailbox and are older than 20 years old from the teachers table.

-- Table Definition: teachers

-- column name	type	comment
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)


-- The Email does not exist when the email field is empty
-- If the query does not return any results, nothing will be returned
-- Example
-- Sample I:

-- Table content : teachers

-- | id | name | email | age | country |
-- | ---- | ---- | ---- | ---- | ---- | ---- | ----
-- | 1 | Eastern Heretic | eastern.heretic@gmail.com | 20 | UK |
-- | 2 | Northern Beggar | northern.beggar@qq.com | 21 | CN |
-- | 3 | Western Venom | | 28 | USA |
-- | 4 | Southern Emperor | southern.emperor@qq.com | 21 | JP |
-- | 5 | Linghu Chong | | 18 | CN |
-- | 6 | Northern Beggar | | 29 | CN |

-- After running your SQL statement, the table should return.

-- | id | name | email | age | country |
-- | :----: | :----: | :----: | :----: | :----: | :----: |
-- | 3 | Western Venom | | 28 | USA |
-- | 6 | Northern Beggar | | 29 | CN |

-- Sample 2:

-- Table Contents : teachers

-- | id | name | email | age | country |
-- | ---- | ---- | ---- | ---- | ---- | ---- | ----
-- | 1 | Eastern Heretic | eastern.heretic@gmail.com | 18 | UK |
-- | 2 | Northern Beggar | northern.beggar@qq.com | 21 | CN |
-- | 3 | Western Venom | western.venom@163.com | 28 | USA |
-- | 4 | Southern Emperor | southern.emperor@qq.com | 21 | JP |
-- | 5 | Linghu Chong | | 18 | USA |

-- After running your SQL statement, the table should return.

-- | id | name | email | age | country |
-- | ----: | :----: | :----: | :----: | :----: | :----: |

-- Since there is no data that matches the criteria in the input sample, only the title is shown here, no data.
SELECT
	*
FROM
	teachers
WHERE
	age > 20 AND
	email IS NULL