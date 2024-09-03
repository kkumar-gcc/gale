package migrations

type MysqlUserDownStub struct{}

func (r *MysqlUserDownStub) Stub() string {
	return `DROP TABLE IF EXISTS users;`
}

type MysqlUserUpStub struct{}

func (r *MysqlUserUpStub) Stub() string {
	return `CREATE TABLE users (
   id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   email VARCHAR(255) NOT NULL UNIQUE,
   password VARCHAR(255) NOT NULL,
   email_verified_at TIMESTAMP NULL DEFAULT NULL,
   created_at DATETIME(3) NOT NULL,
   updated_at DATETIME(3) NOT NULL,
   deleted_at DATETIME(3) NULL DEFAULT NULL,
   KEY idx_users_created_at (created_at),
   KEY idx_users_updated_at (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
}

type MysqlPasswordResetDownStub struct{}

func (r *MysqlPasswordResetDownStub) Stub() string {
	return `DROP TABLE IF EXISTS password_resets;`
}

type MysqlPasswordResetUpStub struct{}

func (r *MysqlPasswordResetUpStub) Stub() string {
	return `CREATE TABLE password_reset_tokens (
   email VARCHAR(255) NOT NULL PRIMARY KEY,
   token VARCHAR(255) NOT NULL,
   created_at DATETIME(3) NOT NULL,
   updated_at DATETIME(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
}
