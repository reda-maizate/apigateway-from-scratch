CREATE TABLE IF NOT EXISTS Notes
(
    id      SERIAL PRIMARY KEY,
    title   VARCHAR(255) NOT NULL,
    content TEXT         NOT NULL
);

CREATE TABLE IF NOT EXISTS Users
(
    id         SERIAL PRIMARY KEY,
    email      VARCHAR(255) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    auth_token VARCHAR(255) NOT NULL
);

INSERT INTO Users (email, password, auth_token)
VALUES ('test@test.com', 'test_password', 'test_token');


INSERT INTO Notes (title, content)
VALUES ('test_title', 'test_content');
