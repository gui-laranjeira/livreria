CREATE TABLE "books" (
  "id" integer PRIMARY KEY,
  "title" varchar(100) NOT NULL,
  "publisher_id" integer NOT NULL,
  "pages" integer NOT NULL,
  "language" varchar(3),
  "edition" integer,
  "year" integer,
  "isbn" varchar(20),
  "owner" varchar(50),
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "active" bool
);

CREATE TABLE "authors" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "publishers" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL,
  "country" varchar
);

CREATE TABLE "categories" (
  "id" integer PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "subcategories" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "category_id" integer
);

CREATE TABLE "books_authors" (
  "book_id" intenger,
  "author_id" integer
);

CREATE TABLE "books_categories" (
  "book_id" integer,
  "subcategory_id" integer
);

ALTER TABLE "books" ADD FOREIGN KEY ("publisher_id") REFERENCES "publishers" ("id");

ALTER TABLE "subcategories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("id") REFERENCES "books_authors" ("book_id");

ALTER TABLE "authors" ADD FOREIGN KEY ("id") REFERENCES "books_authors" ("author_id");

ALTER TABLE "books" ADD FOREIGN KEY ("id") REFERENCES "books_categories" ("book_id");

ALTER TABLE "subcategories" ADD FOREIGN KEY ("id") REFERENCES "books_categories" ("subcategory_id");
