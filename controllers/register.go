package controllers

import (
	"crypto/rand"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/middlewares"
	"itfest-backend-2.0/models"
	"itfest-backend-2.0/types"
)

func generateUserCode() (string, error) {
	codes := make([]byte, 6)
	if _, err := rand.Read(codes); err != nil {
		return "", err
	}

	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	return string(codes), nil
}

func RegisterHandler(c echo.Context) error {
	db := configs.DB
	config := configs.Config.GetMetadata()
	response := models.Response[string]{}

	// 0. Process body request
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		response.Message = "ERROR: BAD REQUEST"
		return c.JSON(http.StatusBadRequest, response)
	}

	// Search if username existed
	result := models.User{}
	username := request["username"].(string)
	condition := models.User{Username: username}
	if err := db.Find(&condition, &result).Error; err == nil {
		response.Message = "ERROR: USERNAME EXISTED"
		return c.JSON(http.StatusBadRequest, response)
	}

	if err != gorm.ErrRecordNotFound {
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}

	// TODO
	// // 2. Create new usercode
	// if userCode, err := generateUserCode(); err != nil {
	// 	response.Message = "ERROR: INTERNAL SERVER ERROR"
	// 	return c.JSON(http.StatusInternalServerError, response)
	// }

	// 3. Check if usercode existed
	// condition = models.User{Usercode: userCode}
	// err = db.Find(&condition, &result).Error;
	// while err != nil {
	// 	// 2. Create new usercode
	// 	if userCode, err := generateUserCode(); err != nil {
	// 		response.Message = "ERROR: INTERNAL SERVER ERROR"
	// 		return c.JSON(http.StatusInternalServerError, response)
	// 	}
	// }

	// Create User
	encryptedString := []byte(request["password"].(string))
	newUser := models.User{
		Username: request["username"].(string),
		Name:     request["name"].(string),
		Usercode: "123456",
		Password: encryptedString,
		Role:     types.User,
		Point:    0,
	}
	if err = db.Create(&newUser).Error; err != nil {
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
