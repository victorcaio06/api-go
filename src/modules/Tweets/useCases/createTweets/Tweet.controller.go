package createTweets

import (
	"fundamentals/src/modules/Tweets/entities"

	"net/http"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []entities.Tweet
}

func CreateTweetController() *tweetController {
	return &tweetController{}
}

func (twt *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, twt.tweets)
}

func (twt *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		return
	}

	twt.tweets = append(twt.tweets, *tweet)

	ctx.JSON(http.StatusOK, twt.tweets)
}

func (twt *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for index, tweet := range twt.tweets {
		if tweet.ID ==	id {
			twt.tweets = append(twt.tweets[0:index], twt.tweets[index+1:]...)
			return
		}
	}
	
	ctx.JSON(http.StatusNotFound, gin.H{
		"ERROR": "Tweet not found!",
	})
}
