package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	if userType != role {
		return errors.New("Unauthorized to access this resource")
	}
	return nil
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	if userType == "USER" && uid != userId {
		return errors.New("Unauthorized to access this resource")
	}
	return CheckUserType(c, userType)
}
