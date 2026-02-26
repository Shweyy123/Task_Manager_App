package handlers

import (
	"encoding/json"
	"net/http"
	"task-manager/config"
	"task-manager/middleware"
	"task-manager/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Check if email already exists
	var existsID int
	err := config.DB.QueryRow("SELECT id FROM users WHERE email=?", user.Email).Scan(&existsID)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"message": "User already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = config.DB.Exec("INSERT INTO users(name,email,password) VALUES(?,?,?)",
		user.Name, user.Email, hashedPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	var dbUser models.User
	err := config.DB.QueryRow("SELECT id, password FROM users WHERE email=?", user.Email).Scan(&dbUser.ID, &dbUser.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := &middleware.Claims{
		UserID: dbUser.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(middleware.JwtKey)

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
