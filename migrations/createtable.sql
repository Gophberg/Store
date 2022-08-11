DROP TABLE IF EXISTS store;

CREATE TABLE IF NOT EXISTS store (
    id bigserial not null primary key,
    title text,
    content text,
    photo text[],
    price numeric(10, 2),
    createdate timestamp
);

INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Title', 'Some content', '{"./path/to/photo.jpg", "./some/one.jpg"}', 123.45, TIMESTAMP '2022-08-05 19:32:00');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 525.65, TIMESTAMP '2022-07-05 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 625.65, TIMESTAMP '2022-07-01 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 725.65, TIMESTAMP '2022-07-02 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 825.65, TIMESTAMP '2022-07-03 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 925.65, TIMESTAMP '2022-07-04 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 425.65, TIMESTAMP '2022-07-05 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 325.65, TIMESTAMP '2022-07-06 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 225.65, TIMESTAMP '2022-07-07 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 125.65, TIMESTAMP '2022-07-08 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 25.65, TIMESTAMP '2022-07-09 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 15.65, TIMESTAMP '2022-07-10 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 5.65, TIMESTAMP '2022-07-11 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 35.65, TIMESTAMP '2022-07-12 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 45.65, TIMESTAMP '2022-07-13 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 55.65, TIMESTAMP '2022-07-14 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 65.65, TIMESTAMP '2022-07-15 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 75.65, TIMESTAMP '2022-07-16 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 85.65, TIMESTAMP '2022-07-17 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 1.65, TIMESTAMP '2022-07-18 19:33:59');
INSERT INTO store (title, content, photo, price, createdate) VALUES ('Some Another Title', 'Another awesome content', '{"./path/to/photo2.jpg", "./another/photo.jpg", "./new/one.jpg"}', 2.65, TIMESTAMP '2022-07-19 19:33:59');
