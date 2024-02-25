create table person(
    id serial primary key,
    name varchar,
    document varchar
);

INSERT INTO person(name, document) VALUES
('Neymar', '123421'),
('Messi', '234543');