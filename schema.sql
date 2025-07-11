-- Users Table (Optional - for extended auth system)
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  name TEXT NOT NULL,
  role TEXT NOT NULL DEFAULT 'admin'
);

-- Books Table
CREATE TABLE IF NOT EXISTS books (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  author TEXT NOT NULL,
  isbn TEXT NOT NULL,
  quantity INT NOT NULL,
  category TEXT NOT NULL
);

-- Lendings Table
CREATE TABLE IF NOT EXISTS lendings (
  id SERIAL PRIMARY KEY,
  book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
  borrower TEXT NOT NULL,
  borrow_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  return_date TIMESTAMP
);

-- Example for future extension (optional)
-- Categories Table
-- CREATE TABLE IF NOT EXISTS categories (
--   id SERIAL PRIMARY KEY,
--   name TEXT NOT NULL
-- );

-- Optional: Book status table
-- CREATE TABLE IF NOT EXISTS book_status (
--   id SERIAL PRIMARY KEY,
--   book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
--   available_qty INT NOT NULL,
--   borrowed_qty INT NOT NULL
-- );
