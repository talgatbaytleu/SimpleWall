package logic

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"auth-service/internal/dal"
	"auth-service/pkg/models"
)

var urlTestDB = "postgres://tbaitleu:talgat9595@localhost:5432/sw_users_auth_test"

func TestMain(m *testing.M) {
	// Подключаемся к тестовой БД
	dal.InitTestDB(urlTestDB)

	// Закрываем БД после выполнения всех тестов

	defer dal.CloseTestDB()

	code := m.Run()
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	userDalTest := dal.NewUserDal(dal.TestDB)
	userLogicTest := NewUserLogic(userDalTest)

	defer userDalTest.TruncateAllUsers()

	tests := []struct {
		name        string
		user        models.User
		expectError bool
	}{
		{
			name: "Valid user",
			user: models.User{
				Username: "validuser",
				Password: "ValidP@ss123!",
			},
			expectError: false,
		},
		{
			name: "Duplicate username",
			user: models.User{
				Username: "validuser", // Используем того же user'а, которого только что создали
				Password: "AnotherP@ss123!",
			},
			expectError: true,
		},
		{
			name: "Weak password",
			user: models.User{
				Username: "weakuser",
				Password: "12345",
			},
			expectError: true,
		},
		{
			name: "Empty password",
			user: models.User{
				Username: "emptyuser",
				Password: "",
			},
			expectError: true,
		},
		{
			name: "Invalid JSON",
			user: models.User{
				Username: "",
				Password: "",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.user)
			err := userLogicTest.CreateUser(bytes.NewReader(body))

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got nil for case: %s", tt.name)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v for case: %s", err, tt.name)
				}
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	userDalTest := dal.NewUserDal(dal.TestDB)
	userLogicTest := NewUserLogic(userDalTest)

	defer userDalTest.TruncateAllUsers()
	// Создаем пользователя
	testUser := models.User{
		Username: "testuser",
		Password: "SecurePass123!",
	}

	createBody, _ := json.Marshal(testUser)
	err := userLogicTest.CreateUser(bytes.NewReader(createBody))
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	tests := []struct {
		name        string
		user        models.User
		expectError bool
	}{
		{
			name: "Valid login",
			user: models.User{
				Username: "testuser",
				Password: "SecurePass123!",
			},
			expectError: false,
		},
		{
			name: "Invalid password",
			user: models.User{
				Username: "testuser",
				Password: "WrongPassword!",
			},
			expectError: true,
		},
		{
			name: "Non-existent user",
			user: models.User{
				Username: "nouser",
				Password: "SomePassword123!",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loginBody, _ := json.Marshal(tt.user)
			token, err := userLogicTest.LoginUser(bytes.NewReader(loginBody))

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got nil for case: %s", tt.name)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v for case: %s", err, tt.name)
				}
				if token == "" {
					t.Errorf("Expected JWT token but got empty string")
				}
			}
		})
	}
}
