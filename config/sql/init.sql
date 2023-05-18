create table wuso.user
(
    id              bigint                              not null comment 'user id',
    email           varchar(255)                                not null unique comment 'user email',
    name            varchar(255)                                not null comment 'user name',
    password        varchar(255)                                not null comment 'password',
    created_at      timestamp default current_timestamp not null comment 'create time',
    updated_at      timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    deleted_at      timestamp default null              null comment 'user delete time',
    constraint id
        primary key (id)
);