create table Clients
(
    id            serial primary key,
    name          varchar(255) not null,
    age           int          not null,
    email         varchar(50) unique,
    phone_number  varchar(50)  not null unique,
    password_hash varchar(255) not null
);

create table Roles
(
    id_role serial primary key,
    name    varchar(50) not null
);

create table Payment_methods
(
    id_payment_method serial primary key,
    number            int not null
);

create table Foods
(
    id_food     serial primary key,
    name        varchar(50) NOT NULL,
    weight      varchar(50) NOT NULL,
    description varchar(50) NOT NULL,
    compound    varchar(50) NOT NULL,
    number_food int         NOT NULL
);

create table Employees
(
    id_employee   serial primary key,
    login         varchar(50)  not null unique,
    email         varchar(50) unique,
    name          varchar(50)  not null,
    surname       varchar(50)  not null,
    middle_name   varchar(50)  not null,
    password_hash varchar(255) not null
--     id_role       int          not null,
--     foreign key (id_role) references Roles (id_role) on delete cascade
);

create table Orders
(
    id_order          serial primary key,
    uuid              varchar(32) not null,
    id_payment_method int         not null,
    id_client         int         not null,
    foreign key (id_payment_method) references Payment_methods (id_payment_method) on delete cascade,
    foreign key (id_client) references Clients (id) on delete cascade
);

create table Tables
(
    id_table     serial primary key,
    number_table int NOT NULL,
    id_client    int NOT NULL,
    id_order     int NOT NULL,
    foreign key (id_client) references Clients (id) on delete cascade,
    foreign key (id_order) references Orders (id_order) on delete cascade
);

create table Food_order
(
    id_food_order serial not null,
    id_food       int    not null,
    id_order      int    not null,
    foreign key (id_food) references Foods (id_food) on delete cascade,
    foreign key (id_order) references Orders (id_order) on delete cascade
);