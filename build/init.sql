create table parameters (
    id integer primary key, 
    funds integer, 
    btc integer
);

insert into parameters values (1, 10000, 10);

create table persons (
    person_id uuid primary key,
    name text not null,
    weight integer
);

create table card_brands (
    brand text primary key
);

create table cards (
    card_id uuid primary key,
    brand text references card_brands(brand) on update cascade
);

create table orders (
    order_id uuid primary key,
    person_id uuid references persons(person_id)
);

create table payments (
    order_id uuid primary key references orders(order_id),
    card_id uuid references cards(card_id)
);

insert into persons values ('f3bf75a9-ea4c-4f57-9161-cfa8f96e2d0b', 'Jerry', 1);

insert into card_brands values ('VISA'), ('AMEX');

insert into cards values ('3224ebc0-0a6e-4e22-9ce8-c6564a1bb6a1', 'VISA');

insert into orders values ('722b694c-984c-4208-bddd-796553cf83e1', 'f3bf75a9-ea4c-4f57-9161-cfa8f96e2d0b');

insert into payments values ('722b694c-984c-4208-bddd-796553cf83e1', '3224ebc0-0a6e-4e22-9ce8-c6564a1bb6a1');