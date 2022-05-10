-- 2627 · Troubleshoot the current database table locks and view the table lock analysis
-- # Description
-- Now we need to troubleshoot the current database table lock situation, please write SQL statement to view the table lock analysis

-- Example
-- After entering your SQL statement, the result should be returned:

-- Variable_name
-- 'Table_locks_immediate'
-- 'Table_locks_waited'
-- 'Table_open_cache_hits'
-- 'Table_open_cache_misses'
-- 'Table_open_cache_overflows'

-- SHOW VARIABLES 
SHOW VARIABLES like '%Table_locks%' 

-- SHOW VARIABLES 
SHOW STATUS like '%Table_open_cache%';


-- # Open_tables & Opened_tables
-- show global status like 'open%_tables'; 
-- +---------------+--------+
-- | Variable_name | Value  |
-- +---------------+--------+
-- | Open_tables   | 384    | 打开的缓存表数量
-- | Opened_tables | 768975 | 打开的所有表数量
-- table_open_cache 这个参数设置多少合适呢？
-- —肯定不是越大越好，太大了会占用太多文件描述符，描述符用尽会导致无法建立新连接。那这么设置这个值呢？
-- 1. 根据 Open_tables，这个参数接近table_open_cache，说明缓冲池快满了
-- 2. Opened_tables还一直在增加，说明还有新的表打开没有被缓存
-- 总结，同时符合上面两条说明你的表打开缓存太小了，需要适量增加。

-- # Table_locks_immediate
-- 可以立即授予表锁请求的次数。

-- # Table_locks_waited
-- 无法立即授予表锁请求的次数，需要等待。如果这个值很高，并且存在性能问题，那么应该首先优化查询，然后拆分表或表，或者使用复制。

-- # Table_open_cache_hits
-- 打开表缓存查找的命中次数。

-- # Table_open_cache_misses
-- 打开表缓存查找失败的次数。

-- # Table_open_cache_overflows
-- 打开表缓存的溢出数。这是在打开或关闭表之后，缓存实例有一个未使用的条目，且实例的大小大于table_open_cache / table_open_cache_instances的次数。


