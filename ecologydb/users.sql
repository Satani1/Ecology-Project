create table ecologydb.users
(
    ID          int auto_increment
        primary key,
    firstName   varchar(30) not null,
    lastName    varchar(30) not null,
    rating      int         null,
    dateCreated datetime    not null,
    email       varchar(40) null,
    username    varchar(30) null
);

