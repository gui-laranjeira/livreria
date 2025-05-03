CREATE TABLE "books" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar(100) NOT NULL,
  "publisher_id" INT NOT NULL,
  "pages" INT NOT NULL,
  "language" varchar(3),
  "edition" INT,
  "year" INT,
  "isbn" varchar(20),
  "owner" varchar(50),
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "active" bool
);

CREATE TABLE "authors" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "publishers" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "country" varchar
);

CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "subcategories" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "category_id" INT
);

CREATE TABLE "books_authors" (
  "id" SERIAL PRIMARY KEY,
  "book_id" INT,
  "author_id" INT
);

CREATE TABLE "books_categories" (
  "id" SERIAL PRIMARY KEY,
  "book_id" INT,
  "subcategory_id" INT
);

ALTER TABLE "books" ADD FOREIGN KEY ("publisher_id") REFERENCES "publishers" ("id");

ALTER TABLE "subcategories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "books_authors"
    ADD CONSTRAINT fk_book FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "books_authors"
    ADD CONSTRAINT fk_author FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "books_categories"
    ADD CONSTRAINT fk_book FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "books_categories"
    ADD CONSTRAINT fk_subcategory FOREIGN KEY ("subcategory_id") REFERENCES "subcategories" ("id");