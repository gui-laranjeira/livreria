INSERT INTO "publishers" ("id", "name", "country")
VALUES
    (1, 'Secker & Warburg', 'UK'),
    (2, 'Editora Globo', 'Brazil');

INSERT INTO "books" ("id", "title", "publisher_id", "pages", "language", "edition", "year", "isbn", "owner", "created_at", "updated_at", "deleted_at", "active")
VALUES
    (1, '1984', 1, 328, 'EN', 1, 1949, '9780451524935', 'George Orwell', CURRENT_TIMESTAMP, NULL, NULL, TRUE),
    (2, 'Dom Casmurro', 2, 256, 'PT', 1, 1899, '9788520932128', 'Machado de Assis', CURRENT_TIMESTAMP, NULL, NULL, TRUE);

INSERT INTO "authors" ("id", "name")
VALUES
    (1, 'George Orwell'),
    (2, 'Machado de Assis');

INSERT INTO "categories" ("id", "name")
VALUES
    (1, 'Fiction'),
    (2, 'Classic');

INSERT INTO "subcategories" ("id", "name", "category_id")
VALUES
    (1, 'Dystopian', 1),
    (2, 'Brazilian Literature', 2);

INSERT INTO "books_authors" ("id", "book_id", "author_id")
VALUES
    (1, 1, 1),
    (2, 2, 2);

INSERT INTO "books_categories" ("id", "book_id", "subcategory_id")
VALUES
    (1, 1, 1),
    (2, 2, 2);