CREATE TABLE cabin_menu (
    id bigserial primary key ,
    name VARCHAR(100),
    type VARCHAR(100),
    path VARCHAR(100),
    method VARCHAR(100),
    created_on int,
    modified_on int,
    deleted_on int
);

INSERT INTO cabin_menu VALUES ('1', 'system', 'table of Contents', '/system', 'GET', null, null, '0');
INSERT INTO cabin_menu VALUES ('2', 'menu', 'menu', '/api/v1/menus', 'GET', null, null, '0');
INSERT INTO cabin_menu VALUES ('3', 'create_menu', 'button', '/api/v1/menus', 'POST', null, null, '0');
INSERT INTO cabin_menu VALUES ('4', 'update_menu', 'button', '/api/v1/menus/:id', 'PUT', null, null, '0');
INSERT INTO cabin_menu VALUES ('5', 'delete_menu', 'button', '/api/v1/menus/:id', 'DELETE', null, null, '0');
INSERT INTO cabin_menu VALUES ('6', 'user', 'menu', '/api/v1/users', 'GET', null, null, '0');
INSERT INTO cabin_menu VALUES ('7', 'create_user', 'button', '/api/v1/users', 'POST', null, null, '0');
INSERT INTO cabin_menu VALUES ('8', 'update_user', 'button', '/api/v1/users/:id', 'PUT', null, null, '0');
INSERT INTO cabin_menu VALUES ('9', 'delete_user', 'button', '/api/v1/users/:id', 'DELETE', null, null, '0');
INSERT INTO cabin_menu VALUES ('10', 'role', 'menu', '/api/v1/roles', 'GET', null, null, '0');
INSERT INTO cabin_menu VALUES ('11', 'create_role', 'button', '/api/v1/roles', 'POST', null, null, '0');
INSERT INTO cabin_menu VALUES ('12', 'update_role', 'button', '/api/v1/roles/:id', 'PUT', null, null, '0');
INSERT INTO cabin_menu VALUES ('13', 'delete_role', 'button', '/api/v1/roles/:id', 'DELETE', null, null, '0');

CREATE TABLE cabin_role (
    id bigserial primary key ,
    name varchar(50) DEFAULT '',
    created_on int,
    modified_on int,
    deleted_on int
);

INSERT INTO cabin_role VALUES ('1', 'Development Department', null, null, '0');
INSERT INTO cabin_role VALUES ('2', 'Operation and Maintenance Department', null, null, '0');
INSERT INTO cabin_role VALUES ('3', 'Testing Division', null, null, '0');

CREATE TABLE cabin_role_menu (
    id bigserial primary key ,
    role_id int DEFAULT NULL,
    menu_id int DEFAULT NULL ,
    deleted_on int DEFAULT '0'
);

INSERT INTO cabin_role_menu VALUES ('1', '2', '1', '0');
INSERT INTO cabin_role_menu VALUES ('2', '2', '2', '0');
INSERT INTO cabin_role_menu VALUES ('3', '2', '3', '0');
INSERT INTO cabin_role_menu VALUES ('4', '2', '4', '0');
INSERT INTO cabin_role_menu VALUES ('5', '2', '5', '0');
INSERT INTO cabin_role_menu VALUES ('6', '2', '6', '0');
INSERT INTO cabin_role_menu VALUES ('7', '2', '7', '0');
INSERT INTO cabin_role_menu VALUES ('8', '2', '8', '0');
INSERT INTO cabin_role_menu VALUES ('9', '2', '9', '0');
INSERT INTO cabin_role_menu VALUES ('10', '2', '10', '0');
INSERT INTO cabin_role_menu VALUES ('11', '2', '11', '0');
INSERT INTO cabin_role_menu VALUES ('12', '2', '12', '0');
INSERT INTO cabin_role_menu VALUES ('13', '2', '13', '0');

CREATE TABLE cabin_user_role (
    id bigserial primary key ,
    user_id int,
    role_id int,
    deleted_on int
);

INSERT INTO cabin_user_role VALUES ('1', '2', '2', '0');