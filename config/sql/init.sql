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

create table wuso.friendship
(
    id          bigint  not null    comment 'primary key',
    user_id1    bigint  not null    comment 'user_id1',
    user_id2    bigint  not null    comment 'user_id2',
    created_at      timestamp default current_timestamp not null comment 'create time',
    updated_at      timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    deleted_at      timestamp default null              null comment 'user delete time',
    constraint id
        primary key (id),
    index user_id1_idx (user_id1),
    index user_id2_idx (user_id2)
);

create table wuso.friend_req
(
    id              bigint  not null    comment 'primary key',
    sender_id       bigint  not null    comment 'user_id1',
    receiver_id     bigint  not null    comment 'user_id2',
    status          char(1) not null    check(status = '0' OR status = '1' OR status = '2') comment 'friend status, 0 = need confirm, 1 = accepted, 2 = refused',
    created_at      timestamp default current_timestamp not null comment 'create time',
    updated_at      timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    deleted_at      timestamp default null              null comment 'user delete time',
    constraint id
        primary key (id),
    index sender_id_idx (sender_id),
    index receiver_id_idx (receiver_id)
);