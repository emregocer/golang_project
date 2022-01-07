CREATE TABLE "categories" (
  "id" serial PRIMARY KEY,
  "name" varchar(32) UNIQUE NOT NULL,
  "description" varchar(128) NOT NULL,
  "is_deleted" boolean DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);
