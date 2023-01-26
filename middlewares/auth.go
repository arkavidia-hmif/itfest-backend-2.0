package middlewares

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/types"
)

type AuthRole string

const (
	Admin AuthRole = "Admin"
	Team  AuthRole = "Team"
)

type AuthClaims struct {
	jwt.RegisteredClaims
	ID   uint       `json:"id"`
	Role types.Role `json:"role"`
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("from middleware one")

		// config := configs.Config.GetMetadata()
		// response := models.Response[string]{}

		authHeader := c.Get("Authorization")
		fmt.Println(authHeader)
		// if !strings.Contains(authHeader, "Bearer") {
		// 	response.Message = "ERROR: NO TOKEN PROVIDED"
		// 	c.JSON(http.StatusUnauthorized, response)
		// }

		// authString := strings.Replace(authHeader, "Bearer ", "", -1)
		// authClaim := AuthClaims{}
		// authToken, err := jwt.ParseWithClaims(authString, &authClaim, func(authToken *jwt.Token) (interface{}, error) {
		// 	if method, ok := authToken.Method.(*jwt.SigningMethodHMAC); !ok || method != config.JWTSigningMethod {
		// 		return nil, fmt.Errorf("ERROR: SIGNING METHOD INVALID")
		// 	}
		// 	return config.JWTSignatureKey, nil
		// })
		// if err != nil {
		// 	response.Message = "ERROR: TOKEN CANNOT BE PARSED"
		// 	c.JSON(http.StatusInternalServerError, response)
		// }
		// if !authToken.Valid {
		// 	response.Message = "ERROR: CLAIMS INVALID"
		// 	c.JSON(http.StatusBadRequest, response)
		// }

		// c.Set("id", authClaim.ID)
		// c.Set("role", authClaim.Role)

		return next(c)
	}

	// c.Next()

}
