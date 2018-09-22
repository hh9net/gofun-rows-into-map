drop database test;

create database test;

use test;

create table test_table(
	id int,
	version int,
	name varchar(50),
	update_time timestamp(6)
);

select * from test_table;

insert into test_table values (1, 1, "Joe", current_timestamp(6));
insert into test_table values (2, 1, "Tom", current_timestamp(6));
insert into test_table values (1, 2, "Joseph", current_timestamp(6));