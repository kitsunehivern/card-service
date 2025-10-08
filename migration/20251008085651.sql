-- Set comment to schema: "public"
COMMENT ON SCHEMA "public" IS 'A standard public schema';
-- Add new schema named "private"
CREATE SCHEMA "private";
-- Create enum type "card_status"
CREATE TYPE "public"."card_status" AS ENUM ('requested', 'active', 'blocked', 'retired', 'closed');
-- Create "cards" table
CREATE TABLE "public"."cards" (
  "id" uuid NOT NULL,
  "user_id" character varying(64) NOT NULL,
  "credit" integer NOT NULL,
  "debit" integer NOT NULL,
  "status" "public"."card_status" NOT NULL,
  "updated_at" timestamp NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_user_id" to table: "cards"
CREATE INDEX "idx_user_id" ON "public"."cards" ("user_id");
-- Set comment to column: "id" on table: "cards"
COMMENT ON COLUMN "public"."cards"."id" IS 'The ID of the card';
-- Set comment to column: "user_id" on table: "cards"
COMMENT ON COLUMN "public"."cards"."user_id" IS 'The ID of the owner';
-- Set comment to column: "credit" on table: "cards"
COMMENT ON COLUMN "public"."cards"."credit" IS 'The credit of the card';
-- Set comment to column: "debit" on table: "cards"
COMMENT ON COLUMN "public"."cards"."debit" IS 'The debit of the card';
-- Set comment to column: "status" on table: "cards"
COMMENT ON COLUMN "public"."cards"."status" IS 'The status of the card (requested, active, blocked, ...)';
-- Set comment to column: "updated_at" on table: "cards"
COMMENT ON COLUMN "public"."cards"."updated_at" IS 'The last time when the card was updated';
