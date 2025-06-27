package utils

import (
	"HabitMuse/internal/constants"
	"HabitMuse/internal/users"
	"github.com/gin-gonic/gin"
	"log"
)

func GetUserByCtx(c *gin.Context) *users.User {
	rawUser, exists := c.Get(constants.UserContextKey)

	if !exists {
		log.Println("user not found in context")
		return nil
	}

	log.Println(rawUser)
	user, ok := rawUser.(users.User)
	if !ok {
		log.Println("failed to cast user from context")
		return nil
	}
	return &user
}
