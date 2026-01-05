CREATE TABLE users
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    login VARCHAR(32) NOT NULL,
    password VARCHAR(256) NOT NULL,
    fio VARCHAR(64) NOT NULL,
    birthday VARCHAR(16) DEFAULT "",
    data_reg VARCHAR(16) NOT NULL
);

CREATE INDEX users_login ON users (login);
CREATE INDEX users_fio ON users (fio);

CREATE TABLE accounts
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account VARCHAR(32) NOT NULL DEFAULT "",
    balance FLOAT NOT NULL DEFAULT 0.00
);

CREATE INDEX account_balance ON accounts (balance);

CREATE TABLE transactions
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_sender INTEGER NOT NULL,
    name VARCHAR(128) NOT NULL "",
    amount FLOAT NOT NULL DEFAULT 0.00,
    user_recipient INTEGER NOT NULL,
    date_time VARCHAR(32) NOT NULL DEFAULT "",
);

CREATE INDEX name_index ON transactions (name);
CREATE INDEX amount_index ON transactions (amount);

CREATE TABLE users_accounts
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    account_id INTEGER NOT NULL,
);

CREATE INDEX user_index ON users_accounts (user_index)

CREATE TABLE users_trans
(
    user_id INTEGER NOT NULL,
    trans_id INTEGER NOT NULL,
);

CREATE INDEX trans_index ON users_trans (trans_id);
CREATE INDEX user_trans_index ON users_trans (user_id);