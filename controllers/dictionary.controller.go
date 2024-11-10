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
//	@Router			/dict/cities [get]
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

// ListEthnos godoc
//
//	@Summary		Lists all ethnos with pagination, auth required
//	@Description	Retrieves all ethnos, supports pagination
//	@Tags			Ethnos
//	@Produce		json
//	@Param			page	query		string	false	"Page number"
//	@Param			limit	query		string	false	"Items per page"
//	@Success		200		{object}	SuccessPageResponse[EthnosResponse[]]
//	@Failure		502		{object}	ErrorResponse
//	@Router			/dict/ethnos [get]
func (pc *DictionaryController) ListEthnos(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var ethnosList []Ethnos

	dbQuery := pc.DB.
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&ethnosList)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: results.Error.Error()})
		return
	}

	response := make([]EthnosResponse, len(ethnosList))
	for i, ethnos := range ethnosList {
		response[i] = *utils.MapEthnos(&ethnos)
	}

	ctx.JSON(http.StatusOK, SuccessPageResponse[[]EthnosResponse]{
		Status:  "success",
		Data:    response,
		Results: len(ethnosList),
		Page:    intPage,
	})
}

// ListBodyTypes godoc
//
//	@Summary		Lists all body types with pagination, auth required
//	@Description	Retrieves all body types, supports pagination
//	@Tags			Body Types
//	@Produce		json
//	@Param			page	query		string	false	"Page number"
//	@Param			limit	query		string	false	"Items per page"
//	@Success		200		{object}	SuccessPageResponse[BodyTypeResponse[]]
//	@Failure		502		{object}	ErrorResponse
//	@Router			/dict/bodies [get]
func (pc *DictionaryController) ListBodyTypes(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var bodyTypes []BodyType

	dbQuery := pc.DB.
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&bodyTypes)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: results.Error.Error()})
		return
	}

	response := make([]BodyTypeResponse, len(bodyTypes))
	for i, bodyType := range bodyTypes {
		response[i] = *utils.MapBodyType(&bodyType)
	}

	ctx.JSON(http.StatusOK, SuccessPageResponse[[]BodyTypeResponse]{
		Status:  "success",
		Data:    response,
		Results: len(bodyTypes),
		Page:    intPage,
	})
}

// ListBodyArts godoc
//
//	@Summary		Lists all body arts with pagination, auth required
//	@Description	Retrieves all body arts, supports pagination
//	@Tags			Body Arts
//	@Produce		json
//	@Param			page	query		string	false	"Page number"
//	@Param			limit	query		string	false	"Items per page"
//	@Success		200		{object}	SuccessPageResponse[BodyArtResponse[]]
//	@Failure		502		{object}	ErrorResponse
//	@Router			/dict/arts [get]
func (pc *DictionaryController) ListBodyArts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var bodyArts []BodyArt

	dbQuery := pc.DB.
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&bodyArts)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: results.Error.Error()})
		return
	}

	response := make([]BodyArtResponse, len(bodyArts))
	for i, bodyArt := range bodyArts {
		response[i] = *utils.MapBodyArt(&bodyArt)
	}

	ctx.JSON(http.StatusOK, SuccessPageResponse[[]BodyArtResponse]{
		Status:  "success",
		Data:    response,
		Results: len(bodyArts),
		Page:    intPage,
	})
}

// ListHairColors godoc
//
//	@Summary		Lists all hair colors with pagination, auth required
//	@Description	Retrieves all hair colors, supports pagination
//	@Tags			Hair Colors
//	@Produce		json
//	@Param			page	query		string	false	"Page number"
//	@Param			limit	query		string	false	"Items per page"
//	@Success		200		{object}	SuccessPageResponse[HairColorResponse[]]
//	@Failure		502		{object}	ErrorResponse
//	@Router			/dict/colors [get]
func (pc *DictionaryController) ListHairColors(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var hairColors []HairColor

	dbQuery := pc.DB.
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&hairColors)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: results.Error.Error()})
		return
	}

	response := make([]HairColorResponse, len(hairColors))
	for i, hairColor := range hairColors {
		response[i] = *utils.MapHairColor(&hairColor)
	}

	ctx.JSON(http.StatusOK, SuccessPageResponse[[]HairColorResponse]{
		Status:  "success",
		Data:    response,
		Results: len(hairColors),
		Page:    intPage,
	})
}

// ListIntimateHairCuts godoc
//
//	@Summary		Lists all intimate hair cuts with pagination, auth required
//	@Description	Retrieves all intimate hair cuts, supports pagination
//	@Tags			Intimate Hair Cuts
//	@Produce		json
//	@Param			page	query		string	false	"Page number"
//	@Param			limit	query		string	false	"Items per page"
//	@Success		200		{object}	SuccessPageResponse[IntimateHairCutResponse[]]
//	@Failure		502		{object}	ErrorResponse
//	@Router			/dict/cuts [get]
func (pc *DictionaryController) ListIntimateHairCuts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var intimateHairCuts []IntimateHairCut

	dbQuery := pc.DB.
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&intimateHairCuts)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: results.Error.Error()})
		return
	}

	response := make([]IntimateHairCutResponse, len(intimateHairCuts))
	for i, hairCut := range intimateHairCuts {
		response[i] = *utils.MapIntimateHairCut(&hairCut)
	}

	ctx.JSON(http.StatusOK, SuccessPageResponse[[]IntimateHairCutResponse]{
		Status:  "success",
		Data:    response,
		Results: len(intimateHairCuts),
		Page:    intPage,
	})
}
