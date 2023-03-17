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
