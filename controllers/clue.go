package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
	"itfest-backend-2.0/types"
)

type CreateClueRequest struct {
	Usercode string `json:"usercode" form:"usercode" query:"usercode"`
	Text     string `json:"text" form:"text" query:"text"`
}

type ClueResponse struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

func getRandomClueId() uint {
	return 1
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

	return c.JSON(http.StatusOK, "hi")
}
