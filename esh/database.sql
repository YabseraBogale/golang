create table Job(
    job_id int not null primary key,
    job_title text not null,
    job_description text not null
);

create table Employee(
    id int not null primary key,
    firstname varchar(30) not null,
    middlename varchar(30) not null,
    lastname varchar(30) not null,
    phonenumber int not null,
    email text not null,
    job_id int references Job
);