CREATE TABLE IF NOT EXISTS Users
(
    uuid       TEXT PRIMARY KEY,
    email      VARCHAR(255) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    auth_token VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Notes
(
    uuid       TEXT PRIMARY KEY,
    title      VARCHAR(255) NOT NULL,
    content    TEXT         NOT NULL,
    created_by TEXT REFERENCES Users (uuid)
);

CREATE TABLE IF NOT EXISTS Permissions
(
    uuid     TEXT PRIMARY KEY,
    service  VARCHAR(255) NOT NULL,
    resource VARCHAR(255) NOT NULL,
    action   VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS UserPermissions
(
    user_uuid       TEXT REFERENCES Users (uuid),
    permission_uuid TEXT REFERENCES Permissions (uuid),
    PRIMARY KEY (user_uuid, permission_uuid)
);


INSERT INTO Users (uuid, email, password, auth_token)
VALUES ('fd1c3669-50c2-4f6e-adec-525bb27c9f77', 'test@test.com', 'test_password',
        'ADWuByBkAa7Mj1K5TYM7WC75FIEBpO9P4lW1ZzaKWxSjEwd4');


INSERT INTO Notes (uuid, title, content, created_by)
VALUES ('a573a8ab-a6d2-4107-a504-8d9f6f76c016', 'test_title', 'test_content', 'fd1c3669-50c2-4f6e-adec-525bb27c9f77');


INSERT INTO Permissions (uuid, service, resource, action)
VALUES ('a573a8ab-a6d2-4107-a504-8d9f6f76c016', 'notes', 'note', 'read');


INSERT INTO UserPermissions (user_uuid, permission_uuid)
VALUES ('fd1c3669-50c2-4f6e-adec-525bb27c9f77', 'a573a8ab-a6d2-4107-a504-8d9f6f76c016');