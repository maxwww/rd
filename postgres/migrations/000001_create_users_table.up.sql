CREATE TABLE users
(
    id            serial       not null unique,
    telegram_id   integer      not null unique,
    is_bot        boolean      not null,
    first_name    varchar(255) not null,
    last_name     varchar(255),
    user_name     varchar(255),
    language_code varchar(255),
    notify        boolean      not null default false
);
