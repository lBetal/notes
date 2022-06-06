CREATE TABLE users(
    id serial not null unique, 
    name varchar(255) not null, 
    surname varchar(255) not null, 
    phone_number int not null, 
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE device
(
    id                serial       not null unique,
    phone_model       varchar(255) not null,
    phone_number      BIGINT       not null,
    identification    BIGINT       not null,
    imei_code         BIGINT       not null
);

CREATE TABLE users_devices
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    device_id int references device (id) on delete cascade not null
);

CREATE TABLE items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);


CREATE TABLE devices_items
(
    id      serial                                           not null unique,
    item_id int references items (id) on delete cascade not null,
    device_id int references device (id) on delete cascade not null
);