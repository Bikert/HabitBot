package http

import (
	"HabitMuse/internal/config"
	"HabitMuse/internal/constants"
	"HabitMuse/internal/users"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"net/http"
	"net/url"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}

		var httpErr *HTTPError
		if ok := ginErrorAs(err, &httpErr); ok {
			log.Printf("Handled error: %s (code=%d)", httpErr.Message, httpErr.Code)
			c.JSON(httpErr.Code, gin.H{"error": httpErr.Message})
			return
		}

		log.Printf("Unhandled error: %v", err.Err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}

func ginErrorAs(err *gin.Error, target interface{}) bool {
	return err != nil && err.Err != nil && errorAs(err.Err, target)
}

func errorAs(err error, target interface{}) bool {
	switch t := target.(type) {
	case **HTTPError:
		var e *HTTPError
		if errors.As(err, &e) {
			*t = e
			return true
		}
	}
	return false
}

func LogRequestBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Body == nil {
			c.Next()
			return
		}

		log.Println("=== Request Headers ===")
		for key, values := range c.Request.Header {
			for _, value := range values {
				log.Printf("%s: %s\n", key, value)
			}
		}

		log.Println("=== END ===")
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("Failed to read request body:", err)
			c.Next()
			return
		}

		log.Printf("Request Body: %s", string(bodyBytes))
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()
	}
}

func ValidationToken(userService users.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		fmt.Println("Path = ", path)

		var token, telegramInitData string
		token = config.Get().TGToken
		telegramInitData = c.Request.Header.Get("x-telegram-init-data")
		log.Println("telegramInitData = ", telegramInitData)
		if telegramInitData == "" {
			log.Println("Telegram Init Data is empty")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Telegram Init Data is empty"})
			return
		}
		isValid, err := tgBotAPI.ValidateWebAppData(token, telegramInitData)
		if !isValid || err != nil {
			log.Println("Invalid token err = %w", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		tgUser, err := getUserFromTgInitData(telegramInitData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Init User exception"})
		}
		user, err := userService.Get(tgUser.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		}

		c.Set(constants.UserContextKey, user)
		c.Next()
	}
}

func getUserFromTgInitData(telegramInitData string) (*tgBotAPI.User, error) {
	initData, err := url.ParseQuery(telegramInitData)
	if err != nil {
		return nil, err
	}
	userStr := initData.Get("user")
	if userStr == "" {
		return nil, errors.New("telegram init data is empty")
	}
	var user tgBotAPI.User
	err = json.Unmarshal([]byte(userStr), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
