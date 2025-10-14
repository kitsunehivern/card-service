-- Modify "users" table
ALTER TABLE "public"."users" DROP COLUMN "password_salt";
-- Set comment to column: "password_hash" on table: "users"
COMMENT ON COLUMN "public"."users"."password_hash" IS 'The hash value (+ salt value) of the password';
