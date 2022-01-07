CREATE TABLE "movies" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(64) NOT NULL,
  "plot" varchar(256) NOT NULL,
  "is_deleted" boolean DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);