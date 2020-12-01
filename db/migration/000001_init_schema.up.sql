create table abouts(
                       id bigserial primary key ,
                       title varchar(255) not null ,
                       meta_title varchar(255),
                       description varchar(255),
                       image varchar(255),
                       detail varchar(255),
                       created_at date,
                       created_by varchar(255),
                       modified_date date,
                       modified_by varchar(255),
                       meta_keyword varchar(255),
                       meta_description varchar(255),
                       status bit
);

create table contacts(
                         id bigserial primary key ,
                         content varchar(255),
                         status bit
);

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

create table categories(
                           id bigserial primary key ,
                           name varchar(255) not null ,
                           meta_title varchar(255) not null ,
                           display_order int,
                           created_at date,
                           created_by varchar(255) not null ,
                           modified_date date,
                           modified_by varchar(255),
                           meta_keyword varchar(255),
                           meta_description varchar(255),
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

create table tags(
                     id bigserial primary key,
                     name varchar(255)
);

create table new_tag(
                        tag_id bigserial primary key,
                        name varchar(255)
);

create table product_category(
                                 id bigserial primary key ,
                                 name varchar(255),
                                 meta_title varchar(255),
                                 display_order int,
                                 title varchar(255),
                                 created_at date,
                                 created_by varchar(255),
                                 modified_date date,
                                 modified_by varchar(255),
                                 description varchar(255),
                                 status bit
);

create table products(
                         id bigserial primary key ,
                         code varchar(255),
                         name varchar(255),
                         title varchar(255),
                         description varchar(255),
                         images varchar(255),
                         price decimal(18,0),
                         vat bit,
                         category_id int,
                         detail text,
                         created_date date,
                         created_by varchar(255),
                         modified_date date,
                         modified_by varchar(255),
                         status bit,
                         top_hot date,
                         view_counts int
);

create table sliders(
                        id bigserial primary key ,
                        image varchar(255),
                        display_order int,
                        link varchar(255),
                        description varchar(255),
                        created_at date,
                        created_by varchar(255),
                        modified_at date,
                        modified_by varchar(255),
                        status bit
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
                      password varchar(255),
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