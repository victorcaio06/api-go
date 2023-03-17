package routes

import (
	"fundamentals/src/modules/Tweets/useCases/createTweets"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := createTweets.CreateTweetController()
	v1 := router.Group("/v1")
	{
		v1.GET("/tweets", tweetController.FindAll)
		v1.POST("/tweet", tweetController.Create)
	}

	return v1
}
