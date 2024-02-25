create table people(
    id serial primary key,
    name varchar,
    document varchar
);

INSERT INTO people(name, document) VALUES
('Neymar', '123421'),
('Messi', '234543');