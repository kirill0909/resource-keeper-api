CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null,
    date_creation timestamp    not null,
    last_update   timestamp    not null
);

CREATE TABLE users_resources
(
    id                     serial                                                 not null unique,
    user_id                int           references users (id) on delete cascade  not null,
    resource_name          varchar(255)                                           not null,
    resource_login         varchar(255)                                           not null,
    resource_password_enc  varchar(255)                                           not null,
    date_creation          timestamp                                              not null,
    last_update            timestamp                                              not null
);
