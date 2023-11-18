CREATE TABLE IF NOT EXISTS Notes
(
    uuid    TEXT PRIMARY KEY,
    title   VARCHAR(255) NOT NULL,
    content TEXT         NOT NULL
);

CREATE TABLE IF NOT EXISTS Users
(
    uuid       TEXT PRIMARY KEY,
    email      VARCHAR(255) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    auth_token VARCHAR(255) NOT NULL
);

INSERT INTO Users (uuid, email, password, auth_token)
VALUES ('fd1c3669-50c2-4f6e-adec-525bb27c9f77', 'test@test.com', 'test_password',
        'ADWuByBkAa7Mj1K5TYM7WC75FIEBpO9P4lW1ZzaKWxSjEwd4');


INSERT INTO Notes (uuid, title, content)
VALUES ('a573a8ab-a6d2-4107-a504-8d9f6f76c016', 'test_title', 'test_content');
