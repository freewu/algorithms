-- 2026 · Insert the current time into the table (accurate to milliseconds)
-- # Description
-- Write an SQL statement to insert the current time (to the millisecond) into the records table records.

-- Table definition: records (records table)

-- column name	type	comment
-- now_time	datetime(3)	now_time

-- The inserted record field type should match the table definition field type

-- Example
-- Table content: records (records table)

-- now_time
-- According to the table there is no data, the type of now_time is datetime(3)

-- After running your SQL statement, we will execute some statements to compare the inserted data with the current time of the system and return the result of the comparison

-- Returns.

-- answer
-- true
-- If true is returned, then the current time is correct
-- If false is returned, the inserted time is incorrect

INSERT INTO `records` (
	`now_time`
)
VALUES (
	CURRENT_TIMESTAMP(3)
)
-- NOW() 只能到秒级

-- current_timestamp() yyyy-mm-dd hh:ii:ss -- 到毫秒级
-- now() yyyy-mm-dd hh:ii:ss
-- curdate() yyyy-mm-dd
-- current_date()
-- curtime() hh:ii:ss
-- current_time()