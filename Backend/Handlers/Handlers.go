package HandlerFunctions

import (
	"github.com/gin-gonic/gin"
)
func Login(context *gin.Context){
	context.JSON(200, gin.H{
		"message": "This is the logind endpoint placeholder",
	})
}

func Register(context *gin.Context){
	context.JSON(200, gin.H{
		"message": "This is the registration endpoint placeholder",
	})
}
