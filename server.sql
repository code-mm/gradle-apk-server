drop table info;
drop table clientid;
drop table channel;

create table info
(
    _id                  integer primary key auto_increment,
    clientid             text    not null,
    channel              text    not null,
    applicationid        text    not null,
    versionname          text    not null,
    maxsdkversion        integer default 28,
    minsdkversion        integer default 16,
    targetsdkversion     integer default 26,
    versioncode          integer not null,
    maxsdkversion_enable bool    default false
);

create table clientid
(
    _id      integer primary key auto_increment,
    clientid text not null
);


insert into clientid
values (1, 'txapp');
insert into clientid
values (2, 'tgapp');
insert into clientid
values (3, 'kdapp');
insert into clientid
values (4, 'lhapp');
insert into clientid
values (5, 'lsapp');
insert into clientid
values (6, 'kbapp');


create table channel
(
    _id     integer primary key auto_increment,
    channel text not null
);
insert into channel
values (1, 'chujian');
insert into channel
values (2, 'chujianthail');
insert into channel
values (3, 'huawei');
insert into channel
values (4, 'xiaomi');
insert into channel
values (5, 'samsung');
insert into channel
values (6, 'vivo');
insert into channel
values (7, 'oppo');
insert into channel
values (8, 'dangle');
insert into channel
values (9, 'ysdk');

insert into info (`_id`, clientid, channel, applicationid, versionname, maxsdkversion, minsdkversion, targetsdkversion,
                  versioncode, maxsdkversion_enable)
values (1, 'kdapp', 'chujian', 'com.chujian.app', '1.0.0.0', 29, 16, 28, 10000, false);


create table client_channel
(
    _id      integer primary key auto_increment,
    clientid text,
    channel  text
);


insert into client_channel (`_id`, clientid, channel)values (1,'txapp','chujian');
insert into client_channel (`_id`, clientid, channel)values (2,'txapp','vivo');
insert into client_channel (`_id`, clientid, channel)values (3,'txapp','oppo');
insert into client_channel (`_id`, clientid, channel)values (4,'txapp','xiaomi');
insert into client_channel (`_id`, clientid, channel)values (5,'txapp','samsung');
insert into client_channel (`_id`, clientid, channel)values (6,'txapp','huawei');


select * from client_channel where clientid='txapp';


select *
from info;


select *
from info
where clientid = 'kdapp'
  and channel = 'chujian';


update info set applicationid='com.chujian.app.q1' where clientid='kdapp' and channel='chujian'