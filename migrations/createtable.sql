DROP TABLE IF EXISTS store;

CREATE TABLE IF NOT EXISTS store (
    id bigserial not null primary key,
    title text,
    content text,
    photo text,
    price numeric(10, 2)
);

INSERT INTO store (title, content, photo, price) VALUES ('Some Title','Some content' ,'./path/to/photo.jpg', 123.45);
INSERT INTO store (title, content, photo, price) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 525.65);
