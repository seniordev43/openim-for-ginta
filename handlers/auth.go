package handlers

import (
	"fmt"
	"net/http"
	"openim/models"
	"openim/storage"
	"openim/utils"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if storage.UserExists(email) {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	hashedPassword := utils.HashPassword(password)
	user := models.User{
		Email:    email,
		Password: hashedPassword,
	}

	storage.SaveUser(user)
	fmt.Fprintf(w, "User registered successfully")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := storage.GetUser(email)
	if err != nil || !utils.CheckPasswordHash(password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.CreateJWTToken(email)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})

	fmt.Fprintf(w, "Login successful")
}
