CREATE TABLE "user_movie" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "movie_id" int NOT NULL,
  "favourited_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "user_movie" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_movie" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("id");

CREATE INDEX ON "user_movie" ("user_id", "movie_id");