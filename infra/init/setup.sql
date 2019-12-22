create database gochat;
\connect gochat;

drop table if exists user_table;
drop table if exists todo_table;
drop table if exists session_table;

create table user_table (
    id         serial primary key,
    uuid       varchar(64) not null unique,
    name       varchar(255),
    email      varchar(255) not null unique,
    password   varchar(255) not null,
    created_at timestamp not null   
);

create table todo_table (
    id         serial primary key,
    uuid       varchar(64) not null unique,
    title      varchar(255),
    body       text,
    deadline   date,
    user_id    integer references user_table(id),
    created_at timestamp not null  
);

create table session_table (
    id         serial primary key,
    uuid       varchar(64) not null unique,
    email      varchar(255),
    user_id    integer references user_table(id),
    created_at timestamp not null   
)