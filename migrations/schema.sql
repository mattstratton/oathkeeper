CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "users" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"email" TEXT,
"provider" TEXT NOT NULL,
"provider_id" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "talks" (
"id" TEXT PRIMARY KEY,
"user_id" char(36) NOT NULL,
"title" TEXT NOT NULL,
"abstract" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "submissions" (
"id" TEXT PRIMARY KEY,
"status" TEXT NOT NULL,
"event_id" char(36) NOT NULL,
"speaker_id" char(36) NOT NULL,
"talk_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "events" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"cfp_open_date" DATETIME NOT NULL,
"cfp_close_date" DATETIME NOT NULL,
"event_start_date" DATETIME NOT NULL,
"event_end_date" DATETIME NOT NULL,
"cfp_url" TEXT NOT NULL,
"event_url" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "speakers" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"bio" TEXT NOT NULL,
"twitter" text,
"github" text,
"homepage" text,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
