create table marks
(
    mark_id     int auto_increment
        primary key,
    dateCreated datetime      not null,
    type        int           not null,
    description varchar(100)  null,
    addr        varchar(100)  not null,
    status      int default 1 not null
);

