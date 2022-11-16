package utilities

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Email string `json:"email"`
}

func GenerateUsername() (string, error) {
	var user User
	resp, err := http.Get("https://random-data-api.com/api/v2/users")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&user)

	return user.Email, nil
}
