package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(repo *Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		email := strings.TrimSpace(strings.ToLower(req.Email)) //convert to lowercase and trim spaces
		password := strings.TrimSpace(req.Password)

		if email == "" || !strings.Contains(email, "@") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
			return
		}

		if len(password) < 8 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters long"})
			return
		}

		hash, err := bcrypt.GenerateFromPassword(
			[]byte(password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		err = repo.CreateUser(email, string(hash))
		if err != nil {
			if errors.Is(err, ErrUserAlreadyExists) {
				c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			}
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "user registered successfully",
		})
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(repo *Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		email := strings.TrimSpace(strings.ToLower(req.Email))
		password := strings.TrimSpace(req.Password)

		if email == "" || !strings.Contains(email, "@") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
			return
		}

		if password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
			return
		}

		user, err := repo.GetUserByEmail(email)

		if err != nil {
			if errors.Is(err, ErrUserNotFound) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
			}
			return
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(user.PasswordHash),
			[]byte(password),
		)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":      user.ID,
			"email":   user.Email,
			"role":    user.Role,
			"message": "login successful",
		})
	}
}

type storeRefreshTokenRequest struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
	Expiry time.Time `json:"expiresAt"`
}

func StoreRefreshToken(repo *Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req storeRefreshTokenRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if req.UserID <= 0 || strings.TrimSpace(req.Token) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id or token"})
			return
		}

		if err := repo.StoreRefreshToken(req.UserID, req.Token, req.Expiry); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store refresh token"})
			return
		}

		if req.Expiry.Before(time.Now()) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Expiry time must be in the future"})
			return
		}
		
		c.Status(http.StatusCreated)
	}
}

type validateRefreshTokenRequest struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

func ValidateRefreshToken(repo *Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req validateRefreshTokenRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if req.UserID <= 0 || strings.TrimSpace(req.Token) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id or token"})
			return
		}

		valid, err := repo.IsRefreshTokenValid(req.UserID, req.Token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate refresh token"})
			return
		}

		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"valid": true})
	}
}

type revokeRefreshTokenRequest struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

func RevokeRefreshToken(repo *Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req revokeRefreshTokenRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if req.UserID <= 0 || strings.TrimSpace(req.Token) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id or token"})
			return
		}

		if err := repo.RevokeRefreshToken(req.UserID, req.Token); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revoke refresh token"})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
