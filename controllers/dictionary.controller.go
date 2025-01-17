package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/ivegotanidea/golang-gorm-postgres/models"
	"github.com/ivegotanidea/golang-gorm-postgres/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type DictionaryController struct {
	DB *gorm.DB
}

func NewDictionaryController(DB *gorm.DB) DictionaryController {
	return DictionaryController{DB}
}

func (pc *DictionaryController) CreateDict(ctx *gin.Context) {
	var payload CreateDictRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid JSON payload: " + err.Error()})
		return
	}

	dictType := payload.Type

	currentUser := ctx.MustGet("currentUser").(User)
	now := time.Now()

	switch dictType {
	case "city":
		var data City
		if err := mapToStruct(payload.Data, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid City data: " + err.Error()})
			return
		}
		data.CreatedAt = now
		data.CreatedBy = currentUser.ID
		if err := pc.DB.Create(&data).Error; err != nil {

			if strings.Contains(err.Error(), "duplicate") {
				ctx.JSON(http.StatusConflict, ErrorResponse{Status: "error", Message: "Duplicating entity"})
			} else {
				ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: "Internal server error"})
			}
			return
		}
	case "ethnos":
		var data Ethnos
		if err := mapToStruct(payload.Data, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid Ethnos data: " + err.Error()})
			return
		}
		data.CreatedAt = now
		data.CreatedBy = currentUser.ID
		if err := pc.DB.Create(&data).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: err.Error()})
			return
		}
	case "profileTag":
		var data ProfileTag
		if err := mapToStruct(payload.Data, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid ProfileTag data: " + err.Error()})
			return
		}
		data.CreatedAt = now
		data.CreatedBy = currentUser.ID
		if err := pc.DB.Create(&data).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: err.Error()})
			return
		}
	case "userTag":
		var data UserTag
		if err := mapToStruct(payload.Data, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid UserTag data: " + err.Error()})
			return
		}
		data.CreatedAt = now
		data.CreatedBy = currentUser.ID
		if err := pc.DB.Create(&data).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: err.Error()})
			return
		}
	case "body":
		var data BodyType
		if err := mapToStruct(payload.Data, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid BodyType data: " + err.Error()})
			return
		}
		data.CreatedAt = now
		data.CreatedBy = currentUser.ID
		if err := pc.DB.Create(&data).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: err.Error()})
			return
		}
	case "art":
		var data BodyArt
		if err := mapToStruct(payload.Data, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid BodyArt data: " + err.Error()})
			return
		}
		data.CreatedAt = now
		data.CreatedBy = currentUser.ID
		if err := pc.DB.Create(&data).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: err.Error()})
			return
		}
	case "color":
		var data HairColor
		if err := mapToStruct(payload.Data, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid HairColor data: " + err.Error()})
			return
		}
		data.CreatedAt = now
		data.CreatedBy = currentUser.ID
		if err := pc.DB.Create(&data).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: err.Error()})
			return
		}
	case "cut":
		var data IntimateHairCut
		if err := mapToStruct(payload.Data, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid IntimateHairCut data: " + err.Error()})
			return
		}
		data.CreatedAt = now
		data.CreatedBy = currentUser.ID
		if err := pc.DB.Create(&data).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: err.Error()})
			return
		}
	default:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Unsupported dictionary type: " + dictType})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func mapToStruct(input any, output any) error {
	bytes, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, output)
}

// ListDict godoc
//
//	@Summary		Lists all dict objects with pagination, auth required
//	@Description	Retrieves all dict objects, supports pagination
//	@Tags			Dict
//	@Produce		json
//	@Param			page	query		string	false	"Page number"
//	@Param			limit	query		string	false	"Items per page"
//	@Success		200		{object}	SuccessPageResponse
//	@Failure		502		{object}	ErrorResponse
//	@Router			/dict [get]
func (pc *DictionaryController) ListDict(ctx *gin.Context) {

	var dictType = ctx.DefaultQuery("type", "")

	if dictType == "city" {
		pc.ListCities(ctx)
	} else if dictType == "ethnos" {
		pc.ListEthnos(ctx)
	} else if dictType == "body" {
		pc.ListBodyTypes(ctx)
	} else if dictType == "art" {
		pc.ListBodyArts(ctx)
	} else if dictType == "color" {
		pc.ListHairColors(ctx)
	} else if dictType == "cut" {
		pc.ListIntimateHairCuts(ctx)
	} else if dictType == "userTag" {
		pc.ListUserTags(ctx)
	} else if dictType == "profileTag" {
		pc.ListProfileTags(ctx)
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, nil)
	}
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
//	@Param			sex		query		string	female	"Sex"
//	@Success		200		{object}	SuccessPageResponse[EthnosResponse[]]
//	@Failure		502		{object}	ErrorResponse
//	@Router			/dict/ethnos [get]
func (pc *DictionaryController) ListEthnos(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")
	var sex = ctx.DefaultQuery("sex", "female")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var ethnosList []Ethnos

	dbQuery := pc.DB.
		Where("sex = ?", sex).
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

// ListProfileTags godoc
//
//	@Summary		Lists all profile tags with pagination, auth required
//	@Description	Retrieves all intimate hair cuts, supports pagination
//	@Tags			Profile Tags
//	@Produce		json
//	@Param			page	query		string	false	"Page number"
//	@Param			limit	query		string	false	"Items per page"
//	@Success		200		{object}	SuccessPageResponse[ProfileTagResponse[]]
//	@Failure		502		{object}	ErrorResponse
//	@Router			/dict/profile/tags [get]
func (pc *DictionaryController) ListProfileTags(ctx *gin.Context) {

	// to create proper profile tag that would be shown in api response
	// you shouldn't leave flags property empty
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var profileTags []ProfileTag

	dbQuery := pc.DB.
		Where("flags != '{}'").
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&profileTags)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: results.Error.Error()})
		return
	}

	response := make([]ProfileTagResponse, len(profileTags))
	for i, tag := range profileTags {
		response[i] = *utils.MapProfileTag(&tag)
	}

	ctx.JSON(http.StatusOK, SuccessPageResponse[[]ProfileTagResponse]{
		Status:  "success",
		Data:    response,
		Results: len(profileTags),
		Page:    intPage,
	})
}

// ListUserTags godoc
//
//	@Summary		Lists all user tags with pagination, auth required
//	@Description	Retrieves all intimate hair cuts, supports pagination
//	@Tags			User Tags
//	@Produce		json
//	@Param			page	query		string	false	"Page number"
//	@Param			limit	query		string	false	"Items per page"
//	@Success		200		{object}	SuccessPageResponse[UserTagResponse[]]
//	@Failure		502		{object}	ErrorResponse
//	@Router			/dict/user/tags [get]
func (pc *DictionaryController) ListUserTags(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var userTags []UserTag

	dbQuery := pc.DB.
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&userTags)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: results.Error.Error()})
		return
	}

	response := make([]UserTagResponse, len(userTags))
	for i, tag := range userTags {
		response[i] = UserTagResponse{
			ID:      tag.ID,
			Name:    tag.Name,
			AliasRu: tag.AliasRu,
			AliasEn: tag.AliasEn,
		}
	}

	ctx.JSON(http.StatusOK, SuccessPageResponse[[]UserTagResponse]{
		Status:  "success",
		Data:    response,
		Results: len(userTags),
		Page:    intPage,
	})
}
