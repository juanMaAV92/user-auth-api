CREATE TABLE users
(
    id         varchar     NOT NULL,
    email      varchar     NOT NULL,
    password   varchar     NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);