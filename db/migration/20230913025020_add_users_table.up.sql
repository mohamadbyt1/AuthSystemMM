CREATE TABLE If NOT EXISTS "users"(
    "id" serial primary key,
    "userName" varchar  not null unique,
    "password" varchar not null,
    "role" varchar DEFAULT "user")

CREATE TABLE IF NOT EXISTS "products"(
    "id" serial primary key,
    "name" varchar not null,
    "category" varchar not null,
    "price" float not null,
    "stock" integer not null)

