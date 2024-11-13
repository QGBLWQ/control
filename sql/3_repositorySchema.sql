-- 提供：插入file，查询这个用户的仓库，根据id返回url,根据id删除file

-- +migrate Up
DROP TABLE IF EXISTS user_files;

-- 创建 files 表
CREATE TABLE files (
    file_id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,   -- file_id 为主键，且自动递增
    user_id INT UNSIGNED NOT NULL,                      -- user_id 外键，指向用户表（users）
    filename VARCHAR(255) NOT NULL,            -- 文件名字段，最大长度 255
    upload_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 上传时间，默认为当前时间戳
    FOREIGN KEY (user_id) REFERENCES users(id)   -- 外键约束，指向 users 表中的 id 字段
);

-- +migrate Down
-- 删除 files 表
DROP TABLE files;
