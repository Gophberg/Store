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
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 625.65, TIMESTAMP '2022-08-01 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 725.65, TIMESTAMP '2022-08-02 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 825.65, TIMESTAMP '2022-08-03 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 925.65, TIMESTAMP '2022-08-04 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 425.65, TIMESTAMP '2022-08-05 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 325.65, TIMESTAMP '2022-08-06 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 225.65, TIMESTAMP '2022-08-07 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 125.65, TIMESTAMP '2022-08-08 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 25.65, TIMESTAMP '2022-08-09 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 15.65, TIMESTAMP '2022-08-10 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 5.65, TIMESTAMP '2022-08-11 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 35.65, TIMESTAMP '2022-08-12 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 45.65, TIMESTAMP '2022-08-13 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 55.65, TIMESTAMP '2022-08-14 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 65.65, TIMESTAMP '2022-08-15 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 75.65, TIMESTAMP '2022-08-16 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 85.65, TIMESTAMP '2022-08-17 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 1.65, TIMESTAMP '2022-08-18 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title','Another awesome content' ,'./path/to/photo2.jpg', 2.65, TIMESTAMP '2022-08-19 19:33:59');
