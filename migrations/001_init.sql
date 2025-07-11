-- ===========================
-- Table: books
-- ===========================
CREATE TABLE IF NOT EXISTS books (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  author TEXT NOT NULL,
  isbn TEXT UNIQUE NOT NULL,
  quantity INT NOT NULL DEFAULT 0,
  category TEXT
);

-- ===========================
-- Table: lendings
-- ===========================
CREATE TABLE IF NOT EXISTS lendings (
  id SERIAL PRIMARY KEY,
  book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
  borrower TEXT NOT NULL,
  borrow_date DATE NOT NULL,
  return_date DATE
);


CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL
);
