package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ivegotanidea/golang-gorm-postgres/initializers"
	. "github.com/ivegotanidea/golang-gorm-postgres/models"
	"net/http"
)

func AbacMiddleware(obj string, act string) gin.HandlerFunc {

	return func(c *gin.Context) {

		user, exists := c.Get("currentUser")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		currentUser, ok := user.(User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting user"})
			c.Abort()
			return
		}

		initializers.Enforcer.EnableLog(true)

		hasProfileStr := "false"
		if currentUser.HasProfile {
			hasProfileStr = "true"
		}

		// Check if the user has permission
		ok, err := initializers.Enforcer.Enforce(
			currentUser.Role, // sub
			obj,              // obj
			act,              // act
			currentUser.Tier, // tier
			hasProfileStr)    // hasProfile

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
