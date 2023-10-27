CREATE TABLE persons
(
	id serial primary key,
	name varchar(255) not null,
	surname varchar(255)    not null,
	patronymic varchar(255),
	age numeric not null, 
    gender varchar(255) not null,
    nationality varchar(255) not null
);