---create tabale---
create table if not exists customers(
	customer_id int primary key,
	customer_name char(50) not null
);

---create tabale---
create table if not exists products(
	prodcuts_id int primary key,
	procuct_name char(50) not null
);

---create tabale---
create table if not exists orders(
	orders_id int primary key,
	order_date date not null,
	total float not null,
	customer_id int not null,
	prodcuts_id int not null,
	constraint fk_customer foreign key (customer_id) references customers (customer_id),
	constraint fk_product foreign key (prodcuts_id) references products (prodcuts_id)
);


insert into customers (customer_id, customer_name) values (1, 'Bambang');
insert into customers (customer_id, customer_name) values (2, 'Ahmad');

insert into products (prodcuts_id, procuct_name) values (1, 'Hp Pavilion gaming 15')
--alter table orders add constraint fk_customer foreign key (customer_id) references cutomers (customer_id);
--add column products_id int not null constraint;

--alter table orders add constraint 