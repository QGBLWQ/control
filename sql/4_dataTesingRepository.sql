-- 提供：插入file，查询这个用户的仓库，根据id返回url,根据id删除file

-- +migrate Up


-- 插入文件数据
INSERT INTO files (file_id, user_id, filename, upload_time)
VALUES (1, 1, 'test', NOW());


-- +migrate Down
-- 删除 这两个数据

DELETE FROM files WHERE file_id = 1;
