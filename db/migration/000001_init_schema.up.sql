CREATE TABLE "todos" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "completed" boolean default False
);