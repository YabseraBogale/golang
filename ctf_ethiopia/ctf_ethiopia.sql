create table ctf_ethiopia_user(
    userid int auto increment primary key,
    username varchar(20) not null,
    email varchar(20) not null,
    password text not null
)