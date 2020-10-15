create table carts(  -- 购物车
    id varchar(100) primary key, -- id
    total_count int not null, -- 总数量
    total_amount double(11,2) not null, -- 总金额
    user_id int not null, -- 用户id
    foreign key (user_id) references users(id)
)