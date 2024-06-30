package service

import (
	"time"

	"rutube/internal/models"
	"rutube/internal/repository"
)

type BirthdayService struct {
	repo *repository.UserRepository
}

func NewBirthdayService(repo *repository.UserRepository) *BirthdayService {
	return &BirthdayService{repo: repo}
}

func (s *BirthdayService) GetTodaysBirthdays() ([]models.User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var todaysBirthdays []models.User
	today := time.Now()
	for _, user := range users {
		// Проверяет, не является ли Birthday пустой строкой, и преобразуйте ее в time.Time
		if user.Birthday != "" {
			birthday, err := time.Parse("2006-01-02T15:04:05Z", user.Birthday)
			if err != nil {
				// Skip users with invalid birthday format
				continue
			}

			if birthday.Month() == today.Month() && birthday.Day() == today.Day() && user.IsSubscribed {
				todaysBirthdays = append(todaysBirthdays, user)
			}
		}
	}

	//fmt.Printf("Found %d из %d users with birthday today\n", len(todaysBirthdays), len(users))
	return todaysBirthdays, nil
}

func (s *BirthdayService) Subscribe(userID int) error {
	return s.repo.Subscribe(userID)
}

func (s *BirthdayService) Unsubscribe(userID int) error {
	return s.repo.Unsubscribe(userID)
}
