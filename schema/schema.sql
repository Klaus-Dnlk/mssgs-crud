
-- CREATE TABLE IF NOT EXISTS public.messages
-- (
--     id bigint NOT NULL DEFAULT nextval('messages_id_seq'::regclass),
--     message text COLLATE pg_catalog."default",
--     sender_id text,
--     recipient_id text,
--     CONSTRAINT messages_pkey PRIMARY KEY (id),
--     CONSTRAINT recipient_fk FOREIGN KEY (recipient_id)
--         REFERENCES public.users (name) MATCH SIMPLE
--         ON UPDATE NO ACTION
--         ON DELETE NO ACTION
--         NOT VALID,
--     CONSTRAINT sender_fk FOREIGN KEY (sender_id)
--         REFERENCES public.users (name) MATCH SIMPLE
--         ON UPDATE NO ACTION
--         ON DELETE NO ACTION
--         NOT VALID
-- );

-- CREATE TABLE IF NOT EXISTS public.users
-- (
--     id bigint NOT NULL DEFAULT nextval('users_id_seq'::regclass),
--     name text COLLATE pg_catalog."default",
--     CONSTRAINT users_pkey PRIMARY KEY (id),
--     CONSTRAINT idx_users_name UNIQUE (name),
--     CONSTRAINT uk_name UNIQUE (name)
--         INCLUDE(name)
-- )

CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null unique
);

CREATE TABLE messages
(
    id serial not null unique,
    sender_id varchar(255) not null references users(name),
    receiver_id varchar(255) not null references users(name),
    message varchar(255) not null
);