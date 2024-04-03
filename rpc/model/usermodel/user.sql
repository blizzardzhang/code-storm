create table storm_user
(
    id          bigint        not null,
    account     varchar(100)  ,
    password    varchar(100)  ,
    name        varchar(5)   ,
    real_name   varchar(5)    ,
    phone       varchar(11)   ,
    address     varchar(100)  ,
    email       varchar(100)  ,
    create_time datetime      ,
    create_by   bigint        ,
    update_by   bigint        ,
    update_time datetime      ,
    status      int default 1 not null,
    is_deleted  int default 0 not null,
    primary key (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

