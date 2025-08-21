package repositories

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T)(*gorm.DB, sqlmock.Sqlmock){
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dialector := postgres.New(postgres.Config{Conn: db})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	return gormDB, mock
}

func TestLoginWithEmailSuccess(t *testing.T) {
	db, mock := setupMockDB(t)
	userRepo := NewUserRepository(db)

	email := "test@example.com"
	rows := sqlmock.NewRows([]string{"user_id", "email", "password", "role", "permission"}).
		AddRow(1, email, "hashedpass", `["admin"]`, `["read"]`)


	mock.ExpectQuery(`SELECT .* FROM "users" WHERE email = .*`).
		WithArgs(email, 1).
		WillReturnRows(rows)

	user, err := userRepo.Login(context.Background(), email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
}

func TestLogin_NotFound(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewUserRepository(db)

	email := "notfound@example.com"
	mock.ExpectQuery(`SELECT .* FROM "users" WHERE email = .*`).
		WithArgs(email).
		WillReturnError(gorm.ErrRecordNotFound)

	userResult, err := repo.Login(context.Background(), email)
	assert.Error(t, err)
	assert.Nil(t, userResult)
}
