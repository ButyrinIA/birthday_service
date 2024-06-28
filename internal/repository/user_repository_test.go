package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"rutube/internal/repository"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "email", "birthday", "is_subscribed"}).
		AddRow(1, "user1", "user1@example.com", "2024-06-19", true).
		AddRow(2, "user2", "user2@example.com", "2024-06-19", false)

	mock.ExpectQuery("SELECT id, username, email, birthday, is_subscribed FROM users").WillReturnRows(rows)

	userRepo := repository.NewUserRepository(db)
	users, err := userRepo.GetAllUsers()

	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, users[0].Username, "user1")
	assert.Equal(t, users[1].Username, "user2")
}

// Additional tests for other methods can be added here
