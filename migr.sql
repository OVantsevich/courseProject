create table if not exists users
(
    id                serial
        constraint users_pk
            primary key,
    user_login        varchar(100)                              not null
        constraint users_user_login_unique
            unique,
    user_password     varchar(200)                              not null,
    user_name         varchar(20)                               not null,
    surname           varchar(50)                               not null,
    is_deleted        boolean      default false                not null,
    creation_date     timestamp(6) default CURRENT_TIMESTAMP(6) not null,
    modification_date timestamp(6) default CURRENT_TIMESTAMP(6) not null
);

alter table users
    owner to postgres;

create unique index if not exists users_id_uindex
    on users (id);

