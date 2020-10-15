create table order_items(
id int primary key,
count int not null,
amount double(11,2) not null,
title varchar(100) not null,
author varchar(100) not null,
price double(11,2) not null,
img_path varchar(100) not null,
order_id varchar(100) not null,
foreign key(order_id) references orders(id)
)