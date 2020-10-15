create table cart_items( -- 创建购物项
id int primary key auto_increment, -- cartItem id
count int not null, -- 图书数量
amount double(11,2) not null, -- 金额小计
book_id int not null, -- 图书的id
cart_id varchar(100) not null, -- 购物车id
foreign key(book_id) references books(id), -- 购物车中购物项的图书id与books中的id关联
foreign key(cart_id) references carts(id) -- 购物车中购物项的图书id与books中的id关联
)