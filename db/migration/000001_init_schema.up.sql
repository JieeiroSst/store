create table feed_backs(
    id bigserial primary key ,
    name varchar(255) not null,
    phone varchar(255) not null,
    email varchar(255) not null ,
    address varchar(255) not null,
    content varchar(255) not null ,
    status bit,
    created_at date
);

create table menues(
    id bigserial primary key,
    text varchar(255),
    link varchar(255),
    display_order int,
    target varchar(255),
    status bit
);


create table tags(
    id bigserial primary key,
    name varchar(255)
);

create table new_tag(
    tag_id bigserial primary key,
    name varchar(255)
);

create table system_config(
    id bigserial primary key ,
    name varchar(255),
    type varchar(255),
    value varchar(255),
    status bit
);

create table users(
    id bigserial primary key ,
    username varchar(255),
    password varchar(255)
);

create table profile (
    userId bigserial primary key,
    first_name varchar(255),
    last_name varchar(255),
    address varchar(255),
    phone varchar(255),
    created_at date,
    created_by varchar(255),
    modified_date date,
    modified_by varchar(255),
    status bit
);

create table news(
    id bigserial primary key ,
    title varchar(255),
    meta_title varchar(255),
    description varchar(255),
    image varchar(255),
    category_id int,
    detail text,
    created_at date,
    created_by varchar(255),
    modified_data varchar(255),
    modified_by varchar(255),
    meta_key_word varchar(255),
    meta_description varchar(255),
    status bit,
    top_hot date,
    view_count int,
    tag_id int
);