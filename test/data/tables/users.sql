CREATE TABLE users (
    usr_id          integer                 NOT NULL DEFAULT nextval('users_usr_id_seq'::regclass),
    usr_first_name  character varying,
    usr_last_name   character varying,
    usr_username    character varying       NOT NULL,
    usr_email       character varying       NOT NULL,
    usr_password    character varying(256)  NOT NULL,
    usr_picture     character varying,
    usr_created_at  date                    NOT NULL,
    usr_updated_at  date,
    CONSTRAINT usr_pk
        PRIMARY KEY (usr_id)
)
