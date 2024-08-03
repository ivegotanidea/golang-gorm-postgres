package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"net/http"
)

func AbacMiddleware(obj string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exists := c.Get("currentUserID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		// Fetch the user's tier
		tier, exists := c.Get("currentUserTier")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User tier not found"})
			c.Abort()
			return
		}

		// Check if the user has permission
		ok, err := initializers.Enforcer.Enforce(tier, obj, act, tier)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while checking permissions"})
			c.Abort()
			return
		}

		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}
