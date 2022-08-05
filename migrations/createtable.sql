-- DROP TABLE IF EXISTS transactions;

CREATE TABLE IF NOT EXISTS store (
    id bigserial not null primary key,
    userid bigserial,
    useremail text,
    amount numeric(10, 2),
    currency text,
    creationdate timestamp,
    updatedate timestamp,
    status text
);

INSERT INTO transactions (userid, useremail, amount, currency, creationdate, updatedate, status) VALUES (1, 'joe@mail.edu', 123.45, 'USD', TIMESTAMP '2022-06-19 21:00:00', TIMESTAMP '2022-06-19 21:00:01', 'New');
INSERT INTO transactions (userid, useremail, amount, currency, creationdate, updatedate, status) VALUES (3, 'jane@mail.edu', 345.67, 'RUR', TIMESTAMP '2022-06-19 21:00:03', TIMESTAMP '2022-06-19 21:00:04', 'Canceled');
