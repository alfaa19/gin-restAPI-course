package controller

import (
	"errors"
	"log"
	"net/http"
	"time"

	config "github.com/alfaa19/gin-restAPI-course/config/mysql"
	"github.com/alfaa19/gin-restAPI-course/helper"
	"github.com/alfaa19/gin-restAPI-course/model"
	"github.com/alfaa19/gin-restAPI-course/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var req request.User

	// Check if datas was binding already
	if c.GetHeader("Content-Type") == "application/json" && c.MustGet("binding_initialized") == nil {
		if err := c.ShouldBindJSON(&req); err != nil {
			helper.ResponseErrorJSON(c, http.StatusBadRequest, err)
			return
		}
		c.Set("binding_initialized", true)
	} else {
		if err := c.ShouldBind(&req); err != nil {
			helper.ResponseErrorJSON(c, http.StatusBadRequest, err)
			return
		}
	}

	var existingUser model.User

	// Check user already exists
	if err := config.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, errors.New("User already exists"))
		return
	}

	// Hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var newUser = model.User{
		Username: req.Username,
		Role:     "user",
		Password: string(hashedPassword),
	}

	// Insert into database
	if err := config.DB.Create(&newUser).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusUnprocessableEntity, err)
		return
	}

	helper.ResponseSuccessJSON(c, "Success registering new user", "")
}

func Login(c *gin.Context) {
	var req request.User

	if c.GetHeader("Content-Type") == "application/json" && c.MustGet("binding_initialized") == nil {
		if err := c.ShouldBindJSON(&req); err != nil {
			helper.ResponseErrorJSON(c, http.StatusBadRequest, err)
			return
		}
		c.Set("binding_initialized", true)
	} else {
		if err := c.ShouldBind(&req); err != nil {
			helper.ResponseErrorJSON(c, http.StatusBadRequest, err)
			return
		}
	}

	if err := req.Validate(); err != nil {
		helper.ResponseValidationErrorJson(c, "Error Validation", err.(validation.Errors))
		return
	}

	var existingUser model.User

	//Get existing user
	if err := config.DB.Where("username = ?", req.Username).First(&existingUser).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, errors.New("User not found"))
		return
	}

	//compare password
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(req.Password)); err != nil {
		helper.ResponseErrorJSON(c, http.StatusUnauthorized, errors.New("Wrong Password"))
		return
	}

	//generate token
	token, err := generateToken(req.Username, existingUser.Role)
	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusUnauthorized, err)
	}

	helper.ResponseSuccessJSON(c, "login succsess", gin.H{"token": token})
}

func generateToken(username, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token valid for 1 hour
	})

	// Generate token using secret key
	secret := "secret-wawawawawaaw"
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	log.Println(role)
	return tokenString, nil
}
