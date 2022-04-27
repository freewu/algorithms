-- 1923 · Increasing Number of Infections
-- # Description
-- Write an SQL query to find the IDs of all dates with a higher number of new cases in the United States than the previous day's date.

-- Table definition: new_cases

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- date	date	date
-- increased_count	int	The number of new infections
-- Return result No Order Required

-- Example
-- Example 1:

-- Table content: new_cases

-- id	date	increased_count
-- 1	2020-12-25	100994
-- 2	2020-12-26	216858
-- 3	2020-12-27	152102
-- 4	2020-12-28	189044
-- After running your SQL statement, it should return:

-- id
-- 2
-- 4
-- Explanation：
-- 2020-12-26 The number of new cases in the United States is higher than the previous day (100994 -> 216858)
-- 2020-12-28 The number of new cases in the United States is higher than the previous day (152102 -> 189044)
-- Example 2:

-- Table content: new_cases

-- id	date	increased_count
-- 1	2011-12-25	100994
-- 2	2011-12-26	296858
-- 3	2011-12-27	152102
-- 4	2011-12-28	149044
-- After running your SQL statement, it should return:

-- id
-- 2
-- Explanation：
-- 2011-12-26 The number of new cases in the United States is higher than the previous day (100994 -> 296858)
SELECT 
	a.id AS id
FROM
	new_cases AS a
WHERE
	a.increased_count > (
		SELECT 
			b.increased_count
		FROM 
			new_cases AS b
		WHERE
			b.date = DATE_ADD(a.date,INTERVAL -1 DAY)
	)

