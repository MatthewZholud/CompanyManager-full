CREATE TABLE company(
    company_id Serial primary key,
    name varchar(255) not null,
    legal_form varchar(255)
);
CREATE TABLE employees(
    employee_id serial primary key,
    name varchar(255) not null,
    secondName varchar(255),
    surname varchar(255),
    photoUrl varchar(255) not null,
    hireDate timestamp without time zone,
    position varchar(255),
    company_id int REFERENCES company(company_id) ON DELETE CASCADE
);
