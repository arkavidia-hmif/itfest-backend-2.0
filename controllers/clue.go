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
	Code string `json:"code" form:"code" query:"code"`
}

type ClueResponse struct {
	ID             uint   `json:"id"`
	Clue           string `json:"text"`
	RemainingTries uint   `json:"remaining_tries"`
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

	// if game already done
	if game.CurrentClueId == 9999 {
		response.Message = "SUCCESS: GAME IS DONE!"
		return c.JSON(http.StatusOK, response)
	}

	result := models.Clue{}
	clueId := game.CurrentClueId
	if err := db.First(&result, clueId).Error; err != nil {
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "SUCCESS"
	response.Data = ClueResponse{
		ID:             result.ID,
		Clue:           result.Text,
		RemainingTries: game.RemainingTries,
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
	if err := db.Where(&models.Game{UserID: id}).Find(&game).Error; err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, response)
	}

	// if game already done
	if game.CurrentClueId == 9999 {
		response.Message = "SUCCESS: GAME IS DONE!"
		return c.JSON(http.StatusOK, response)
	}

	clue := models.Clue{}
	if err := db.First(&clue, game.CurrentClueId).Error; err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, response)
	}

	// if code is true
	if request.Code == clue.Usercode {
		if err := services.NewGame(game); err != nil {
			response.Message = err.Error()
			return c.JSON(http.StatusInternalServerError, response)
		}

		admin := models.User{}
		if err := db.Where(models.User{Role: types.Admin}).First(&admin).Error; err != nil {
			response.Message = err.Error()
			return c.JSON(http.StatusInternalServerError, response)
		}

		points, perr := strconv.Atoi(os.Getenv("CLUE_SUCCESS_POINT"))
		if perr != nil {
			response.Message = perr.Error()
			return c.JSON(http.StatusInternalServerError, response)
		}

		_, err := services.GrantPoint(admin.ID, id, uint(points))
		if err != nil {
			response.Message = err.Error()
			return c.JSON(http.StatusInternalServerError, response)
		}

		response.Message = "SUCCESS: ANSWER IS CORRECT"
		return c.JSON(http.StatusOK, response)
	}

	// if code is wrong and remaining tries is 1
	if game.RemainingTries == 1 {
		if err := services.NewGame(game); err != nil {
			response.Message = err.Error()
			return c.JSON(http.StatusInternalServerError, response)
		}

		response.Message = "SUCCESS: TRIED 3 TIMES"
		return c.JSON(http.StatusOK, response)
	}

	// code is wrong, reduce try -1
	if err := db.First(&models.Game{}, game.ID).Updates(models.Game{
		RemainingTries: game.RemainingTries - 1,
	}).Error; err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "SUCCESS: ANSWER IS WRONG"
	return c.JSON(http.StatusOK, response)
}
