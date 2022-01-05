DROP TABLE IF EXISTS "breweries" CASCADE;
DROP TABLE IF EXISTS "beers" CASCADE;
DROP TABLE IF EXISTS "users" CASCADE;
DROP TABLE IF EXISTS "profiles" CASCADE;
DROP TABLE IF EXISTS "beer_comments" CASCADE;
DROP TABLE IF EXISTS "brewery_comments" CASCADE;
DROP TABLE IF EXISTS "beer_tags" CASCADE;
DROP TABLE IF EXISTS "brewery_tags" CASCADE;

DROP TYPE IF EXISTS "vote_tag";
DROP TYPE IF EXISTS "beer_style";
DROP TYPE IF EXISTS "user_role";

-- CreateEnum
CREATE TYPE "vote_tag" AS ENUM ('UPVOTE', 'DOWNVOTE', 'COMMENT', 'CREATE', 'UPDATE', 'REMOVE');

-- CreateEnum
CREATE TYPE "beer_style" AS ENUM ('LAGER', 'PALE_ALE', 'IPA', 'DOUBLE_IPA', 'TRIPLE_IPA', 'AMBER_ALE', 'STOUT');

-- CreateEnum
CREATE TYPE "user_role" AS ENUM ('DRINKING_BUDDY', 'BARTENDER', 'BREWMASTER');


-- CreateTable
CREATE TABLE "breweries"
(
    "id"         TEXT          NOT NULL,
    "created_at" TIMESTAMP(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3)  NOT NULL,
    "image_url"  VARCHAR(1024) NOT NULL,
    "name"       VARCHAR(255)  NOT NULL,
    "street_address"           TEXT          NOT NULL,
    "street_address_extended"  TEXT,
    "city"                     VARCHAR(255)  NOT NULL,
    "state"                    VARCHAR(2)    NOT NULL,
    "zip_code"                 VARCHAR(5)    NOT NULL,
    "zip_code_extension"       VARCHAR(5),

    CONSTRAINT "breweries_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "beers"
(
    "id"         TEXT            NOT NULL,
    "created_at" TIMESTAMP(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3)    NOT NULL,
    "image_url"  VARCHAR(1024)   NOT NULL,
    "name"       VARCHAR(255)    NOT NULL,
    "ibu"        INTEGER         NOT NULL,
    "abv"        DECIMAL(65, 30) NOT NULL,
    "brewery_id" TEXT            NOT NULL,
    "style"      "beer_style"    NOT NULL,

    CONSTRAINT "beers_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "users"
(
    "id"         TEXT         NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "email"      VARCHAR(32)  NOT NULL,
    "__"   VARCHAR      NOT NULL,
    "salt"       VARCHAR(255) NOT NULL,
    "password"   VARCHAR(255) NOT NULL,
    "verified"   BOOLEAN      NOT NULL,
    "role"       "user_role"  NOT NULL,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "profiles"
(
    "id"         TEXT         NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "user_id"    TEXT         NOT NULL,

    CONSTRAINT "profiles_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "beer_comments"
(
    "id"         TEXT         NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "text"       VARCHAR(512) NOT NULL,
    "beer_id"    TEXT         NOT NULL,
    "profile_id" TEXT         NOT NULL,

    CONSTRAINT "beer_comments_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "brewery_comments"
(
    "id"         TEXT         NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "text"       VARCHAR(512) NOT NULL,
    "brewery_id" TEXT         NOT NULL,
    "profile_id" TEXT         NOT NULL,

    CONSTRAINT "brewery_comments_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "beer_tags"
(
    "id"         TEXT         NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "tag"        "vote_tag"   NOT NULL,
    "beer_id"    TEXT         NOT NULL,
    "profile_id" TEXT         NOT NULL,

    CONSTRAINT "beer_tags_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "brewery_tags"
(
    "id"         TEXT         NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "tag"        "vote_tag"   NOT NULL,
    "beer_id"    TEXT         NOT NULL,
    "profile_id" TEXT         NOT NULL,

    CONSTRAINT "brewery_tags_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "profiles_user_id_key" ON "profiles" ("user_id");

-- AddForeignKey
ALTER TABLE "beers"
    ADD CONSTRAINT "beers_brewery_id_fkey" FOREIGN KEY ("brewery_id") REFERENCES "breweries" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "profiles"
    ADD CONSTRAINT "profiles_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "beer_comments"
    ADD CONSTRAINT "beer_comments_beer_id_fkey" FOREIGN KEY ("beer_id") REFERENCES "beers" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "beer_comments"
    ADD CONSTRAINT "beer_comments_profile_id_fkey" FOREIGN KEY ("profile_id") REFERENCES "profiles" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "brewery_comments"
    ADD CONSTRAINT "brewery_comments_brewery_id_fkey" FOREIGN KEY ("brewery_id") REFERENCES "breweries" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "brewery_comments"
    ADD CONSTRAINT "brewery_comments_profile_id_fkey" FOREIGN KEY ("profile_id") REFERENCES "profiles" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "beer_tags"
    ADD CONSTRAINT "beer_tags_beer_id_fkey" FOREIGN KEY ("beer_id") REFERENCES "beers" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "beer_tags"
    ADD CONSTRAINT "beer_tags_profile_id_fkey" FOREIGN KEY ("profile_id") REFERENCES "profiles" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "brewery_tags"
    ADD CONSTRAINT "brewery_tags_beer_id_fkey" FOREIGN KEY ("beer_id") REFERENCES "breweries" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "brewery_tags"
    ADD CONSTRAINT "brewery_tags_profile_id_fkey" FOREIGN KEY ("profile_id") REFERENCES "profiles" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;


-- Seed Data
DO
$$
    DECLARE
        current_datetime timestamp := current_timestamp;
    BEGIN
        INSERT INTO public.breweries (id, created_at, updated_at, image_url, name, street_address, street_address_extended, city, state, zip_code, zip_code_extension)
        VALUES (gen_random_uuid(), current_datetime, current_datetime, 'asdf', 'Fall River Brewery', '1707 Whistling Dr', '', 'Redding', 'CA', '96003', '');
    END;
$$;
COMMIT