package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"rutube/internal/models"
)

func TestUserRepository_GetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "email", "birthday", "is_subscribed"}).
		AddRow(1, "user1", "user1@example.com", "2000-01-01", true).
		AddRow(2, "user2", "user2@example.com", "2000-02-02", false)

	mock.ExpectQuery("SELECT id, username, email, birthday, is_subscribed FROM users").
		WillReturnRows(rows)

	repo := NewUserRepository(db)
	users, err := repo.GetAllUsers()

	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "user1", users[0].Username)
	assert.Equal(t, "user2", users[1].Username)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	user := models.User{
		Username:     "user1",
		Password:     "password",
		Email:        "user1@example.com",
		Birthday:     "2000-01-01",
		IsSubscribed: true,
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.Username, user.Password, user.Email, user.Birthday, user.IsSubscribed).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewUserRepository(db)
	err = repo.CreateUser(user)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Subscribe(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userID := 1

	mock.ExpectExec("UPDATE users SET is_subscribed = true WHERE id = \\$1").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewUserRepository(db)
	err = repo.Subscribe(userID)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Unsubscribe(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userID := 1

	mock.ExpectExec("UPDATE users SET is_subscribed = false WHERE id = \\$1").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewUserRepository(db)
	err = repo.Unsubscribe(userID)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	username := "user1"
	row := sqlmock.NewRows([]string{"id", "username", "password", "email", "birthday", "is_subscribed"}).
		AddRow(1, "user1", "password", "user1@example.com", "2000-01-01", true)

	mock.ExpectQuery("SELECT id, username, password, email, birthday, is_subscribed FROM users WHERE username = \\$1").
		WithArgs(username).
		WillReturnRows(row)

	repo := NewUserRepository(db)
	user, err := repo.GetUserByUsername(username)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "user1", user.Username)
	assert.NoError(t, mock.ExpectationsWereMet())
}
