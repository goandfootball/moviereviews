CREATE TABLE IF NOT EXISTS users (
    usr_id          serial              NOT NULL,
    usr_first_name  VARCHAR(50),
    usr_last_name   VARCHAR(50),
    usr_username    VARCHAR(50)         UNIQUE,
    usr_email       VARCHAR(50)         UNIQUE,
    usr_password    VARCHAR(256)        NOT NULL,
    usr_picture     CHARACTER VARYING,
    usr_created_at  DATE                NOT NULL,
    usr_updated_at  DATE,
    CONSTRAINT usr_pk
        PRIMARY KEY (usr_id)
);
