CREATE TABLE users
(
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

CREATE TABLE photo
(
    id          serial       not null unique,
    path       varchar(255) not null
);


CREATE TABLE devices_photos
(
    id      serial                                           not null unique,
    photo_id int references photo (id) on delete cascade not null,
    device_id int references device (id) on delete cascade not null
);

CREATE TABLE video
(
    id          serial       not null unique,
    path       varchar(255) not null
);


CREATE TABLE devices_videos
(
    id      serial                                           not null unique,
    video_id int references video (id) on delete cascade not null,
    device_id int references device (id) on delete cascade not null
);

CREATE TABLE audio
(
    id          serial       not null unique,
    path       varchar(255) not null
);


CREATE TABLE devices_audios
(
    id      serial                                           not null unique,
    audio_id int references audio (id) on delete cascade not null,
    device_id int references device (id) on delete cascade not null
);