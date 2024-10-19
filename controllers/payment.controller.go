package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ivegotanidea/golang-gorm-postgres/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type PaymentController struct {
	DB                     *gorm.DB
	PaymentProviderApiKey  string
	PaymentProviderBaseURL string
}

func NewPaymentController(DB *gorm.DB, apiKey string, baseUrl string) PaymentController {
	return PaymentController{DB, apiKey, baseUrl}
}

// PaymentWebhook godoc
// @Summary Webhook for payment updates
// @Description Receives payment updates and updates the payment status in the database.
// @Tags Payments
// @Accept json
// @Produce json
// @Param body body models.Payment true "Payment Update"
// @Success 200 {object} models.SuccessResponse{data=string} "payment updated"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /payments/webhook [post]
func (pc *PaymentController) PaymentWebhook(ctx *gin.Context) {
	var paymentUpdate models.Payment

	// Bind the incoming JSON data to the payment model
	if err := ctx.ShouldBindJSON(&paymentUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid data",
		})
		return
	}

	// Update payment in the database
	if err := pc.DB.Model(&models.Payment{}).Where("id = ?", paymentUpdate.ID).Updates(paymentUpdate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to update payment",
		})
		return
	}

	// Return success response
	ctx.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   "payment updated",
	})
}

// GetPaymentHistory godoc
// @Summary Get payment history for a user
// @Description Retrieves the payment history for a specified user between two dates.
// @Tags Payments
// @Accept json
// @Produce json
// @Param userID path string true "User ID"
// @Param start_date query string true "Start Date in RFC3339 format"
// @Param end_date query string true "End Date in RFC3339 format"
// @Success 200 {object} models.SuccessResponse{data=[]models.Payment}
// @Failure 500 {object} models.ErrorResponse
// @Router /payments/history/{userID} [get]
func (pc *PaymentController) GetPaymentHistory(ctx *gin.Context) {
	// Get userID from path and date range from query parameters
	userID := ctx.Param("userID")
	startDate, _ := time.Parse(time.RFC3339, ctx.Query("start_date"))
	endDate, _ := time.Parse(time.RFC3339, ctx.Query("end_date"))

	var payments []models.Payment

	// Fetch payments from the database for the given user and date range
	if err := pc.DB.Where("user_id = ? AND payment_date BETWEEN ? AND ?", userID, startDate, endDate).Find(&payments).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve payments",
		})
		return
	}

	// Return success response with the payment data
	ctx.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   payments,
	})
}

// ListPayments godoc
// @Summary List all payments
// @Description Retrieves all payments, sorted by payment date in descending order with pagination.
// @Tags Payments
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Limit per page" default(10)
// @Success 200 {object} models.SuccessResponse{data=[]models.Payment}
// @Failure 500 {object} models.ErrorResponse
// @Router /payments [get]
func (pc *PaymentController) ListPayments(ctx *gin.Context) {
	// Pagination parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var payments []models.Payment

	// Retrieve payments with sorting and pagination
	result := pc.DB.Order("payment_date DESC").
		Limit(limit).Offset(offset).Find(&payments)

	// Check if any error occurred during the query
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve payments",
		})
		return
	}

	// Return the payments in the response
	ctx.JSON(http.StatusOK, models.SuccessPageResponse{
		Status:  "success",
		Results: len(payments),
		Page:    page,
		Limit:   limit,
		Data:    payments,
	})
}

// GetMyPayments godoc
// @Summary Get current user's payments
// @Description Retrieves the payments made by the current user, sorted by payment date in descending order with pagination.
// @Tags Payments
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Limit per page" default(10)
// @Success 200 {object} models.SuccessResponse{data=[]models.Payment}
// @Failure 500 {object} models.ErrorResponse
// @Router /payments/me [get]
func (pc *PaymentController) GetMyPayments(ctx *gin.Context) {
	// Get the current user from context (assumes middleware has set currentUser)
	currentUser := ctx.MustGet("currentUser").(models.User)

	// Pagination parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var payments []models.Payment

	// Retrieve payments for the current user with sorting and pagination
	result := pc.DB.Where("user_id = ?", currentUser.ID).
		Order("payment_date DESC").
		Limit(limit).Offset(offset).Find(&payments)

	// Check if any error occurred during the query
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve user payments",
		})
		return
	}

	// Return the user's payments in the response
	ctx.JSON(http.StatusOK, models.SuccessPageResponse{
		Status:  "success",
		Page:    page,
		Limit:   limit,
		Results: len(payments),
		Data:    payments,
	})
}
