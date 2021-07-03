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

insert into
   tasks(id,detail,isAvaliable)
values
   (0,"運動させる",1);

insert into
   tasks(id,detail,isAvaliable)
values
   (1,"音楽を流す",1);

insert into
   tasks(id,detail,isAvaliable)
values
   (2,"報酬を与える",0);

insert into
   tasks(id,detail,isAvaliable)
values
   (3,"瞑想させる",0);