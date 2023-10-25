CREATE TYPE "users_role_enum" AS ENUM ('admin', 'basic-user');

CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "role" users_role_enum DEFAULT 'basic-user',
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX "email" ON "users" ("email");

INSERT INTO users (name, email, password, role)
VALUES
    ('admin', 'admin@localhost', '$2a$10$OicMfEavvx92wZl.uB8K6e9AO/94SvMBUc5DFjhRq5xy4oLdAfbB2', 'admin');
