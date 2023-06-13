# go_xorm_cloud_disk
学习地址: 
[【【项目实战】基于Go-zero、Xorm的网盘系统】](https://www.bilibili.com/video/BV1cr4y1s7H4)

## 建表语句

```mysql
create table user_basic
(
    id         int auto_increment comment '用户表ID'
        primary key,
    identity   varchar(36)  null comment '用户唯一标识',
    name       varchar(256) null comment '用户登录名称',
    password   varchar(32)  null comment '用户密码',
    email      varchar(100) null comment '注册邮箱',
    created_at datetime     null comment '用户创建时间',
    updated_at datetime     null comment '用户信息更新时间',
    deleted_at datetime     null comment '用户删除时间'
)
    comment '用户表';
```

```mysql
create table user_repository
(
    id                  int auto_increment
        primary key,
    identity            varchar(36)  null,
    user_identity       varchar(36)  null,
    parent_id           int          null comment '用户上传父级ID',
    repository_identity varchar(36)  null comment '关联中心存储池的标识',
    hash                varchar(32)  null comment '文件的唯一标识',
    name                varchar(255) null comment '文件名称',
    ext                 varchar(30)  null comment '扩展名(显示文件夹类型等信息)',
    created_at          datetime     null comment '文件创建时间',
    updated_at          datetime     null comment '文件修改时间',
    deleted_at          datetime     null comment '文件删除时间'
)
    comment '个人存储池';
```

```mysql
create table repository_pool
(
    id         int auto_increment
        primary key,
    identity   varchar(36)  null,
    hash       varchar(32)  null comment '文件的唯一标识',
    name       varchar(255) null comment '文件名称',
    ext        varchar(30)  null comment '扩展名',
    size       int          null comment '文件大小',
    path       varchar(255) null comment '文件路径',
    created_at datetime     null comment '文件创建时间',
    updated_at datetime     null comment '文件修改时间',
    deleted_at datetime     null comment '文件删除时间'
)
    comment '存储池';
```





