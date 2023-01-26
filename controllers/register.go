package controllers

import (
	"crypto/rand"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/middlewares"
	"itfest-backend-2.0/models"
	"itfest-backend-2.0/types"
)

type UserRequest struct {
	Username string `json:"username" form:"username" query:"username"`
	Name     string `json:"name" form:"name" query:"name"`
	Password string `json:"password" form:"password" query:"password"`
}

func generateUserCode() string {
	codes := make([]byte, 6)
	if _, err := rand.Read(codes); err != nil {
		return ""
	}

	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	return string(codes)
}

func Testing(c echo.Context) error {
	return c.JSON(http.StatusOK, "HAI")
}

func RegisterHandler(c echo.Context) error {
	db := configs.DB.GetConnection()
	config := configs.Config.GetMetadata()
	response := models.Response[string]{}

	request := UserRequest{}
	if err := c.Bind(&request); err != nil {
		response.Message = "ERROR: BAD REQUEST"
		return c.JSON(http.StatusBadRequest, response)
	}

	// Search if username existed
	user := models.User{}
	username := request.Username
	condition := models.User{Username: username}
	if err := db.Where(&condition).Find(&user).Error; err != nil {
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}
	if user.Username != "" {
		response.Message = "ERROR: USERNAME EXISTED"
		return c.JSON(http.StatusBadRequest, response)
	}

	// Create new usercode
	userCode := generateUserCode()
	condition = models.User{Usercode: userCode}
	if err := db.Where(&condition).Find(&user).Error; err != nil {
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}

	// Loop until usercode is unique
	for user.Username != "" {
		userCode := generateUserCode()
		condition = models.User{Usercode: userCode}
		if err := db.Where(&condition).Find(&user).Error; err != nil {
			response.Message = "ERROR: INTERNAL SERVER ERROR"
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	// Create User
	encryptedString := []byte(request.Password)
	newUser := models.User{
		Username: request.Username,
		Name:     request.Name,
		Usercode: userCode,
		Password: encryptedString,
		Role:     types.User,
		Point:    0,
	}
	if err := db.Create(&newUser).Error; err != nil {
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}

	authClaims := middlewares.AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.ApplicationName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.LoginExpirationDuration)),
		},
		ID: newUser.ID,
	}

	unsignedAuthToken := jwt.NewWithClaims(config.JWTSigningMethod, authClaims)
	signedAuthToken, err := unsignedAuthToken.SignedString(config.JWTSignatureKey)
	if err != nil {
		response.Message = "ERROR: JWT SIGNING ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "SUCCESS"
	response.Data = signedAuthToken
	return c.JSON(http.StatusCreated, response)
}
