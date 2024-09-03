package migrations

type SqliteUserDownStub struct{}

func (r *SqliteUserDownStub) Stub() string {
	return `DROP TABLE IF EXISTS users;`
}

type SqliteUserUpStub struct{}

func (r *SqliteUserUpStub) Stub() string {
	return `CREATE TABLE users (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   name TEXT NOT NULL,
   email TEXT NOT NULL UNIQUE,
   password TEXT NOT NULL,
   email_verified_at DATETIME DEFAULT NULL,
   created_at DATETIME NOT NULL,
   updated_at DATETIME NOT NULL,
   deleted_at DATETIME DEFAULT NULL
);`
}

type SqlitePasswordResetDownStub struct{}

func (r *SqlitePasswordResetDownStub) Stub() string {
	return `DROP TABLE IF EXISTS password_reset_tokens;`
}

type SqlitePasswordResetUpStub struct{}

func (r *SqlitePasswordResetUpStub) Stub() string {
	return `CREATE TABLE password_reset_tokens (
   email TEXT PRIMARY KEY,
   token TEXT NOT NULL,
   created_at DATETIME NOT NULL,
   updated_at DATETIME NOT NULL
);`
}
