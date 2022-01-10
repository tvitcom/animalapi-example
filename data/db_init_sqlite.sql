-- Database: `grocery`

PRAGMA foreign_keys=on;
PRAGMA main.secure_delete=on;
PRAGMA journal_mode = PERSIST;

-- Table structure for table `animal`

CREATE TABLE "animal" (
  "id"  INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "name"  TEXT NOT NULL,
  "dob" INTEGER DEFAULT 0,
  "owner" TEXT NOT NULL,
  "kind"  TEXT NOT NULL
);
