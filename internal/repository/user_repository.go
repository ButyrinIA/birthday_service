package repository

import (
	"database/sql"
	"rutube/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT id, username, email, birthday, is_subscribed FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Birthday, &user.IsSubscribed); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) CreateUser(user models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (username, password, email, birthday, is_subscribed) VALUES ($1, $2, $3, $4, $5)",
		user.Username, user.Password, user.Email, user.Birthday, user.IsSubscribed)
	return err
}

func (r *UserRepository) Subscribe(userID int) error {
	_, err := r.DB.Exec("UPDATE users SET is_subscribed = true WHERE id = $1", userID)
	return err
}

func (r *UserRepository) Unsubscribe(userID int) error {
	_, err := r.DB.Exec("UPDATE users SET is_subscribed = false WHERE id = $1", userID)
	return err
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	row := r.DB.QueryRow("SELECT id, username, password, email, birthday, is_subscribed FROM users WHERE username = $1", username)

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Birthday, &user.IsSubscribed); err != nil {
		return nil, err
	}
	return &user, nil
}
