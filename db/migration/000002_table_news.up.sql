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