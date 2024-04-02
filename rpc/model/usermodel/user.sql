create table storm_user
(
    id          bigint        not null,
    account     varchar(100)  null,
    password    varchar(100)  null,
    name        varchar(5)    null,
    real_name   varchar(5)    null,
    phone       varchar(11)   null,
    address     varchar(100)  null,
    email       varchar(100)  null,
    create_time datetime      null,
    create_by   bigint        null,
    update_by   bigint        null,
    update_time datetime      null,
    status      int default 1 not null,
    is_deleted  int default 0 not null,
    primary key (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

