DROP TABLE IF EXISTS store;

CREATE TABLE IF NOT EXISTS store (
    id bigserial not null primary key,
    title text,
    photo text,
    price numeric(10, 2)
);

INSERT INTO store (title, photo, price) VALUES ('Some Title', './path/to/photo.jpg', 123.45);
INSERT INTO store (title, photo, price) VALUES ('Some Title', './path/to/photo.jpg', 345.67);
