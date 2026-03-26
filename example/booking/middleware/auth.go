package middleware

// executes before the actual function

import (
	"booking/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	fmt.Println(userId)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized !"})
		return
	}

	context.Set("userId", userId)
	context.Next()

}
