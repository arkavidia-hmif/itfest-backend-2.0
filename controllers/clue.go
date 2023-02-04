package controllers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
	"itfest-backend-2.0/services"
	"itfest-backend-2.0/types"
)

type CreateClueRequest struct {
	Usercode string `json:"usercode" form:"usercode" query:"usercode"`
	Text     string `json:"text" form:"text" query:"text"`
}

type SubmitClueRequest struct {
	ID   string `json:"id" form:"id" query:"id"`
	Code string `json:"code" form:"code" query:"code"`
}

type ClueResponse struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

func GetTriesHandler(c echo.Context) error {
	db := configs.DB.GetConnection()
	response := models.Response[uint]{}

	userID := c.Get("id").(uint)
	result := models.Game{}
	condition := models.Game{UserID: userID}
	if err := db.Where(&condition).Find(&result).Error; err != nil {
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "SUCCESS"
	response.Data = result.RemainingTries
	return c.JSON(http.StatusOK, response)
}

// for admin only
func CreateClueHandler(c echo.Context) error {
	db := configs.DB.GetConnection()
	role := c.Get("role")
	response := models.Response[string]{}

	if role != types.Admin {
		response.Message = "ERROR: UNAUTHORIZED"
		return c.JSON(http.StatusUnauthorized, response)
	}

	request := CreateClueRequest{}
	if err := c.Bind(&request); err != nil {
		response.Message = "ERROR: BAD REQUEST"
		return c.JSON(http.StatusBadRequest, response)
	}

	newClue := models.Clue{
		Usercode: request.Usercode,
		Text:     request.Text,
	}
	if err := db.Create(&newClue).Error; err != nil {
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "SUCCESS"
	return c.JSON(http.StatusCreated, response)

}

func ClueHandler(c echo.Context) error {
	db := configs.DB.GetConnection()
	response := models.Response[ClueResponse]{}
	id := c.Get("id").(uint)

	game := models.Game{}
	condition := models.Game{UserID: id}
	if err := db.Where(&condition).Find(&game).Error; err != nil {
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}

	result := models.Clue{}
	clueId := game.CurrentClueId
	if err := db.First(&result, clueId).Error; err != nil {
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "SUCCESS"
	response.Data = ClueResponse{
		ID:   result.ID,
		Text: result.Text,
	}
	return c.JSON(http.StatusOK, response)
}

func SubmitClueHandler(c echo.Context) error {
	db := configs.DB.GetConnection()
	response := models.Response[string]{}
	id := c.Get("id").(uint)

	request := SubmitClueRequest{}
	if err := c.Bind(&request); err != nil {
		response.Message = "ERROR: BAD REQUEST"
		return c.JSON(http.StatusBadRequest, response)
	}

	game := models.Game{}
	if err := db.First(&game, request.ID).Error; err != nil {
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}

	clue := models.Clue{}
	if err := db.First(&clue, game.CurrentClueId).Error; err != nil {
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}

	// if code is true
	if request.Code == clue.Usercode {
		points, perr := strconv.Atoi(os.Getenv("CLUE_SUCCESS_POINT"))

		if perr != nil {
			return perr
		}

		admin := models.User{}
		if err := db.Where(models.User{Role: types.Admin}).First(&admin).Error; err != nil {
			return err
		}

		_, err := services.GrantPoint(admin.ID, id, uint(points))
		if err != nil {
			return err
		}

		
	}

	if game.RemainingTries == 1 {

	}

	return c.JSON(http.StatusOK, "hi")
}
