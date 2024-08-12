package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserFromRequest(c *gin.Context) (uuid.UUID, error) {
	user, exists := c.Get("userID")

	if !exists {
		// return null uuid
		return uuid.Nil, errors.New("user id not found")
	}

	//convert user to uuid
	userIDStr, ok := user.(string)

	if !ok {
		return uuid.Nil, errors.New("invalid user id")
	}

	userID, err := uuid.Parse(userIDStr)

	if err != nil {
		return uuid.Nil, errors.New("invalid user id")
	}

	return userID, nil

}
