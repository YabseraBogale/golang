create table if not exists Job(
    job_id int GENERATED ALWAYS AS IDENTITY,
    job_title text not null,
    office text not null,
    Department text not null,
    job_description text not null
);

create table if not exists Employee(
    employee_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    firstname text not null,
    middlename text not null,
    lastname text not null,
    phonenumber text not null,
    fyda_id text not null,
    email text not null,
    hire_date date not null,
    salary float not null,
    year_experince int not null,
    office text not null,
    emergency_firstname text not null,
    emergency_middlename text not null,
    emergency_lastname text not null,
    emergency_phonenumber text not null,
    emergency_email text not null,
    emergency_fyda_id text not null,
    Department text not null,
    job_title text not null
);

create table if not exists Item(
    item_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	employee_id uuid references Employee,
    item_name text not null,
    item_description text not null,
    item_quantity int not null,
    item_status text not null,
    item_date date not null
);

create table purchase_request(
    item_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	employee_id uuid references Employee,
    item_name text not null,
    item_description text not null,
    item_quantity int not null,
    item_status text not null,
    item_purchase_request text not null,
    item_date date not null
)

create table Inventory(
    inventory_iD INTEGER PRIMARY KEY AUTOINCREMENT,
    item_id uuid references Item,
    Quantity INTEGER NOT NULL DEFAULT 0,
    LastUpdated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   

)
