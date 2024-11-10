package controllers

import (
	"github.com/gin-gonic/gin"
	. "github.com/ivegotanidea/golang-gorm-postgres/models"
	"github.com/ivegotanidea/golang-gorm-postgres/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type DictionaryController struct {
	DB *gorm.DB
}

func NewDictionaryController(DB *gorm.DB) DictionaryController {
	return DictionaryController{DB}
}

// ListCities godoc
//
//	@Summary		Lists all cities with pagination, auth required
//	@Description	Retrieves all cities, supports pagination
//	@Tags			Cities
//	@Produce		json
//	@Param			page	query		string	false	"Page number"
//	@Param			limit	query		string	false	"Items per page"
//	@Success		200		{object}	SuccessPageResponse[CityResponse[]]
//	@Failure		502		{object}	ErrorResponse
//	@Router			/profiles/all [get]
func (pc *DictionaryController) ListCities(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var cities []City

	dbQuery := pc.DB.
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&cities)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: results.Error.Error()})
		return
	}

	response := make([]CityResponse, len(cities))
	for i, city := range cities {
		response[i] = *utils.MapCity(&city) // Assuming you have the mapDictionary function
	}

	ctx.JSON(http.StatusOK, SuccessPageResponse[[]CityResponse]{
		Status:  "success",
		Data:    response,
		Results: len(cities),
		Page:    intPage,
	})
}
