package migrations

type PostgresUserDownStub struct{}

func (r *PostgresUserDownStub) Stub() string {
	return `DROP TABLE IF EXISTS users;`
}

type PostgresUserUpStub struct{}

func (r *PostgresUserUpStub) Stub() string {
	return `CREATE TABLE users (
   id BIGSERIAL PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   email VARCHAR(255) NOT NULL UNIQUE,
   password VARCHAR(255) NOT NULL,
   email_verified_at TIMESTAMP WITH TIME ZONE NULL,
   created_at TIMESTAMPTZ NOT NULL,
   updated_at TIMESTAMPTZ NOT NULL,
   deleted_at TIMESTAMPTZ NULL DEFAULT NULL
);`
}

type PostgresPasswordResetDownStub struct{}

func (r *PostgresPasswordResetDownStub) Stub() string {
	return `DROP TABLE IF EXISTS password_reset_tokens;`
}

type PostgresPasswordResetUpStub struct{}

func (r *PostgresPasswordResetUpStub) Stub() string {
	return `CREATE TABLE password_reset_tokens (
   email VARCHAR(255) PRIMARY KEY,
   token VARCHAR(255) NOT NULL,
   created_at TIMESTAMPTZ NOT NULL,
   updated_at TIMESTAMPTZ NOT NULL
);`
}
