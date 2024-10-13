package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/models"
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
			Status:  "fail",
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

func (pc *PaymentController) GetPaymentHistory(ctx *gin.Context) {
	userID := ctx.Param("userID")
	startDate, _ := time.Parse(time.RFC3339, ctx.Query("start_date"))
	endDate, _ := time.Parse(time.RFC3339, ctx.Query("end_date"))

	var payments []models.Payment
	if err := pc.DB.Where("user_id = ? AND payment_date BETWEEN ? AND ?", userID, startDate, endDate).Find(&payments).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payments"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"payments": payments})
}

func (pc *PaymentController) ListPayments(ctx *gin.Context) {
	// Pagination parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var payments []models.Payment
	result := pc.DB.Order("payment_date DESC"). // Sort by payment_date in descending order
							Limit(limit).Offset(offset).Find(&payments)

	// Check if any error occurred during the query
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to retrieve payments"})
		return
	}

	// Return the payments in the response
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(payments), "data": payments})
}

func (pc *PaymentController) GetMyPayments(ctx *gin.Context) {
	// Get the current user from context (assumes middleware has set currentUser)
	currentUser := ctx.MustGet("currentUser").(models.User)

	// Pagination parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var payments []models.Payment
	result := pc.DB.Where("user_id = ?", currentUser.ID). // Filter payments by user_id
								Order("payment_date DESC"). // Sort by payment_date in descending order
								Limit(limit).Offset(offset).Find(&payments)

	// Check if any error occurred during the query
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to retrieve user payments"})
		return
	}

	// Return the user's payments in the response
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(payments), "data": payments})
}
