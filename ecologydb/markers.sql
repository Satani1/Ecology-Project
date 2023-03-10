create table ecologydb.markers
(
    mark_id     int auto_increment
        primary key,
    name        varchar(30) null,
    latitude    float       null,
    longitude   float       null,
    address     text        null,
    description text        null
);

