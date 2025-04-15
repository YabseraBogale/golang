drop table Job
drop table Emergency_Contact
drop table Employee
drop table item
create table if not exists Job(
    job_id int not null primary key,
    job_title text not null,
    job_description text not null
);

create table if not exists Emergency_Contact(
    emergency_contact_id int not null primary key,
    firstname varchar(30) not null,
    middlename varchar(30) not null,
    lastname varchar(30) not null,
    phonenumber varchar(30) not null,
    email varchar(30) not null,
    fyda_id text not null

);

create table if not exists Employee(
    employee_id int not null primary key,
    firstname varchar(30) not null,
    middlename varchar(30) not null,
    lastname varchar(30) not null,
    phonenumber int not null,
    fyda_id text not null,
    email text not null,
    hire_date date not null,
    emergency_contact_id int references Emergency_Contact,
    job_id int references Job
);

create table if not exists item(
    employee_id int references Employee,
    item_id varchar(30) not null primary key,
    item_name varchar(30) not null,
    item_description text not null,
    item_quantity int not null,
    item_status varchar(30) not null,
    item_date date not null
)