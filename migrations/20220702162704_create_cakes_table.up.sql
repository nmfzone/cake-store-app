CREATE TABLE IF NOT EXISTS cakes
(
    id bigint unsigned not null primary key,
    title varchar(255) not null,
    description varchar(255),
    rating tinyint,
    image text,
    created_at datetime not null default now(),
    updated_at datetime not null default now()
);
