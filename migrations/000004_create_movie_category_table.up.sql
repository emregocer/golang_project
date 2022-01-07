CREATE TABLE "movie_category" (
  "id" bigserial PRIMARY KEY,
  "movie_id" bigint NOT NULL,
  "category_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "movie_category" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
ALTER TABLE "movie_category" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("id");

CREATE INDEX ON "movie_category" ("movie_id", "category_id");