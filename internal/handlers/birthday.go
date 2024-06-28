package handlers

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rutube/internal/auth"
	"rutube/internal/models"
	"rutube/internal/service"
)

type BirthdayHandler struct {
	birthdayService *service.BirthdayService
	authService     *service.AuthService
}

func NewBirthdayHandler(birthdayService *service.BirthdayService, authService *service.AuthService) *BirthdayHandler {
	return &BirthdayHandler{
		birthdayService: birthdayService,
		authService:     authService,
	}
}

func (h *BirthdayHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неправильный запрос", http.StatusBadRequest)
		fmt.Println("Ошибка декодирования JSON:", err)
		return
	}

	birthday, err := user.GetBirthday()
	if err != nil {
		http.Error(w, "Неправильный формат даты", http.StatusBadRequest)
		fmt.Println("Ошибка парсинга даты:", err)
		return
	}

	// Хэшируем пароль перед сохранением
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Ошибка хэширования пароля:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Преобразование строки даты в формат времени
	user.Birthday = birthday.Format("2006-01-02")

	if err := h.authService.RegisterUser(user); err != nil {
		fmt.Println("Ошибка при регистрации пользователя:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BirthdayHandler) GetTodaysBirthdays(w http.ResponseWriter, r *http.Request) {
	users, err := h.birthdayService.GetTodaysBirthdays()
	if err != nil {
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}
	// Вывод списка пользователей в консоль
	for _, user := range users {
		fmt.Printf("Имя: %s, День рождения: %s\n", user.Username, user.Birthday)
	}

	// Возвращаем ответ клиенту в формате JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
	}

}

func (h *BirthdayHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	var request struct {
		UserID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Неправильный запрос", http.StatusBadRequest)
		return
	}

	if err := h.birthdayService.Subscribe(request.UserID); err != nil {
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BirthdayHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	var request struct {
		UserID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Неправильный запрос", http.StatusBadRequest)
		return
	}

	if err := h.birthdayService.Unsubscribe(request.UserID); err != nil {
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (h *BirthdayHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Неправильный запрос", http.StatusBadRequest)
		return
	}

	token, err := auth.AuthenticateUser(h.authService, credentials.Username, credentials.Password)
	if err != nil {
		if err == service.ErrInvalidUsername {
			http.Error(w, "Неверное имя пользователя", http.StatusUnauthorized)
		} else if err == service.ErrInvalidPassword {
			http.Error(w, "Неверный пароль", http.StatusUnauthorized)
		} else {
			http.Error(w, "Неверные учетные данные", http.StatusUnauthorized)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
