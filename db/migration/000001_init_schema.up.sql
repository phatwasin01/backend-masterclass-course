CREATE TABLE "users" (
  "user_id" varchar PRIMARY KEY,
  "email" varchar NOT NULL,
  "display_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "organizers" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "phone" varchar(10)
);

CREATE TABLE "events" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "organizer_id" bigint NOT NULL,
  "price" int NOT NULL,
  "amount" int NOT NULL,
  "amount_sold" int NOT NULL DEFAULT 0,
  "description" text,
  "is_closed" boolean NOT NULL DEFAULT false,
  "amount_redeem" int NOT NULL DEFAULT 0,
  "start_time" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
  CONSTRAINT sold_check CHECK ("amount_sold" <= "amount")
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "event_id" bigint NOT NULL,
  "amount" int NOT NULL,
  "sum_price" int NOT NULL,
  "payment" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tickets" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "event_id" bigint NOT NULL,
  "order_id" bigint NOT NULL,
  "is_redeemed" boolean NOT NULL DEFAULT false,
  "ticket_uuid" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("user_id");

CREATE INDEX ON "events" ("organizer_id");

CREATE INDEX ON "orders" ("user_id");

CREATE INDEX ON "tickets" ("user_id");

CREATE INDEX ON "tickets" ("ticket_uuid");

CREATE INDEX ON "tickets" ("event_id");

CREATE INDEX ON "tickets" ("order_id");

CREATE INDEX ON "tickets" ("user_id", "event_id");

ALTER TABLE "events" ADD FOREIGN KEY ("organizer_id") REFERENCES "organizers" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
