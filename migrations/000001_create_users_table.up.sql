CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(32) UNIQUE NOT NULL,
  "password" varchar(64) NOT NULL,
  "email" varchar(32) UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);
