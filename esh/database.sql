create table if not exists Job(
    job_id int not null primary key,
    job_title text not null,
    job_description text not null
);

create table if not exists Emergency_Contact(
    emergency_contact_id int not null primary key,
    firstname text not null,
    middlename text not null,
    lastname text not null,
    phonenumber text not null,
    email text not null,
    fyda_id text not null

);

create table if not exists Employee(
    employee_id int not null primary key,
    firstname text not null,
    middlename text not null,
    lastname text not null,
    phonenumber int not null,
    fyda_id text not null,
    email text not null,
    hire_date date not null,
    emergency_contact_id int references Emergency_Contact,
    job_id int references Job
);

create table if not exists Candate(
    candate_id int not null,
    firstname text not null,
    middlename text not null,
    lastname text not null,
    phonenumber int not null,
    fyda_id text not null,
    email text not null,
    hire_date date not null,
    cv text null,
    job_id int references Job                        
)

create table if not exists Item(
    employee_id int references Employee,
    item_id text not null primary key,
    item_name text not null,
    item_description text not null,
    item_quantity int not null,
    item_status text not null,
    item_date date not null
)