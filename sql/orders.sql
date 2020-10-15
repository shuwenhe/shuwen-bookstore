create table orders(
id varchar(100) primary key,
create_time datetime not null,
total_count int not null,
total_amount double
(11,2) not null,
state int not null,
user_id int,
foreign key
(user_id) references users
(id)
)