package repositories

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/models/user"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
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

func TestRegisterSuccess(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewUserRepository(db)

	email := "test1@gmail.com"
	password := "1234555"
	rolesJSON := datatypes.JSON([]byte(`["user","editor"]`))
	permissionsJSON := datatypes.JSON([]byte(`["read", "write"]`))

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).
    WithArgs(
        sqlmock.AnyArg(), // email
        sqlmock.AnyArg(), // password_hash
        sqlmock.AnyArg(), // is_active
        sqlmock.AnyArg(), // create_at
        sqlmock.AnyArg(), // roles
        sqlmock.AnyArg(), // permissions
        sqlmock.AnyArg(), // refresh_token
        sqlmock.AnyArg(), // id
    ).
    WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	user := &user.Users{
		UserID:     1,
		Email:      email,
		Password:   password,
		Role:       rolesJSON,
		Permission: permissionsJSON,
		IsActive:   false,
	}

	// Call the Register method
	err := repo.Register(context.Background(), user)
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestRegister_EmailAlreadyExists(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewUserRepository(db)

	email := "test1@gmail.com"
	password := "1234555"
	rolesJSON := datatypes.JSON([]byte(`["user"]`))
	permissionsJSON := datatypes.JSON([]byte(`["read"]`))

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).
		WithArgs(
			sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
			sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		).
		WillReturnError(errors.New("duplicate key value violates unique constraint"))
	mock.ExpectRollback()

	user := &user.Users{
		UserID:     1,
		Email:      email,
		Password:   password,
		Role:       rolesJSON,
		Permission: permissionsJSON,
		IsActive:   false,
	}

	err := repo.Register(context.Background(), user)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate")
}

func TestRegister_NilUser(t *testing.T) {
	db, _ := setupMockDB(t)
	repo := NewUserRepository(db)

	err := repo.Register(context.Background(), nil)
	assert.Error(t, err)
}

