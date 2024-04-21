CREATE TABLE IF NOT EXISTS products (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT,
  description TEXT,
  product_status TEXT,
  created_at INTEGER NOT NULL,
  updated_at INTEGER NOT NULL,
  reserved_at INTEGER
);

