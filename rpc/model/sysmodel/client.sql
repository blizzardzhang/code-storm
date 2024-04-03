create table storm_client
(
    id              bigint        not null,
    name            varchar(50)   ,
    client_id       varchar(50)   ,
    client_secret   varchar(100)  ,
    client_key      varchar(100)  ,
    additional_info text          ,
    grant_type      varchar(255)  ,
    expire_in       bigint        ,
    refresh_after   bigint        ,
    remark          varchar(100)  ,
    status          int default 1 not null,
    is_deleted      int default 0 not null,
    primary key (id)
)
    comment '客户端';

