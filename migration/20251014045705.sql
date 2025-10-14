-- Set comment to schema: "public"
COMMENT ON SCHEMA "public" IS 'A standard public schema';
-- Add new schema named "private"
CREATE SCHEMA "private";
-- Create enum type "card_status"
CREATE TYPE "public"."card_status" AS ENUM ('requested', 'active', 'blocked', 'expired', 'closed');
-- Create enum type "card_type"
CREATE TYPE "public"."card_type" AS ENUM ('gold', 'diamond', 'platinum');
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" integer NOT NULL,
  "name" character varying(64) NOT NULL,
  "phone_number" character varying(16) NOT NULL,
  "password_hash" character varying(64) NOT NULL,
  "password_salt" character varying(64) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_phone_number" to table: "users"
CREATE INDEX "idx_phone_number" ON "public"."users" ("phone_number");
-- Set comment to column: "id" on table: "users"
COMMENT ON COLUMN "public"."users"."id" IS 'The ID of the user';
-- Set comment to column: "name" on table: "users"
COMMENT ON COLUMN "public"."users"."name" IS 'The full name of the user';
-- Set comment to column: "phone_number" on table: "users"
COMMENT ON COLUMN "public"."users"."phone_number" IS 'The phone number of the user';
-- Set comment to column: "password_hash" on table: "users"
COMMENT ON COLUMN "public"."users"."password_hash" IS 'The hash value of the password';
-- Set comment to column: "password_salt" on table: "users"
COMMENT ON COLUMN "public"."users"."password_salt" IS 'The salt value of the password';
-- Set comment to column: "created_at" on table: "users"
COMMENT ON COLUMN "public"."users"."created_at" IS 'The time when the user was first created';
-- Set comment to column: "updated_at" on table: "users"
COMMENT ON COLUMN "public"."users"."updated_at" IS 'The time when the user was last updated';
-- Create "cards" table
CREATE TABLE "public"."cards" (
  "id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "type" "public"."card_type" NOT NULL,
  "credit" integer NOT NULL,
  "debit" integer NOT NULL,
  "expiration_date" timestamp NOT NULL,
  "status" "public"."card_status" NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_expiration_date" to table: "cards"
CREATE INDEX "idx_expiration_date" ON "public"."cards" ("expiration_date");
-- Create index "idx_user_id" to table: "cards"
CREATE INDEX "idx_user_id" ON "public"."cards" ("user_id");
-- Set comment to column: "id" on table: "cards"
COMMENT ON COLUMN "public"."cards"."id" IS 'The ID of the card';
-- Set comment to column: "user_id" on table: "cards"
COMMENT ON COLUMN "public"."cards"."user_id" IS 'The ID of the owner';
-- Set comment to column: "type" on table: "cards"
COMMENT ON COLUMN "public"."cards"."type" IS 'The type of the card';
-- Set comment to column: "credit" on table: "cards"
COMMENT ON COLUMN "public"."cards"."credit" IS 'The credit of the card';
-- Set comment to column: "debit" on table: "cards"
COMMENT ON COLUMN "public"."cards"."debit" IS 'The debit of the card';
-- Set comment to column: "expiration_date" on table: "cards"
COMMENT ON COLUMN "public"."cards"."expiration_date" IS 'The expiration date of the card';
-- Set comment to column: "status" on table: "cards"
COMMENT ON COLUMN "public"."cards"."status" IS 'The status of the card (requested, active, blocked, ...)';
-- Set comment to column: "created_at" on table: "cards"
COMMENT ON COLUMN "public"."cards"."created_at" IS 'The time when the card was first created';
-- Set comment to column: "updated_at" on table: "cards"
COMMENT ON COLUMN "public"."cards"."updated_at" IS 'The time when the card was last updated';
