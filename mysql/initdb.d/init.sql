use backend_db;

-- TODO: API設計
 CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    detail VARCHAR(128) NOT NULL COMMENT 'Task Content',
    isAvailable BOOLEAN NOT NULL
 );

 CREATE TABLE counts (
    id SERIAL,
    count INT NOT NULL,
    FOREIGN KEY(id) REFERENCES tasks(id) ON DELETE CASCADE
);
