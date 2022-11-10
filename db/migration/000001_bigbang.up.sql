CREATE TABLE "booth" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" varchar(255) NOT NULL,
  "logo_url" varchar(1000) NOT NULL,
  "description" varchar(1000) NOT NULL,
  "twitter_url" varchar(1000) NOT NULL,
  "instagram_url" varchar(1000) NOT NULL
);

CREATE TABLE "rating" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "price" integer NOT NULL,
  "quality" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "booth_id" uuid NOT NULL
);

ALTER TABLE "rating" ADD FOREIGN KEY ("booth_id") REFERENCES "booth" ("id");