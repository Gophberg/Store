DROP TABLE IF EXISTS store;

CREATE TABLE IF NOT EXISTS store (
    id bigserial not null primary key,
    title text,
    content text,
    photo text,
    price numeric(10, 2),
    createdate timestamp
);

INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Title','Some content' ,'./path/to/photo.jpg', 123.45, TIMESTAMP '2022-08-05 19:32:00');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 525.65, TIMESTAMP '2022-08-05 19:33:59');
