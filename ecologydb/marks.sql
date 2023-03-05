create table marks
(
    mark_id     int auto_increment
        primary key,
    dateCreated datetime      null,
    type        int           null,
    description text          null,
    addr        varchar(100)  not null,
    status      int default 1 not null,
    name        varchar(50)   null
);

