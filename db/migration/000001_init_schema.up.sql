create table feed_backs(
    id bigserial primary key ,
    name varchar(255) not null,
    phone varchar(255) not null,
    email varchar(255) not null ,
    address varchar(255) not null,
    content varchar(255) not null ,
    created_at date default now()
);

create table menues(
    id bigserial primary key,
    text varchar(255),
    link varchar(255),
    target varchar(255)
);


create table tags(
    id bigserial primary key,
    name varchar(255)
);

create table new_tags(
    id bigserial primary key,
    tag_id int,
    new_id int
);

create table system_configs(
    id bigserial primary key ,
    name varchar(255),
    type varchar(255),
    value varchar(255)
);

create table users(
    id bigserial primary key ,
    username varchar(255),
    password varchar(255),
    permission varchar(100)
);

create table profiles (
    id bigserial primary key ,
    user_id int,
    first_name varchar(255),
    last_name varchar(255),
    address varchar(255),
    phone varchar(255),
    created_at date default now(),
    created_by varchar(255)
);

create table news(
    id bigserial primary key ,
    title varchar(255),
    description varchar(255),
    image varchar(255),
    detail text,
    created_at date default now(),
    created_by varchar(255),
    view_count int,
    tag_id int,
    content text,
    active bool default false
);