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
   tasks(detail,isAvailable)
values
   ("運動させる",1);

insert into
   tasks(detail,isAvailable)
values
   ("音楽を流す",1);

insert into
   tasks(detail,isAvailable)
values
   ("報酬を与える",0);

insert into
   tasks(detail,isAvailable)
values
   ("瞑想させる",0);

insert into
   counts(count)
values
   (0);

insert into
   counts(count)
values
   (0);

insert into
   counts(count)
values
   (0);

insert into
   counts(count)
values
   (0);