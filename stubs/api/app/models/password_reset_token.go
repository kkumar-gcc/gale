package models

type PasswordResetToken struct{}

func (m *PasswordResetToken) Stub() string {
	return `package models

import "github.com/goravel/framework/database/orm"

type PasswordResetToken struct {
	Email string ` + "`gorm:\"primaryKey\"`" + `
	Token string
	orm.Timestamps
}
`
}
