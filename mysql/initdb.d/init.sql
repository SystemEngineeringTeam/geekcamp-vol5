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
   tasks(id,detail,isAvailable)
values
   (0,"運動させる",1);

insert into
   tasks(id,detail,isAvailable)
values
   (1,"音楽を流す",1);

insert into
   tasks(id,detail,isAvailable)
values
   (2,"報酬を与える",0);

insert into
   tasks(id,detail,isAvailable)
values
   (3,"瞑想させる",0);

insert into
   counts(id,count)
values
   (0,0)

insert into
   counts(id,count)
values
   (1,0)

insert into
   counts(id,count)
values
   (2,0)

insert into
   counts(id,count)
values
   (3,0)