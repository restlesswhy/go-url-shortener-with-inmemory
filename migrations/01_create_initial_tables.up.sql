CREATE TABLE urls
(
    short_url varchar(255) not null unique,
    long_url text not null unique
);