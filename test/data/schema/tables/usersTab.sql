CREATE TABLE users (
    usr_id          SERIAL                 /*NOT NULL DEFAULT nextval('users_usr_id_seq'::regclass)*/,
    usr_first_name  VARCHAR(50),
    usr_last_name   VARCHAR(50),
    usr_username    VARCHAR(50)         NOT NULL,
    usr_email       VARCHAR(50)         NOT NULL,
    usr_password    VARCHAR(256)        NOT NULL,
    usr_picture     CHARACTER VARYING,
    usr_created_at  DATE                NOT NULL,
    usr_updated_at  DATE,
    CONSTRAINT usr_pk
        PRIMARY KEY (usr_id)
);
