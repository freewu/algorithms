-- 2081 · Insert the current date into the table
-- # Description
-- Write an SQL statement to insert the current date into the record table records.

-- Table definition: records (records table)

-- column name	type	comment
-- now_time	date	current time

-- The inserted record field type should match the table definition field type.
-- Example
-- Table content: records (records table)

-- now_time
-- There is no data in the data table, the type of now_time is `date

-- After running your SQL statement, we will execute some statements that compare the inserted data with the current date of the system and return the result of the comparison

-- Returns.

-- answer
-- true
-- If true is returned, then the current time is correct
-- If false is returned, the inserted time is incorrect

-- use DATE()
INSERT INTO records (`now_time`) VALUES ( DATE(NOW()) )

-- DATE_FORMAT
INSERT INTO records (`now_time`) VALUES ( DATE_FORMAT(NOW(),'%Y-%m-%d') )